"""Unit tests for the CLI interface."""

{% if cookiecutter.cli_framework == "typer" -%}
from typer.testing import CliRunner
from {{ cookiecutter.project_slug }}.cli import app

runner = CliRunner()


def test_cli_version() -> None:
    """Test CLI version option."""
    result = runner.invoke(app, ["--version"])
    assert result.exit_code == 0
    assert "v0.1.0" in result.stdout


def test_cli_greet_command() -> None:
    """Test CLI greet command."""
    result = runner.invoke(app, ["greet", "Alice"])
    assert result.exit_code == 0
    assert "Hello, Alice!" in result.stdout


def test_cli_greet_with_option() -> None:
    """Test CLI greet command with custom greeting."""
    result = runner.invoke(app, ["greet", "Bob", "--greeting", "Welcome"])
    assert result.exit_code == 0
    assert "Welcome, Bob!" in result.stdout


{% elif cookiecutter.cli_framework == "click" -%}
from click.testing import CliRunner
from {{ cookiecutter.project_slug }}.cli import cli

runner = CliRunner()


def test_cli_version() -> None:
    """Test CLI version option."""
    result = runner.invoke(cli, ["--version"])
    assert result.exit_code == 0
    assert "0.1.0" in result.stdout


def test_cli_greet_command() -> None:
    """Test CLI greet command."""
    result = runner.invoke(cli, ["greet", "Alice"])
    assert result.exit_code == 0
    assert "Hello, Alice!" in result.stdout


def test_cli_greet_with_option() -> None:
    """Test CLI greet command with custom greeting."""
    result = runner.invoke(cli, ["greet", "Bob", "--greeting", "Welcome"])
    assert result.exit_code == 0
    assert "Welcome, Bob!" in result.stdout


{% elif cookiecutter.cli_framework == "argparse" -%}
import pytest
from {{ cookiecutter.project_slug }}.cli import main


def test_cli_greet_command(capsys: pytest.CaptureFixture[str]) -> None:
    """Test CLI greet command with argparse."""
    main(["greet", "Alice"])
    captured = capsys.readouterr()
    assert "Hello, Alice!" in captured.out


def test_cli_greet_with_option(capsys: pytest.CaptureFixture[str]) -> None:
    """Test CLI greet command with custom greeting using argparse."""
    main(["greet", "Bob", "-g", "Welcome"])
    captured = capsys.readouterr()
    assert "Welcome, Bob!" in captured.out


def test_cli_help(capsys: pytest.CaptureFixture[str]) -> None:
    """Test CLI help output on missing command."""
    with pytest.raises(SystemExit) as excinfo:
        main([])
    assert excinfo.value.code == 1


{% endif -%}
