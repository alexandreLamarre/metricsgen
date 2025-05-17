package main

import "github.com/alexandreLamarre/metricsgen/pkg/metricsgen"

func main() {
	cfg := &metricsgen.Config{
		Attributes: map[string]metricsgen.Attribute{
			"pid": {
				Name:        "pid",
				Description: "Process ID",
				Type:        "int",
				Required:    true,
			},
			"pid.gid": {
				Name:        "pid.gid",
				Description: "Process Group ID",
				Type:        "int",
				Required:    true,
			},
		},
		Metrics: map[string]metricsgen.Metric{
			"bpf.tcp.rx": {
				Name:          "bpf.tcp.rx",
				Short:         "TCP received bytes",
				Long:          "collects from bpf tracepoints the total received bytes for the TCP protocol. Data is associated per pid, etc,etc",
				Unit:          "bytes",
				ValueType:     "int64",
				MetricTypeSum: &metricsgen.MetricTypeSum{},
				Attributes:    []string{"pid"},
			},
			"bpf.tcp.tx": {
				Name:          "bpf.tcp.tx",
				Short:         "TCP transmitted bytes",
				Unit:          "bytes",
				ValueType:     "int64",
				MetricTypeSum: &metricsgen.MetricTypeSum{},
				Attributes:    []string{"pid", "pid.gid"},
			},
		},
	}
	if err := cfg.Gen(); err != nil {
		panic(err)
	}
}
