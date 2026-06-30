package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary YAML config file
	data := []byte(`
app:
  name: "test-app"
  env: "test"
  log_level: "debug"
server:
  host: "localhost"
  port: 8081
`)
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(data); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Load configuration
	cfg, err := Load(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Assertions
	if cfg.App.Name != "test-app" {
		t.Errorf("expected app.name to be 'test-app', got '%s'", cfg.App.Name)
	}
	if cfg.Server.Port != 8081 {
		t.Errorf("expected server.port to be 8081, got %d", cfg.Server.Port)
	}
}

func TestLoadInvalid(t *testing.T) {
	_, err := Load("non-existent-file.yaml")
	if err == nil {
		t.Error("expected error when loading non-existent file, got nil")
	}
}
