# Backend development

Run commands from `api/`. Use Go 1.27.1 (pinned in CI, compatible with the minimum
in `go.mod`), GNU Make, curl, and tar. Race tests also require a C compiler and CGO
(available on the CI Linux runner).

| Command | Verification |
| --- | --- |
| `make fmt` | Apply gofmt to backend Go files |
| `make fmt-check` | Fail on unformatted Go files |
| `make vet` | Run go vet |
| `make test` | Run unit tests |
| `make test-race` | Run tests with the race detector |
| `make build` | Compile all backend packages |
| `make lint` | Run pinned golangci-lint |
| `make vuln` | Run pinned govulncheck against the current Go vulnerability database |
| `make check` | Formatting check, vet, tests, lint, and build |

Before committing, run `make check test-race vuln`. CI runs the same checks for
backend pull requests and pushes to main, plus local Compose syntax validation.
The workflow can also be started manually. Go module/build caches are managed by
setup-go using `api/go.sum`.

Tools install automatically into ignored `.tools/` versioned directories. The
Makefile pins golangci-lint's release and official installer commit; its installer
checks release checksums. govulncheck is installed at an exact module version,
without adding application dependencies. Updating tools requires changing these
pins and verifying them against the module's Go version. Network access is needed
for first installation and for the vulnerability database; a scan failure fails CI.

## Tests and error handling

Use standard `testing`, table-driven cases, and small handwritten fakes for unit
tests. HTTP tests use `net/http/httptest`. Database integration tests may adopt
testcontainers-go with PostgreSQL in KAN-19 or later. No mocking framework is
needed for the upcoming extraction interface; testify/require can be considered
when actual tests warrant it.

Application errors use `*apperror.AppError`. Deep layers return typed/wrapped
errors and preserve causes; API responses expose only the public message. The
application/transport boundary should log an error once. Current startup/shutdown
logs use the standard log package and Gin supplies request logging/recovery.
Structured slog initialization and consistent boundary error logging are deferred
to a focused follow-up so this baseline does not introduce partial logging policy.

## Local PostgreSQL

The root `.env.example` provides Compose variables, not the API's full runtime
configuration. Set `DATABASE_URL` in the API process environment separately.
Validate the merged Compose file without starting containers:

```sh
docker compose --env-file ../.env.example -f ../.docker/compose.local.yaml config --quiet
```

This quality baseline does not require a running database, API container, migration
runner, Swagger, or external AI credentials.
