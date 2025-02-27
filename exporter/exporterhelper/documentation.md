[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# exporterhelper

## Internal Telemetry

The following telemetry is emitted by this component.

### otelcol_exporter_enqueue_failed_log_records

Number of log records failed to be added to the sending queue. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {records} | Sum | Int | true |

### otelcol_exporter_enqueue_failed_metric_points

Number of metric points failed to be added to the sending queue. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {datapoints} | Sum | Int | true |

### otelcol_exporter_enqueue_failed_spans

Number of spans failed to be added to the sending queue. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {spans} | Sum | Int | true |

### otelcol_exporter_queue_capacity

Fixed capacity of the retry queue (in batches) [alpha]

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {batches} | Gauge | Int |

### otelcol_exporter_queue_size

Current size of the retry queue (in batches) [alpha]

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| {batches} | Gauge | Int |

### otelcol_exporter_send_failed_log_records

Number of log records in failed attempts to send to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {records} | Sum | Int | true |

### otelcol_exporter_send_failed_metric_points

Number of metric points in failed attempts to send to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {datapoints} | Sum | Int | true |

### otelcol_exporter_send_failed_spans

Number of spans in failed attempts to send to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {spans} | Sum | Int | true |

### otelcol_exporter_sent_log_records

Number of log record successfully sent to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {records} | Sum | Int | true |

### otelcol_exporter_sent_metric_points

Number of metric points successfully sent to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {datapoints} | Sum | Int | true |

### otelcol_exporter_sent_spans

Number of spans successfully sent to destination. [alpha]

| Unit | Metric Type | Value Type | Monotonic |
| ---- | ----------- | ---------- | --------- |
| {spans} | Sum | Int | true |
