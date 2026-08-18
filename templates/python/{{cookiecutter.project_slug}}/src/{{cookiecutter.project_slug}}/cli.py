"""Command Line Interface for {{ cookiecutter.project_name }}."""

{% if cookiecutter.cli_framework == "typer" -%}
from typing import Optional
import typer
from {{ cookiecutter.project_slug }} import __version__
from {{ cookiecutter.project_slug }}.core import Greeter

app = typer.Typer(
    name="{{ cookiecutter.project_slug }}",
    help="{{ cookiecutter.project_description }}",
    no_args_is_help=True,
)


def version_callback(value: bool) -> None:
    if value:
        typer.echo(f"{{ cookiecutter.project_slug }} v{__version__}")
        raise typer.Exit()


@app.callback()
def main_callback(
    version: Optional[bool] = typer.Option(
        None,
        "--version",
        "-v",
        callback=version_callback,
        is_eager=True,
        help="Show version and exit.",
    ),
) -> None:
    """Entry point callback for global options."""
    pass


@app.command()
def greet(
    name: str = typer.Argument(..., help="Name of the person to greet."),
    greeting: str = typer.Option("Hello", "--greeting", "-g", help="Greeting prefix."),
) -> None:
    """Send a friendly greeting."""
    greeter = Greeter(greeting=greeting)
    message = greeter.format_message(name)
    typer.echo(message)


def main() -> None:
    """CLI execution entrypoint."""
    app()


{% elif cookiecutter.cli_framework == "click" -%}
import click
from {{ cookiecutter.project_slug }} import __version__
from {{ cookiecutter.project_slug }}.core import Greeter


@click.group()
@click.version_option(version=__version__, prog_name="{{ cookiecutter.project_slug }}")
def cli() -> None:
    """{{ cookiecutter.project_description }}"""
    pass


@cli.command()
@click.argument("name")
@click.option("--greeting", "-g", default="Hello", help="Greeting prefix.")
def greet(name: str, greeting: str) -> None:
    """Send a friendly greeting."""
    greeter = Greeter(greeting=greeting)
    click.echo(greeter.format_message(name))


def main() -> None:
    """CLI execution entrypoint."""
    cli()


{% elif cookiecutter.cli_framework == "argparse" -%}
import argparse
import sys
from typing import Optional, Sequence
from {{ cookiecutter.project_slug }} import __version__
from {{ cookiecutter.project_slug }}.core import Greeter


def create_parser() -> argparse.ArgumentParser:
    """Create and configure argument parser."""
    parser = argparse.ArgumentParser(
        prog="{{ cookiecutter.project_slug }}",
        description="{{ cookiecutter.project_description }}",
    )
    parser.add_argument(
        "-v", "--version", action="version", version=f"%(prog)s {__version__}"
    )
    subparsers = parser.add_subparsers(dest="command", help="Available subcommands")

    greet_parser = subparsers.add_parser("greet", help="Send a friendly greeting")
    greet_parser.add_argument("name", help="Name of the person to greet")
    greet_parser.add_argument(
        "-g", "--greeting", default="Hello", help="Greeting prefix"
    )

    return parser


def main(argv: Optional[Sequence[str]] = None) -> None:
    """CLI execution entrypoint."""
    parser = create_parser()
    args = parser.parse_args(argv)

    if args.command == "greet":
        greeter = Greeter(greeting=args.greeting)
        print(greeter.format_message(args.name))
    else:
        parser.print_help()
        sys.exit(0 if args.command is not None else 1)


{% endif -%}
if __name__ == "__main__":
    main()
