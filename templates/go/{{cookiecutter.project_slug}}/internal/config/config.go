package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
}

type AppConfig struct {
	Name     string `yaml:"name"`
	Env      string `yaml:"env"`
	LogLevel string `yaml:"log_level"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Load loads the configuration from a YAML file
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config file: %w", err)
	}

	// Validate config fields
	if cfg.App.Name == "" {
		return nil, fmt.Errorf("app.name is required")
	}
	if cfg.Server.Port <= 0 {
		return nil, fmt.Errorf("server.port must be positive")
	}

	return &cfg, nil
}
