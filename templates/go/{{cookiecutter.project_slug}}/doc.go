// Package {{ cookiecutter.project_slug.replace('-', '_') }} provides core functionality for {{ cookiecutter.project_name }}.
//
// Usage:
//
//	import "{{ cookiecutter.go_module }}"
//
//	client := {{ cookiecutter.project_slug.replace('-', '_') }}.New(
//	    {{ cookiecutter.project_slug.replace('-', '_') }}.WithGreeting("Hello"),
//	)
//	msg := client.Greet("Gopher")
package {{ cookiecutter.project_slug.replace('-', '_') }}
