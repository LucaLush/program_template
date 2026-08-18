"""Package executable entrypoint for `python -m {{ cookiecutter.project_slug }}`."""

from {{ cookiecutter.project_slug }}.cli import main

if __name__ == "__main__":
    main()
