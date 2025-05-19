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

	m.MetricBpfTcpConnlat.Record(ctx, 0, 1, 1)

	m.MetricBpfTcpRx.Record(ctx, 0, 1)

	m.MetricBpfTcpTx.Record(ctx, 0, 1, 1, metrics.WithBpfTcpTxCpuId(1))
}
