# {{ cookiecutter.project_name }}

{{ cookiecutter.project_description }}

---

## Features

- **Standard PEP 621 / PEP 517 Packaging**: Configured with `pyproject.toml` and `hatchling`.
- **`src/` Layout**: Reliable package structure avoiding import ambiguity.
- **Fast Linting & Formatting**: Powered by [Ruff](https://github.com/astral-sh/ruff).
- **Static Type Checking**: Strict type checking with [Mypy](https://github.com/python/mypy) and `py.typed` marker.
- **Testing & Coverage**: Unit testing configured via [pytest](https://docs.pytest.org/) and `pytest-cov`.
{% if cookiecutter.project_type == "CLI Application" -%}
- **CLI Interface**: Integrated command-line entrypoint using `{{ cookiecutter.cli_framework }}`.
{% endif -%}
{% if cookiecutter.add_dockerfile == "yes" -%}
- **Container Support**: Multi-stage minimal `Dockerfile`.
{% endif -%}

---

## Quick Start

### 1. Requirements

- Python >={{ cookiecutter.python_version }}
- Recommended tool: [`uv`](https://github.com/astral-sh/uv) or `pip` / `virtualenv`

### 2. Setup Virtual Environment & Install

```bash
# Using standard venv + pip
python -m venv .venv
source .venv/bin/activate  # On Windows: .venv\Scripts\activate
make install-dev
```

### 3. Development Workflow

| Task | Command |
|---|---|
| Run tests | `make test` |
| Run tests with coverage | `make test-cov` |
| Code Linting | `make lint` |
| Code Formatting | `make format` |
| Type Checking | `make type-check` |
| Build Distribution Packages | `make build` |

---

{% if cookiecutter.project_type == "CLI Application" -%}
## Usage

You can invoke the CLI directly after installation:

```bash
{{ cookiecutter.project_slug }} --help
# Or run as a module:
python -m {{ cookiecutter.project_slug }} --help
```
{% else -%}
## Usage

Import the package in Python:

```python
import {{ cookiecutter.project_slug }}

print({{ cookiecutter.project_slug }}.__version__)
```
{% endif -%}

---

## License

MIT
