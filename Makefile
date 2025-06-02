PROFILE_FLAGS :=
ifeq ($(PROFILE),true)
	PROFILE_FLAGS := -cpuprofile=cpu.prof -memprofile=mem.prof
endif

BENCH_TARGET?=otel

build:
	go build -o ./bin/metricsgen cmd/metricsgen/main.go
generate: build
	cd examples && LOG_LEVEL=error go generate ./...
	cd tests && LOG_LEVEL=error go generate ./...
test:
	go test -timeout=30s ./...
	cd examples && go test -timeout=30s ./...
	cd tests && go test -timeout=30s ./...
benchmark: generate
	cd tests && go test -benchmem -benchtime=2s -bench=Benchmark $(PROFILE_FLAGS) -timeout=30s ./benchmark/$(BENCH_TARGET)
