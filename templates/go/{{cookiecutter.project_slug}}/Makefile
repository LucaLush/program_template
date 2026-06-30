# Project variables
BINARY_NAME = {{ cookiecutter.project_slug }}
MODULE_NAME = {{ cookiecutter.go_module }}
BUILD_DIR = build
CONFIG_DIR = configs
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "v0.1.0")
COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME = $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# Go build flags
LDFLAGS = -ldflags "-X $(MODULE_NAME)/internal/version.Version=$(VERSION) \
                    -X $(MODULE_NAME)/internal/version.Commit=$(COMMIT) \
                    -X $(MODULE_NAME)/internal/version.BuildTime=$(BUILD_TIME) \
                    -w -s"

# Default target
.PHONY: all
all: build

# Help target
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Development targets:"
	@echo "  run           Run the application locally"
	@echo "  test          Run unit tests"
	@echo "  test-cov      Run unit tests and generate coverage report"
	@echo "  lint          Run golangci-lint"
	@echo "  fmt           Format Go code"
	@echo "  tidy          Run go mod tidy"
	@echo ""
	@echo "Build targets:"
	@echo "  build         Build for the default target ({{ cookiecutter.default_build_target }})"
	@echo "  build-amd64   Build for linux/amd64"
	@echo "  build-arm64   Build for linux/arm64"
	@echo "  build-all     Build for both amd64 and arm64"
	@echo ""
	@echo "Pack targets (outputs tar.gz to $(BUILD_DIR)/):"
	@echo "  pack          Pack the default build target"
	@echo "  pack-amd64    Pack linux/amd64"
	@echo "  pack-arm64    Pack linux/arm64"
	@echo "  pack-all      Pack both targets"
	@echo ""
	@echo "Clean targets:"
	@echo "  clean         Clean up build artifacts"

# Development targets
.PHONY: run
run:
	go run $(LDFLAGS) cmd/$(BINARY_NAME)/main.go --config $(CONFIG_DIR)/config.yaml

.PHONY: test
test:
	go test -v -race ./...

.PHONY: test-cov
test-cov:
	mkdir -p $(BUILD_DIR)
	go test -v -race -coverprofile=$(BUILD_DIR)/coverage.out ./...
	go tool cover -html=$(BUILD_DIR)/coverage.out -o $(BUILD_DIR)/coverage.html
	@echo "Coverage report generated at $(BUILD_DIR)/coverage.html"

.PHONY: lint
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint is not installed. Please install it to lint code."; \
		exit 1; \
	fi

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy

# Build targets
.PHONY: build
{% if cookiecutter.default_build_target == 'all' -%}
build: build-all
{% elif cookiecutter.default_build_target == 'amd64' -%}
build: build-amd64
{% else -%}
build: build-arm64
{% endif %}

.PHONY: build-amd64
build-amd64:
	@echo "Building $(BINARY_NAME) for linux/amd64..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/amd64/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go
	@if [ -d "$(CONFIG_DIR)" ]; then cp -r $(CONFIG_DIR)/* $(BUILD_DIR)/amd64/; fi

.PHONY: build-arm64
build-arm64:
	@echo "Building $(BINARY_NAME) for linux/arm64..."
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/arm64/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go
	@if [ -d "$(CONFIG_DIR)" ]; then cp -r $(CONFIG_DIR)/* $(BUILD_DIR)/arm64/; fi

.PHONY: build-all
build-all: build-amd64 build-arm64

# Pack targets
.PHONY: pack
{% if cookiecutter.default_build_target == 'all' -%}
pack: pack-all
{% elif cookiecutter.default_build_target == 'amd64' -%}
pack: pack-amd64
{% else -%}
pack: pack-arm64
{% endif %}

.PHONY: pack-amd64
pack-amd64: build-amd64
	@echo "Packing $(BINARY_NAME) for linux/amd64..."
	@mkdir -p $(BUILD_DIR)/tmp-amd64/$(BINARY_NAME)
	@cp $(BUILD_DIR)/amd64/$(BINARY_NAME) $(BUILD_DIR)/tmp-amd64/$(BINARY_NAME)/
	@if [ -d "$(CONFIG_DIR)" ]; then cp -r $(CONFIG_DIR)/* $(BUILD_DIR)/tmp-amd64/$(BINARY_NAME)/; fi
	@if [ -f "README.md" ]; then cp README.md $(BUILD_DIR)/tmp-amd64/$(BINARY_NAME)/; fi
	@if [ -f "$(BINARY_NAME).service" ]; then cp $(BINARY_NAME).service $(BUILD_DIR)/tmp-amd64/$(BINARY_NAME)/; fi
	@cd $(BUILD_DIR)/tmp-amd64 && tar -czf ../$(BINARY_NAME)-linux-amd64.tar.gz $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)/tmp-amd64
	@echo "Successfully created $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64.tar.gz"

.PHONY: pack-arm64
pack-arm64: build-arm64
	@echo "Packing $(BINARY_NAME) for linux/arm64..."
	@mkdir -p $(BUILD_DIR)/tmp-arm64/$(BINARY_NAME)
	@cp $(BUILD_DIR)/arm64/$(BINARY_NAME) $(BUILD_DIR)/tmp-arm64/$(BINARY_NAME)/
	@if [ -d "$(CONFIG_DIR)" ]; then cp -r $(CONFIG_DIR)/* $(BUILD_DIR)/tmp-arm64/$(BINARY_NAME)/; fi
	@if [ -f "README.md" ]; then cp README.md $(BUILD_DIR)/tmp-arm64/$(BINARY_NAME)/; fi
	@if [ -f "$(BINARY_NAME).service" ]; then cp $(BINARY_NAME).service $(BUILD_DIR)/tmp-arm64/$(BINARY_NAME)/; fi
	@cd $(BUILD_DIR)/tmp-arm64 && tar -czf ../$(BINARY_NAME)-linux-arm64.tar.gz $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)/tmp-arm64
	@echo "Successfully created $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64.tar.gz"

.PHONY: pack-all
pack-all: pack-amd64 pack-arm64

# Clean targets
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
