package metricsgen

import (
	"fmt"
	"slices"
	"strings"

	"github.com/alexandreLamarre/metricsgen/pkg/templates"
	"github.com/alexandreLamarre/metricsgen/pkg/util"
	"github.com/samber/lo"
)

var (
	validAttributeTypes = []string{
		"int",
		"int64",
		"string",
		"float64",
		"bool",
		"[]int",
		"[]int64",
		"[]float64",
		"[]bool",
		"[]string",
	}

	validMetricTypes = []string{
		"int",
		"int64",
		"float",
		"float64",
	}
)

type Config struct {
	Attributes map[string]*Attribute `yaml:"attributes"`
	Metrics    map[string]*Metric    `yaml:"metrics"`
}

func (c *Config) Sanitize() error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}
	for aId, attr := range c.Attributes {
		attr.Name = aId
		c.Attributes[aId] = attr
	}

	for mId, metr := range c.Metrics {
		metr.Name = mId
		c.Metrics[mId] = metr
	}
	return nil
}

func (c *Config) Validate() error {
	if c == nil {
		return fmt.Errorf("config is nil")
	}
	if len(c.Attributes) == 0 && len(c.Metrics) == 0 {
		return fmt.Errorf("config must have at least one attribute or metric")
	}

	for _, attribute := range c.Attributes {
		if err := attribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Attribute struct {
	Name        string
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
}

func (a Attribute) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("attribute has an empty name")
	}
	if !slices.Contains(validAttributeTypes, a.Type) {
		return fmt.Errorf("invalid type : `%s`, must be one of %s", a.Type, strings.Join(validAttributeTypes, ","))
	}
	return nil
}

func (a Attribute) ToTemplateDefinition() templates.AttributeDef {
	return templates.AttributeDef{
		Name:        a.Name,
		Field:       util.OtelStringToCamelCaseField(a.Name),
		CamelCase:   util.OtelStringToCamelCase(a.Name),
		Constructor: util.ValueTypeToAttributeConstructor(a.Type),
		ValueType:   a.Type,
		Description: a.Description,
	}
}

func (a Attribute) ToDocsTemplateDefinition(required bool) templates.DocAttribute {
	return templates.DocAttribute{
		Name:        a.Name,
		Description: a.Description,
		ValueType:   a.Type,
		Required:    required,
	}
}

type Metric struct {
	Name               string
	Short              string `yaml:"short"`
	Long               string `yaml:"long"`
	Unit               string `yaml:"unit"`
	ValueType          string `yaml:"metric_type"`
	*MetricTypeCounter `yaml:"counter,omitempty"`
	*MetricTypeGauge   `yaml:"gauge,omitempty"`
	*MetricTypeHist    `yaml:"histogram,omitempty"`
	*MetricTypeExpHist `yaml:"exponential_histogram,omitempty"`
	Attributes         []string `yaml:"attributes"`
	OptionAttributes   []string `yaml:"optional_attributes"`
}

func (m Metric) Validate(attrTable map[string]Attribute) error {
	if m.Name == "" {
		return fmt.Errorf("metric has an empty name")
	}
	if !slices.Contains(validMetricTypes, m.ValueType) {
		return fmt.Errorf("invalid value type : `%s`, must be one of : %s", m.ValueType, strings.Join(validMetricTypes, ","))
	}
	count := 0
	if m.MetricTypeExpHist != nil {
		count += 1
	}
	if m.MetricTypeGauge != nil {
		count += 1
	}
	if m.MetricTypeCounter != nil {
		count += 1
	}
	if m.MetricTypeHist != nil {
		count += 1
	}
	if count == 0 {
		return fmt.Errorf("no metric types declared")
	}
	if count > 1 {
		return fmt.Errorf("multiple metric types declated for metric")
	}
	if util.HasDuplicateStrings(m.Attributes) {
		return fmt.Errorf("duplicate attribute registered to metric")
	}

	for _, attr := range m.Attributes {
		if _, ok := attrTable[attr]; !ok {
			return fmt.Errorf("no defined attribute for : `%s`", attr)
		}
	}

	for _, attr := range m.OptionAttributes {
		if _, ok := attrTable[attr]; !ok {
			return fmt.Errorf("no defined attribute for : `%s`", attr)
		}
	}

	return nil
}

func (m Metric) ToTemplateDefinition(attrTable map[string]*Attribute) templates.MetricConfig {
	requiredAttrs := AttributesForMetric(m.Attributes, attrTable)
	optionalAttrs := AttributesForMetric(m.OptionAttributes, attrTable)

	return templates.MetricConfig{
		Name:        m.Name,
		Description: m.Short,
		Units:       m.Unit,
		ValueType:   otelValueType(m.ValueType),
		Value:       goValueType(m.ValueType),
		MetricType:  m.metricValueType(),
		RequiredAttributes: lo.Map(requiredAttrs, func(a Attribute, _ int) templates.AttributeDef {
			return a.ToTemplateDefinition()
		}),
		OptionalAttributes: lo.Map(optionalAttrs, func(a Attribute, _ int) templates.AttributeDef {
			return a.ToTemplateDefinition()
		}),
	}
}

func (m Metric) ToDocsTemplateDefinition(attrTable map[string]*Attribute) templates.DocMetric {
	requireAttrs := AttributesForMetric(m.Attributes, attrTable)
	optAttrs := AttributesForMetric(m.OptionAttributes, attrTable)
	ret := []templates.DocAttribute{}
	for _, attr := range requireAttrs {
		ret = append(ret, attr.ToDocsTemplateDefinition(true))
	}
	for _, attr := range optAttrs {
		ret = append(ret, attr.ToDocsTemplateDefinition(false))
	}

	slices.SortFunc(ret, func(a, b templates.DocAttribute) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})

	return templates.DocMetric{
		Name:       m.Name,
		Link:       MarkdownLinkAnchor(m.Name),
		Short:      m.Short,
		Long:       m.Long,
		Unit:       m.Unit,
		MetricType: m.metricValueType(),
		ValueType:  m.ValueType,
		Attributes: ret,
	}
}

func (c *Config) ToTemplateDefinition() map[string]templates.MetricConfig {
	ret := map[string]templates.MetricConfig{}
	for _, m := range c.Metrics {
		structName := util.OtelStringToCamelCase(m.Name)
		ret[structName] = m.ToTemplateDefinition(c.Attributes)
	}
	return ret
}

func (c *Config) ToDocsTemplateDefinition() templates.DocConfig {
	ret := []templates.DocMetric{}
	for _, m := range c.Metrics {
		ret = append(ret, m.ToDocsTemplateDefinition(c.Attributes))
	}

	slices.SortFunc(ret, func(a, b templates.DocMetric) int {
		if a.Name < b.Name {
			return -1
		}
		if a.Name > b.Name {
			return 1
		}
		return 0
	})

	return templates.DocConfig{
		Metrics: ret,
	}
}

func (m Metric) metricValueType() string {
	if m.MetricTypeCounter != nil {
		return "Counter"
	}
	if m.MetricTypeGauge != nil {
		return "Gauge"
	}
	if m.MetricTypeHist != nil {
		return "Histogram"
	}
	if m.MetricTypeExpHist != nil {
		return "Histogram"
	}
	panic("unregisted metric type")
}

type MetricTypeCounter struct {
	ValueType   string `yaml:"value_type"`
	Aggregation string `yaml:"aggregation"`
}

type MetricTypeGauge struct {
	ValueType   string `yaml:"value_type"`
	Aggregation string `yaml:"aggregation"`
}

type MetricTypeHist struct {
	ValueType   string `yaml:"value_type"`
	Aggregation string `yaml:"aggregation"`
}

type MetricTypeExpHist struct {
}

func AttributesForMetric(attrs []string, attrTable map[string]*Attribute) []Attribute {
	ret := []Attribute{}
	for _, attr := range attrs {
		ret = append(ret, *attrTable[attr])
	}
	return ret
}

func goValueType(input string) string {
	if strings.HasPrefix(input, "int") {
		return "int64"
	}
	if strings.HasPrefix(input, "float") {
		return "float64"
	}
	panic(fmt.Sprintf("invalid input type : %s", input))
}

func otelValueType(input string) string {
	if strings.HasPrefix(input, "int") {
		return "Int64"
	}
	if strings.HasPrefix(input, "float") {
		return "Float64"
	}
	panic(fmt.Sprintf("invalid input type : %s", input))
}
