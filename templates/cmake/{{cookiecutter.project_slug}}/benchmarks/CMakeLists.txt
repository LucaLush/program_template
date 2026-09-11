# -----------------------------------------------------------------------------
# Benchmark Target Definition
# -----------------------------------------------------------------------------

# Add Google Benchmark dependencies
{% if cookiecutter.use_vcpkg == "yes" -%}
find_package(benchmark CONFIG REQUIRED)
{%- else -%}
include(FetchContent)
FetchContent_Declare(
    googlebenchmark
    URL https://github.com/google/benchmark/archive/refs/tags/v1.9.1.tar.gz
)
set(BENCHMARK_ENABLE_TESTING OFF CACHE BOOL "" FORCE)
set(BENCHMARK_ENABLE_INSTALL OFF CACHE BOOL "" FORCE)
set(BENCHMARK_ENABLE_WERROR OFF CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googlebenchmark)
{%- endif %}

# Benchmark executable
add_executable(benchmarks bench_main.cpp)

# Link benchmark libraries
target_link_libraries(benchmarks PRIVATE benchmark::benchmark)

# If generating a library, link it to benchmark its functions
{% if cookiecutter.project_type == "Library" -%}
target_link_libraries(benchmarks PRIVATE ${PROJECT_NAME})
{%- endif %}
