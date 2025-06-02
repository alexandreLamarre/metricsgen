package util

import (
	"github.com/prometheus/otlptranslator"
)

type PrometheusNameGenerator struct {
	withoutCounterSuffixes bool
	namespace              string
	withoutUnits           bool

	translator *otlptranslator.MetricNamer
}

func (g *PrometheusNameGenerator) Apply(opts ...PrometheusNameOption) {
	for _, opt := range opts {
		opt(g)
	}
}

func NewPrometheusNameGenerator(opts ...PrometheusNameOption) *PrometheusNameGenerator {
	g := &PrometheusNameGenerator{
		namespace:    "",
		withoutUnits: false,
	}

	g.Apply(opts...)

	v := &otlptranslator.MetricNamer{
		Namespace:          g.namespace,
		WithMetricSuffixes: !g.withoutUnits,
		UTF8Allowed:        false,
	}
	g.translator = v
	return g
}

type PrometheusNameOption func(g *PrometheusNameGenerator)

func WithPrometheusNameNamespace(namespace string) PrometheusNameOption {
	return func(g *PrometheusNameGenerator) {
		g.namespace = namespace
	}
}

func WithPrometheusNameCounterSuffix(withoutCounterSuffixes bool) PrometheusNameOption {
	return func(g *PrometheusNameGenerator) {
		g.withoutCounterSuffixes = withoutCounterSuffixes
	}
}

func WithPrometheusNameWithoutUnits(withoutUnits bool) PrometheusNameOption {
	return func(g *PrometheusNameGenerator) {
		g.withoutUnits = withoutUnits
	}
}

func (g *PrometheusNameGenerator) GetPrometheusName(name, unit string, isCounter bool) string {
	return g.translator.Build(otlptranslator.Metric{
		Name: name,
		Unit: unit,
		Type: func() otlptranslator.MetricType {
			if isCounter {
				return otlptranslator.MetricTypeMonotonicCounter
			} else {
				return otlptranslator.MetricTypeGauge
			}
		}(),
	})
}

func GetPrometheusLabel(attrK string) string {
	return otlptranslator.NormalizeLabel(attrK)
}
