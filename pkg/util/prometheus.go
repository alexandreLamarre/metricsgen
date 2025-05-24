package util

import (
	"strings"

	"github.com/prometheus/common/model"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var unitSuffixes = map[string]string{
	// Time
	"d":   "days",
	"h":   "hours",
	"min": "minutes",
	"s":   "seconds",
	"ms":  "milliseconds",
	"us":  "microseconds",
	"ns":  "nanoseconds",

	// Bytes
	"By":   "bytes",
	"KiBy": "kibibytes",
	"MiBy": "mebibytes",
	"GiBy": "gibibytes",
	"TiBy": "tibibytes",
	"KBy":  "kilobytes",
	"MBy":  "megabytes",
	"GBy":  "gigabytes",
	"TBy":  "terabytes",

	// SI
	"m": "meters",
	"V": "volts",
	"A": "amperes",
	"J": "joules",
	"W": "watts",
	"g": "grams",

	// Misc
	"Cel": "celsius",
	"Hz":  "hertz",
	"1":   "ratio",
	"%":   "percent",
}

// prometheus counters MUST have a _total suffix by default:
// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.20.0/specification/compatibility/prometheus_and_openmetrics.md
const counterSuffix = "total"

// convertsToUnderscore returns true if the character would be converted to an
// underscore when the escaping scheme is underscore escaping. This is meant to
// capture any character that should be considered a "delimiter".
// ref :https://github.com/open-telemetry/opentelemetry-go/blob/e57879908fbd71f0fe64c4905c579d5827a34779/exporters/prometheus/exporter.go#L521-L526
func ConvertsToUnderscore(b rune) bool {
	return (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') && b != ':' && (b < '0' || b > '9')
}

type PrometheusNameGenerator struct {
	withoutCounterSuffixes bool
	namespace              string
	withoutUnits           bool
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

func (g *PrometheusNameGenerator) GetPrometheusName(m metricdata.Metrics, isCounter bool) string {
	name := m.Name
	name = model.EscapeName(name, model.NameEscapingScheme)
	addCounterSuffix := !g.withoutCounterSuffixes && isCounter
	if addCounterSuffix {
		// Remove the _total suffix here, as we will re-add the total suffix
		// later, and it needs to come after the unit suffix.
		name = strings.TrimSuffix(name, counterSuffix)
		// If the last character is an underscore, or would be converted to an underscore, trim it from the name.
		// an underscore will be added back in later.
		if ConvertsToUnderscore(rune(name[len(name)-1])) {
			name = name[:len(name)-1]
		}
	}
	if g.namespace != "" {
		name = g.namespace + name
	}
	if suffix, ok := unitSuffixes[m.Unit]; ok && !g.withoutUnits && !strings.HasSuffix(name, suffix) {
		name += "_" + suffix
	}
	if addCounterSuffix {
		name += "_" + counterSuffix
	}
	return name
}

func GetPrometheusLabel(attrK string) string {
	return model.EscapeName(attrK, model.NameEscapingScheme)
}
