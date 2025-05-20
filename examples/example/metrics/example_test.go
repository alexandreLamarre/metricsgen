package metrics_test

import (
	"context"
	"testing"

	// "time"

	"github.com/alexandreLamarre/metricsgen/examples/example/metrics"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
)

func TestExampleMetrics(t *testing.T) {
	ctx := context.TODO()
	m, err := metrics.NewMetrics(otel.Meter("tcp"))
	require.NoError(t, err)

	m.MetricDummyTcpConnlat.Record(ctx, 0, 1, 1, metrics.WithDummyTcpConnlatCpuId(1))

	m.MetricDummyTcpRx.Record(ctx, 0, 1, metrics.EnumOff)

	m.MetricDummyTcpTx.Record(ctx, 0, 1, 1, 1, metrics.EnumIdle)
}
