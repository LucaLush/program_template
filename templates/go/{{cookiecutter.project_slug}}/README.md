# {{ cookiecutter.project_name }}

{% if cookiecutter.project_type == "Library/Package" -%}
A Go library module designed for modularity, clean APIs, and easy integration into other Go applications.

## Prerequisites

- [Go](https://go.dev/doc/install) (version `{{ cookiecutter.go_version }}` or higher)
- [make](https://www.gnu.org/software/make/)

## Project Structure

```text
├── doc.go                  # Package-level documentation comments
├── {{ cookiecutter.project_slug }}.go  # Core public library logic and APIs
├── {{ cookiecutter.project_slug }}_test.go # Unit tests and benchmarks
├── example_test.go         # Executable godoc example tests
├── Makefile                # Task runner for tests, benchmarks, linting
└── go.mod                  # Go module definition
```

## Installation

```bash
go get {{ cookiecutter.go_module }}
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"{{ cookiecutter.go_module }}"
)

func main() {
	client := {{ cookiecutter.project_slug.replace('-', '_') }}.New(
		{{ cookiecutter.project_slug.replace('-', '_') }}.WithGreeting("Hello"),
	)
	msg, err := client.Greet("Gopher")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
```

## Development & Testing

### Run Tests & Benchmarks

- **Run all unit tests:**
  ```bash
  make test
  ```
- **Run tests with visual HTML coverage report:**
  ```bash
  make test-cov
  ```
  The report will be located at `build/coverage.html`.
- **Run performance benchmarks:**
  ```bash
  make bench
  ```

### Linting & Formatting

- **Format source code:**
  ```bash
  make fmt
  ```
- **Run code linter (`golangci-lint`):**
  ```bash
  make lint
  ```

{% else -%}
An industrial-grade Go project template built using Clean Architecture principles.

## Prerequisites

- [Go](https://go.dev/doc/install) (version `{{ cookiecutter.go_version }}` or higher)
- [make](https://www.gnu.org/software/make/)

## Project Structure

```text
├── cmd/
│   └── {{ cookiecutter.project_slug }}/
│       └── main.go         # Application entry point
├── configs/
│   └── config.yaml         # Application configuration options
├── internal/
│   ├── app/
│   │   └── app.go          # Core server setup & graceful shutdown logic
│   ├── config/
│   │   ├── config.go       # Structs and loaders for configuration
│   │   └── config_test.go  # Unit testing for configuration loader
│   └── version/
│       └── version.go      # Embedded version variables injected at compile-time
├── tests/
│   └── integration_test.go # Integration test boilerplate
├── Dockerfile              # Containerized multi-stage build setup
├── Makefile                # Task runner for builds, tests, linting, packaging
└── go.mod                  # Go module definition
```

## Getting Started

### Running Locally

To run the application locally:

```bash
make run
```

This starts the HTTP server configured in `configs/config.yaml`. By default, it runs on port `8080` with endpoints:
- `http://localhost:8080/healthz`
- `http://localhost:8080/readyz`

### Run Tests & Coverage

To run Go unit tests:

```bash
make test
```

To run tests and view the visual HTML coverage report:

```bash
make test-cov
```

The report will be located at `build/coverage.html`.

### Linting & Formatting

Format the source files:

```bash
make fmt
```

Run the `golangci-lint` code analyzers (requires `golangci-lint` to be installed):

```bash
make lint
```

## Compilation & Packaging

### Multi-Architecture Builds

This Makefile supports cross-compiling for `linux/amd64` and `linux/arm64`.

- **Build default platform:**
  ```bash
  make build
  ```
- **Build amd64 only:**
  ```bash
  make build-amd64
  ```
  Generates binary at `build/amd64/{{ cookiecutter.project_slug }}`.
- **Build arm64 only:**
  ```bash
  make build-arm64
  ```
  Generates binary at `build/arm64/{{ cookiecutter.project_slug }}`.
- **Build all platforms:**
  ```bash
  make build-all
  ```

### Packaging for Release

The pack targets bundle the compiled binary, the `configs/` folder, and the `README.md` file into a compressed `tar.gz` archive located under the `build/` directory.

- **Pack default target:**
  ```bash
  make pack
  ```
- **Pack amd64 release:**
  ```bash
  make pack-amd64
  ```
  Produces `build/{{ cookiecutter.project_slug }}-linux-amd64.tar.gz`.
- **Pack arm64 release:**
  ```bash
  make pack-arm64
  ```
  Produces `build/{{ cookiecutter.project_slug }}-linux-arm64.tar.gz`.
- **Pack all releases:**
  ```bash
  make pack-all
  ```

### Cleanup

To clean up Go build binaries and packed archives:

```bash
make clean
```
{% endif -%}
