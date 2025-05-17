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
