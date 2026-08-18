"""Core business logic and library functions for {{ cookiecutter.project_name }}."""

from dataclasses import dataclass
from typing import Optional


@dataclass
class Greeter:
    """A sample greeter class demonstrating typed object-oriented design."""

    greeting: str = "Hello"

    def format_message(self, name: str, punctuation: Optional[str] = "!") -> str:
        """Format a personalized greeting message.

        Args:
            name: Name of the recipient.
            punctuation: Optional punctuation mark to end the greeting with.

        Returns:
            The formatted greeting string.
        """
        p = punctuation if punctuation is not None else ""
        return f"{self.greeting}, {name.strip()}{p}"


def greet(name: str) -> str:
    """Convenience function to generate a default greeting.

    Args:
        name: Name of the recipient.

    Returns:
        Formatted greeting string.
    """
    greeter = Greeter()
    return greeter.format_message(name)
