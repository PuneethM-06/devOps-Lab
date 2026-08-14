# DAY 79 - PROMETHEUS

1. ### WHAT IS MONITORING 
- It is continously watching a system and checking its know health indicators 
- Example:
```
CPU > 80%
        ↓
Something might be wrong
        ↓
Trigger an alert
```

2. ### MONITORING LIMITATION 
- Monitoring may say something is failing, or something is not right but it does not tell which application, which request etc and that is done by observability

3. ### OBSERVABILITY
- Monitoring might tell you something is wrong, observability will tell you why it is wrong 
- There are 3 pillars in observability
```
                 OBSERVABILITY
                        │
          ┌─────────────┼─────────────┐
          │             │             │
          ▼             ▼             ▼
       Metrics         Logs         Traces
```
- **METRICS** - WHAT IS HAPPENING *EXAMPLE: CPU USAGE % IS ABOVE 90%*
- **LOGS** - WHAT HAPPENED - *Example: ERROR - Database connection timedout*
- **TRACES** - WHERE AND WHY DID IT HAPPEN
| Monitoring            | Observability                                 |
| --------------------- | --------------------------------------------- |
| Watches known signals | Helps investigate system behavior             |
| Detects problems      | Helps understand problems                     |
| "Something is wrong"  | "Why is it wrong?"                            |
| Metrics and alerts    | Metrics + logs + traces and their correlation |

4. ### METRICS vs LOGS vs TRACES
1. **METRICS**
- **Are numerical measurements collected overtime**
- Example: How many errors are failing : 250

2. **LOGS**
- **Logs gives detailed information about a certain event**
```
INFO  User requested /payment

INFO  Connecting to database

ERROR Database connection timeout
```
3. **TRACES**
- Tracing becomes useful when a request travels through multiple microservices
```
Request ID: abc123

API Service
    │ 50ms
    ▼
Payment Service
    │ 200ms
    ▼
Database
    │ 5 seconds ❌
    ▼
Timeout
```
5. ### WHAT EXACTLY IS A METRIC
- **A metric is a numerical measurement of something in a system, recorded over time.**
- A metric has 4 important values
    1. Metric name 
    2. Labels 
    3. Value
    4. Timestamp
1. **METRIC NAME**
- This tells what are we measuring 
- Example: https_requests_total 

2. **LABELS**
- Labels add dimension to the metric
- Example: `method="GET", endpoint="/"`
`http_requests_total{method="POST", endpoint="/login"} = 50`

3. **VALUES**
- This the metric value given by the tool 

4. **Timestamp**
- **Prometheus also records when that value was recorded**
```
10:00 → 100
10:15 → 120
10:30 → 150
```

6. ### LABELES AND TIME SERIES 
- Labels give extra context to the metric 
- Example:
```
http_requests_total{
    method="GET",
    endpoint="/",
    status="200"
} 450
```
- Now we know there were 450 successful GET requests
- **A unique combination of metric name and labels identifies a unique time series.**
- Example:
```
Time Series 1

http_requests_total
method="GET"
status="200"

Time Series 2

http_requests_total
method="GET"
status="500"
```
- **The metric name is the same, but each unique combination of labels creates a separate time series.**

### WHAT IS PROMETHEUS AND WHAT DOES IT SOLVE 
- Prometheus is a** monitoring and observability system** that:
    1. **Collect  metric** from application and system 
    2. Stores those metrics as time-series data 
    3. Query metrics using **PromQL**
    4. Create **Alert rules** 
```
Application
    │
    │ exposes metrics
    ▼
 /metrics
    ▲
    │
    │ Prometheus collects them
    │
Prometheus
    │
    ├── Stores metrics
    │
    ├── PromQL queries
    │
    └── Alert rules
```
- **Time Series Database** - **Prometheus has its own built-in time-series database for storing metric samples.**

### PROMETHEUS ARCHITECTURE AND MAIN COMPONENTS
```                 ┌──────────────────┐
                 │    Flask App     │
                 │                  │
                 │    /metrics      │
                 └────────▲─────────┘
                          │
                          │ 1. Scrape
                          │
                 ┌────────┴─────────┐
                 │   Prometheus     │
                 │                  │
                 │  ┌────────────┐  │
                 │  │ Scraper    │  │
                 │  ├────────────┤  │
                 │  │ TSDB       │  │
                 │  ├────────────┤  │
                 │  │ PromQL     │  │
                 │  ├────────────┤  │
                 │  │ Alert Rules│  │
                 │  └────────────┘  │
                 └────────┬─────────┘
                          │
                 ┌────────┴─────────┐
                 │                  │
                 ▼                  ▼
              PromQL             Alertmanager
              Queries                 │
                                      ▼
                                  Slack / Email
```
1. **TARGET**
- A target is simply something Prometheus monitor

2. **SCRAPING**
- Prometheus periodically sends a request to the target 

3. **TSDB - Time series Database**
- The scraped metric values from the target is stored in TSDB

4. **PromQL**
- Prometheus query language 
- It is to query and fetch the details 

5. **ALERT RULES**
- Prometheus can continuously evaluate rules 
```
IF error rate > 5%
FOR 5 minutes
```

### PULL MODEL AND SCRAPPING 
1. **PULL MODEL**
- It is done by Prometheus, It pulls the metrics from the application 
- It sends a `GET /metrics`, collects the value and stores it in TSDB

2. **SCRAPE INTERVAL**
- It defines how often prometheus collects metrics
```
10:00:00 → Scrape
10:00:15 → Scrape
10:00:30 → Scrape
10:00:45 → Scrape
```

- **WHY IS PULL MODEL USEFUL??**
1. Prometheus controls collections 
2. Prometheus can detect unhealthy targets
3. Targets dont need to know prometheus internals 

## PROMETHEUS METRIC TYPES
1. **COUNTER**
- A counter only increases
- Example:
```
Requests received:

0 → 1 → 2 → 3 → 4 → 5
```
- It represents something that keeps accumulating 
- Counters can reset; when the application restarts the counters can reset 

2. **GAUGE**
- A gauage can go up and down
```
Active users:

10 → 15 → 8 → 20 → 5
```
3. **HISTOGRAM**
- A histogram is used when we want distributed values 
- Example: How long are HTTP requests taking?
- Prometheu drops these values to buckets, and then they are populated in histogram 

4. **SUMMARY**
- Summary can observe values such as:
```
Request duration
Response size
Count
Sum
Quantiles
```
### INSTRUMENTATION
- It answers; How does our flask application actually create these metrics and expose them to prometheus
- **It is the process of adding a code or library to an application so it can create metric about its behaviour is called as instrumentation**
```
Flask Application
       │
       │ uses
       ▼
prometheus_client
       │
       ▼
Creates metrics
       │
       ▼
/metrics endpoint
       ▲
       │
       │ Prometheus scrapes
       │
Prometheus
```
### TWO TYPES OF METRICS
1. **DEFAULT METRICS**
- Prometheus client already exposes some process or runtime metrics before creating anything 

2. **CUSTOM METRICS**
- These are metric we explicitly define 
```
prometheus_client → creates/records/exposes metrics
Prometheus        → scrapes and stores metrics in TSDB
```

### EXPORTERS
- **Collect or translates metrics from another system and exposes them in a format prometheus an scrape**

### PROMETHEUS CONFIG
```
                prometheus.yml
                       │
                       ▼
            scrape_configs
                       │
                       ▼
              job: go-app
                       │
                       ▼
             target: go-app:8080
                       │
                       │ every 15 seconds
                       ▼
         GET /metrics
                       │
                       ▼
                 Go Application
                       │
                       ▼
                    Metrics
                       │
                       ▼
                 Prometheus TSDB
```
- We make use of yaml file to say prometheus to get the metric from where 
```
Target
→ The actual thing being scraped

Job
→ Logical group/configuration for one or more targets
```
### PROMETHEUS UP AND DOWN 
- **UP** - If prometheus can scrape the target successfully
- **down** - If prometheu fails to connect to target for various reasons like wrong config then it is called as DOWN
```
JOB        TARGET           STATE

go-app     go-app:8080      UP
```
### SERVICE DISCOVERY 
- **static configuration**
```
static_configs:
  - targets:
      - "go-app:8080"
```
- pods maybe created and deleted and resulting in change of IP addressed too and this is solved by service discovery 
- **Prometheus can dynamically discover targets.**
```
Kubernetes
     │
     │ "Here are the current Pods/Services"
     ▼
Prometheus
     │
     ▼
Discovers targets automatically
```
