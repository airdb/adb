# adb
Airdb Development Builder

## Quick Install

```
go install -ldflags -X=github.com/airdb/adb/internal/adblib.BuildTime=$(date +%s) github.com/airdb/adb@latest
```

## Download

```
https://github.com/airdb/adb/releases/latest/download/adb
https://github.com/airdb/adb/releases/latest/download/adb-darwin
https://github.com/airdb/adb/releases/latest/download/adb-linux-amd64.zip


https://github.com/airdb/adb/releases/download/v1.0.0/adb
```

## Cowsay

`adb cowsay` supports the built-in figures and also recursively loads custom
ASCII figures from `~/.config/bar/`. Relative file paths under that directory
become figure names for `adb cowsay --list` and `adb cowsay -f <figure>`.
When no message is provided, `adb cowsay` also looks for fortunes in
`~/.config/bar/fortune.txt`, using one non-empty trimmed line per fortune
before falling back to the built-in defaults.

## Password

`adb passwd` generates a random 16-character alphanumeric password by default.
Use `adb passwd --complex` to require lower, upper, digit, and symbol
characters, and `adb passwd -n 24` to control the length.



## API CLI

Read a local Swagger 2.0 or OpenAPI 3.x YAML/JSON file without generating client code:

```bash
adb api list --spec /tmp/apis.swagger.yaml
adb api list --spec /tmp/apis.swagger.yaml --json
adb api list --spec /tmp/apis.swagger.yaml UserAuthService_Welcome

adb api example --spec /tmp/apis.swagger.yaml \
  --operation CompanyProfile_CompanyProfile > request.json
```

`list` shows operation IDs, paths, methods, and required parameters. An optional
operationId argument (or `--operation`) filters the list and displays the
description, deprecation status, and copyable example/preview/call commands;
duplicate IDs show all matching routes. `--json` outputs only metadata. A selector such as `"GET /info"` selects one route. `example`
prints only the JSON request body, using an explicit example when provided and
otherwise generating data from the schema. `--seed` defaults to `1` for
reproducibility. Schemas with no concrete properties may generate `{}`; generated
data does not guarantee valid business IDs or database state.

Preview a generated request without sending it:

```bash
adb api call --spec /tmp/apis.swagger.yaml \
  --operation CompanyProfile_CompanyProfile \
  --base-url http://localhost:8080 \
  --path name=test-company --generate --dry-run
```

Send your edited JSON file:

```bash
adb api call --spec /tmp/apis.swagger.yaml \
  --operation CompanyProfile_CompanyProfile \
  --base-url http://localhost:8080 \
  --path name=test-company --data @request.json \
  --header "Authorization: Bearer $API_TOKEN"
```

- Select by operationId, or by `--operation "POST /v1/company/{name}"` when an ID
  is missing or duplicated. Ambiguous IDs produce an error instead of picking an endpoint.
- Repeat `--path name=value`, `--query name=value`, and `--header 'Name: value'`
  as needed. Required path, query, header, cookie, and body inputs are checked.
  Supply cookies through a `Cookie` header.
- `--data` accepts inline JSON, `@file.json`, or `-` for stdin. `--generate`
  creates a body directly from the spec and cannot be combined with `--data`.
- `--base-url` replaces the complete server URL **including its base path**.
  Without it, the spec's server URL is used; relative or missing servers require
  an override. Paths always follow the spec, even if the service differs.
- Response bodies go to stdout; HTTP status goes to stderr. Non-2xx responses
  return a nonzero exit code. Redirects are not followed. `--timeout` defaults to
  `30s`. This checks HTTP success, not the response schema or business behavior.
- This version supports JSON bodies and common scalar/simple/form parameter
  encodings, not multipart uploads, deepObject queries, automatic authentication,
  Protobuf/gRPC calls, or external references. Bundle external references into
  the spec first. Requests and responses are limited to 32 MiB.


# For Template Maintaining

```bash
go get github.com/rakyll/statik
statik -include='*' -src gin-template -f
```
