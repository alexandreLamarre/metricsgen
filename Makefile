build:
	go build -o ./bin/metricsgen main.go
generate: build
	LOG_LEVEL=error go generate ./...
benchmark: generate
	go test -benchmem -benchtime=2s -bench=Benchmark ./examples/benchmark/
