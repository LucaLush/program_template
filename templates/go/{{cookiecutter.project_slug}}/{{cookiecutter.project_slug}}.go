package {{ cookiecutter.project_slug.replace('-', '_') }}

import (
	"errors"
	"fmt"
	"strings"
)

// ErrEmptyName is returned when a provided name is empty or only whitespace.
var ErrEmptyName = errors.New("name cannot be empty")

// Option defines a functional option for configuring a Greeter instance.
type Option func(*Greeter)

// WithGreeting sets a custom greeting prefix.
func WithGreeting(greeting string) Option {
	return func(g *Greeter) {
		if greeting != "" {
			g.greeting = greeting
		}
	}
}

// Greeter handles formatting greeting messages.
type Greeter struct {
	greeting string
}

// New creates a new Greeter configured with the provided options.
func New(opts ...Option) *Greeter {
	g := &Greeter{
		greeting: "Hello",
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Greet formats a friendly greeting message. Returns ErrEmptyName if name is blank.
func (g *Greeter) Greet(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", ErrEmptyName
	}
	return fmt.Sprintf("%s, %s!", g.greeting, trimmed), nil
}

// Greet is a package-level helper that generates a default greeting.
func Greet(name string) (string, error) {
	return New().Greet(name)
}
