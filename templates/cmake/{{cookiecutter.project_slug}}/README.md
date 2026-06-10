# {{ cookiecutter.project_name }}

A modern C++ {{ cookiecutter.project_type.lower() }} project.

## Requirements

* CMake >= {{ cookiecutter.cmake_minimum_version }}
* C++ compiler supporting C++{{ cookiecutter.cpp_standard }}
{%- if cookiecutter.use_vcpkg == "yes" %}
* vcpkg package manager (ensure the `VCPKG_ROOT` environment variable is set)
{%- endif %}
{%- if cookiecutter.use_doxygen == "yes" %}
* Doxygen (optional, for documentation generation)
{%- endif %}

## Building the Project

### Using CMake Presets (Recommended)

1. Configure the project:
   ```bash
   cmake --preset default
   ```
2. Build the project:
   ```bash
   cmake --build --preset default
   ```
{%- if cookiecutter.add_testing == "yes" %}
3. Run tests:
   ```bash
   ctest --preset default
   ```
{%- endif %}

### Traditional CMake

1. Create a build directory and configure:
   ```bash
   cmake -B build -S . {% if cookiecutter.use_vcpkg == "yes" %}-DCMAKE_TOOLCHAIN_FILE=$VCPKG_ROOT/scripts/buildsystems/vcpkg.cmake{% endif %}
   ```
2. Build:
   ```bash
   cmake --build build
   ```
{%- if cookiecutter.add_testing == "yes" %}
3. Run tests:
   ```bash
   cd build && ctest
   ```
{%- endif %}

{%- if cookiecutter.use_doxygen == "yes" %}
## Generating Documentation

To build documentation, run the following target:
```bash
cmake --build build --target doc
```
The HTML documentation will be generated in `docs/html/index.html`.
{%- endif %}

## Installation

To install the project files to a system or specific directory:

```bash
# Build the project first
cmake --build build --config Release

# Install to a custom prefix directory
cmake --install build --prefix /path/to/install/directory
```

{% if cookiecutter.project_type == "Library" -%}
### Consuming the Library in Downstream Projects

Once installed, downstream C++ CMake projects can easily find and link to this library:

```cmake
# In the downstream project's CMakeLists.txt
find_package({{ cookiecutter.project_name }} REQUIRED)

add_executable(my_app main.cpp)
target_link_libraries(my_app PRIVATE {{ cookiecutter.project_name }}::{{ cookiecutter.project_name }})
```
Ensure you add the library's installation prefix to the downstream project's `CMAKE_PREFIX_PATH` if it was installed to a non-standard location:
```bash
cmake -B build -S . -DCMAKE_PREFIX_PATH=/path/to/install/directory
```
{%- endif %}

