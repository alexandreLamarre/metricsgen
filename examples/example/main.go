package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/alexandreLamarre/metricsgen/examples/example/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize prometheus exporter: %v", err)
	}
	// Set global meter provider
	provider := metricsdk.NewMeterProvider(metricsdk.WithReader(exporter))
	otel.SetMeterProvider(provider)

	// Get a meter
	meter := otel.Meter("example")

	m, err := metrics.NewMetrics(meter)
	if err != nil {
		log.Fatal(err)
	}

	go recordDummyMetrics(context.Background(), m)

	// Expose metrics on /metrics
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Serving metrics at :2222/metrics")
	log.Fatal(http.ListenAndServe(":2222", nil))
}

func recordDummyMetrics(ctx context.Context, m metrics.Metrics) {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Println("exiting recording dummy metrics")
		case <-t.C:
			for x := range 16 {
				m.MetricDummyTcpConnlat.Record(
					ctx,
					rand.Float64(),
					rand.Int(),
					rand.Int(),
					metrics.WithDummyTcpConnlatCpuId(x),
				)

				m.MetricDummyTcpRx.Record(
					ctx,
					rand.Int63(),
					rand.Int(),
				)

				m.MetricDummyTcpTx.Record(
					ctx,
					rand.Int63(),
					rand.Int(),
					rand.Int(),
					x,
				)
			}
		}
	}
}
