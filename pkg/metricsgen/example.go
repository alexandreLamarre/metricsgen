package metricsgen

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/otel/metric"
)

type metricBpfTcpRx struct {
	data pmetric.Metric
}

type metricBpfTcpRxOptions struct {
	hello string
}

func defaultHello() string {
	return "world"
}

func defaultMetricBpfTcpRxOptions() *metricBpfTcpRxOptions {
	return &metricBpfTcpRxOptions{
		hello: defaultHello(),
	}
}

type metricBpfTcpRxOption func(*metricBpfTcpRxOptions)

func WithHello(hello string) metricBpfTcpRxOption {
	return func(o *metricBpfTcpRxOptions) {
		o.hello = hello
	}
}

type metricBpfTcpRxRecorder interface {
	Record(ctx context.Context, value int64, options ...metric.AddOption)
}

type pdataDriver struct {
	data pmetric.Metric
}

var _ metricBpfTcpRxRecorder = &pdataDriver{}

type otelDriver struct {
	metric.Int64Counter
}

var _ metricBpfTcpRxRecorder = &otelDriver{}

func (d *otelDriver) Record(ctx context.Context, value int64, options ...metric.AddOption) {
	d.Int64Counter.Add(ctx, value, options...)
}

type promDriver struct {
	data prometheus.Counter
}

func (d *promDriver) Record(ctx context.Context, value int64, options ...metric.AddOption) {
	d.data.Add(float64(value))
}

var _ metricBpfTcpRxRecorder = &promDriver{}

func (d *pdataDriver) Record(ctx context.Context, value int64, options ...metric.AddOption) {
	dp := d.data.Sum().DataPoints().AppendEmpty()
	dp.SetIntValue(value)
}

func New() {
	prometheus.NewCounter(prometheus.CounterOpts{})
}
