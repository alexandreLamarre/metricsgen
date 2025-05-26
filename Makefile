PROFILE_FLAGS :=
ifeq ($(PROFILE),true)
	PROFILE_FLAGS := -cpuprofile=cpu.prof -memprofile=mem.prof
endif

BENCH_TARGET?=otel

build:
	go build -o ./bin/metricsgen main.go
generate: build
	LOG_LEVEL=error go generate ./...
benchmark: generate
	go test -benchmem -benchtime=2s -bench=Benchmark $(PROFILE_FLAGS) -timeout=30s ./examples/benchmark/$(BENCH_TARGET)
