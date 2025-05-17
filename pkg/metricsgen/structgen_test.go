package metricsgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type structGenTc struct {
	name     string
	desc     string
	attrs    []attributeDef
	expected string
}

func TestStructGen(t *testing.T) {
	tcs := []structGenTc{
		{
			name:  "example",
			desc:  "",
			attrs: []attributeDef{},
			expected: `type example struct {
}
`,
		},
		{
			name:  "example",
			desc:  "this is an empty struct",
			attrs: []attributeDef{},
			expected: `//example this is an empty struct
type example struct {
}
`,
		},
		{
			name: "example",
			desc: "this is an empty struct",
			attrs: []attributeDef{
				{
					field:    "Hello",
					attrType: "string",
					pointer:  false,
				},
				{
					field:    "hello",
					attrType: "int",
					pointer:  true,
				},
				{
					attrType: "Config",
					pointer:  false,
				},
				{
					attrType: "Config2",
					pointer:  true,
				},
			},
			expected: `//example this is an empty struct
type example struct {
	Hello string
	hello *int
	Config
	*Config2
}
`,
		},
	}

	for _, tc := range tcs {

		s := NewStructWriter(
			tc.name,
			tc.desc,
			tc.attrs,
		)

		generated := s.Generate()
		require.Equal(t, tc.expected, generated)
	}
}
