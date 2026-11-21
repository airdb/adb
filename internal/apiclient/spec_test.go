package apiclient

import (
	"context"
	"encoding/json/v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSwaggerLoadAndExample(t *testing.T) {
	s, err := Load("testdata/swagger.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Operations) != 3 {
		t.Fatalf("operations=%d", len(s.Operations))
	}
	o, err := s.Find("createItem")
	if err != nil {
		t.Fatal(err)
	}
	if o.BaseURL != "https://example.invalid/api" {
		t.Fatal(o.BaseURL)
	}
	data, err := o.Example(1)
	if err != nil {
		t.Fatal(err)
	}
	var value map[string]any
	if err := json.Unmarshal(data, &value); err != nil {
		t.Fatal(err)
	}
	if value["name"] != "demo" || value["state"] != "pending" {
		t.Fatalf("example=%s", data)
	}
	if _, ok := value["details"].(map[string]any); !ok {
		t.Fatalf("nested ref missing: %s", data)
	}
	again, err := o.Example(1)
	if err != nil {
		t.Fatal(err)
	}
	if string(again) != string(data) {
		t.Fatal("seed is not reproducible")
	}
	if _, err = s.Find("getItem"); err == nil || !strings.Contains(err.Error(), "ambiguous") {
		t.Fatalf("expected ambiguity: %v", err)
	}
	if _, err = s.Find("GET /items/{id}"); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Find("missing"); err == nil {
		t.Fatal("expected missing operation")
	}
}

func TestOpenAPIServersAndExamples(t *testing.T) {
	s, err := Load("testdata/openapi.yaml")
	if err != nil {
		t.Fatal(err)
	}
	o, err := s.Find("createItem")
	if err != nil {
		t.Fatal(err)
	}
	if o.BaseURL != "https://test.example.invalid/root" {
		t.Fatal(o.BaseURL)
	}
	for _, p := range o.Parameters {
		if p.Name == "q" && p.Required {
			t.Fatal("operation did not override path-level parameter")
		}
	}
	data, err := o.Example(42)
	if err != nil {
		t.Fatal(err)
	}
	var value map[string]any
	if err = json.Unmarshal(data, &value); err != nil {
		t.Fatal(err)
	}
	if value["name"] != "explicit" {
		t.Fatal(string(data))
	}
	alt, _ := s.Find("alternate")
	if alt.BaseURL != "https://operation.example.invalid/prefix" {
		t.Fatal(alt.BaseURL)
	}
	form, _ := s.Find("uploadForm")
	if _, err = form.Example(1); err == nil {
		t.Fatal("expected unsupported body")
	}
	if _, err = form.Request(context.Background(), RequestOptions{}); err == nil {
		t.Fatal("expected unsupported body")
	}
}

func TestRequestParameters(t *testing.T) {
	s, err := Load("testdata/swagger.yaml")
	if err != nil {
		t.Fatal(err)
	}
	o, _ := s.Find("createItem")
	opts := RequestOptions{
		Path:    map[string]string{"id": "a/b ?#%"},
		Query:   url.Values{"limit": {"10"}, "tags": {"a", "b"}, "search": {"a b&c"}},
		Headers: http.Header{"X-Test": {"yes"}}, Data: []byte(`{"name":"test"}`),
	}
	req, err := o.Request(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if req.URL.EscapedPath() != "/api/items/a%2Fb%20%3F%23%25" {
		t.Fatal(req.URL.String())
	}
	if req.URL.Query().Get("search") != "a b&c" || len(req.URL.Query()["tags"]) != 2 {
		t.Fatal(req.URL.String())
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Fatal(req.Header)
	}
	body, _ := io.ReadAll(req.Body)
	if string(body) != string(opts.Data) {
		t.Fatal(string(body))
	}
	opts.BaseURL = "http://localhost:1234/custom"
	req, err = o.Request(context.Background(), opts)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(req.URL.String(), "http://localhost:1234/custom/items/") {
		t.Fatal(req.URL.String())
	}
	for _, tc := range []struct {
		name   string
		change func(*RequestOptions)
		want   string
	}{
		{"path", func(o *RequestOptions) { o.Path = nil }, "unresolved path"},
		{"query", func(o *RequestOptions) { o.Query = nil }, "missing required query"},
		{"header", func(o *RequestOptions) { o.Headers = nil }, "missing required header"},
		{"body", func(o *RequestOptions) { o.Data = nil }, "request body is required"},
		{"invalid JSON", func(o *RequestOptions) { o.Data = []byte(`{`) }, "invalid JSON"},
		{"base", func(o *RequestOptions) { o.BaseURL = "file:///tmp/test" }, "absolute http(s)"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			copy := opts
			tc.change(&copy)
			_, err := o.Request(context.Background(), copy)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error=%v want=%s", err, tc.want)
			}
		})
	}
}

func TestLoadJSONAndInvalidSpecs(t *testing.T) {
	file := filepath.Join(t.TempDir(), "openapi.json")
	if err := os.WriteFile(file, []byte(`{"openapi":"3.0.3","info":{"title":"test","version":"1"},"paths":{"/health":{"get":{"responses":{"200":{"description":"OK"}}}}}}`), 0600); err != nil {
		t.Fatal(err)
	}
	s, err := Load(file)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = s.Find("GET /health"); err != nil {
		t.Fatal(err)
	}
	for _, data := range []string{"not a spec", "openapi: 3.0.3\ninfo: {title: test, version: '1'}\n", "swagger: '2.0'\npaths:\n  /x:\n    post:\n      parameters:\n        - in: body\n          name: body\n          schema: {$ref: '#/definitions/Missing'}\n"} {
		if err := os.WriteFile(file, []byte(data), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := Load(file); err == nil {
			t.Errorf("expected invalid spec error: %s", data)
		}
	}
}
