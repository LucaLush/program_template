"""Pytest fixtures and test configurations."""

import pytest
from {{ cookiecutter.project_slug }}.core import Greeter


@pytest.fixture
def default_greeter() -> Greeter:
    """Fixture providing a default Greeter instance."""
    return Greeter()


@pytest.fixture
def custom_greeter() -> Greeter:
    """Fixture providing a custom Greeter instance."""
    return Greeter(greeting="Hi")
