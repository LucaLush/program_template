#include <benchmark/benchmark.h>
{% if cookiecutter.project_type == "Library" -%}
#include "{{ cookiecutter.project_slug }}/lib.hpp"
{%- endif %}
#include <string>
#include <vector>

// Benchmark string construction
static void BM_StringCreation(benchmark::State& state) {
    for (auto _ : state) {
        std::string empty_string;
        benchmark::DoNotOptimize(empty_string);
    }
}
BENCHMARK(BM_StringCreation);

{% if cookiecutter.project_type == "Library" -%}
// Benchmark library greeting function
static void BM_LibraryGreet(benchmark::State& state) {
    for (auto _ : state) {
        auto msg = {{ cookiecutter.project_slug }}::get_greet_message();
        benchmark::DoNotOptimize(msg);
    }
}
BENCHMARK(BM_LibraryGreet);
{%- else -%}
// Benchmark dynamic vector allocations
static void BM_VectorPushBack(benchmark::State& state) {
    for (auto _ : state) {
        std::vector<int> v;
        v.reserve(state.range(0));
        for (int i = 0; i < state.range(0); ++i) {
            v.push_back(i);
        }
        benchmark::DoNotOptimize(v.data());
    }
}
BENCHMARK(BM_VectorPushBack)->Range(8, 8 << 10);
{%- endif %}

BENCHMARK_MAIN();
