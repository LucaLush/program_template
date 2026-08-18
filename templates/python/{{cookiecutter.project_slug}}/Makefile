.PHONY: help install install-dev lint format type-check test test-cov build clean run

help: ## Show help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-16s\033[0m %s\n", $$1, $$2}'

install: ## Install package in production mode
	pip install .

install-dev: ## Install package in editable mode with development dependencies
	pip install -e ".[dev]"

lint: ## Run code linter (Ruff)
	ruff check src tests

format: ## Run code formatter (Ruff)
	ruff format src tests
	ruff check --fix src tests

type-check: ## Run static type checking (mypy)
	mypy src tests

test: ## Run unit tests with pytest
	pytest

test-cov: ## Run unit tests with coverage report
	pytest --cov=src --cov-report=html --cov-report=term

build: clean ## Build source and wheel distribution packages
	python -m build

clean: ## Clean build artifacts, caches and temporary files
	rm -rf build/ dist/ *.egg-info .eggs/
	rm -rf .pytest_cache .coverage htmlcov .mypy_cache .ruff_cache
	find . -type d -name "__pycache__" -exec rm -rf {} +
	find . -type f -name "*.pyc" -delete

{% if cookiecutter.project_type == "CLI Application" -%}
run: ## Run CLI application
	python -m {{ cookiecutter.project_slug }} --help
{% else -%}
run: ## Run demo test
	python -c "import {{ cookiecutter.project_slug }}; print('{{ cookiecutter.project_slug }} version:', {{ cookiecutter.project_slug }}.__version__)"
{% endif -%}
