package conformance

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alexandreLamarre/metricsgen/examples/prometheus"
	"github.com/alexandreLamarre/metricsgen/tests/conformance/metrics/otel"
	promsdk "github.com/prometheus/client_golang/prometheus"
	promtestutil "github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"
	expprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func TestMetricsCounterConformanceAllValues(t *testing.T) {
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

func TestMetricsCounterConformanceEmptyValues(t *testing.T) {
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

func TestMetricsCounterConformanceOptionalAllValues(t *testing.T) {
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

func TestMetricsCounterConformanceEmptyAllValues(t *testing.T) {
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
