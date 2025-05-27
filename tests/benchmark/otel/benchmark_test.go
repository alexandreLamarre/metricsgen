package benchmark_test

import (
	"context"
	"testing"

	benchmark "github.com/alexandreLamarre/metricsgen/tests/benchmark/otel"
	promsdk "github.com/prometheus/client_golang/prometheus"
	expprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func otelMeter(name string) metric.Meter {
	registry := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(registry),
	)
	if err != nil {
		panic(err)
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)

	return meterProvider.Meter(name)
}

func BenchmarkNoLabels(b *testing.B) {
	meter := otelMeter("nolabels")

	metrics, err := benchmark.NewMetrics(meter)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricNoLabelCounter.Record(ctx, 1)
	}

}

func BenchmarkOneLabel(b *testing.B) {
	meter := otelMeter("onelabel")

	metrics, err := benchmark.NewMetrics(meter)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricOneLabelCounter.Record(ctx, 1, "a")
	}
}

func BenchmarkFourLabel(b *testing.B) {
	meter := otelMeter("fourlabels")

	metrics, err := benchmark.NewMetrics(meter)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricFourLabelCounter.Record(ctx, 1, "a", "b", "c", "d")
	}
}

func BenchmarkEightLabel(b *testing.B) {
	meter := otelMeter("fourlabels")

	metrics, err := benchmark.NewMetrics(meter)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricEightLabelCounter.Record(ctx, 1, "a", "b", "c", "d", "e", "f", "g", "h")
	}
}

func BenchmarkSplitLabels(b *testing.B) {
	meter := otelMeter("splitLabels")
	metrics, err := benchmark.NewMetrics(meter)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		metrics.MetricSplitLabelCounter.Record(
			ctx, 1, "a", "b", "c", "d",
			benchmark.WithSplitLabelCounterStringLabel5("e"),
			benchmark.WithSplitLabelCounterStringLabel6("f"),
			benchmark.WithSplitLabelCounterStringLabel7("g"),
			benchmark.WithSplitLabelCounterStringLabel8("h"),
		)
	}
}
