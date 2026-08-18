# Build stage
FROM python:{{ cookiecutter.python_version }}-slim AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

COPY pyproject.toml README.md ./
COPY src/ ./src/

RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir build && \
    python -m build --wheel && \
    pip install --no-cache-dir --prefix=/install dist/*.whl

# Runtime stage
FROM python:{{ cookiecutter.python_version }}-slim AS runner

WORKDIR /app

RUN useradd -m -u 1000 appuser && \
    chown -R appuser:appuser /app

COPY --from=builder /install /usr/local

USER appuser

{% if cookiecutter.project_type == "CLI Application" -%}
ENTRYPOINT ["{{ cookiecutter.project_slug }}"]
CMD ["--help"]
{% else -%}
CMD ["python", "-c", "import {{ cookiecutter.project_slug }}; print('{{ cookiecutter.project_slug }} ready')"]
{% endif -%}
