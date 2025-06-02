# metricsgen

![Build](https://github.com/alexandreLamarre/metricsgen/actions/workflows/ci.yaml/badge.svg)

A go generate tool for generating type safe metric instrumentation scaffholding and matching metric documentation.

## Supported drivers

Can be set via the `--driver` flag

- `prometheus`: instrumentation using `github.com/prometheus/client_golang`
- `otel` : instrumentation using `go.opentelemetry.io/otel`

:warning: Generated prometheus doc names may not be consistent with otel sdk's `go.opentelemetry.io/otel/exporters/prometheus`, see https://github.com/open-telemetry/opentelemetry-go/issues/6704

## Install

Two installation methods are supported:

- Install from github releases 
- using go's 1.24 tool directive.

using tool directive:

```go.mod
require(
    github.com/alexandreLamarre/metricsgen v0.4.0
)

tool github.com/alexandreLamarre/metricsgen
```

```sh
go install tool
```

## Example Usage

```go
//go:generate metricsgen metrics.yaml
package example
```

`metrics.yaml`
```yaml
attributes:
    label:
        description: sample opaque label
        type: string
    state:
        description: state of something
        type: string
        enum : [on,off]
    label.optional:
        type : string
        description: sample optional opaque label
metrics:
    example.measurement:
        metric_type: float
        counter:
        attributes: [label, state]
        optional_attributes : [label.optional]
```

Which allows you to create type-safe metrics instrumentation:

```go
meter := otel.Meter("example-metrics")
metrics, _ := example.NewMetrics(meter)
// records a value for the metric `example.measurement`
metrics.ExampleMeasurement.Record(
    context.TODO(), 
    0.1, 
    // sets label=labelA
    "labelA",
    // sets state="on"
    example.EnumStateOn,
    // sets label.optional=labelB
    metrics.WithExampleMeasurementLabelOptional("labelB"),
)
```

Check `examples` folder for more comprehensive examples

## Attributes

Attributes are defined as follows:
```yaml
attributes:
    # this is the attribute ID
    your.attribute:
        # required must be a valid otel attribute type, see examples
        type: string
        # optional description for documentation
        description: Some information here
        # optional : generate type safe enums for recording data
        enum : [A,B,C]
```

## Metrics

Metrics are defined as follows:
```yaml
metrics:
    # this is the metric name, following naming otel convetions
    your.metric:
      # optional unit
      unit : By
      # optional short description
      short: Some metric
      # optional long description
      long: "Collected using advanced system info, only available on the following linux distros: Ubuntu"
      # optional : list of required attributes when recording a metric
      attributes:
        - foo
      # optional : list of optional attributes when recording a metric
      optional_attributes:
        - system.specific.label
```

You must instantiate a specific metric type, the currently supported ones are  `counter`, `gauge` and `histogram`:
```yaml
metrics:
    your.metric:
        counter:
            # required, can only be int,float,int64,float64
            value_type: int
```

histograms come with default bucket boundaries, but you can define your own like so:
```yaml
metrics:
    your.histogram.metric:
        histogram:
            # required, can only be int,float,int64,float64
            value_type: int
            buckets: [0.0, 1.2, 5.6, 1000, 100000.7]
```

## Coming Soon

- Generated prometheus code utils for type-safe [Perses](https://perses.dev/) dashboards.
- Customizing some cli output options
- Type-safe garbage-collected metrics / type-safe observable implementation
- (Maybe) pdata generation for use directly in open-telemetry collector