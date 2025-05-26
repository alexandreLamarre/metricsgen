package benchmark_test

import (
	"testing"

	benchmark "github.com/alexandreLamarre/metricsgen/examples/benchmark/prometheus"
	promsdk "github.com/prometheus/client_golang/prometheus"
)

func promReg() *promsdk.Registry {

	return promsdk.NewRegistry()
}

func BenchmarkNoLabels(b *testing.B) {
	reg := promReg()

	metrics, err := benchmark.NewPrometheusMetrics(reg)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricNoLabelCounter.Add(1)
	}

}

func BenchmarkOneLabel(b *testing.B) {
	reg := promReg()

	metrics, err := benchmark.NewPrometheusMetrics(reg)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricOneLabelCounter.Add(1, "a")
	}
}

func BenchmarkFourLabel(b *testing.B) {
	reg := promReg()

	metrics, err := benchmark.NewPrometheusMetrics(reg)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricFourLabelCounter.Add(1, "a", "b", "c", "d")
	}
}

func BenchmarkEightLabel(b *testing.B) {
	reg := promReg()

	metrics, err := benchmark.NewPrometheusMetrics(reg)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricEightLabelCounter.Add(1, "a", "b", "c", "d", "e", "f", "g", "h")
	}
}

func BenchmarkSplitLabels(b *testing.B) {
	//TODO
	reg := promReg()

	metrics, err := benchmark.NewPrometheusMetrics(reg)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricSplitLabelCounter.Add(
			1, "a", "b", "c", "d",
			benchmark.WithSplitLabelCounterStringLabel5("e"),
			benchmark.WithSplitLabelCounterStringLabel6("f"),
			benchmark.WithSplitLabelCounterStringLabel7("g"),
			benchmark.WithSplitLabelCounterStringLabel8("h"),
		)
	}
}
