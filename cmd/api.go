package cmd

import (
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/airdb/adb/internal/apiclient"
	"github.com/spf13/cobra"
)

const apiBodyLimit = 32 << 20

func newAPICommand() *cobra.Command {
	var specFile string
	api := &cobra.Command{
		Use:   "api",
		Short: "Inspect OpenAPI specs, generate JSON examples, and call HTTP APIs",
		Example: `  adb api list --spec swagger.yaml
  adb api example --spec swagger.yaml --operation createCompany
  adb api call --spec swagger.yaml --operation createCompany --base-url http://localhost:8080 --data @request.json
  adb api call --spec swagger.yaml --operation "POST /companies" --base-url http://localhost:8080 --generate --dry-run`,
	}
	api.PersistentFlags().StringVar(&specFile, "spec", "", "Local Swagger 2 / OpenAPI 3 YAML or JSON file")
	_ = api.MarkPersistentFlagRequired("spec")
	load := func() (*apiclient.Spec, error) { return apiclient.Load(specFile) }
	var listJSON bool
	var listOperation string
	list := &cobra.Command{Use: "list [operationId]", Short: "List operations and show per-operation request help",
		Long: "List operations without sending requests. Select an operationId or 'METHOD /path' to show its description, deprecation status, and example/call commands. Use --json for machine-readable metadata.",
		Example: `  adb api list --spec swagger.yaml
  adb api list --spec swagger.yaml createCompany
  adb api list --spec swagger.yaml "GET /info"
  adb api list --spec swagger.yaml --operation createCompany --json`,
		Args: cobra.MaximumNArgs(1), RunE: func(cmd *cobra.Command, args []string) error {
			selector := listOperation
			if len(args) == 1 {
				if cmd.Flags().Changed("operation") {
					return fmt.Errorf("use either an operation argument or --operation")
				}
				selector = args[0]
			}
			spec, err := load()
			if err != nil {
				return err
			}
			operations := spec.Operations
			if selector != "" {
				operations = nil
				for _, operation := range spec.Operations {
					if operation.ID == selector || operation.Method+" "+operation.Path == selector {
						operations = append(operations, operation)
					}
				}
				if len(operations) == 0 {
					return fmt.Errorf("operation %q not found; use 'adb api list'", selector)
				}
			}
			if listJSON {
				return printAPIJSON(cmd.OutOrStdout(), operations)
			}
			out := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 4, 2, ' ', 0)
			fmt.Fprintln(out, "OPERATION_ID\tMETHOD\tPATH\tREQUIRED_PARAMETERS")
			for _, o := range operations {
				var required []string
				for _, p := range o.Parameters {
					if p.Required {
						required = append(required, p.In+":"+p.Name)
					}
				}
				if o.BodyRequired {
					required = append(required, "body")
				}
				fmt.Fprintf(out, "%s\t%s\t%s\t%s\n", o.ID, o.Method, o.Path, strings.Join(required, ", "))
			}
			if err := out.Flush(); err != nil {
				return err
			}
			if selector != "" {
				for _, operation := range operations {
					if err := printAPIOperationHelp(cmd.OutOrStdout(), specFile, operation); err != nil {
						return err
					}
				}
			} else {
				_, err = fmt.Fprintln(cmd.OutOrStdout(), "\nInspect an operation: adb api list --spec <file> <operationId>\nMore help: adb api list --help")
				return err
			}
			return nil
		}}
	list.Flags().BoolVar(&listJSON, "json", false, "Print operation metadata as JSON")
	list.Flags().StringVar(&listOperation, "operation", "", "Filter by operationId or 'METHOD /path'")
	var exampleOperation string
	var seed int64
	example := &cobra.Command{Use: "example", Short: "Generate a JSON request body from examples or schema", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := load()
		if err != nil {
			return err
		}
		operation, err := spec.Find(exampleOperation)
		if err != nil {
			return err
		}
		data, err := operation.Example(seed)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout(), string(data))
		return err
	}}
	example.Flags().StringVar(&exampleOperation, "operation", "", "operationId or 'METHOD /path'")
	_ = example.MarkFlagRequired("operation")
	example.Flags().Int64Var(&seed, "seed", 1, "Random seed for reproducible generated examples")
	var operation, baseURL, data string
	var paths, queries, headers []string
	var timeout time.Duration
	var dryRun, generate bool
	var callSeed int64
	call := &cobra.Command{Use: "call", Short: "Send one HTTP request for an operation", Args: cobra.NoArgs, RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := load()
		if err != nil {
			return err
		}
		op, err := spec.Find(operation)
		if err != nil {
			return err
		}
		if timeout <= 0 {
			return fmt.Errorf("--timeout must be positive")
		}
		pathValues, err := apiPairs(paths)
		if err != nil {
			return err
		}
		queryValues, err := apiPairs(queries)
		if err != nil {
			return err
		}
		pathMap := make(map[string]string)
		for k, v := range pathValues {
			if len(v) != 1 {
				return fmt.Errorf("path parameter %q supplied more than once", k)
			}
			pathMap[k] = v[0]
		}
		headerValues := make(http.Header)
		for _, value := range headers {
			name, value, ok := strings.Cut(value, ":")
			if !ok || strings.TrimSpace(name) == "" {
				return fmt.Errorf("invalid header; use --header 'Name: value'")
			}
			headerValues.Add(strings.TrimSpace(name), strings.TrimSpace(value))
		}
		var body []byte
		if generate {
			if cmd.Flags().Changed("data") {
				return fmt.Errorf("use either --data or --generate")
			}
			body, err = op.Example(callSeed)
		} else if cmd.Flags().Changed("data") {
			body, err = readAPIData(data, cmd.InOrStdin())
		}
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()
		req, err := op.Request(ctx, apiclient.RequestOptions{BaseURL: baseURL, Path: pathMap, Query: queryValues, Headers: headerValues, Data: body})
		if err != nil {
			return err
		}
		if dryRun {
			return printAPIJSON(cmd.OutOrStdout(), struct {
				Method  string         `json:"method"`
				URL     string         `json:"url"`
				Headers http.Header    `json:"headers"`
				Body    jsontext.Value `json:"body,omitzero"`
			}{req.Method, req.URL.String(), req.Header, body})
		}
		client := &http.Client{Timeout: timeout, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("call API: %w", err)
		}
		defer resp.Body.Close()
		fmt.Fprintf(cmd.ErrOrStderr(), "HTTP %s\n", resp.Status)
		response, err := io.ReadAll(io.LimitReader(resp.Body, apiBodyLimit+1))
		if err != nil {
			return fmt.Errorf("read response: %w", err)
		}
		if len(response) > apiBodyLimit {
			return fmt.Errorf("response exceeds 32 MiB")
		}
		if _, err = cmd.OutOrStdout().Write(response); err != nil {
			return err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return fmt.Errorf("API returned HTTP %s", resp.Status)
		}
		return nil
	}}
	call.Flags().StringVar(&operation, "operation", "", "operationId or 'METHOD /path'")
	_ = call.MarkFlagRequired("operation")
	call.Flags().StringVar(&baseURL, "base-url", "", "Override the complete server URL, including any base path")
	call.Flags().StringVar(&data, "data", "", "JSON body, @file.json, or - for stdin")
	call.Flags().StringArrayVar(&paths, "path", nil, "Path parameter name=value (repeatable)")
	call.Flags().StringArrayVar(&queries, "query", nil, "Query parameter name=value (repeatable, including array items)")
	call.Flags().StringArrayVarP(&headers, "header", "H", nil, "Request header 'Name: value' (repeatable)")
	call.Flags().DurationVar(&timeout, "timeout", 30*time.Second, "Total request timeout")
	call.Flags().BoolVar(&dryRun, "dry-run", false, "Print request JSON without sending it")
	call.Flags().BoolVar(&generate, "generate", false, "Generate the JSON body from the spec")
	call.Flags().Int64Var(&callSeed, "seed", 1, "Seed for --generate")
	api.AddCommand(list, example, call)
	return api
}

func init() { rootCmd.AddCommand(newAPICommand()) }

func apiPairs(pairs []string) (url.Values, error) {
	values := make(url.Values)
	for _, pair := range pairs {
		key, value, ok := strings.Cut(pair, "=")
		if !ok || key == "" {
			return nil, fmt.Errorf("invalid parameter %q; expected name=value", pair)
		}
		values.Add(key, value)
	}
	return values, nil
}

func readAPIData(data string, stdin io.Reader) ([]byte, error) {
	if data == "-" {
		return readAPIBody(stdin)
	}
	if strings.HasPrefix(data, "@") {
		file, err := os.Open(strings.TrimPrefix(data, "@"))
		if err != nil {
			return nil, fmt.Errorf("read body file: %w", err)
		}
		defer file.Close()
		return readAPIBody(file)
	}
	if len(data) > apiBodyLimit {
		return nil, fmt.Errorf("request body exceeds 32 MiB")
	}
	return []byte(data), nil
}
func readAPIBody(reader io.Reader) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(reader, apiBodyLimit+1))
	if err != nil {
		return nil, err
	}
	if len(data) > apiBodyLimit {
		return nil, fmt.Errorf("request body exceeds 32 MiB")
	}
	return data, nil
}
func printAPIJSON(out io.Writer, value any) error {
	data, err := json.Marshal(value, json.Deterministic(true), jsontext.WithIndent("  "))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, string(data))
	return err
}
