# Metrics
- [eight.label.counter](#eightlabelcounter) : counter with 8 labels
- [four.label.counter](#fourlabelcounter) : counter with 4 labels
- [no.label.counter](#nolabelcounter) : counter with no labels
- [one.label.counter](#onelabelcounter) : counter with 1 label
- [split.label.counter](#splitlabelcounter) : counter with 4 labels and 4 optional labels


## eight.label.counter

counter with 8 labels



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| eight_label_counter_total |  | Counter | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| stringLabel | stringLabel | A string label | string | ✅ |
| stringLabel2 | stringLabel2 | A string label | string | ✅ |
| stringLabel3 | stringLabel3 | A string label | string | ✅ |
| stringLabel4 | stringLabel4 | A string label | string | ✅ |
| stringLabel5 | stringLabel5 | A string label | string | ✅ |
| stringLabel6 | stringLabel6 | A string label | string | ✅ |
| stringLabel7 | stringLabel7 | A string label | string | ✅ |
| stringLabel8 | stringLabel8 | A string label | string | ✅ |


## four.label.counter

counter with 4 labels



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| four_label_counter_total |  | Counter | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| stringLabel | stringLabel | A string label | string | ✅ |
| stringLabel2 | stringLabel2 | A string label | string | ✅ |
| stringLabel3 | stringLabel3 | A string label | string | ✅ |
| stringLabel4 | stringLabel4 | A string label | string | ✅ |


## no.label.counter

counter with no labels



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| no_label_counter_total |  | Counter | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |


## one.label.counter

counter with 1 label



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| one_label_counter_total |  | Counter | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| stringLabel | stringLabel | A string label | string | ✅ |


## split.label.counter

counter with 4 labels and 4 optional labels



| Prometheus name | Unit | Metric Type | ValueType |
| --------------- |  ---- | ------------ | --------- |
| split_label_counter_total |  | Counter | int64|

### Attributes

| Name | Prometheus label | Description | Type | Required |
|------| ---------------- |-------------|------| ------- |
| stringLabel | stringLabel | A string label | string | ✅ |
| stringLabel2 | stringLabel2 | A string label | string | ✅ |
| stringLabel3 | stringLabel3 | A string label | string | ✅ |
| stringLabel4 | stringLabel4 | A string label | string | ✅ |
| stringLabel5 | stringLabel5 | A string label | string | ❌ |
| stringLabel6 | stringLabel6 | A string label | string | ❌ |
| stringLabel7 | stringLabel7 | A string label | string | ❌ |
| stringLabel8 | stringLabel8 | A string label | string | ❌ |

