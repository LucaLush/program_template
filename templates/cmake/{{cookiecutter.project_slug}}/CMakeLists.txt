cmake_minimum_required(VERSION {{ cookiecutter.cmake_minimum_version }})

project({{ cookiecutter.project_slug }}
    VERSION 0.1.0
    DESCRIPTION "{{ cookiecutter.project_name }}"
    LANGUAGES CXX
)

# Configure C++ standard
set(CMAKE_CXX_STANDARD {{ cookiecutter.cpp_standard }})
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Compiler-specific warning flags
if(MSVC)
    add_compile_options(/W4 /WX)
else()
    add_compile_options(-Wall -Wextra -Wpedantic -Werror)
endif()

{% if cookiecutter.project_type == "Executable" -%}
# Executable target
add_executable(${PROJECT_NAME} src/main.cpp)

target_include_directories(${PROJECT_NAME}
    PRIVATE
        ${CMAKE_CURRENT_SOURCE_DIR}/src
)

# Install rules (Executable)
include(GNUInstallDirs)
install(TARGETS ${PROJECT_NAME}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
)
{%- else -%}
# Library target
add_library(${PROJECT_NAME} src/lib.cpp)

target_include_directories(${PROJECT_NAME}
    PUBLIC
        $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
        $<INSTALL_INTERFACE:include>
    PRIVATE
        ${CMAKE_CURRENT_SOURCE_DIR}/src
)

# Install & Export rules (Library - Industry Standard)
include(GNUInstallDirs)
include(CMakePackageConfigHelpers)

# Install library binary
install(TARGETS ${PROJECT_NAME}
    EXPORT ${PROJECT_NAME}Targets
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
    INCLUDES DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}
)

# Install library headers
install(DIRECTORY include/
    DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}
)

# Generate and install CMake export targets file
install(EXPORT ${PROJECT_NAME}Targets
    FILE ${PROJECT_NAME}Targets.cmake
    NAMESPACE ${PROJECT_NAME}::
    DESTINATION ${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}
)

# Generate and install CMake package Config and Version files
write_basic_package_version_file(
    "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}ConfigVersion.cmake"
    VERSION ${PROJECT_VERSION}
    COMPATIBILITY SameMajorVersion
)

configure_package_config_file(
    "${CMAKE_CURRENT_SOURCE_DIR}/cmake/Config.cmake.in"
    "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}Config.cmake"
    INSTALL_DESTINATION ${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}
)

install(
    FILES
        "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}Config.cmake"
        "${CMAKE_CURRENT_BINARY_DIR}/${PROJECT_NAME}ConfigVersion.cmake"
    DESTINATION
        ${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}
)
{%- endif %}

# -----------------------------------------------------------------------------
# Doxygen Configuration
# -----------------------------------------------------------------------------
{% if cookiecutter.use_doxygen == "yes" -%}
find_package(Doxygen QUIET)
if(DOXYGEN_FOUND)
    set(DOXYGEN_IN_FILE ${CMAKE_CURRENT_SOURCE_DIR}/docs/Doxyfile.in)
    set(DOXYGEN_OUT_FILE ${CMAKE_CURRENT_BINARY_DIR}/Doxyfile)

    # Configure Doxyfile with build-specific directories
    configure_file(${DOXYGEN_IN_FILE} ${DOXYGEN_OUT_FILE} @ONLY)

    # Add custom target to build docs
    add_custom_target(doc
        COMMAND ${DOXYGEN_EXECUTABLE} ${DOXYGEN_OUT_FILE}
        WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
        COMMENT "Generating API documentation with Doxygen"
        VERBATIM
    )
else()
    message(STATUS "Doxygen not found, 'doc' target will not be available")
endif()
{%- endif %}

# -----------------------------------------------------------------------------
# Testing Configuration
# -----------------------------------------------------------------------------
{% if cookiecutter.add_testing == "yes" -%}
enable_testing()
add_subdirectory(tests)
{%- endif %}
