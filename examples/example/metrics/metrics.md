# Metrics
- [bpf.tcp.connlat](#bpftcpconnlat) : TCP connection latency ms
- [bpf.tcp.rx](#bpftcprx) : TCP received bytes
- [bpf.tcp.tx](#bpftcptx) : TCP transmitted bytes


## bpf.tcp.connlat

TCP connection latency ms



| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| ms | Histogram | float|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| pid | Process ID | int | ✅ |
| pid.gid | Process Group ID | int | ✅ |


## bpf.tcp.rx

TCP received bytes

collects from bpf tracepoints the total received bytes for the TCP protocol. Data is associated per pid, etc,etc

| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| bytes | Gauge | int64|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| pid | Process ID | int | ✅ |


## bpf.tcp.tx

TCP transmitted bytes



| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| bytes | Counter | int64|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| pid | Process ID | int | ✅ |
| pid.gid | Process Group ID | int | ✅ |
| cpu.id | cpu id in the range [0, numCPU] | int | ❌ |

