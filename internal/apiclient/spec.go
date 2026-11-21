// Package apiclient reads API descriptions without generating client source code.
package apiclient

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pb33f/go-yaml"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/pb33f/libopenapi/orderedmap"
	"github.com/pb33f/libopenapi/renderer"
)

type Parameter struct {
	Name             string `json:"name"`
	In               string `json:"in"`
	Required         bool   `json:"required"`
	Style            string `json:"style,omitempty"`
	Explode          bool   `json:"explode"`
	CollectionFormat string `json:"collectionFormat,omitempty"`
}

type Operation struct {
	ID              string      `json:"operationId"`
	Method          string      `json:"method"`
	Path            string      `json:"path"`
	Summary         string      `json:"summary,omitempty"`
	Description     string      `json:"description,omitempty"`
	Deprecated      bool        `json:"deprecated"`
	BaseURL         string      `json:"baseUrl,omitempty"`
	Parameters      []Parameter `json:"parameters,omitempty"`
	BodyRequired    bool        `json:"bodyRequired"`
	ContentType     string      `json:"contentType,omitempty"`
	body            any
	unsupportedBody bool
}

type Spec struct{ Operations []Operation }

// Load resolves in-document references. External references must be bundled first.
func Load(filename string) (*Spec, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read spec: %w", err)
	}
	doc, err := libopenapi.NewDocument(data)
	if err != nil {
		return nil, fmt.Errorf("parse spec: %w", err)
	}
	spec := &Spec{}
	switch {
	case doc.GetVersion() == "2.0":
		model, err := doc.BuildV2Model()
		if err != nil {
			return nil, fmt.Errorf("build Swagger model (bundle external references first): %w", err)
		}
		m := model.Model
		baseURL := m.BasePath
		if m.Host != "" {
			scheme := "https"
			if len(m.Schemes) > 0 {
				scheme = m.Schemes[0]
			}
			baseURL = scheme + "://" + m.Host + m.BasePath
		}
		if m.Paths == nil {
			return nil, fmt.Errorf("spec has no paths")
		}
		for path, item := range m.Paths.PathItems.FromOldest() {
			for method, op := range item.GetOperations().FromOldest() {
				o := Operation{ID: op.OperationId, Method: strings.ToUpper(method), Path: path, Summary: op.Summary, Description: op.Description, Deprecated: op.Deprecated, BaseURL: baseURL}
				consumes := op.Consumes
				if len(consumes) == 0 {
					consumes = m.Consumes
				}
				if len(op.Schemes) > 0 && m.Host != "" {
					o.BaseURL = op.Schemes[0] + "://" + m.Host + m.BasePath
				}
				for _, p := range append(append([]*v2.Parameter{}, item.Parameters...), op.Parameters...) {
					if p.In == "body" {
						o.BodyRequired = p.Required != nil && *p.Required
						o.body = &mockBody{Schema: p.Schema}
						o.ContentType = jsonMediaType(consumes)
						o.unsupportedBody = o.ContentType == ""
					} else {
						o.addParameter(Parameter{Name: p.Name, In: p.In, Required: p.Required != nil && *p.Required, CollectionFormat: p.CollectionFormat})
						if p.In == "formData" {
							o.unsupportedBody = true
						}
					}
				}
				spec.Operations = append(spec.Operations, o)
			}
		}
	case strings.HasPrefix(doc.GetVersion(), "3."):
		model, err := doc.BuildV3Model()
		if err != nil {
			return nil, fmt.Errorf("build OpenAPI model (bundle external references first): %w", err)
		}
		m := model.Model
		if m.Paths == nil {
			return nil, fmt.Errorf("spec has no paths")
		}
		for path, item := range m.Paths.PathItems.FromOldest() {
			for method, op := range item.GetOperations().FromOldest() {
				servers := op.Servers
				if len(servers) == 0 {
					servers = item.Servers
				}
				if len(servers) == 0 {
					servers = m.Servers
				}
				o := Operation{ID: op.OperationId, Method: strings.ToUpper(method), Path: path, Summary: op.Summary, Description: op.Description, Deprecated: op.Deprecated != nil && *op.Deprecated, BaseURL: serverURL(servers)}
				for _, p := range append(append([]*v3.Parameter{}, item.Parameters...), op.Parameters...) {
					style := p.Style
					if style == "" {
						if p.In == "query" || p.In == "cookie" {
							style = "form"
						} else {
							style = "simple"
						}
					}
					explode := style == "form"
					if p.Explode != nil {
						explode = *p.Explode
					}
					o.addParameter(Parameter{Name: p.Name, In: p.In, Required: p.Required != nil && *p.Required, Style: style, Explode: explode})
				}
				if op.RequestBody != nil {
					o.BodyRequired = op.RequestBody.Required != nil && *op.RequestBody.Required
					o.unsupportedBody = true
					if op.RequestBody.Content != nil {
						var types []string
						for ct := range op.RequestBody.Content.FromOldest() {
							types = append(types, ct)
						}
						o.ContentType = jsonMediaType(types)
						if o.ContentType != "" {
							o.body = op.RequestBody.Content.GetOrZero(o.ContentType)
							o.unsupportedBody = false
						}
					}
				}
				spec.Operations = append(spec.Operations, o)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported spec version %q; expected Swagger 2.0 or OpenAPI 3.x", doc.GetVersion())
	}
	sort.Slice(spec.Operations, func(i, j int) bool {
		a, b := spec.Operations[i], spec.Operations[j]
		if a.Path == b.Path {
			return a.Method < b.Method
		}
		return a.Path < b.Path
	})
	return spec, nil
}

func (o *Operation) addParameter(p Parameter) {
	for i, old := range o.Parameters {
		if old.In == p.In && old.Name == p.Name {
			o.Parameters[i] = p
			return
		}
	}
	o.Parameters = append(o.Parameters, p)
}

func (s *Spec) Find(selector string) (Operation, error) {
	var matches []Operation
	for _, o := range s.Operations {
		if o.Method+" "+o.Path == selector {
			return o, nil
		}
		if o.ID != "" && o.ID == selector {
			matches = append(matches, o)
		}
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	if len(matches) > 1 {
		var choices []string
		for _, o := range matches {
			choices = append(choices, fmt.Sprintf("%q", o.Method+" "+o.Path))
		}
		return Operation{}, fmt.Errorf("ambiguous operationId %q; select %s", selector, strings.Join(choices, " or "))
	}
	return Operation{}, fmt.Errorf("operation %q not found; use 'adb api list'", selector)
}

func (o Operation) Example(seed int64) ([]byte, error) {
	if o.unsupportedBody {
		return nil, fmt.Errorf("%s %s: only JSON request bodies are supported", o.Method, o.Path)
	}
	if o.body == nil {
		return nil, fmt.Errorf("%s %s has no request body; use --path and --query for parameters", o.Method, o.Path)
	}
	generator := renderer.NewMockGenerator(renderer.JSON)
	generator.SetSeed(seed)
	generator.SetPretty()
	generator.DisableRequiredCheck()
	data, err := generator.GenerateMock(o.body, "")
	if err != nil {
		return nil, fmt.Errorf("generate example: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("request body has no usable example or schema")
	}
	return data, nil
}

// The renderer accepts media-type-like objects for both Swagger and OpenAPI.
type mockBody struct {
	Example  *yaml.Node
	Examples *orderedmap.Map[string, *base.Example]
	Schema   *base.SchemaProxy
}

func jsonMediaType(types []string) string {
	if len(types) == 0 {
		return "application/json"
	}
	for _, ct := range types {
		if ct == "application/json" {
			return ct
		}
	}
	sort.Strings(types)
	for _, ct := range types {
		if strings.HasSuffix(ct, "+json") {
			return ct
		}
	}
	return ""
}

func serverURL(servers []*v3.Server) string {
	if len(servers) == 0 {
		return ""
	}
	s := servers[0]
	url := s.URL
	if s.Variables != nil {
		for name, variable := range s.Variables.FromOldest() {
			url = strings.ReplaceAll(url, "{"+name+"}", variable.Default)
		}
	}
	return url
}

// HasJSONBody reports whether this operation supports JSON body examples.
func (o Operation) HasJSONBody() bool { return o.body != nil && !o.unsupportedBody }

// HasRequestBody includes unsupported form bodies as well as JSON bodies.
func (o Operation) HasRequestBody() bool { return o.body != nil || o.unsupportedBody }
