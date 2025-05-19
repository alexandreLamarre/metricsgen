package main

import (
	"bytes"
	_ "embed"
	"os"
)

func main() {
	var b bytes.Buffer

	if err := metricsGenTemplate.Execute(&b, GenConfig{
		PackageName: "main",
		ImportDefs: []ImportDef{
			{
				Dependency: "context",
			},
			{
				Alias:      "otelmetricsdk",
				Dependency: "go.opentelemetry.io/otel/metric",
			},
			{
				Alias:      "otelattribute",
				Dependency: "go.opentelemetry.io/otel/attribute",
			},
		},
		Metrics: map[string]MetricConfig{
			"BpfTcpRx": {
				Name:        "bpf.tcp.rx",
				Description: "TCP receive bytes",
				Units:       "bytes",
				ValueType:   "Int64",
				MetricType:  "Counter",
				Value:       "int64",
				RequiredAttributes: []AttributeDef{
					{
						Name:        "pid",
						Field:       "pid",
						ValueType:   "string",
						Constructor: "String",
						CamelCase:   "Pid",
						Description: "process ID",
					},
				},
				OptionalAttributes: []AttributeDef{
					{
						Name:        "pid.namespace",
						Field:       "pidNamespace",
						ValueType:   "string",
						Constructor: "String",
						CamelCase:   "PidNamespace",
						Description: "process Namespace",
					},
				},
			},
			"BpfTcpTx": {
				Name:        "bpf.tcp.tx",
				Description: "TCP transmit bytes",
				Units:       "bytes",
				ValueType:   "Int64",
				MetricType:  "Gauge",
				Value:       "int64",
				RequiredAttributes: []AttributeDef{
					{
						Name:        "pid",
						Field:       "pid",
						ValueType:   "string",
						Constructor: "String",
						CamelCase:   "Pid",
						Description: "Process id",
					},
					{
						Name:        "pid.gid",
						Field:       "pidGid",
						ValueType:   "string",
						Constructor: "String",
						CamelCase:   "PidGid",
						Description: "Process group thread ID",
					},
					{
						Name:        "cpu.id",
						Field:       "cpuId",
						ValueType:   "int",
						Constructor: "Int",
						CamelCase:   "CpuId",
						Description: "Cpu Identifier in the range [0, numCpus]",
					},
				},
			},
		},
	}); err != nil {
		panic(err)
	}

	if err := os.WriteFile("metricsgen2.go", b.Bytes(), 0644); err != nil {
		panic(err)
	}

	var b2 bytes.Buffer

	if err := docsGenTemplate.Execute(&b2, DocConfig{
		Metrics: []DocMetric{
			{
				Name:       "bpf.tcp.rx",
				Short:      "TCP received bytes",
				Long:       "collects from bpf tracepoints the total received bytes for the TCP protocol. Data is associated per pid, etc,etc",
				Link:       "bpftcprx",
				Unit:       "bytes",
				MetricType: "Gauge",
				ValueType:  "int64",
				Attributes: []DocAttribute{
					{
						Name:        "pid",
						Description: "process id",
						ValueType:   "int",
						Required:    true,
					},
				},
			},
		},
	}); err != nil {
		panic(err)
	}

	if err := os.WriteFile("metrics.md", b2.Bytes(), 0644); err != nil {
		panic(err)
	}
}
