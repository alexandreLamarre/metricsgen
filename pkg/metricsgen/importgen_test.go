package metricsgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type importTestCase struct {
	importDefs []importDef
	expected   string
}

func TestImportGen(t *testing.T) {

	tcs := []importTestCase{
		{
			importDefs: []importDef{},
			expected:   "",
		},
		{
			importDefs: []importDef{
				{

					dependency: "github.com/example",
				},
			},
			expected: `import (
	"github.com/example"
)
`,
		},
		{
			importDefs: []importDef{
				{

					dependency: "github.com/example",
				},
				{
					dependency: "strings",
				},
			},
			expected: `import (
	"github.com/example"
	"strings"
)
`,
		},
		{
			importDefs: []importDef{
				{
					alias:      "ex",
					dependency: "github.com/example",
				},
				{
					alias:      "s",
					dependency: "strings",
				},
			},
			expected: `import (
	ex "github.com/example"
	s "strings"
)
`,
		},
	}
	for _, tc := range tcs {
		i := NewImportWriter(tc.importDefs)
		generated := i.Generate()
		require.Equal(t, tc.expected, generated)
	}

}
