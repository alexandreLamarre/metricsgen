package main_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alexandreLamarre/metricsgen/tests/conformance/metrics/otel"
	"github.com/alexandreLamarre/metricsgen/tests/conformance/metrics/prometheus"
	promsdk "github.com/prometheus/client_golang/prometheus"
	promtestutil "github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
	expprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

// this doubles as a conformance test for counters
func TestMetricsConformanceAllAttributeTypes(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

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

	pm.MetricExampleCounter.Add(1,
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

	optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedFmt := `# HELP "example_counter_total" Example Counter
# TYPE "example_counter_total" counter
{"example_counter_total","example_bool"="false","example_boolSlice"="[true false true]","example_enum"="on","example_enum2"="on","example_float"="0.1","example_floatSlice"="[1.2,2.3]","example_int"="1","example_int64"="14","example_int64Slice"="[1,2,3]","example_intSlice"="[1,2]","example_string"="a","example_stringSlice"="[\"a\",\"b\",\"c\"]"%s} 1
`

	expectedPromRd := strings.NewReader(fmt.Sprintf(expectedFmt, ""))
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_counter_total")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")

	expectedOtelRd := strings.NewReader(fmt.Sprintf(expectedFmt, optionalOtelStuff))
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_counter_total")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")
}

func TestMetricsConformanceAllAttributeTypesEmptyValues(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

	om.MetricExampleCounter.Record(ctx, 1,
		"",
		0,
		0,
		false,
		0,
		[]float64{},
		[]bool{},
		[]int{},
		[]int64{},
		[]string{},
		otel.EnumExampleEnumOn,
		otel.EnumExampleEnum2On,
	)

	pm.MetricExampleCounter.Add(1,
		"",
		0,
		0,
		false,
		0,
		[]float64{},
		[]bool{},
		[]int{},
		[]int64{},
		[]string{},
		prometheus.EnumExampleEnumOn,
		prometheus.EnumExampleEnum2On,
	)

	optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedFmt := `# HELP "example_counter_total" Example Counter
# TYPE "example_counter_total" counter
	example_counter_total{example_bool="false",example_boolSlice="[]",example_enum="on",example_enum2="on",example_float="0",example_floatSlice="[]",example_int="0",example_int64="0",example_int64Slice="[]",example_intSlice="[]",example_string="",example_stringSlice="[]"%s} 1
`
	expectedOtelRd := strings.NewReader(fmt.Sprintf(expectedFmt, optionalOtelStuff))
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_counter_total")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")

	expectedPromRd := strings.NewReader(fmt.Sprintf(expectedFmt, ""))
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_counter_total")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")
}

func TestMetricsConformanceAllOptionalAttributes(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

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

	optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedFmt := `# HELP "example_counter_optional_total" Example Counter
# TYPE "example_counter_optional_total" counter
{"example_counter_optional_total","example_bool"="false","example_boolSlice"="[true false true]","example_float"="0.1","example_floatSlice"="[1.2,2.3]","example_int"="1","example_int64"="14","example_int64Slice"="[1,2,3]","example_intSlice"="[1,2]","example_string"="a","example_stringSlice"="[\"a\",\"b\",\"c\"]"%s} 1
`

	expectedOtelRd := strings.NewReader(fmt.Sprintf(expectedFmt, optionalOtelStuff))
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_counter_optional_total")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")

	expectedPromRd := strings.NewReader(fmt.Sprintf(expectedFmt, ""))
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_counter_optional_total")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")
}

func TestMetricsConformanceAllOptionalAttributesEmptyValues(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

	om.MetricExampleCounterOptional.Record(ctx, 1,
		otel.WithExampleCounterOptionalExampleString(""),
		otel.WithExampleCounterOptionalExampleInt(0),
		otel.WithExampleCounterOptionalExampleFloat(0.0),
		otel.WithExampleCounterOptionalExampleBool(false),
		otel.WithExampleCounterOptionalExampleInt64(0),
		otel.WithExampleCounterOptionalExampleFloatSlice([]float64{}),
		otel.WithExampleCounterOptionalExampleBoolSlice([]bool{}),
		otel.WithExampleCounterOptionalExampleIntSlice([]int{}),
		otel.WithExampleCounterOptionalExampleInt64Slice([]int64{}),
		otel.WithExampleCounterOptionalExampleStringSlice([]string{}),
	)

	pm.MetricExampleCounterOptional.Add(1,
		prometheus.WithExampleCounterOptionalExampleString(""),
		prometheus.WithExampleCounterOptionalExampleInt(0),
		prometheus.WithExampleCounterOptionalExampleFloat(0.0),
		prometheus.WithExampleCounterOptionalExampleBool(false),
		prometheus.WithExampleCounterOptionalExampleInt64(0),
		prometheus.WithExampleCounterOptionalExampleFloatSlice([]float64{}),
		prometheus.WithExampleCounterOptionalExampleBoolSlice([]bool{}),
		prometheus.WithExampleCounterOptionalExampleIntSlice([]int{}),
		prometheus.WithExampleCounterOptionalExampleInt64Slice([]int64{}),
		prometheus.WithExampleCounterOptionalExampleStringSlice([]string{}),
	)
	optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedFmt := `# HELP "example_counter_optional_total" Example Counter
# TYPE "example_counter_optional_total" counter
	example_counter_optional_total{example_bool="false",example_boolSlice="[]",example_float="0",example_floatSlice="[]",example_int="0",example_int64="0",example_int64Slice="[]",example_intSlice="[]",example_string="",example_stringSlice="[]"%s} 1
`

	expectedOtelRd := strings.NewReader(fmt.Sprintf(expectedFmt, optionalOtelStuff))
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_counter_optional_total")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")

	expectedPromRd := strings.NewReader(fmt.Sprintf(expectedFmt, ""))
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_counter_optional_total")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")
}

func TestMetricsGaugeConformance(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

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

	pm.MetricExampleGauge.Add(1,
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

	optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedFmt := `# HELP "example_gauge" Example Gauge
# TYPE "example_gauge" gauge
{"example_gauge","example_bool"="false","example_boolSlice"="[true false true]","example_float"="0.1","example_floatSlice"="[1.2,2.3]","example_int"="1","example_int64"="14","example_int64Slice"="[1,2,3]","example_intSlice"="[1,2]","example_string"="a","example_stringSlice"="[\"a\",\"b\",\"c\"]"%s} 1
`

	expectedOtelRd := strings.NewReader(fmt.Sprintf(expectedFmt, optionalOtelStuff))
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_gauge")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")
	expectedPromRd := strings.NewReader(fmt.Sprintf(expectedFmt, ""))
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_gauge")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")
}

func TestMetricsHistogramConformance(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

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

	// optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	// default buckets are set as : []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	expectedOtelFmt := `# HELP example_histogram_milliseconds Example Histogram
# TYPE example_histogram_milliseconds histogram
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.005"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.01"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.025"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.05"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.1"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.25"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="0.5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="1"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="2.5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="10"} 1
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="+Inf"} 1
{"example_histogram_milliseconds_sum",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version=""} 10
{"example_histogram_milliseconds_count",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version=""} 1
`

	expectedPromFmt := `# HELP example_histogram_milliseconds Example Histogram
# TYPE example_histogram_milliseconds histogram
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.005"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.01"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.025"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.05"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.1"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.25"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="0.5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="1"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="2.5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="5"} 0
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="10"} 1
{"example_histogram_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="+Inf"} 1
{"example_histogram_milliseconds_sum",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]"} 10
{"example_histogram_milliseconds_count",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]"} 1
`

	expectedOtelRd := strings.NewReader(expectedOtelFmt)
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_histogram_milliseconds")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")

	expectedPromRd := strings.NewReader(expectedPromFmt)
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_histogram_milliseconds")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")

}

func TestMetricsHistogramConformanceCustomBuckets(t *testing.T) {
	ctx := context.TODO()
	otelReg := promsdk.NewRegistry()
	promReg := promsdk.NewRegistry()
	exporter, err := expprom.New(
		expprom.WithRegisterer(otelReg),
	)
	require.NoError(t, err)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(exporter),
	)
	om, err := otel.NewMetrics(meterProvider.Meter("conformance"))
	require.NoError(t, err)
	require.NotNil(t, om)

	pm, err := prometheus.NewPrometheusMetrics(promReg)
	require.NoError(t, err)
	require.NotNil(t, pm)

	om.MetricExampleHistogramCustomized.Record(
		ctx,
		3,
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
		3,
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

	// optionalOtelStuff := `,otel_scope_name="conformance",otel_scope_version=""`

	expectedOtelFmt := `# HELP example_histogram_customized_milliseconds Example Exponential Histogram
# TYPE example_histogram_customized_milliseconds histogram
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="1"} 0
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="2"} 0
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="3"} 1
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="4"} 1
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version="",le="+Inf"} 1
{"example_histogram_customized_milliseconds_sum",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version=""} 3
{"example_histogram_customized_milliseconds_count",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",otel_scope_name="conformance",otel_scope_version=""} 1
`

	expectedPromFmt := `# HELP example_histogram_customized_milliseconds Example Exponential Histogram
# TYPE example_histogram_customized_milliseconds histogram
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="1"} 0
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="2"} 0
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="3"} 1
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="4"} 1
{"example_histogram_customized_milliseconds_bucket",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]",le="+Inf"} 1
{"example_histogram_customized_milliseconds_sum",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]"} 3
{"example_histogram_customized_milliseconds_count",example_bool="false",example_boolSlice="[true false true]",example_float="0.1",example_floatSlice="[1.2,2.3]",example_int="1",example_int64="14",example_int64Slice="[1,2,3]",example_intSlice="[1,2]",example_string="a",example_stringSlice="[\"a\",\"b\",\"c\"]"} 1
`

	expectedOtelRd := strings.NewReader(expectedOtelFmt)
	compareOtelErr := promtestutil.CollectAndCompare(otelReg, expectedOtelRd, "example_histogram_customized_milliseconds")
	require.NoError(t, compareOtelErr, "failed to match otel sdk metrics")

	expectedPromRd := strings.NewReader(expectedPromFmt)
	comparePromErr := promtestutil.CollectAndCompare(promReg, expectedPromRd, "example_histogram_customized_milliseconds")
	require.NoError(t, comparePromErr, "failed to match prom sdk metrics")
}
