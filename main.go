package main

import "github.com/alexandreLamarre/metricsgen/pkg/metricsgen"

func main() {
	if err := metricsgen.Gen(&metricsgen.Config{
		Attributes: map[string]metricsgen.Attribute{
			"pid": {
				Name:         "pid",
				Description:  "Process ID",
				Type:         "int",
				Required:     true,
				DefaultValue: "0",
			},
		},
		Metrics: map[string]metricsgen.Metric{
			"bpf.tcp.rx": {
				Name:          "bpf.tcp.rx",
				Description:   "TCP received bytes",
				Unit:          "bytes",
				ValueType:     "int64",
				MetricTypeSum: &metricsgen.MetricTypeSum{},
				Attributes:    []string{"pid"},
			},
			"bpf.tcp.tx": {
				Name:          "bpf.tcp.tx",
				Description:   "TCP received bytes",
				Unit:          "bytes",
				ValueType:     "int64",
				MetricTypeSum: &metricsgen.MetricTypeSum{},
				Attributes:    []string{"pid"},
			},
		},
	}); err != nil {
		panic(err)
	}
}
