package main

import (
	"context"
	"log"
	"net/http"

	"github.com/alexandreLamarre/metricsgen/tests/conformance/metrics/otel"
	"github.com/alexandreLamarre/metricsgen/tests/conformance/metrics/prometheus"

	promsdk "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	expprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	if err != nil {
		panic(err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)

	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	if err != nil {
		panic(err)
	}

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	if err != nil {
		panic(err)
	}

	recordOtelMetrics(om)
	recordPromMetrics(pm)

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics/prometheus", func(w http.ResponseWriter, r *http.Request) {
		promhttp.HandlerFor(promReg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	})

	mux.HandleFunc("/metrics/otel", func(w http.ResponseWriter, r *http.Request) {
		promhttp.HandlerFor(otelReg, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	})
	srv := &http.Server{
		Handler: mux,
		Addr:    ":2222",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func recordOtelMetrics(om otel.Metrics) {
	ctx := context.TODO()
	om.MetricExampleCounter.Record(
		ctx, 1,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
		otel.EnumExampleEnumOn,
		otel.EnumExampleEnum2On,
	)

	om.MetricExampleCounterOptional.Record(ctx, 1,
		otel.WithExampleCounterOptionalExampleString("a"),
		otel.WithExampleCounterOptionalExampleInt(1),
		otel.WithExampleCounterOptionalExampleFloat(0.1),
		otel.WithExampleCounterOptionalExampleBool(false),
		otel.WithExampleCounterOptionalExampleInt64(14),
		otel.WithExampleCounterOptionalExampleFloatSlice([]float64{1.2, 2.3}),
		otel.WithExampleCounterOptionalExampleBoolSlice([]bool{true, false, true}),
		otel.WithExampleCounterOptionalExampleIntSlice([]int{1, 2}),
		otel.WithExampleCounterOptionalExampleInt64Slice([]int64{1, 2, 3}),
		otel.WithExampleCounterOptionalExampleStringSlice([]string{"a", "b", "c"}),
	)

	om.MetricExampleGauge.Record(
		ctx, 1,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)

	om.MetricExampleHistogram.Record(
		ctx,
		10,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)

	om.MetricExampleHistogramCustomized.Record(
		ctx,
		10,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)
}

func recordPromMetrics(pm prometheus.PrometheusMetrics) {
	pm.MetricExampleCounter.Add(
		1,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
		prometheus.EnumExampleEnumOn,
		prometheus.EnumExampleEnum2On,
	)

	pm.MetricExampleCounterOptional.Add(1,
		prometheus.WithExampleCounterOptionalExampleString("a"),
		prometheus.WithExampleCounterOptionalExampleInt(1),
		prometheus.WithExampleCounterOptionalExampleFloat(0.1),
		prometheus.WithExampleCounterOptionalExampleBool(false),
		prometheus.WithExampleCounterOptionalExampleInt64(14),
		prometheus.WithExampleCounterOptionalExampleFloatSlice([]float64{1.2, 2.3}),
		prometheus.WithExampleCounterOptionalExampleBoolSlice([]bool{true, false, true}),
		prometheus.WithExampleCounterOptionalExampleIntSlice([]int{1, 2}),
		prometheus.WithExampleCounterOptionalExampleInt64Slice([]int64{1, 2, 3}),
		prometheus.WithExampleCounterOptionalExampleStringSlice([]string{"a", "b", "c"}),
	)

	pm.MetricExampleGauge.Add(
		1,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)

	pm.MetricExampleHistogram.Observe(
		10,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)

	pm.MetricExampleHistogramCustomized.Observe(
		10,
		"a",
		1,
		0.1,
		false,
		14,
		[]float64{1.2, 2.3},
		[]bool{true, false, true},
		[]int{1, 2},
		[]int64{1, 2, 3},
		[]string{"a", "b", "c"},
	)
}
