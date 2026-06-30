import os
import shutil

def remove_file(filepath):
    if os.path.exists(filepath):
        os.remove(filepath)

def remove_dir(dirpath):
    if os.path.exists(dirpath):
        shutil.rmtree(dirpath)

add_linter = "{{ cookiecutter.add_linter }}"
add_dockerfile = "{{ cookiecutter.add_dockerfile }}"

if add_linter == "no":
    remove_file(".golangci.yml")

if add_dockerfile == "no":
    remove_file("Dockerfile")
