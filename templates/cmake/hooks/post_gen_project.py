import os
import shutil

def remove_file(filepath):
    if os.path.exists(filepath):
        os.remove(filepath)

def remove_dir(dirpath):
    if os.path.exists(dirpath):
        shutil.rmtree(dirpath)

# Retrieve Cookiecutter variables
project_type = "{{ cookiecutter.project_type }}"
use_vcpkg = "{{ cookiecutter.use_vcpkg }}"
use_doxygen = "{{ cookiecutter.use_doxygen }}"
add_testing = "{{ cookiecutter.add_testing }}"

# Cleanup vcpkg
if use_vcpkg == "no":
    remove_file("vcpkg.json")

# Cleanup doxygen
if use_doxygen == "no":
    remove_dir("docs")

# Cleanup testing
if add_testing == "no":
    remove_dir("tests")

# Cleanup source code structure depending on project type
if project_type == "Executable":
    remove_dir("include")
    remove_file(os.path.join("src", "lib.cpp"))
elif project_type == "Library":
    remove_file(os.path.join("src", "main.cpp"))
