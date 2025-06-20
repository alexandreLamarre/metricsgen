# Metrics
- [example.counter](#examplecounter) : Example Counter
- [example.counter.optional](#examplecounteroptional) : Example Counter
- [example.exponential_histogram](#exampleexponentialhistogram) : Example Exponential Histogram
- [example.gauge](#examplegauge) : Example Gauge
- [example.gauge.optional](#examplegaugeoptional) : Example Gauge
- [example.histogram](#examplehistogram) : Example Histogram
- [example.histogram.optional](#examplehistogramoptional) : Example Histogram


## example.counter

Example Counter

Some extra details about what the Example Counter represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_counter_unit_total | unit | Counter | int|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.enum | example_enum | Example enum | string | ✅ |
| example.enum2 | example_enum2 | Example enum 2 | string | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |


## example.counter.optional

Example Counter

Some extra details about what the Example Counter represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_counter_optional_unit_total | unit | Counter | int|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ❌ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ❌ |
| example.float | example_float | Example float value | float64 | ❌ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ❌ |
| example.int | example_int | Example int value | int | ❌ |
| example.int64 | example_int64 | Example int64 value | int64 | ❌ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ❌ |
| example.intSlice | example_intSlice | Example int slice value | []int | ❌ |
| example.string | example_string | Example string value | string | ❌ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ❌ |


## example.exponential_histogram

Example Exponential Histogram



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_exponential_histogram_milliseconds | ms | Histogram | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |


## example.gauge

Example Gauge

Some extra details about what the Example Gauge represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_gauge_unit | unit | Gauge | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |


## example.gauge.optional

Example Gauge

Some extra details about what the Example Gauge represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_gauge_optional_unit | unit | Gauge | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |


## example.histogram

Example Histogram

Some extra details about what the Example Histogram represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_histogram_milliseconds | ms | Histogram | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |


## example.histogram.optional

Example Histogram

Some extra details about what the Example Histogram represents and how it is collected.

| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| example_histogram_optional_milliseconds | ms | Histogram | float|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| example.bool | example_bool | Example boolean value | bool | ✅ |
| example.boolSlice | example_boolSlice | Example bool slice value | []bool | ✅ |
| example.float | example_float | Example float value | float64 | ✅ |
| example.floatSlice | example_floatSlice | Example float slice value | []float64 | ✅ |
| example.int | example_int | Example int value | int | ✅ |
| example.int64 | example_int64 | Example int64 value | int64 | ✅ |
| example.int64Slice | example_int64Slice | Example int64 slice value | []int64 | ✅ |
| example.intSlice | example_intSlice | Example int slice value | []int | ✅ |
| example.string | example_string | Example string value | string | ✅ |
| example.stringSlice | example_stringSlice | Example int slice value | []string | ✅ |

