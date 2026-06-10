#include <gtest/gtest.h>
{% if cookiecutter.project_type == "Library" -%}
#include "{{ cookiecutter.project_slug }}/lib.hpp"
{%- endif %}

TEST(SampleTest, BasicMath) {
    EXPECT_STRNE("hello", "world");
    EXPECT_EQ(1 + 1, 2);
}

{% if cookiecutter.project_type == "Library" -%}
TEST(GreetTest, ValidateLibraryOutput) {
    std::string expected = "Hello from {{ cookiecutter.project_name }}!";
    EXPECT_EQ({{ cookiecutter.project_slug }}::get_greet_message(), expected);
}
{%- endif %}
