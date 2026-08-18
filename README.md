# Project Templates Monorepo

This repository is a monorepo containing various project templates designed to be instantiated using [Cookiecutter](https://github.com/cookiecutter/cookiecutter).

By storing different types of project templates in a single repository, we keep templates unified and easy to maintain.

---

## How to Use

Cookiecutter (since version 1.7.0) supports generating projects from templates stored in subdirectories of a repository using the `--directory` option.

### Prerequisites

First, ensure you have Cookiecutter installed:
```bash
pip install cookiecutter
# Or via brew/apt depending on your OS
```

### 1. CMake C++ Project Template

A modern CMake-based C++ template featuring target-based configuration, vcpkg dependency management, Doxygen integration, and unit testing presets.

#### Instantiation

To generate a new C++ project using this template, run:

**Using Git Repository:**
```bash
cookiecutter https://github.com/LucaLush/program_template.git --directory="templates/cmake"
```

**Using Local Directory:**
```bash
cookiecutter /path/to/program_template --directory="templates/cmake"
```

#### Template Configuration Options

During instantiation, Cookiecutter will prompt you for the following configurations:

| Prompt | Description | Default | Choices / Format |
|---|---|---|---|
| `project_name` | Human-readable name of your project | `My Awesome Project` | Any text |
| `project_slug` | Code-safe project identifier (automatically generated) | `my_awesome_project` | lowercase_with_underscores |
| `cmake_minimum_version` | Minimum required CMake version | `3.20` | e.g. `3.15`, `3.22` |
| `cpp_standard` | C++ standard version to target | `17` | `11`, `14`, `17`, `20`, `23` |
| `project_type` | Build output target type | `Executable` | `Executable`, `Library` |
| `use_vcpkg` | Enable vcpkg package manager in manifest mode | `yes` | `yes`, `no` |
| `use_doxygen` | Configure Doxygen documentation generation target | `yes` | `yes`, `no` |
| `add_testing` | Add testing support (using CTest & GoogleTest) | `yes` | `yes`, `no` |

### 2. Go Project Template

An industrial-grade Go service template featuring structured config, structured logging (`slog`), multi-architecture builds, robust testing targets, and an automated package pipeline.

#### Instantiation

To generate a new Go project using this template, run:

**Using Git Repository:**
```bash
cookiecutter https://github.com/LucaLush/program_template.git --directory="templates/go"
```

**Using Local Directory:**
```bash
cookiecutter /path/to/program_template --directory="templates/go"
```

#### Template Configuration Options

During instantiation, Cookiecutter will prompt you for the following configurations:

| Prompt | Description | Default | Choices / Format |
|---|---|---|---|
| `project_name` | Human-readable name of your project | `My Go Project` | Any text |
| `project_slug` | Code-safe project identifier (automatically generated) | `my_go_project` | lowercase_with_underscores |
| `go_module` | Go module import path | `my_go_project` | e.g. `github.com/user/repo` |
| `go_version` | Target Go version | `1.21` | e.g. `1.20`, `1.21`, `1.22` |
| `project_type` | Type of Go application template | `CLI/Daemon` | `CLI/Daemon`, `Web Service`, `Library/Package` |
| `default_build_target` | Default build platform architecture | `all` | `all`, `amd64`, `arm64` |
| `add_linter` | Include `golangci-lint` configuration (`.golangci.yml`) | `yes` | `yes`, `no` |
| `add_dockerfile` | Include multi-stage container build `Dockerfile` | `yes` | `yes`, `no` |
| `add_systemd` | Include systemd service configuration file (`<slug>.service`) | `yes` | `yes`, `no` |

### 3. Python Project Template

An enterprise-ready, PEP 621-compliant modern Python template featuring `src/` layout, `hatchling` build backend, `ruff` linting and formatting, strict `mypy` type checking, `pytest` coverage testing, multi-stage `Dockerfile`, and flexible support for both CLI applications (Typer/Click/Argparse) and pure libraries/packages.

#### Instantiation

To generate a new Python project using this template, run:

**Using Git Repository:**
```bash
cookiecutter https://github.com/LucaLush/program_template.git --directory="templates/python"
```

**Using Local Directory:**
```bash
cookiecutter /path/to/program_template --directory="templates/python"
```

#### Template Configuration Options

During instantiation, Cookiecutter will prompt you for the following configurations:

| Prompt | Description | Default | Choices / Format |
|---|---|---|---|
| `project_name` | Human-readable name of your project | `My Python Project` | Any text |
| `project_slug` | Code-safe project identifier (automatically generated) | `my_python_project` | lowercase_with_underscores |
| `project_description` | Short description of the project | `A modern Python project.` | Any text |
| `author_name` | Author or maintainer name | `Developer` | Any text |
| `author_email` | Author or maintainer email address | `dev@example.com` | Email format |
| `python_version` | Minimum required Python version | `3.10` | `3.10`, `3.11`, `3.12`, `3.13` |
| `project_type` | Type of Python project | `CLI Application` | `CLI Application`, `Library/Package` |
| `cli_framework` | CLI framework (when `project_type` is `CLI Application`) | `typer` | `typer`, `argparse`, `click` |
| `add_dockerfile` | Include multi-stage container build `Dockerfile` | `yes` | `yes`, `no` |
| `add_docs` | Include MkDocs documentation template | `no` | `yes`, `no` |
| `add_github_actions` | Include GitHub Actions CI workflow (`.github/workflows/ci.yml`) | `yes` | `yes`, `no` |

---

## Contributing

To add a new template, create a folder under `templates/` with your template name (e.g. `templates/python`), containing:
1. `cookiecutter.json`
2. The template project folder (e.g. `{{cookiecutter.project_slug}}/`)
3. Optional `hooks/` folder for pre/post-generation scripts.

