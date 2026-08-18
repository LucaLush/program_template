import os
import shutil

def remove_file(filepath):
    if os.path.exists(filepath):
        os.remove(filepath)

def remove_dir(dirpath):
    if os.path.exists(dirpath):
        shutil.rmtree(dirpath)

project_type = "{{ cookiecutter.project_type }}"
add_dockerfile = "{{ cookiecutter.add_dockerfile }}"
add_docs = "{{ cookiecutter.add_docs }}"
add_github_actions = "{{ cookiecutter.add_github_actions }}"
project_slug = "{{ cookiecutter.project_slug }}"

if project_type == "Library/Package":
    remove_file(os.path.join("src", project_slug, "cli.py"))
    remove_file(os.path.join("src", project_slug, "__main__.py"))
    remove_file(os.path.join("tests", "test_cli.py"))

if add_dockerfile == "no":
    remove_file("Dockerfile")
    remove_file(".dockerignore")

if add_docs == "no":
    remove_dir("docs")
    remove_file("mkdocs.yml")

if add_github_actions == "no":
    remove_dir(".github")
