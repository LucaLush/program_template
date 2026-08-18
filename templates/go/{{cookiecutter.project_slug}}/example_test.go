package {{ cookiecutter.project_slug.replace('-', '_') }}_test

import (
	"fmt"
	"log"

	"{{ cookiecutter.go_module }}"
)

func Example() {
	msg, err := {{ cookiecutter.project_slug.replace('-', '_') }}.Greet("Gopher")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	// Output: Hello, Gopher!
}

func ExampleNew() {
	client := {{ cookiecutter.project_slug.replace('-', '_') }}.New(
		{{ cookiecutter.project_slug.replace('-', '_') }}.WithGreeting("Welcome"),
	)
	msg, err := client.Greet("Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	// Output: Welcome, Alice!
}
