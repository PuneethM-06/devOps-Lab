# Interview.md — Day 81

# Grafana — Dashboards from Prometheus Data

---

## 1. What is Grafana?

Grafana is a visualization and observability tool used to display data from different data sources, such as Prometheus.

It allows us to visualize metrics using:

- Graphs
- Panels
- Dashboards
- Gauges
- Tables
- Stat panels

Grafana does not collect metrics in the same way Prometheus does.

```text
Prometheus → Collects and stores metrics
Grafana    → Queries and visualizes metrics
```

---

## 2. Why use Grafana if Prometheus already has a UI?

Prometheus has its own UI where we can run PromQL queries and view basic graphs.

However, Grafana provides:

- Better visualizations
- Custom dashboards
- Multiple panel types
- Centralized monitoring
- Support for multiple data sources
- Reusable and dynamic dashboards

Grafana is useful for creating dashboards that help teams quickly understand what is happening in their systems.

---

## 3. Prometheus vs Grafana

```text
Prometheus
    ↓
Collects metrics
Stores metrics
Provides PromQL
Evaluates queries and alerts

Grafana
    ↓
Uses Prometheus as a data source
Sends queries to Prometheus
Receives query results
Visualizes the data
```

Simple way to remember:

> Prometheus collects and stores the metrics. Grafana queries and visualizes them.

---

# 4. Complete Grafana and Prometheus Data Flow

```text
Application / Exporter
        ↓
Exposes metrics
        ↓
Prometheus
        ↓
Scrapes metrics
        ↓
Stores metrics in TSDB
```

When a Grafana panel needs data:

```text
Grafana Panel
        ↓
PromQL Query
        ↓
Prometheus
        ↓
Queries Prometheus TSDB
        ↓
Returns query result
        ↓
Grafana
        ↓
Visualizes the data
```

Example query:

```promql
rate(http_requests_total[5m])
```

Complete flow:

```text
Grafana
   ↓
Sends PromQL query
   ↓
Prometheus
   ↓
Queries its TSDB
   ↓
Returns metric data
   ↓
Grafana
   ↓
Displays the result in a panel
```

Important:

> Grafana does not directly query the Prometheus TSDB. Grafana sends the query to Prometheus, and Prometheus queries its TSDB.

---

# 5. What is a Dashboard?

A Dashboard is a collection of multiple panels that provide an overview of a system or application.

Example:

```text
Application Monitoring Dashboard

├── Request Rate
├── Error Rate
├── P95 Latency
├── CPU Usage
└── Service Health
```

Simple definition:

> A Dashboard is a collection of panels.

---

# 6. What is a Panel?

A Panel is a single visualization inside a Grafana dashboard.

Example:

```text
Dashboard

├── Panel → Request Rate
├── Panel → Error Rate
├── Panel → P95 Latency
└── Panel → CPU Usage
```

Each panel usually contains:

```text
Data Source
    ↓
Query
    ↓
Visualization
```

Example:

```text
Prometheus
    ↓
rate(http_requests_total[5m])
    ↓
Time Series Panel
```

Simple definition:

> A Panel is a single visualization of a metric or data.

---

# 7. Time Series Panel

A Time Series panel is used when we want to see how a metric changes over time.

Examples:

- Request rate
- Error rate
- CPU usage
- Memory usage
- P95 latency

Example:

```text
Metric changes over time

100 ┤       ╭──╮
 80 ┤   ╭───╯  ╰──╮
 60 ┤───╯         ╰──
    └──────────────────
             Time →
```

Use a Time Series panel when:

> The trend, spikes, increases, decreases, or fluctuations over time are important.

---

# 8. Stat Panel

A Stat panel is used to display a single important value.

Examples:

```text
Current Error Rate

2.4%
```

```text
CPU Usage

72%
```

Use a Stat panel when:

> You want to quickly see one important value.

A Stat panel does not show how the metric changed over time like a Time Series panel.

---

# 9. Gauge Panel

A Gauge panel is used to display a value within a defined range or against thresholds.

Examples:

- CPU usage
- Memory usage
- Disk usage

Example:

```text
0% ───────────●────────── 100%
              72%
```

Use a Gauge when:

> The value has a meaningful minimum, maximum, or threshold.

---

# 10. Table Panel

A Table panel displays data in rows and columns.

Example:

```text
Instance        CPU Usage

server-1        45%
server-2        78%
server-3        92%
```

Use a Table when:

> You want to display and compare multiple instances, services, or values.

---

# 11. Visualization Cheat Sheet

```text
Metric changes over time
        ↓
Time Series

Single important value
        ↓
Stat

Value within a defined range
        ↓
Gauge

Data in rows and columns
        ↓
Table
```

---

# 12. Request Rate Panel

## Query

```promql
sum(rate(http_requests_total[5m]))
```

## What it does

```text
http_requests_total
        ↓
Counter containing total requests

rate(...[5m])
        ↓
Calculates average requests per second
over the last 5 minutes

sum(...)
        ↓
Adds request rates from all matching instances
```

Example:

```text
pod-1 → 20 req/sec
pod-2 → 30 req/sec
pod-3 → 50 req/sec

Total → 100 req/sec
```

## Visualization

Use:

> Time Series

Reason:

Request rate changes over time, and we want to see:

- Traffic trends
- Spikes
- Increases
- Decreases

Important:

> The request Counter generally increases, but the request rate itself can increase or decrease.

---

# 13. Error Rate Panel

## Query

```promql
(
  sum(rate(http_requests_total{status=~"5.."}[5m]))
  /
  sum(rate(http_requests_total[5m]))
) * 100
```

## What it does

```text
5xx request rate
──────────────── × 100
Total request rate
```

The numerator calculates the total 5xx error request rate.

The denominator calculates the total request rate.

Example:

```text
5xx requests   → 5 req/sec
Total requests → 100 req/sec

Error Rate:

(5 / 100) × 100 = 5%
```

## Visualization

Use:

> Time Series

Reason:

We want to see:

- Error spikes
- When errors started
- How long the problem lasted
- Whether the error rate returned to normal

If the error rate changes:

```text
1% → 15% → 1%
```

A Time Series panel shows when the spike occurred.

We can then investigate using:

- Logs
- Deployments
- CPU and memory metrics
- Database metrics
- Other application metrics

A Stat panel might only show the latest value and could hide an earlier spike.

---

# 14. P95 Latency Panel

## Query

```promql
histogram_quantile(
  0.95,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

## What it does

```text
Histogram buckets
        ↓
rate()
        ↓
Calculate recent per-second increase
for each bucket

sum by (le)
        ↓
Combine matching buckets across instances
while preserving bucket boundaries

histogram_quantile(0.95)
        ↓
Estimate P95 latency
```

Example:

```text
P95 = 0.8 seconds
```

This means:

> 95% of requests completed in 0.8 seconds or less.

The remaining 5% took longer.

## Visualization

Use:

> Time Series

Reason:

We want to monitor:

- Latency spikes
- Performance degradation
- Changes after deployments
- Long-term performance trends

Example:

```text
Normal P95 → 200ms

Suddenly → 2 seconds
```

This indicates that the application's request latency has significantly increased.

Possible things to investigate:

- Pod CPU usage
- Memory usage
- Resource limits or throttling
- Increased traffic
- Scaling requirements
- Database latency
- External API latency
- Recent deployments
- Network issues

---

# 15. Service Health Panel

## Query

```promql
up
```

For a specific application:

```promql
up{job="my-app"}
```

## What does `up` mean?

```text
up = 1
```

Prometheus successfully scraped the target.

```text
up = 0
```

Prometheus could not successfully scrape the target.

Example:

```text
instance-1 → 1
instance-2 → 1
instance-3 → 0
```

## Visualization

Use:

> Stat

This provides a quick view of the current scrape status.

Example:

```text
Service Health

UP
```

or:

```text
Service Health

DOWN
```

Important:

```text
up = 1 ≠ Application is healthy
```

`up = 1` only means Prometheus can successfully reach and scrape the target.

The application can still have internal problems.

Example:

```text
Prometheus → Successfully scrapes target → up = 1

Users → Receiving HTTP 500 errors
```

---

# 16. Application Monitoring Dashboard

A simple application monitoring dashboard can contain:

```text
Application Monitoring Dashboard

├── Request Rate
│   └── Time Series
│
├── Error Rate
│   └── Time Series
│
├── P95 Latency
│   └── Time Series
│
└── Service Health
    └── Stat
```

---

# 17. What are Grafana Dashboard Variables?

Grafana variables make dashboards dynamic and reusable.

Instead of creating multiple dashboards for different applications or servers, we can create one dashboard and select the application or instance using a dropdown.

Example:

```text
Application: [ frontend ▼ ]
```

The user can select:

```text
Application: [ backend ▼ ]
```

The same dashboard updates automatically.

---

# 18. `$job` Variable

Suppose Prometheus has:

```text
job="frontend"
job="backend"
job="payment-service"
```

We can create a Grafana variable:

```text
$job
```

Then use it in a PromQL query:

```promql
rate(http_requests_total{job="$job"}[5m])
```

If the user selects:

```text
frontend
```

Grafana effectively queries:

```promql
rate(http_requests_total{job="frontend"}[5m])
```

If the user selects:

```text
backend
```

Grafana effectively queries:

```promql
rate(http_requests_total{job="backend"}[5m])
```

The same dashboard can therefore display metrics for different applications.

---

# 19. `$instance` Variable

We can also dynamically select an instance.

Variable:

```text
$instance
```

Example query:

```promql
up{instance="$instance"}
```

The user can select:

```text
Instance: [ server-1:8080 ▼ ]
```

The dashboard then filters the query using the selected instance.

---

# 20. How Grafana Variables Work

```text
Grafana Variable
       ↓
User selects a value
       ↓
$job or $instance is replaced
       ↓
PromQL query is filtered
       ↓
Dashboard updates
```

---

# 21. Why Use Variables?

Variables help us:

- Reuse dashboards
- Dynamically filter metrics
- Avoid duplicate dashboards
- Switch between applications
- Switch between instances
- Keep dashboards easier to maintain

Instead of:

```text
Dashboard 1 → Frontend
Dashboard 2 → Backend
Dashboard 3 → Payment Service
```

We can use:

```text
One Dashboard
      │
      └── $job dropdown
             │
             ├── Frontend
             ├── Backend
             └── Payment Service
```

Key point:

> Grafana variables make dashboards dynamic and reusable by filtering queries based on selected values such as `$job` or `$instance`.

---

# 22. Grafana Dashboard Best Practices

## Don't overload dashboards

Avoid putting too many unrelated panels on one dashboard.

A good overview dashboard should quickly show the most important metrics.

Example:

```text
Application Overview

├── Request Rate
├── Error Rate
├── P95 Latency
└── Service Health
```

Create separate dashboards for deeper investigation when needed.

---

## Use meaningful panel titles

Bad:

```text
Panel 1
Graph
CPU
```

Better:

```text
HTTP Request Rate
HTTP 5xx Error Rate
P95 Request Latency
Pod CPU Usage
```

The panel title should immediately explain what the user is looking at.

---

## Choose the correct visualization

```text
Changes over time
        ↓
Time Series

Single value
        ↓
Stat

Defined range or threshold
        ↓
Gauge

Rows and columns
        ↓
Table
```

Choose the visualization based on what you are trying to understand, not just what looks good.

---

## Use sensible time ranges

The time range should match the investigation.

Examples:

```text
Troubleshooting an incident
→ Last 15–60 minutes

Daily monitoring
→ Last 6–24 hours

Long-term trends
→ Days or weeks
```

---

## Use variables when appropriate

Instead of creating separate dashboards for every application:

```text
Frontend Dashboard
Backend Dashboard
Payment Dashboard
```

Use one reusable dashboard:

```text
Application Dashboard

Job: [ frontend ▼ ]
```

---

# Must-Know Interview Answers

## What is Grafana?

> Grafana is a visualization tool that queries data sources such as Prometheus and displays metrics using graphs, panels, and dashboards.

## How does Grafana get Prometheus data?

> Grafana sends a PromQL query to Prometheus. Prometheus queries its TSDB, returns the requested data to Grafana, and Grafana visualizes the result.

## What is the difference between a Dashboard and a Panel?

> A Dashboard is a collection of panels, while a Panel is a single visualization of a metric or data.

## When should you use a Time Series panel?

> Use a Time Series panel when you want to see how a metric changes over time, including trends, spikes, increases, and decreases.

## When should you use a Stat panel?

> Use a Stat panel when you want to display a single important value or quick snapshot.

## What are Grafana variables?

> Grafana variables make dashboards dynamic and reusable. Variables such as `$job` or `$instance` can be selected from a dropdown and used to filter PromQL queries.

## Why use `$job` or `$instance`?

> They allow us to reuse the same dashboard for multiple applications or instances instead of creating separate dashboards for each one.