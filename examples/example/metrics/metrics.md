# Metrics
- [dummy.tcp.connlat](#dummytcpconnlat) : TCP connection latency ms
- [dummy.tcp.rx](#dummytcprx) : TCP received bytes
- [dummy.tcp.tx](#dummytcptx) : TCP transmitted bytes


## dummy.tcp.connlat

TCP connection latency ms

Randomly generated tcp connection latency

| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| ms | Histogram | float|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| cpu.id | cpu id in the range [0, numCPU] | int | ❌ |
| pid | Process ID | int | ✅ |
| pid.gid | Process Group ID | int | ✅ |


## dummy.tcp.rx

TCP received bytes

Randomly generated tcp received bytes

| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| bytes | Gauge | int64|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| pid | Process ID | int | ✅ |


## dummy.tcp.tx

TCP transmitted bytes

Randomly generated tcp transmitted bytes

| Unit | Metric Type | ValueType |
| ---- | ------------ | --------- |
| bytes | Counter | int64|

### Attributes

| Name | Description | Type | Required |
|------|-------------|------| ------- |
| cpu.id | cpu id in the range [0, numCPU] | int | ✅ |
| pid | Process ID | int | ✅ |
| pid.gid | Process Group ID | int | ✅ |

