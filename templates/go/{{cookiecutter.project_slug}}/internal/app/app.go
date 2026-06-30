package app

import (
{% if cookiecutter.project_type == 'Web Service' -%}
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{ cookiecutter.go_module }}/internal/config"
{%- else -%}
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{ cookiecutter.go_module }}/internal/config"
{%- endif %}
)

// Run runs the main application loop
func Run(cfg *config.Config) error {
{% if cookiecutter.project_type == 'Web Service' -%}
	// A simple industrial server with graceful shutdown
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ready":true}`))
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Graceful shutdown channel
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	serverErrors := make(chan error, 1)

	// Start server in background
	go func() {
		slog.Info("HTTP server listening", slog.String("addr", srv.Addr))
		serverErrors <- srv.ListenAndServe()
	}()

	// Block until signal or error
	select {
	case err := <-serverErrors:
		if err != http.ErrServerClosed {
			return fmt.Errorf("http server error: %w", err)
		}
	case sig := <-shutdown:
		slog.Info("Shutdown signal received, stopping server...", slog.String("signal", sig.String()))

		// Attempt graceful shutdown with 15s timeout
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Graceful shutdown failed, forcing close", slog.Any("error", err))
			if err := srv.Close(); err != nil {
				return fmt.Errorf("failed to close server: %w", err)
			}
		}
		slog.Info("Server stopped cleanly")
	}
{%- else -%}
	// Graceful shutdown context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	slog.Info("Application started. Running background tasks...")

	// Example background loop representing general work
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			slog.Info("Performing periodic task...")
		case <-ctx.Done():
			slog.Info("Shutting down gracefully...")
			// Perform cleanup here
			return nil
		}
	}
{%- endif %}
	return nil
}
