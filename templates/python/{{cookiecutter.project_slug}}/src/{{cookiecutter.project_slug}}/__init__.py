"""{{ cookiecutter.project_name }} - {{ cookiecutter.project_description }}"""

from {{ cookiecutter.project_slug }}.core import Greeter, greet

__version__ = "0.1.0"
__all__ = ["Greeter", "greet", "__version__"]
