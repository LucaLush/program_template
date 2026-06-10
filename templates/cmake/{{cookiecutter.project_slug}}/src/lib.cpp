#include "{{ cookiecutter.project_slug }}/lib.hpp"

namespace {{ cookiecutter.project_slug }} {

std::string get_greet_message() {
    return "Hello from {{ cookiecutter.project_name }}!";
}

} // namespace {{ cookiecutter.project_slug }}
