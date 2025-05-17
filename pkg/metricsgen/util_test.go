package metricsgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type stringConversionTc struct {
	in       string
	expected string
}

func TestUtils(t *testing.T) {
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
		got := OtelStringToCamelCase(tc.in)
		require.Equal(t, tc.expected, got)
	}

}
