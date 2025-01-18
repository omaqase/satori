PROJECT_DIR := $(CURDIR)
PROJECT_BIN := $(PROJECT_DIR)/bin
GOLANGCI_LINT := golangci-lint.exe
GOLANGCI_LINT_VERSION := v1.63.4

.PHONY: install-linter
install-linter:
	scoop install golangci-lint

.PHONY: lint
lint: install-linter
	$(GOLANGCI_LINT) run ./...

.PHONY: lint-fast
lint-fast: install-linter
	$(GOLANGCI_LINT) run ./... --fast

.PHONY: buf-dep-update
buf-dep-update:
	cd protobuf && buf dep update