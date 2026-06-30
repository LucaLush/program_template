# Build stage
FROM golang:{{ cookiecutter.go_version }}-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache make git

# Copy dependency files and download
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code and build
COPY . .
RUN make build-all

# Final stage (runs the default amd64 binary)
FROM alpine:3.19

WORKDIR /app

# Copy binary and configuration
COPY --from=builder /app/build/amd64/{{ cookiecutter.project_slug }} ./server
COPY --from=builder /app/configs/config.yaml ./configs/config.yaml

EXPOSE 8080

ENTRYPOINT ["./server"]
CMD ["--config", "configs/config.yaml"]
