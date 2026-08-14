# Interview.md — Day 79

# Prometheus Fundamentals

## 1. What is Prometheus, and how does it collect metrics?

Prometheus is an open-source monitoring and observability tool used to collect and store metrics.

It mainly uses a **pull model**. Prometheus periodically sends a request to configured targets, usually to the `/metrics` endpoint, scrapes the exposed metrics, and stores them as time-series data in its TSDB.

```text
Application
     │
     │ exposes metrics
     ▼
 /metrics
     ▲
     │ Prometheus scrapes
     │
Prometheus
     │
     ▼
   TSDB
```

---

## 2. Explain the basic Prometheus architecture.

```text
Application / Exporter
        │
        │ exposes metrics
        ▼
     /metrics
        ▲
        │
        │ Scrape / Pull
        │
   Prometheus
        │
        ├── Stores metrics in TSDB
        │
        ├── PromQL
        │
        └── Alert Rules
```

The application or exporter exposes metrics through `/metrics`.

Prometheus periodically scrapes those metrics and stores the collected samples as time-series data in its TSDB.

The stored data can then be:

- Queried using PromQL
- Visualized using Grafana
- Evaluated using alert rules

---

## 3. What is the difference between monitoring and observability?

### Monitoring

Monitoring is the process of continuously watching applications and systems to understand their current state and detect problems.

Examples:

- CPU usage
- Memory usage
- Request count
- Error rate
- Response time

Monitoring helps answer:

```text
Is the system healthy?
Is CPU usage high?
Are errors increasing?
Is the application slow?
```

### Observability

Observability is the ability to understand what is happening inside a system based on the signals it produces.

It helps answer:

```text
What went wrong?
Where did it go wrong?
Why did it happen?
```

The main observability signals are:

```text
Metrics
Logs
Traces
```

---

## 4. What is the difference between metrics, logs, and traces?

### Metrics

Metrics are numerical measurements collected over time.

Examples:

- Total requests
- CPU usage
- Memory usage
- Error rate
- Request duration

Metrics help understand:

> What is happening?

### Logs

Logs are detailed records of events.

Example:

```text
Payment request failed because database connection timed out.
```

Logs help understand:

> Why something happened?

### Traces

Traces show how a request travels through different services or components.

Example:

```text
User
  ↓
API
  ↓
Authentication Service
  ↓
Payment Service
  ↓
Database
```

Tracing helps understand:

> Where a request became slow or failed.

---

## 5. What is a metric?

A metric is a numerical measurement representing some behavior or state of an application or system.

Examples:

```text
http_requests_total
active_requests
cpu_usage
memory_usage
```

Metrics are collected over time so that we can observe changes and trends.

---

## 6. What are metric names, labels, and values?

Example:

```text
http_requests_total{method="GET",status="200"} 450
```

### Metric Name

```text
http_requests_total
```

It tells us what we are measuring.

### Labels

```text
method="GET"
status="200"
```

Labels add dimensions to the metric.

### Value

```text
450
```

The numerical value currently recorded for that metric.

---

## 7. What is a time series?

A time series is created by:

> Metric name + unique label combination.

Example:

```text
http_requests_total{method="GET",status="200"}
```

This is one time series.

Another combination:

```text
http_requests_total{method="GET",status="500"}
```

creates a different time series.

Prometheus stores values for these time series over time.

Example:

```text
10:00 → 450
10:15 → 500
10:30 → 540
```

---

## 8. What is the Prometheus pull model?

Prometheus uses a pull-based approach.

Prometheus periodically contacts configured targets and requests their metrics.

Example:

```text
Prometheus
    │
    │ GET /metrics
    ▼
Application
    │
    │ Returns metrics
    ▼
Prometheus
```

This process is called **scraping**.

---

## 9. What is a scrape?

A scrape is when Prometheus requests and collects metrics from a configured target.

Example:

```text
Prometheus
    │
    │ GET /metrics
    ▼
Target
```

The returned metrics are stored as time-series samples in Prometheus.

---

## 10. What is `scrape_interval`?

The `scrape_interval` defines how frequently Prometheus scrapes configured targets.

Example:

```yaml
global:
  scrape_interval: 15s
```

This means Prometheus attempts to scrape the target every 15 seconds.

```text
10:00:00 → Scrape
10:00:15 → Scrape
10:00:30 → Scrape
10:00:45 → Scrape
```

---

## 11. What is a target?

A target is an application, system, or endpoint from which Prometheus scrapes metrics.

Example:

```text
go-app:8080
```

Prometheus typically accesses:

```text
http://go-app:8080/metrics
```

---

## 12. What is a job in Prometheus?

A job is a logical group of targets that Prometheus scrapes using the same configuration.

Example:

```yaml
- job_name: "go-app"
```

A job can contain multiple targets:

```text
Job: go-app

Targets:
- go-app-1:8080
- go-app-2:8080
- go-app-3:8080
```

---

## 13. What is `prometheus.yml`?

`prometheus.yml` is the main Prometheus configuration file.

It can define:

- Global configuration
- Scrape intervals
- Jobs
- Targets
- Alert configuration

Example:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: "go-app"
    static_configs:
      - targets:
          - "go-app:8080"
```

---

## 14. What does `static_configs` mean?

Static configuration means targets are manually specified in the Prometheus configuration.

Example:

```yaml
static_configs:
  - targets:
      - "go-app:8080"
```

Prometheus is explicitly told where the target is located.

---

## 15. What is service discovery?

Service discovery allows Prometheus to dynamically discover targets instead of manually defining every target.

This is useful in environments such as Kubernetes where:

- Pods are created
- Pods are deleted
- IP addresses change
- New replicas are added

Simple distinction:

```text
Static Configuration
→ We manually define the targets

Service Discovery
→ Prometheus automatically discovers the targets
```

---

## 16. What are the main Prometheus metric types?

Prometheus has four main metric types:

```text
Counter
Gauge
Histogram
Summary
```

---

## 17. What is a Counter?

A Counter is a cumulative metric that generally only increases.

It can reset when the application restarts.

Examples:

- Total requests
- Total errors
- Total failed logins

Example:

```text
http_requests_total
```

Conceptually:

```text
0 → 1 → 2 → 3 → 4
```

---

## 18. What is a Gauge?

A Gauge represents a value that can increase or decrease.

Examples:

- Active requests
- Memory usage
- Number of connected users
- CPU temperature

Example:

```text
10 → 15 → 8 → 20
```

---

## 19. What is a Histogram?

A Histogram records observations and groups them into configurable buckets.

It is useful for measuring distributions.

A common example is:

```text
HTTP request duration
```

For example:

```text
Less than 0.1 seconds
Less than 0.5 seconds
Less than 1 second
Less than 5 seconds
```

Histograms are especially useful for latency analysis.

---

## 20. What is a Summary?

A Summary also records observations such as request duration.

It can provide quantile information calculated on the client side.

Examples:

```text
50th percentile
90th percentile
99th percentile
```

The key difference is:

```text
Histogram
→ Observations are grouped into buckets

Summary
→ Quantiles are calculated by the application/client
```

---

## 21. What is instrumentation?

Instrumentation means adding code or a client library to an application so that it can measure and expose information about its behavior.

For example:

```text
Total requests
Errors
Active requests
Request duration
```

Conceptually:

```text
Application
      │
      ▼
Prometheus Client Library
      │
      ▼
Creates and updates metrics
      │
      ▼
/metrics
```

---

## 22. What is an exporter?

An exporter is a component that collects or exposes metrics from a system or service in a format that Prometheus can scrape.

Example:

```text
Linux Server
     │
     ▼
Node Exporter
     │
     ▼
/metrics
     ▲
     │
Prometheus
```

Exporters are useful when we don't directly instrument the system ourselves.

---

## 23. What is the difference between instrumentation and an exporter?

### Instrumentation

Instrumentation is used when we can add metrics directly into the application code.

Example:

```text
Go Application
      │
      ▼
Prometheus Client Library
```

### Exporter

An exporter is used to expose metrics from external systems.

Examples:

```text
Linux Server → Node Exporter
Database → Database Exporter
```

Simple distinction:

> Instrumentation adds metrics directly to the application. An exporter exposes metrics from external systems so Prometheus can scrape them.

---

## 24. What does `UP` mean in Prometheus?

`UP` means Prometheus successfully scraped the configured target.

Prometheus exposes the `up` metric:

```text
up = 1
→ Target was successfully scraped

up = 0
→ Scrape failed
```

---

## 25. Does `UP` mean the entire application is healthy?

No.

`UP` only means Prometheus was able to successfully scrape the target's metrics endpoint.

Example:

```text
Prometheus
    │
    ▼
/metrics → Working ✅

User
    │
    ▼
/api/payment → 500 Error ❌
```

The application can be broken while Prometheus still shows the target as `UP`.

---

# Key Interview Questions

## 1. What is Prometheus, and how does it collect metrics?

Prometheus is an open-source monitoring and observability tool that uses a pull model to periodically scrape metrics from configured targets and store them as time-series data in its TSDB.

---

## 2. Explain the Prometheus architecture.

Applications or exporters expose metrics through an endpoint such as `/metrics`.

Prometheus periodically scrapes these targets and stores the collected samples in its TSDB.

The data can then be queried using PromQL, visualized using Grafana, or evaluated using alert rules.

---

## 3. What is the difference between Counter, Gauge, Histogram, and Summary?

- **Counter** → Cumulative value that generally increases
- **Gauge** → Value that can increase or decrease
- **Histogram** → Records observations in buckets
- **Summary** → Records observations and calculates quantiles on the client side

---

## 4. What is the difference between instrumentation and an exporter?

Instrumentation adds metrics directly into an application using code or a client library.

An exporter exposes metrics from external systems such as Linux servers or databases.

---

## 5. What does `UP` mean in Prometheus? Does it guarantee application health?

`UP` means Prometheus successfully scraped the target's metrics endpoint.

It does not guarantee that the entire application is healthy. The application can still have broken APIs, database failures, or return errors to users.

---

# Day 79 — Final Architecture

```text
Application
     │
     │ Instrumentation
     ▼
Prometheus Metrics
     │
     ▼
 /metrics
     ▲
     │
     │ Scrape / Pull
     │
Prometheus
     │
     ├── Stores samples in TSDB
     │
     ├── PromQL
     │      └── Day 80
     │
     └── Alert Rules
            └── Day 84
```

For external systems:

```text
External System
       │
       ▼
    Exporter
       │
       ▼
    /metrics
       ▲
       │
   Prometheus
```

---

# Day 79 — Final Takeaway

```text
Instrument Application
        ↓
Expose Metrics
        ↓
      /metrics
        ↓
Prometheus Scrapes
        ↓
Stores Data in TSDB
        ↓
Time Series Created
        ↓
PromQL Queries → Day 80
Grafana Dashboards → Day 81
OpenTelemetry Tracing → Day 82
Loki Logs → Day 83
Alertmanager → Day 84
```