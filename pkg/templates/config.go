package templates

import (
	_ "embed"
	"text/template"
)

//go:embed templates/metrics.go.tmpl
var metricsGenRawTemplate string

//go:embed templates/docs.md.tmpl
var docsGenRawTemplate string

var metricsGenTemplate = template.Must(template.New("metricsgen").Parse(
	metricsGenRawTemplate,
))

var docsGenTemplate = template.Must(template.New("docsgen").Parse(
	docsGenRawTemplate,
))

type ImportDef struct {
	Alias      string
	Dependency string
}

type AttributeDef struct {
	ValueType   string
	Field       string
	Name        string
	Constructor string
	CamelCase   string
	Description string
	Enum        bool
}

type MetricConfig struct {
	Name        string
	Description string
	Units       string
	//ValueType, one of "Int64" / "Float64"
	ValueType          string
	MetricType         string
	Value              string
	RequiredAttributes []AttributeDef
	OptionalAttributes []AttributeDef
	Buckets            []float64
}

type EnumConfig struct {
	EnumType    string
	Description string
	ValueType   string
	CamelCase   string
	// set only if ValueType="string"
	Values []EnumValue
}

type EnumValue struct {
	ValueCase string
	Value     any
}

type GenConfig struct {
	PackageName string
	ImportDefs  []ImportDef
	Metrics     map[string]MetricConfig
	EnumTypes   []EnumConfig
}

type DocConfig struct {
	Metrics []DocMetric
}

type DocMetric struct {
	Name           string
	PrometheusName string
	Link           string
	Short          string
	Long           string
	Unit           string
	MetricType     string
	ValueType      string
	Attributes     []DocAttribute
}

type DocAttribute struct {
	Name            string
	PrometheusLabel string
	Description     string
	ValueType       string
	Required        bool
}
