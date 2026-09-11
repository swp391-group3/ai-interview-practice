SHELL := /bin/sh
.DEFAULT_GOAL := check

GOLANGCI_VERSION := v2.13.2
GOLANGCI_COMMIT := 27774aaf853a4fd21f1dd5e69439459dc1b26e68
GOVULNCHECK_VERSION := v1.8.0
TOOLS_DIR := $(CURDIR)/.tools
GOLANGCI := $(TOOLS_DIR)/golangci-lint-$(GOLANGCI_VERSION)/golangci-lint
GOVULNCHECK := $(TOOLS_DIR)/govulncheck-$(GOVULNCHECK_VERSION)/govulncheck
GO_FILES := $(shell find . -type d \( -name .tools -o -name tmp -o -name vendor \) -prune -o -type f -name '*.go' -print)

.PHONY: fmt fmt-check vet test test-race build lint vuln check

fmt:
	gofmt -w $(GO_FILES)

fmt-check:
	@set -eu; files="$$(gofmt -l $(GO_FILES))"; if [ -n "$$files" ]; then printf '%s\n' "$$files"; exit 1; fi

vet:
	go vet ./...

test:
	go test ./...

test-race:
	go test -race ./...

build:
	go build ./...

lint: $(GOLANGCI)
	$(GOLANGCI) run ./...

vuln: $(GOVULNCHECK)
	$(GOVULNCHECK) ./...

# Fast feedback; run test-race and vuln separately (as CI does).
check: fmt-check vet test lint build

$(GOLANGCI):
	@mkdir -p "$(dir $(GOLANGCI))"
	@set -eu; installer=$$(mktemp); trap 'rm -f "$$installer"' EXIT; \
		curl --fail --silent --show-error --location \
		"https://raw.githubusercontent.com/golangci/golangci-lint/$(GOLANGCI_COMMIT)/install.sh" -o "$$installer"; \
		sh "$$installer" -b "$(dir $(GOLANGCI))" $(GOLANGCI_VERSION)

$(GOVULNCHECK):
	GOBIN="$(dir $(GOVULNCHECK))" go install golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)
