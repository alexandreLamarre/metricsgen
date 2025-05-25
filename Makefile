PROFILE_FLAGS :=
ifeq ($(PROFILE),true)
	PROFILE_FLAGS := -cpuprofile=cpu.prof -memprofile=mem.prof
endif

build:
	go build -o ./bin/metricsgen main.go
generate: build
	LOG_LEVEL=error go generate ./...
benchmark: generate
	go test -benchmem -benchtime=2s -bench=Benchmark $(PROFILE_FLAGS) ./examples/benchmark/
