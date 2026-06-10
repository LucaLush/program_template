# -----------------------------------------------------------------------------
# Test Target Definition
# -----------------------------------------------------------------------------

# Add GoogleTest dependencies
{% if cookiecutter.use_vcpkg == "yes" -%}
find_package(GTest CONFIG REQUIRED)
{%- else -%}
include(FetchContent)
FetchContent_Declare(
    googletest
    URL https://github.com/google/googletest/archive/refs/tags/v1.14.0.tar.gz
)
# For MSVC runtimes alignment
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)
FetchContent_MakeAvailable(googletest)
{%- endif %}

# Unit test executable
add_executable(unit_tests test_main.cpp)

# Link libraries
{% if cookiecutter.use_vcpkg == "yes" -%}
target_link_libraries(unit_tests PRIVATE GTest::gtest GTest::gtest_main)
{%- else -%}
target_link_libraries(unit_tests PRIVATE gtest_main)
{%- endif %}

# If generating a library, link it to verify its functions
{% if cookiecutter.project_type == "Library" -%}
target_link_libraries(unit_tests PRIVATE ${PROJECT_NAME})
{%- endif %}

# Discover tests and register with CTest
include(GoogleTest)
gtest_discover_tests(unit_tests)
