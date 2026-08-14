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
- 
