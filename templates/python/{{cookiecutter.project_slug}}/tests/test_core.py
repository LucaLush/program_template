"""Unit tests for the core library logic."""

from {{ cookiecutter.project_slug }} import __version__
from {{ cookiecutter.project_slug }}.core import Greeter, greet


def test_version() -> None:
    """Test that package version is defined."""
    assert __version__ == "0.1.0"


def test_greeter_default(default_greeter: Greeter) -> None:
    """Test Greeter with default values."""
    assert default_greeter.format_message("Alice") == "Hello, Alice!"


def test_greeter_custom(custom_greeter: Greeter) -> None:
    """Test Greeter with customized greeting prefix."""
    assert custom_greeter.format_message("Bob") == "Hi, Bob!"


def test_greeter_custom_punctuation(default_greeter: Greeter) -> None:
    """Test Greeter with custom punctuation."""
    assert default_greeter.format_message("World", punctuation=".") == "Hello, World."
    assert default_greeter.format_message("World", punctuation=None) == "Hello, World"


def test_greet_convenience() -> None:
    """Test greet convenience function."""
    assert greet("Developer") == "Hello, Developer!"
