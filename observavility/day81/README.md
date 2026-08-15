# DAY 81 - GRAFANA - DASHBOARD FROM PROMETHEUS DATA 

1. ### WHAT IS GRAFANA 
- **Grafana is a visualization and observability platform that queries data sources such as Prometheus and represents the data through dashboard, panels, graphs, tables etc.**
- Grafana acts as a centralized visualization platform where we can connect multiple data sources, such as Prometheus, and represent the data using dashboards, graphs, gauges, tables, and other panels.

### GRANFANA ARCHITECTURE 
```
Grafana Dashboard
        │
        ▼
Grafana Panel
        │
        │ PromQL Query
        ▼
     Prometheus
        │
        │ Queries historical metric data
        ▼
       TSDB
        │
        │ Returns result
        ▼
     Prometheus
        │
        │ Sends query response
        ▼
      Grafana
        │
        ▼
Visualizes the result as a graph
```
> The application exposes metrics, Prometheus scrapes and stores them in its TSDB. Grafana uses Prometheus as a data source and sends PromQL queries to Prometheus. Prometheus queries the TSDB, returns the results to Grafana, and Grafana visualizes those results using dashboards and panels.

### DASHBOARD vs PANEL 
- **DASHBOARD** - A dashboard is a collection of multiple visualizations that give you an overall view of a system or application 
```
Application Dashboard
│
├── Request Rate
├── Error Rate
├── P95 Latency
├── CPU Usage
└── Memory Usage
```
> **DASHBOARD IS A COLLECTION OF PANELS**
- **PANEL**
- Panel is a indiviual visualization inside a dashboard 

### PANEL VISUALIZATION TYPES
1. ### TIME SERIES
- Used to see how a metric changes over a time 
- Example: 
    - Request rate 
    - CPU usage
    - Memory usage 
```
Request Rate

100 ┤       ╭──╮
 80 ┤   ╭───╯  ╰──╮
 60 ┤───╯         ╰──
    └──────────────────
       Time →
```
2. ### STAT
-  Used to show one important current value
- Example:
```
Current Error Rate
      2.4%
```

3. ### GAUGE
- Used to show a value in a defined range 
- Example
```
CPU Usage

0% ───────────●────────── 100%
              72%
```

4. ### TABLE
- Used to display data in rows and columns 
- Example:
```
Instance        CPU Usage
-------------------------
server-1        45%
server-2        78%
server-3        92%
```
# Grafana Dashboard Panels

## 1. Request Rate Panel

### Query

```promql
sum(rate(http_requests_total[5m]))
```

### What it does

- `http_requests_total` is a Counter.
- `rate(...[5m])` calculates the average requests per second over the last 5 minutes.
- `sum()` adds the request rates from all matching instances.

Example:

```text
pod-1 → 20 req/sec
pod-2 → 30 req/sec

Total → 50 req/sec
```

### Visualization

**Time Series**

Use it because request rate changes over time, and we want to see traffic trends, spikes, increases, and decreases.

---

## 2. Error Rate Panel

### Query

```promql
(
  sum(rate(http_requests_total{status=~"5.."}[5m]))
  /
  sum(rate(http_requests_total[5m]))
) * 100
```

### What it does

Calculates the percentage of HTTP requests that resulted in 5xx errors.

```text
5xx requests
───────────── × 100
total requests
```

Example:

```text
5xx requests → 5 req/sec
Total requests → 100 req/sec

Error Rate → 5%
```

### Visualization

**Time Series**

This helps us see when the error rate increases or spikes.

We can then investigate that time using:

- Logs
- Deployments
- CPU or memory usage
- Database issues
- Other metrics

A Stat panel would only show a single/latest value and could hide previous spikes.

---

## 3. P95 Latency Panel

### Query

```promql
histogram_quantile(
  0.95,
  sum by (le) (
    rate(http_request_duration_seconds_bucket[5m])
  )
)
```

### What it does

```text
Histogram buckets
        ↓
rate()
        ↓
Calculate recent per-second increase

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

### Visualization

**Time Series**

Use it to monitor:

- Latency spikes
- Performance degradation
- Improvements after deployments
- Changes in application performance over time

If P95 normally stays at 200ms and suddenly reaches 2 seconds, we should investigate possible issues such as:

- High CPU or memory usage
- Resource limits or throttling
- Increased traffic
- Need for scaling
- Database latency
- External API issues
- Recent deployments

---

## 4. Service Health Panel

### Query

```promql
up
```

Or for a specific application:

```promql
up{job="my-app"}
```

### What it does

Prometheus automatically provides the `up` metric.

```text
up = 1 → Prometheus successfully scraped the target
up = 0 → Prometheus could not successfully scrape the target
```

Example:

```text
instance-1 → 1
instance-2 → 1
instance-3 → 0
```

### Visualization

**Stat**

Useful for quickly checking the current scrape status.

```text
Service Health

UP
```

or:

```text
Service Health

DOWN
```

### Important

```text
up = 1 ≠ Application is healthy
```

`up = 1` only means Prometheus can successfully reach and scrape the target.

The application can still be failing internally.

Example:

```text
Prometheus → Successfully scrapes target → up = 1

Users → Receiving HTTP 500 errors
```

---

# Application Monitoring Dashboard

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

### GRAFANA DASHBOARD VARIABLES
- **Variables make dashboards dynamic and reusable**
# Grafana Dashboard Variables

## What are Dashboard Variables?

Grafana variables make dashboards dynamic and reusable.

Instead of creating multiple dashboards for different applications or servers, we can create one dashboard and use variables to select what data we want to see.

For example:

Application: [ frontend ▼ ]

The user can select another application:

Application: [ backend ▼ ]

The same dashboard updates automatically.

---

## Example: `$job` Variable

Suppose Prometheus has multiple applications:

```text
job="frontend"
job="backend"
job="payment-service"
```
### BEST PRACTICES
1. Dont overload a dashboard
2. Use meaningful panel titles
3. Choose the correct visualization 
4. Use sensible time ranges
5. Use variables when appropriate 
