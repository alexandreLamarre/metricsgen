package metricsgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type metricInfoTc struct {
	in       *Config
	expected string
}

func TestMetricGenInfo(t *testing.T) {
	tcs := []metricInfoTc{
		{
			in: &Config{
				Metrics: map[string]Metric{
					"hello": {
						Name: "hello",
					},
				},
			},
			expected: `type MetricInfo struct {
	*MetricHello
}
`,
		},
		{
			in: &Config{
				Metrics: map[string]Metric{
					"bpf.tcp.tx": {
						Name: "bpf.tcp.tx",
					},
					"bpf.tcp.rx": {
						Name: "bpf.tcp.rx",
					},
				},
			},
			expected: `type MetricInfo struct {
	*MetricBpfTcpTx
	*MetricBpfTcpRx
}
`,
		},
	}

	for _, tc := range tcs {
		got := tc.in.GenerateMetricInfo()
		require.Equal(t, tc.expected, got)
	}

}

type metricAttributeStructTc struct {
	in       *Config
	expected string
}

func TestMetricsGenAttributeStruct(t *testing.T) {
	tcs := []metricAttributeStructTc{
		{
			in: &Config{
				Metrics: map[string]Metric{
					"bpf.tcp.rx": {
						Name:       "bpf.tcp.rx",
						Attributes: []string{"pid.id"},
					},
					"bpf.tcp.tx": {
						Name:       "bpf.tcp.tx",
						Attributes: []string{"pid.id", "pid.gid"},
					},
				},
				Attributes: map[string]Attribute{
					"pid.id": {
						Name: "pid.id",
					},
					"pid.gid": {
						Name: "pid.gid",
					},
				},
			},
			expected: `type MetricBpfTcpRxAttributes struct {
	PidId string
}
type MetricBpfTcpTxAttributes struct {
	PidId string
	PidGid string
}
`,
		},
	}

	for _, tc := range tcs {
		got := tc.in.GenerateAllMetricAttributesStruct()
		require.Equal(t, tc.expected, got)

	}
}

type metricStructTc struct {
	in       *Config
	expected string
}

func TestMetricsGenMetricsStruct(t *testing.T) {
	tcs := []metricStructTc{
		{
			in: &Config{
				Metrics: map[string]Metric{
					"bpf.tcp.tx": {
						Name:  "bpf.tcp.tx",
						Short: "TCP transmitted bytes per process",
					},
				},
			},
			expected: `//MetricBpfTcpTx TCP transmitted bytes per process
type MetricBpfTcpTx struct {
}
`,
		},
		{
			in: &Config{
				Metrics: map[string]Metric{
					"bpf.tcp.tx": {
						Name:  "bpf.tcp.tx",
						Short: "TCP transmitted bytes per process",
					},
					"bpf.tcp.rx": {
						Name:  "bpf.tcp.rx",
						Short: "TCP received bytes per process",
					},
				},
			},
			expected: `//MetricBpfTcpTx TCP transmitted bytes per process
type MetricBpfTcpTx struct {
}
//MetricBpfTcpRx TCP received bytes per process
type MetricBpfTcpRx struct {
}
`,
		},
	}

	for _, tc := range tcs {
		got := tc.in.GenerateAllMetricsStruct()
		require.Equal(t, tc.expected, got)
	}
}
