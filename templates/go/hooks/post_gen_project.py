import os
import shutil

def remove_file(filepath):
    if os.path.exists(filepath):
        os.remove(filepath)

def remove_dir(dirpath):
    if os.path.exists(dirpath):
        shutil.rmtree(dirpath)

project_type = "{{ cookiecutter.project_type }}"
add_linter = "{{ cookiecutter.add_linter }}"
add_dockerfile = "{{ cookiecutter.add_dockerfile }}"
add_systemd = "{{ cookiecutter.add_systemd }}"
project_slug = "{{ cookiecutter.project_slug }}"

if project_type == "Library/Package":
    # Clean up service and executable related directories/files
    remove_dir("cmd")
    remove_dir("internal")
    remove_dir("configs")
    remove_dir("tests")
    remove_file("Dockerfile")
    remove_file(f"{project_slug}.service")
else:
    # Clean up library root package files
    remove_file("doc.go")
    remove_file(f"{project_slug}.go")
    remove_file(f"{project_slug}_test.go")
    remove_file("example_test.go")

    if add_dockerfile == "no":
        remove_file("Dockerfile")

    if add_systemd == "no":
        remove_file(f"{project_slug}.service")

if add_linter == "no":
    remove_file(".golangci.yml")
