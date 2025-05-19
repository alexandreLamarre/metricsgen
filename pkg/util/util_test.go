package util_test

import (
	"testing"

	"github.com/alexandreLamarre/metricsgen/pkg/util"
	"github.com/stretchr/testify/require"
)

type stringConversionTc struct {
	in       string
	expected string
}

func TestUtilCamelCase(t *testing.T) {
	tcs := []stringConversionTc{
		{
			in:       "pid.namespace",
			expected: "PidNamespace",
		},
		{
			in:       "bpf_helpers.pid",
			expected: "BpfHelpersPid",
		},
		{
			in:       "otel-driver.namespace",
			expected: "OtelDriverNamespace",
		},
	}

	for _, tc := range tcs {
		got := util.OtelStringToCamelCase(tc.in)
		require.Equal(t, tc.expected, got)
	}
}

func TestUtilCamelCaseField(t *testing.T) {
	tcs := []stringConversionTc{
		{
			in:       "pid.namespace",
			expected: "pidNamespace",
		},
		{
			in:       "bpf_helpers.pid",
			expected: "bpfHelpersPid",
		},
		{
			in:       "otel-driver.namespace",
			expected: "otelDriverNamespace",
		},
	}

	for _, tc := range tcs {
		got := util.OtelStringToCamelCaseField(tc.in)
		require.Equal(t, tc.expected, got)
	}
}
