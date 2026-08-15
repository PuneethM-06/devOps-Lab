# Interview.md — Day 80

## PromQL — `rate()`, `histogram_quantile()`, and Prometheus Alerting

---

## 1. What is `rate()` in PromQL?

`rate()` calculates the average per-second rate of increase of a metric over a specified time range.

It is mainly used with Counters because Counters are cumulative metrics that generally increase over time.

Example:

```promql
rate(http_requests_total[5m])
```

This calculates the average number of requests per second over the last 5 minutes.

---

## 2. Why does `rate()` require a time range?

`rate()` needs multiple metric samples over time to calculate how much the Counter has increased.

Example:

```promql
rate(http_requests_total[5m])
```

Prometheus compares the Counter values within the last 5 minutes and calculates the average per-second rate.

Without a range:

```promql
rate(http_requests_total)
```

Prometheus does not have a range of historical samples to calculate the rate.

---

## 3. Why is `rate()` mainly used with Counters?

Counters represent cumulative values that increase over time.

Examples include:

- Total HTTP requests
- Total errors
- Total database queries

`rate()` helps calculate how quickly these Counters are increasing.

Example:

```promql
rate(http_requests_total[5m])
```

A Gauge can increase and decrease, so querying a Gauge directly usually gives the current/latest scraped value.

Example:

```promql
node_memory_MemAvailable_bytes
```

---

## 4. What is the difference between `rate()` and `irate()`?

### `rate()`

```promql
rate(http_requests_total[5m])
```

Calculates the average per-second rate over the selected time range.

It is smoother and less sensitive to short spikes.

### `irate()`

```promql
irate(http_requests_total[5m])
```

Uses the two most recent samples within the selected time range.

It reacts more quickly to recent changes but is more sensitive and noisy.

### Simple Difference

> `rate()` = average rate over a time range.

> `irate()` = recent rate based on the two latest samples.

---

## 5. What does `sum(rate(...))` do?

Example:

```promql
sum(rate(http_requests_total[5m]))
```

`rate()` first calculates the per-second rate for each matching time series.

Then `sum()` adds all those rates together.

Example:

```text
Instance A → 2 requests/sec
Instance B → 3 requests/sec
Instance C → 5 requests/sec
```

Result:

```text
Total → 10 requests/sec
```

---

## 6. What does `sum by (job)` do?

Example:

```promql
sum by (job) (
  rate(http_requests_total[5m])
)
```

`rate()` calculates the per-second rate for each time series.

`sum by (job)` then groups time series with the same `job` label and adds them together.

Example:

```text
frontend, instance=1 → 2 req/sec
frontend, instance=2 → 3 req/sec
backend, instance=1  → 5 req/sec
backend, instance=2  → 4 req/sec
```

Result:

```text
frontend → 5 req/sec
backend  → 9 req/sec
```

Different jobs remain separate.

---

## 7. How do you calculate HTTP 5xx error percentage?

```promql
(
  sum(rate(http_requests_total{status=~"5.."}[5m]))
  /
  sum(rate(http_requests_total[5m]))
) * 100
```

### Top Part

```promql
sum(rate(http_requests_total{status=~"5.."}[5m]))
```

Calculates the total rate of 5xx error requests.

The regex:

```text
5..
```

matches status codes such as:

```text
500
501
502
503
504
```

### Bottom Part

```promql
sum(rate(http_requests_total[5m]))
```

Calculates the total request rate regardless of status code.

### Final Calculation

```text
5xx request rate
-----------------
total request rate
```

Multiplying by `100` converts the result into a percentage.

---

# Histograms and `histogram_quantile()`

## 8. What is a Histogram?

A Histogram is used to observe the distribution of values.

Examples include:

- HTTP request duration
- Response size
- Database query duration

A Histogram exposes metrics such as:

```text
http_request_duration_seconds_bucket
http_request_duration_seconds_sum
http_request_duration_seconds_count
```

---

## 9. What does the `le` label mean?

`le` means:

> Less than or equal to.

Example:

```text
le="0.5"
```

means requests that took:

```text
0.5 seconds or less
```

Example buckets:

```text
le="0.1"  → requests ≤ 0.1s
le="0.5"  → requests ≤ 0.5s
le="1"    → requests ≤ 1s
le="2"    → requests ≤ 2s
le="+Inf" → all requests
```

---

## 10. Why are Histogram buckets cumulative?

Histogram bucket values accumulate.

Example:

```text
le="0.5" → 50
le="1"   → 80
```

This means:

```text
50 requests took ≤ 0.5 seconds
80 requests took ≤ 1 second
```

Requests between `0.5` and `1` second:

```text
80 - 50 = 30 requests
```

---

## 11. Why do we use `rate()` with Histogram buckets?

Histogram bucket metrics are Counters.

Example:

```promql
rate(http_request_duration_seconds_bucket[5m])
```

This calculates how quickly each bucket is increasing per second over the last 5 minutes.

This gives us the recent distribution of request durations instead of the total distribution since the application started.

---

## 12. Why do we use `sum by (le)`?

Example:

```promql
sum by (le) (
  rate(http_request_duration_seconds_bucket[5m])
)
```

Suppose multiple pods expose the same bucket:

```text
pod-1, le="0.5" → 20 req/sec
pod-2, le="0.5" → 30 req/sec
```

The result becomes:

```text
le="0.5" → 50 req/sec
```

`sum by (le)` combines matching buckets while preserving the `le` label.

This is important because `histogram_quantile()` needs the bucket boundaries.

---

## 13. Why is plain `sum()` wrong for Histogram Quantiles?

This:

```promql
sum(
  rate(http_request_duration_seconds_bucket[5m])
)
```

removes the `le` label.

That means Prometheus loses information about which values belong to:

```text
≤ 0.5s
≤ 1s
≤ 2s
```

Without the bucket boundaries, `histogram_quantile()` cannot determine the distribution.

---

## 14. What do P50, P95, and P99 mean?

### P50

```text
50% of observations are at or below this value.
```

This is approximately the median.

### P95

```text
95% of observations are at or below this value.
```

The slowest 5% are above it.

### P99

```text
99% of observations are at or below this value.
```

The slowest 1% are above it.

---

## 15. How do you calculate P95 latency?

```promql
histogram_quantile(
  0.95,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

### Breakdown

```text
http_request_duration_seconds_bucket
        ↓
Get cumulative duration buckets

rate(...[5m])
        ↓
Calculate the recent per-second increase for each bucket

sum by (le)
        ↓
Combine matching buckets while preserving bucket boundaries

histogram_quantile(0.95, ...)
        ↓
Estimate P95 request latency
```

The query estimates the value below which 95% of request durations fall.

---

## 16. P50, P95, and P99 Queries

### P50

```promql
histogram_quantile(
  0.50,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

### P95

```promql
histogram_quantile(
  0.95,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

### P99

```promql
histogram_quantile(
  0.99,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

Simple rule:

```text
0.50 → P50
0.95 → P95
0.99 → P99
```

---

# Prometheus Alerting

## 17. What is a Prometheus alert rule?

An alert rule contains a PromQL expression that Prometheus continuously evaluates.

Example:

```yaml
- alert: HighCPUUsage
  expr: cpu_usage_percent > 90
  for: 5m
```

If the condition remains true for the configured duration, the alert fires.

---

## 18. What does `for` do?

Example:

```yaml
for: 5m
```

The alert condition must remain true continuously for 5 minutes before the alert becomes Firing.

This prevents alerts from firing because of short-lived spikes or temporary issues.

---

## 19. What are the alert states?

```text
Inactive → Pending → Firing
```

### Inactive

The alert condition is false.

### Pending

The condition is true, but the configured `for` duration has not yet completed.

### Firing

The condition has remained true for the complete `for` duration.

If the condition becomes false while Pending, the alert returns to Inactive.

---

## 20. What is the difference between labels and annotations?

### Labels

Labels are used to classify, group, and route alerts.

Example:

```yaml
labels:
  severity: critical
  team: backend
```

They can be used by Alertmanager for routing and grouping.

### Annotations

Annotations provide human-readable information about the alert.

Example:

```yaml
annotations:
  summary: High CPU usage detected
  description: CPU usage has been above 90% for more than 5 minutes.
```

### Simple Rule

> Labels are mainly used by systems.

> Annotations are mainly used by humans.

---

## 21. What is Prometheus responsible for?

Prometheus:

- Scrapes metrics
- Stores metrics
- Evaluates PromQL alert expressions
- Determines alert states
- Sends firing and resolved alerts to Alertmanager

---

## 22. What is Alertmanager responsible for?

Alertmanager:

- Groups alerts
- Deduplicates alerts
- Routes alerts
- Handles silences
- Handles inhibition
- Sends notifications to destinations such as Slack, email, or PagerDuty

---

## 23. Prometheus Alerting Flow

```text
Application / Exporter
        ↓
     Metrics
        ↓
   Prometheus
        ↓
Evaluates Alert Rules
        ↓
Condition becomes true
        ↓
     Pending
        ↓
Condition remains true for `for` duration
        ↓
     Firing
        ↓
Prometheus sends alert
        ↓
   Alertmanager
        ↓
Group / Deduplicate / Route
        ↓
Notification
(Slack / Email / PagerDuty)
```

---

## 24. High Error Rate Alert Example

```yaml
groups:
  - name: application-alerts
    rules:
      - alert: HighErrorRate
        expr: |
          (
            sum(rate(http_requests_total{status=~"5.."}[5m]))
            /
            sum(rate(http_requests_total[5m]))
          ) * 100 > 5
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: High HTTP error rate
          description: More than 5% of HTTP requests have returned 5xx errors for 5 minutes.
```

### Flow

```text
5xx error rate > 5%
        ↓
      Pending
        ↓
Still > 5% for 5 minutes
        ↓
      Firing
        ↓
Prometheus sends alert to Alertmanager
        ↓
Alertmanager groups and routes the alert
        ↓
Slack / Email / PagerDuty
```

---

# Must-Know Interview Answers

## What is `rate()`?

> `rate()` calculates the average per-second rate of increase of a Counter over a specified time range.

## What is the difference between `rate()` and `irate()`?

> `rate()` calculates an average rate over the selected time range, while `irate()` uses the two most recent samples, making it more responsive but noisier.

## What does P95 mean?

> P95 is the estimated value below which 95% of observations fall. The remaining 5% are above that value.

## Why use `sum by (le)` for `histogram_quantile()`?

> It combines matching histogram buckets across multiple time series while preserving the `le` bucket boundaries required by `histogram_quantile()`.

## What does `for: 5m` do?

> It ensures the alert condition remains true continuously for 5 minutes before the alert becomes Firing, preventing alerts caused by temporary spikes.

## What is the difference between Prometheus and Alertmanager?

> Prometheus evaluates alert conditions and determines alert states. Alertmanager groups, deduplicates, routes, silences, and sends notifications.