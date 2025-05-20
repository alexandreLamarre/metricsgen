# metricsgen

![Build](https://github.com/alexandreLamarre/metricsgen/actions/workflows/ci.yaml/badge.svg)

A go generate tool for generating type safe metric instrumentation scaffholding and matching metric documentation.

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

## Coming Soon

- Generated prometheus code utils for type-safe [Perses](https://perses.dev/) dashboards.
- Some form of importing and merging attribute files
- Customizing some cli output options
- Maybe some form of importing and merging metric files
