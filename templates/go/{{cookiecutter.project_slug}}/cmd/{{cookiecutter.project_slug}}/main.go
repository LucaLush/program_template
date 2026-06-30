package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"{{ cookiecutter.go_module }}/internal/app"
	"{{ cookiecutter.go_module }}/internal/config"
	"{{ cookiecutter.go_module }}/internal/version"
)

func main() {
	var (
		configFile  string
		showVersion bool
	)

	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
	flag.BoolVar(&showVersion, "version", false, "Show application version")
	flag.Parse()

	if showVersion {
		fmt.Printf("Version:    %s\n", version.Version)
		fmt.Printf("Commit:     %s\n", version.Commit)
		fmt.Printf("Build Time: %s\n", version.BuildTime)
		os.Exit(0)
	}

	// Initialize configuration
	cfg, err := config.Load(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Set up structured logging
	var level slog.Level
	switch cfg.App.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var logger *slog.Logger
	if cfg.App.Env == "production" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	}
	slog.SetDefault(logger)

	slog.Info("Starting application",
		slog.String("name", cfg.App.Name),
		slog.String("env", cfg.App.Env),
		slog.String("version", version.Version),
	)

	// Run application
	if err := app.Run(cfg); err != nil {
		slog.Error("Application terminated with error", slog.Any("error", err))
		os.Exit(1)
	}
}
