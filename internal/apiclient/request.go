package apiclient

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RequestOptions struct {
	BaseURL string
	Path    map[string]string
	Query   url.Values
	Headers http.Header
	Data    []byte
}

func (o Operation) Request(ctx context.Context, opts RequestOptions) (*http.Request, error) {
	baseURL := opts.BaseURL
	if baseURL == "" {
		baseURL = o.BaseURL
	}
	base, err := url.Parse(baseURL)
	if err != nil || base.Host == "" || (base.Scheme != "http" && base.Scheme != "https") {
		return nil, fmt.Errorf("set --base-url to an absolute http(s) URL (spec server: %q)", o.BaseURL)
	}
	if base.Fragment != "" {
		return nil, fmt.Errorf("base URL must not contain a fragment")
	}
	path := o.Path
	for name, value := range opts.Path {
		marker := "{" + name + "}"
		if !strings.Contains(path, marker) {
			return nil, fmt.Errorf("unknown path parameter %q", name)
		}
		if value == "" {
			return nil, fmt.Errorf("path parameter %q must not be empty", name)
		}
		path = strings.ReplaceAll(path, marker, url.PathEscape(value))
	}
	if strings.ContainsAny(path, "{}") {
		return nil, fmt.Errorf("unresolved path %q; supply --path name=value", path)
	}
	query := base.Query()
	for name, values := range opts.Query {
		query[name] = values
	}
	headers := opts.Headers.Clone()
	if headers == nil {
		headers = make(http.Header)
	}
	for _, p := range o.Parameters {
		var present bool
		switch p.In {
		case "path":
			_, present = opts.Path[p.Name]
		case "query":
			_, present = query[p.Name]
		case "header":
			present = headers.Get(p.Name) != ""
		case "cookie":
			request := &http.Request{Header: headers}
			_, err := request.Cookie(p.Name)
			present = err == nil
		case "formData":
			return nil, fmt.Errorf("form/multipart requests are not supported")
		}
		if p.Required && !present {
			return nil, fmt.Errorf("missing required %s parameter %q", p.In, p.Name)
		}
		if p.In == "query" && present {
			values := query[p.Name]
			switch {
			case p.Style != "" && p.Style != "form":
				return nil, fmt.Errorf("query style %q is not supported for %q", p.Style, p.Name)
			case p.CollectionFormat == "multi" || (p.Style == "form" && p.Explode):
			case p.CollectionFormat == "ssv":
				query.Set(p.Name, strings.Join(values, " "))
			case p.CollectionFormat == "tsv":
				query.Set(p.Name, strings.Join(values, "\t"))
			case p.CollectionFormat == "pipes":
				query.Set(p.Name, strings.Join(values, "|"))
			default:
				if len(values) > 1 {
					query.Set(p.Name, strings.Join(values, ","))
				}
			}
		}
	}
	if o.unsupportedBody {
		return nil, fmt.Errorf("only JSON request bodies are supported")
	}
	if o.BodyRequired && opts.Data == nil {
		return nil, fmt.Errorf("request body is required; use --data @file.json or --generate")
	}
	if opts.Data != nil {
		if o.body == nil {
			return nil, fmt.Errorf("operation has no request body in the spec")
		}
		var value any
		if err := json.Unmarshal(opts.Data, &value); err != nil {
			return nil, fmt.Errorf("invalid JSON request body: %w", err)
		}
		if headers.Get("Content-Type") == "" {
			headers.Set("Content-Type", o.ContentType)
		}
	}
	if headers.Get("Accept") == "" {
		headers.Set("Accept", "application/json")
	}
	escapedPath := strings.TrimRight(base.EscapedPath(), "/") + path
	base.Path, err = url.PathUnescape(escapedPath)
	if err != nil {
		return nil, fmt.Errorf("invalid request path: %w", err)
	}
	base.RawPath = escapedPath
	base.RawQuery = query.Encode()
	request, err := http.NewRequestWithContext(ctx, o.Method, base.String(), bytes.NewReader(opts.Data))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	if host := headers.Get("Host"); host != "" {
		request.Host = host
		request.Header.Del("Host")
	}
	return request, nil
}
