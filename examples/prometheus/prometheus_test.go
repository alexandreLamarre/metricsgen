package prometheus_test

import (
	"testing"

	"github.com/alexandreLamarre/metricsgen/examples/prometheus"
	promsdk "github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestMetrics(t *testing.T) {
	reg := promsdk.NewRegistry()
	err := reg.Register(promsdk.NewCounter(promsdk.CounterOpts{
		Name:      "help_metric",
		Help:      "heeeelp",
		Namespace: "",
		Subsystem: "",
	}))
	require.NoError(t, err)
	m, err := prometheus.NewPrometheusMetrics(reg)
	require.NoError(t, err)
	require.NotNil(t, m)
}
