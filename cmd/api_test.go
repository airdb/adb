package cmd

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const apiTestSpec = "../internal/apiclient/testdata/swagger.yaml"

func executeAPITest(t *testing.T, input string, args ...string) (string, string, error) {
	t.Helper()
	command := newAPICommand()
	var stdout, stderr bytes.Buffer
	command.SetOut(&stdout)
	command.SetErr(&stderr)
	command.SetIn(strings.NewReader(input))
	command.SilenceUsage = true
	command.SilenceErrors = true
	command.SetArgs(args)
	err := command.Execute()
	return stdout.String(), stderr.String(), err
}

func TestAPIListAndExample(t *testing.T) {
	output, _, err := executeAPITest(t, "", "list", "--spec", apiTestSpec)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, "createItem") || !strings.Contains(output, "path:id") || !strings.Contains(output, "header:X-Test") {
		t.Fatal(output)
	}
	output, _, err = executeAPITest(t, "", "list", "--spec", apiTestSpec, "--json")
	if err != nil {
		t.Fatal(err)
	}
	var operations []map[string]any
	if err = json.Unmarshal([]byte(output), &operations); err != nil || len(operations) != 3 {
		t.Fatalf("list=%s error=%v", output, err)
	}
	output, _, err = executeAPITest(t, "", "example", "--spec", apiTestSpec, "--operation", "createItem")
	if err != nil {
		t.Fatal(err)
	}
	var example map[string]any
	if err = json.Unmarshal([]byte(output), &example); err != nil || example["name"] != "demo" {
		t.Fatalf("example=%s error=%v", output, err)
	}
	for _, args := range [][]string{
		{"list"},
		{"example", "--spec", apiTestSpec},
		{"example", "--spec", apiTestSpec, "--operation", "getItem"},
		{"example", "--spec", apiTestSpec, "--operation", "GET /items/{id}"},
	} {
		if _, _, err := executeAPITest(t, "", args...); err == nil {
			t.Fatalf("expected error: %v", args)
		}
	}
}

func apiCallArgs(baseURL string) []string {
	return []string{"call", "--spec", apiTestSpec, "--operation", "createItem", "--base-url", baseURL, "--path", "id=a/b", "--query", "limit=10", "--query", "tags=a", "--query", "tags=b", "-H", "X-Test: yes"}
}

func TestAPICallDataSources(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if r.Method != "POST" || r.URL.EscapedPath() != "/prefix/items/a%2Fb" {
			t.Errorf("request=%s %s", r.Method, r.URL.String())
		}
		if r.URL.Query().Get("limit") != "10" || len(r.URL.Query()["tags"]) != 2 {
			t.Error(r.URL.String())
		}
		if r.Header.Get("X-Test") != "yes" || r.Header.Get("Content-Type") != "application/json" {
			t.Error(r.Header)
		}
		body, _ := io.ReadAll(r.Body)
		var value map[string]any
		if err := json.Unmarshal(body, &value); err != nil || value["name"] == nil {
			t.Errorf("body=%s error=%v", body, err)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"id":123}`)
	}))
	defer server.Close()
	file := filepath.Join(t.TempDir(), "body.json")
	if err := os.WriteFile(file, []byte(`{"name":"file"}`), 0600); err != nil {
		t.Fatal(err)
	}
	for _, source := range []string{`{"name":"inline"}`, "@" + file, "-", "generate"} {
		args := apiCallArgs(server.URL + "/prefix")
		if source == "generate" {
			args = append(args, "--generate")
		} else {
			args = append(args, "--data", source)
		}
		output, stderr, err := executeAPITest(t, `{"name":"stdin"}`, args...)
		if err != nil {
			t.Fatal(err)
		}
		if output != `{"id":123}` || !strings.Contains(stderr, "HTTP 201") {
			t.Fatalf("out=%q err=%q", output, stderr)
		}
	}
	if calls.Load() != 4 {
		t.Fatalf("calls=%d", calls.Load())
	}
}

func TestAPIDryRunAndValidationDoNotSend(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls.Add(1) }))
	defer server.Close()
	args := append(apiCallArgs(server.URL), "--generate", "--dry-run")
	output, _, err := executeAPITest(t, "", args...)
	if err != nil {
		t.Fatal(err)
	}
	var preview struct {
		Method string         `json:"method"`
		URL    string         `json:"url"`
		Body   map[string]any `json:"body"`
	}
	if err = json.Unmarshal([]byte(output), &preview); err != nil {
		t.Fatal(err)
	}
	if preview.Method != "POST" || preview.Body["name"] != "demo" {
		t.Fatal(output)
	}
	for _, extra := range [][]string{
		{"--data", "{"},
		{"--data", "@/does-not-exist"},
		{"--generate", "--data", "{}"},
		{"--generate", "--timeout", "0s"},
		{"--generate", "--query", "invalid"},
		{"--generate", "--path", "unknown=x"},
		{"--generate", "-H", "invalid"},
	} {
		_, _, err := executeAPITest(t, "", append(apiCallArgs(server.URL), extra...)...)
		if err == nil {
			t.Fatalf("expected error: %v", extra)
		}
	}
	if calls.Load() != 0 {
		t.Fatalf("unexpected requests=%d", calls.Load())
	}
}

func TestAPICallHTTPFailureAndRedirect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Mode") == "redirect" {
			w.Header().Set("Location", "/unexpected")
			w.WriteHeader(302)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, `{"error":"invalid"}`)
	}))
	defer server.Close()
	output, _, err := executeAPITest(t, "", append(apiCallArgs(server.URL), "--generate")...)
	if err == nil || !strings.Contains(err.Error(), "422") || output != `{"error":"invalid"}` {
		t.Fatalf("out=%s err=%v", output, err)
	}
	_, _, err = executeAPITest(t, "", append(apiCallArgs(server.URL), "--generate", "-H", "X-Mode: redirect")...)
	if err == nil || !strings.Contains(err.Error(), "302") {
		t.Fatalf("redirect error=%v", err)
	}
}

func TestAPICallTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(time.Second):
		}
	}))
	defer server.Close()
	_, _, err := executeAPITest(t, "", append(apiCallArgs(server.URL), "--generate", "--timeout", "10ms")...)
	if err == nil || (!strings.Contains(err.Error(), context.DeadlineExceeded.Error()) && !strings.Contains(err.Error(), "timeout")) {
		t.Fatalf("error=%v", err)
	}
}

func TestAPIDryRunKeepsEmptyJSONBody(t *testing.T) {
	output, _, err := executeAPITest(t, "", append(apiCallArgs("http://localhost:8080"), "--data", "{}", "--dry-run")...)
	if err != nil {
		t.Fatal(err)
	}
	var preview map[string]any
	if err = json.Unmarshal([]byte(output), &preview); err != nil {
		t.Fatal(err)
	}
	if _, ok := preview["body"]; !ok {
		t.Fatalf("empty JSON body missing: %s", output)
	}
}

func TestAPIListFilter(t *testing.T) {
	for _, tc := range []struct {
		args  []string
		count int
	}{
		{[]string{"createItem"}, 1},
		{[]string{"--operation", "createItem"}, 1},
		{[]string{"getItem"}, 2},
		{[]string{"GET /items/{id}"}, 1},
	} {
		args := append([]string{"list", "--spec", apiTestSpec, "--json"}, tc.args...)
		output, _, err := executeAPITest(t, "", args...)
		if err != nil {
			t.Fatal(err)
		}
		var operations []map[string]any
		if err := json.Unmarshal([]byte(output), &operations); err != nil {
			t.Fatal(err)
		}
		if len(operations) != tc.count {
			t.Fatalf("args=%v output=%s", tc.args, output)
		}
	}
	output, _, err := executeAPITest(t, "", "list", "--spec", apiTestSpec, "createItem")
	if err != nil || !strings.Contains(output, "createItem") || strings.Contains(output, "getItem") {
		t.Fatalf("out=%s err=%v", output, err)
	}
	for _, extra := range [][]string{{"missing"}, {"createItem", "extra"}, {"createItem", "--operation", "getItem"}} {
		if _, _, err := executeAPITest(t, "", append([]string{"list", "--spec", apiTestSpec}, extra...)...); err == nil {
			t.Fatalf("expected error for %v", extra)
		}
	}
}

func TestAPIListOperationHelp(t *testing.T) {
	output, _, err := executeAPITest(t, "", "list", "--spec", apiTestSpec, "createItem")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Summary: Create a test item", "Description: Sends a test JSON document.", "Deprecated: false", " api example ", "'POST /items/{id}'", "--path id=example", "--query limit=example", "--header 'X-Test: example'", "--generate --dry-run", "--data @request.json"} {
		if !strings.Contains(output, want) {
			t.Errorf("help missing %q: %s", want, output)
		}
	}
	output, _, err = executeAPITest(t, "", "list", "--spec", apiTestSpec, "GET /items/{id}")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output, "Deprecated: true") || !strings.Contains(output, "No JSON request body") || strings.Contains(output, "--generate") || strings.Contains(output, "--data @") {
		t.Fatalf("unexpected GET help: %s", output)
	}
	output, _, err = executeAPITest(t, "", "list", "--spec", apiTestSpec, "getItem", "--json")
	if err != nil {
		t.Fatal(err)
	}
	var operations []struct {
		Deprecated bool `json:"deprecated"`
	}
	if err = json.Unmarshal([]byte(output), &operations); err != nil {
		t.Fatalf("JSON mixed with help: %s", output)
	}
	if len(operations) != 2 || !operations[0].Deprecated || !operations[1].Deprecated {
		t.Fatal(output)
	}
	output, _, err = executeAPITest(t, "", "list", "--help")
	if err != nil || !strings.Contains(output, "--operation") || !strings.Contains(output, "createCompany") {
		t.Fatalf("out=%s err=%v", output, err)
	}
}

func TestAPIHelpShellQuotes(t *testing.T) {
	for _, tc := range []struct{ value, want string }{
		{"adb", "adb"},
		{"./output/adb", "./output/adb"},
		{"/tmp/apis.swagger.yaml", "/tmp/apis.swagger.yaml"},
		{"http://localhost:8080", "http://localhost:8080"},
		{"name=example", "name=example"},
		{"", "''"},
		{"POST /v1/user/{name}", "'POST /v1/user/{name}'"},
		{"a b.yaml", "'a b.yaml'"},
		{"a'b.yaml", `'a'"'"'b.yaml'`},
		{"https://example.com/?a=1&b=2", "'https://example.com/?a=1&b=2'"},
		{"$(whoami)", "'$(whoami)'"},
		{"`whoami`", "'`whoami`'"},
	} {
		if got := apiShellQuote(tc.value); got != tc.want {
			t.Errorf("quote(%q) = %q, want %q", tc.value, got, tc.want)
		}
	}
}
