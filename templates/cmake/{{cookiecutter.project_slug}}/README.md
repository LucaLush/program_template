# {{ cookiecutter.project_name }}

A modern C++ {{ cookiecutter.project_type.lower() }} project.

## Requirements

* CMake >= {{ cookiecutter.cmake_minimum_version }} (>= 3.17 for the `Ninja Multi-Config` generator, >= 3.25 when using `CMakePresets.json`)
* [Ninja](https://ninja-build.org/) (the presets use the `Ninja Multi-Config` generator)
* C++ compiler supporting C++{{ cookiecutter.cpp_standard }}
{%- if cookiecutter.use_vcpkg == "yes" %}
* vcpkg package manager (ensure the `VCPKG_ROOT` environment variable is set)
{%- endif %}
{%- if cookiecutter.use_doxygen == "yes" %}
* Doxygen (optional, for documentation generation)
{%- endif %}

## Building the Project

### Using CMake Presets (Recommended)

The presets use the `Ninja Multi-Config` generator: configure once into `build/`, then build any of
`Debug`, `RelWithDebInfo` or `Release` without reconfiguring. `compile_commands.json` is exported to `build/`.

1. Configure the project (only needed once):
   ```bash
   cmake --preset default
   ```
2. Build the desired configuration:
   ```bash
   cmake --build --preset debug           # Debug
   cmake --build --preset relwithdebinfo  # RelWithDebInfo
   cmake --build --preset release         # Release
   ```
   Binaries are placed in a per-configuration subdirectory, e.g. `build/Debug/`.
{%- if cookiecutter.add_testing == "yes" %}
3. Run tests for a built configuration:
   ```bash
   ctest --preset debug
   ```
{%- endif %}

### Traditional CMake

1. Configure with the multi-config generator:
   ```bash
   cmake -B build -S . -G "Ninja Multi-Config" {% if cookiecutter.use_vcpkg == "yes" %}-DCMAKE_TOOLCHAIN_FILE=$VCPKG_ROOT/scripts/buildsystems/vcpkg.cmake{% endif %}
   ```
2. Build a configuration:
   ```bash
   cmake --build build --config Debug
   ```
{%- if cookiecutter.add_testing == "yes" %}
3. Run tests:
   ```bash
   ctest --test-dir build -C Debug
   ```
{%- endif %}

{%- if cookiecutter.add_benchmark == "yes" %}
## Running Benchmarks

Benchmarks measure code performance using Google Benchmark. **Run benchmarks in Release mode** for accurate results:

### Using CMake Presets
```bash
# Build benchmark with the release preset
cmake --build --preset release --target benchmarks

# Execute benchmark binary
./build/benchmarks/Release/benchmarks
```

### Traditional CMake
```bash
cmake -B build -S . -G "Ninja Multi-Config" {% if cookiecutter.use_vcpkg == "yes" %}-DCMAKE_TOOLCHAIN_FILE=$VCPKG_ROOT/scripts/buildsystems/vcpkg.cmake{% endif %}
cmake --build build --config Release --target benchmarks
./build/benchmarks/Release/benchmarks
```
{%- endif %}

{%- if cookiecutter.use_doxygen == "yes" %}
## Generating Documentation

To build documentation, run the following target:
```bash
cmake --build --preset debug --target doc
```
The HTML documentation will be generated in `docs/html/index.html`.
{%- endif %}

## Installation

To install the project files to a system or specific directory:

```bash
# Build the project first
cmake --build --preset release

# Install the Release configuration to a custom prefix directory
cmake --install build --config Release --prefix /path/to/install/directory
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

