package metricsgen

import (
	"errors"
	"fmt"
	"log/slog"
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
	logger     *slog.Logger
}

func (c *Config) SetLogger(l *slog.Logger) {
	c.logger = l
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
	if c.logger == nil {
		return fmt.Errorf("internal error: no logger defined")
	}
	logger := c.logger
	if c == nil {
		logger.Error("config is nil")
		return fmt.Errorf("config is nil")
	}
	if len(c.Attributes) == 0 && len(c.Metrics) == 0 {
		c.logger.Error("config must have at least one attribute or metric")
		return fmt.Errorf("config must have at least one attribute or metric")
	}
	var errs []error
	for _, attribute := range c.Attributes {
		if err := attribute.Validate(logger); err != nil {
			c.logger.With("attribute", attribute.Name).Error(err.Error())
			errs = append(errs, err)
		}
	}

	for _, m := range c.Metrics {
		if err := m.Validate(logger, c.Attributes); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

type Attribute struct {
	Name        string
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Enum        []string `yaml:"enum,omitempty"`
}

var ErrInvalidAttribute = errors.New("invalid attribute")

func (a Attribute) Validate(l *slog.Logger) error {
	if a.Name == "" {
		return fmt.Errorf("attribute has an empty name")
	}
	logger := l.With("attribute", a.Name)
	invalid := false
	if !slices.Contains(validAttributeTypes, a.Type) {
		invalid = true
		logger.With("type", a.Type).Error(
			fmt.Sprintf("invalid type, must be one of %s", strings.Join(validAttributeTypes, ",")),
		)
	}

	if len(a.Enum) > 0 {
		validEnumTypes := []string{"string", "int"}
		if !slices.Contains(validEnumTypes, a.Type) {
			invalid = true
			logger.With("type", a.Type).Error(
				fmt.Sprintf("enum not supported for type, must be : %s", strings.Join(validEnumTypes, ",")),
			)
		}
	}
	if invalid {
		return ErrInvalidAttribute
	}
	return nil
}

func (a Attribute) ToTemplateDefinition() templates.AttributeDef {
	return templates.AttributeDef{
		Name:        a.Name,
		Field:       util.OtelStringToCamelCaseField(a.Name),
		CamelCase:   util.OtelStringToCamelCase(a.Name),
		Constructor: util.ValueTypeToAttributeConstructor(a.Type),
		ValueType:   attributeValueType(a),
		Description: a.Description,
		Enum:        len(a.Enum) > 0,
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
	*MetricTypeCounter `yaml:"counter,omitempty"`
	*MetricTypeGauge   `yaml:"gauge,omitempty"`
	*MetricTypeHist    `yaml:"histogram,omitempty"`
	Attributes         []string `yaml:"attributes"`
	OptionAttributes   []string `yaml:"optional_attributes"`
}

func (m Metric) getValueType() string {
	if m.MetricTypeCounter != nil {
		return m.MetricTypeCounter.ValueType
	}
	if m.MetricTypeGauge != nil {
		return m.MetricTypeGauge.ValueType
	}
	if m.MetricTypeHist != nil {
		return m.MetricTypeHist.ValueType
	}
	return ""
}

var ErrInvalidMetric = errors.New("invalid metric")

func (m Metric) Validate(l *slog.Logger, attrTable map[string]*Attribute) error {
	if m.Name == "" {
		return fmt.Errorf("metric has an empty name")
	}
	invalid := false
	logger := l.With("metric", m.Name)
	count := 0
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
		invalid = true
		logger.Error("no metric types declared")
	}
	if count > 1 {
		invalid = true
		logger.Error("multiple metric types declated for metric")
	}

	if err := m.checkMetricValueType(); err != nil {
		invalid = true
		logger.Error("metrics must have `gauge`,`counter` or `histogram` specs defined")
	}

	if !slices.Contains(validMetricTypes, m.getValueType()) {
		logger.Error(
			fmt.Sprintf("invalid value type : `%s`, must be one of : %s", m.getValueType(), strings.Join(validMetricTypes, ",")),
		)
		invalid = true
	}

	if util.HasDuplicateStrings(m.Attributes) {
		logger.Error("duplicate attribute registered to metric")
		invalid = true
	}

	for _, attr := range m.Attributes {
		if _, ok := attrTable[attr]; !ok {
			invalid = true
			logger.With("attribute", attr).Error("no attribute definition")
		}
	}

	for _, attr := range m.OptionAttributes {
		if _, ok := attrTable[attr]; !ok {
			invalid = true
			logger.With("attribute", attr).Error("no matching optional attribute")
		}
	}

	attrs := lo.Intersect(m.Attributes, m.OptionAttributes)
	if len(attrs) > 0 {
		invalid = true
		for _, attr := range attrs {
			logger.With("attribute", attr).Error("attribute defined as both required and optional")
		}
	}

	if invalid {
		return ErrInvalidMetric
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
		ValueType:   otelValueType(m.getValueType()),
		Value:       goValueType(m.getValueType()),
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
		ValueType:  m.getValueType(),
		Attributes: ret,
	}
}

func (c *Config) ToMetricsTemplateDefinition() map[string]templates.MetricConfig {
	ret := map[string]templates.MetricConfig{}
	for _, m := range c.Metrics {
		structName := util.OtelStringToCamelCase(m.Name)
		ret[structName] = m.ToTemplateDefinition(c.Attributes)
	}
	return ret
}

func attributeValueType(a Attribute) string {
	if len(a.Enum) > 0 {
		return "Enum" + util.OtelStringToCamelCase(a.Name)
	}
	return a.Type
}

func (c *Config) ToEnumTemplateDefinition() []templates.EnumConfig {
	ret := []templates.EnumConfig{}

	for _, attr := range c.Attributes {
		if len(attr.Enum) == 0 {
			continue
		}
		cc := util.OtelStringToCamelCase(attr.Name)
		t := templates.EnumConfig{
			EnumType:  attributeValueType(*attr),
			ValueType: attr.Type,
			CamelCase: cc,
			Values:    []templates.EnumValue{},
		}

		for idx, val := range attr.Enum {
			if attr.Type == "string" {
				t.Values = append(t.Values, templates.EnumValue{
					ValueCase: util.OtelStringToCamelCase(val),
					Value:     fmt.Sprintf(`"%s"`, val),
				})
			}
			if attr.Type == "int" {
				t.Values = append(t.Values, templates.EnumValue{
					ValueCase: util.OtelStringToCamelCase(val),
					Value:     idx,
				})
			}
		}
		ret = append(ret, t)
	}

	slices.SortFunc(ret, func(a, b templates.EnumConfig) int {
		if a.EnumType < b.EnumType {
			return -1
		}
		if a.EnumType > b.EnumType {
			return 1
		}
		return 0
	})
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

func (m Metric) checkMetricValueType() error {
	if m.MetricTypeCounter != nil {
		return nil
	}
	if m.MetricTypeGauge != nil {
		return nil
	}
	if m.MetricTypeHist != nil {
		return nil
	}
	return fmt.Errorf("unknown metric type")
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
	panic("unregisted metric type")
}

type MetricTypeCounter struct {
	ValueType string `yaml:"value_type"`
}

type MetricTypeGauge struct {
	ValueType string `yaml:"value_type"`
}

type MetricTypeHist struct {
	ValueType string `yaml:"value_type"`
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
