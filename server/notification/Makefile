PROJECT_DIR := $(CURDIR)
PROJECT_BIN := $(PROJECT_DIR)/bin
GOLANGCI_LINT := $(PROJECT_BIN)/golangci-lint.exe
GOLANGCI_LINT_VERSION := v1.63.4

.PHONY: prepare
prepare:
	@if not exist "$(PROJECT_BIN)" mkdir "$(PROJECT_BIN)"

.PHONY: install-linter
install-linter: prepare
	powershell -Command "[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; Invoke-WebRequest -URI https://raw.githubusercontent.com/golangci/golangci-lint/master/install.ps1 -OutFile $(PROJECT_BIN)/install.ps1" && \
	powershell -Command "$(PROJECT_BIN)/install.ps1 -b $(PROJECT_BIN) $(GOLANGCI_LINT_VERSION)" && \
	powershell -Command "Remove-Item -Path $(PROJECT_BIN)/install.ps1 -Force"

.PHONY: lint
lint: install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast: install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml

.PHONY: buf-dep-update
buf-dep-update:
	cd protobuf && buf dep update

.PHONY: buf-lint
buf-lint:
	cd protobuf && buf lint

.PHONY: buf-generate
buf-generate:
	cd protobuf && buf generate