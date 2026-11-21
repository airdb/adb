package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/airdb/adb/internal/apiclient"
)

func printAPIOperationHelp(out io.Writer, spec string, operation apiclient.Operation) error {
	var text strings.Builder
	fmt.Fprintf(&text, "\n%s %s\n", operation.Method, operation.Path)
	if operation.Summary != "" {
		fmt.Fprintf(&text, "Summary: %s\n", operation.Summary)
	}
	fmt.Fprintf(&text, "Deprecated: %t\n", operation.Deprecated)
	if operation.Description != "" {
		fmt.Fprintf(&text, "Description: %s\n", operation.Description)
	}
	baseURL := operation.BaseURL
	if server, err := url.Parse(baseURL); err != nil || server.Host == "" || (server.Scheme != "http" && server.Scheme != "https") {
		baseURL = "http://localhost:8080"
	}
	// Method + path is unambiguous even when the spec has duplicate operation IDs.
	shared := " --spec " + apiShellQuote(spec) + " --operation " + apiShellQuote(operation.Method+" "+operation.Path)
	call := apiShellQuote(os.Args[0]) + " api call" + shared + " --base-url " + apiShellQuote(baseURL)
	for _, parameter := range operation.Parameters {
		if !parameter.Required && parameter.In != "path" {
			continue
		}
		switch parameter.In {
		case "path":
			call += " --path " + apiShellQuote(parameter.Name+"=example")
		case "query":
			call += " --query " + apiShellQuote(parameter.Name+"=example")
		case "header":
			call += " --header " + apiShellQuote(parameter.Name+": example")
		case "cookie":
			call += " --header " + apiShellQuote("Cookie: "+parameter.Name+"=example")
		}
	}
	fmt.Fprintln(&text, "\nReplace the example parameter values and base URL for your service.")
	if operation.HasJSONBody() {
		fmt.Fprintln(&text, "\nGenerate a JSON request body (does not send a request):")
		fmt.Fprintf(&text, "  %s api example%s > request.json\n", apiShellQuote(os.Args[0]), shared)
		fmt.Fprintln(&text, "\nPreview a generated request (does not send):")
		fmt.Fprintln(&text, "  "+call+" --generate --dry-run")
		fmt.Fprintln(&text, "\nSend your edited request.json:")
		fmt.Fprintln(&text, "  "+call+" --data @request.json")
	} else if !operation.HasRequestBody() {
		fmt.Fprintln(&text, "\nNo JSON request body is defined; api example generates request bodies only.")
		fmt.Fprintln(&text, "\nPreview the request (does not send):")
		fmt.Fprintln(&text, "  "+call+" --dry-run")
		fmt.Fprintln(&text, "\nSend the request:")
		fmt.Fprintln(&text, "  "+call)
	} else {
		fmt.Fprintln(&text, "\nJSON example/call commands are unavailable for this request body format.")
	}
	_, err := io.WriteString(out, text.String())
	return err
}

func apiShellQuote(value string) string {
	if value != "" && strings.IndexFunc(value, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' || strings.ContainsRune("_@%+=:,./-", r))
	}) == -1 {
		return value
	}
	return "'" + strings.ReplaceAll(value, "'", `'"'"'`) + "'"
}
