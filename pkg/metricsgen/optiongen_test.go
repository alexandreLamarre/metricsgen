package metricsgen

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type optionGenTc struct {
	name        string
	optionAttrs []optionAttr
	expected    string
}

func TestOptionGen(t *testing.T) {
	tcs := []optionGenTc{
		{
			name:        "Example",
			optionAttrs: []optionAttr{},
			expected: `type ExampleOptions struct {
}
type ExampleOption func(o *ExampleOptions)

func (o *ExampleOptions) Apply(opts ...ExampleOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func defaultExampleOptions() *ExampleOptions {
	return &ExampleOptions{
	}
}
`,
		},
		{
			name: "Example",
			optionAttrs: []optionAttr{
				{
					attributeDef: attributeDef{
						field:    "hello",
						attrType: "string",
					},
				},
				{
					attributeDef: attributeDef{
						attrType: "Config",
						pointer:  true,
					},
				},
			},
			expected: `type ExampleOptions struct {
	hello string
	*Config
}
type ExampleOption func(o *ExampleOptions)

func (o *ExampleOptions) Apply(opts ...ExampleOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func defaultExampleOptions() *ExampleOptions {
	return &ExampleOptions{
	}
}

func WithExampleHello(val string) ExampleOption {
	return func(o *ExampleOptions) {
		o.hello = val
	}
}


func WithExampleConfig(val *Config) ExampleOption {
	return func(o *ExampleOptions) {
		o.Config = val
	}
}

`,
		},
		{
			name: "Example",
			optionAttrs: []optionAttr{
				{
					defaultValue: `"a"`,
					attributeDef: attributeDef{
						field:    "hello",
						attrType: "string",
					},
				},
				{
					attributeDef: attributeDef{
						attrType: "Config",
						pointer:  true,
					},
				},
			},
			expected: `type ExampleOptions struct {
	hello string
	*Config
}
type ExampleOption func(o *ExampleOptions)

func (o *ExampleOptions) Apply(opts ...ExampleOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func defaultExampleOptions() *ExampleOptions {
	return &ExampleOptions{
		hello : "a",
	}
}

func WithExampleHello(val string) ExampleOption {
	return func(o *ExampleOptions) {
		o.hello = val
	}
}


func WithExampleConfig(val *Config) ExampleOption {
	return func(o *ExampleOptions) {
		o.Config = val
	}
}

`,
		},
		{
			name: "Example",
			optionAttrs: []optionAttr{
				{
					defaultValue: `"a"`,
					attributeDef: attributeDef{
						field:    "hello",
						attrType: "string",
					},
				},
				{
					defaultValue: `&Config{}`,
					attributeDef: attributeDef{
						attrType: "Config",
						pointer:  true,
					},
				},
			},
			expected: `type ExampleOptions struct {
	hello string
	*Config
}
type ExampleOption func(o *ExampleOptions)

func (o *ExampleOptions) Apply(opts ...ExampleOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func defaultExampleOptions() *ExampleOptions {
	return &ExampleOptions{
		hello : "a",
		Config : &Config{},
	}
}

func WithExampleHello(val string) ExampleOption {
	return func(o *ExampleOptions) {
		o.hello = val
	}
}


func WithExampleConfig(val *Config) ExampleOption {
	return func(o *ExampleOptions) {
		o.Config = val
	}
}

`,
		},
	}

	for _, tc := range tcs {
		o := NewOptionWriter(tc.name, tc.optionAttrs)
		generated := o.Generate()
		os.WriteFile(tc.name+".tmp", []byte(generated), 0644)
		require.Equal(t, tc.expected, generated)
	}
}
