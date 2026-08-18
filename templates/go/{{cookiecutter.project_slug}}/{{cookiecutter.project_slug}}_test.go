package {{ cookiecutter.project_slug.replace('-', '_') }}_test

import (
	"errors"
	"testing"

	"{{ cookiecutter.go_module }}"
)

func TestGreeter_Default(t *testing.T) {
	g := {{ cookiecutter.project_slug.replace('-', '_') }}.New()
	got, err := g.Greet("Gopher")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Hello, Gopher!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreeter_WithOptions(t *testing.T) {
	g := {{ cookiecutter.project_slug.replace('-', '_') }}.New(
		{{ cookiecutter.project_slug.replace('-', '_') }}.WithGreeting("Hi"),
	)
	got, err := g.Greet("Developer")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Hi, Developer!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreeter_EmptyName(t *testing.T) {
	g := {{ cookiecutter.project_slug.replace('-', '_') }}.New()
	_, err := g.Greet("   ")
	if !errors.Is(err, {{ cookiecutter.project_slug.replace('-', '_') }}.ErrEmptyName) {
		t.Errorf("expected ErrEmptyName, got %v", err)
	}
}

func TestPackageGreet(t *testing.T) {
	got, err := {{ cookiecutter.project_slug.replace('-', '_') }}.Greet("World")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkGreeter_Greet(b *testing.B) {
	g := {{ cookiecutter.project_slug.replace('-', '_') }}.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = g.Greet("Gopher")
	}
}
