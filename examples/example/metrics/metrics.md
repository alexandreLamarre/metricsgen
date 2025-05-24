# Metrics
- [dummy.tcp.connlat](#dummytcpconnlat) : TCP connection latency ms
- [dummy.tcp.rx](#dummytcprx) : TCP received bytes
- [dummy.tcp.tx](#dummytcptx) : TCP transmitted bytes


## dummy.tcp.connlat

TCP connection latency ms

Randomly generated tcp connection latency

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| dummy_tcp_connlat_milliseconds | ms | Histogram | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| cpu.id | cpu_id | cpu id in the range [0, numCPU] | int | ❌ |
| pid | pid | Process ID | int | ✅ |
| pid.gid | pid_gid | Process Group ID | int | ✅ |
| random.int | random_int | random enum int | int | ❌ |


## dummy.tcp.rx

TCP received bytes

Randomly generated tcp received bytes

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| dummy_tcp_rx_bytes | By | Gauge | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| cpu.mode | cpu_mode | cpu state | string | ❌ |
| pid | pid | Process ID | int | ✅ |
| random.int | random_int | random enum int | int | ✅ |


## dummy.tcp.tx

TCP transmitted bytes

Randomly generated tcp transmitted bytes

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| dummy_tcp_tx_bytes_total | By | Counter | int|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| cpu.id | cpu_id | cpu id in the range [0, numCPU] | int | ✅ |
| cpu.mode | cpu_mode | cpu state | string | ✅ |
| pid | pid | Process ID | int | ✅ |
| pid.gid | pid_gid | Process Group ID | int | ✅ |

