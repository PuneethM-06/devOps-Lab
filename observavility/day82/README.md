# DAY 82 - OpenTelementry 

1. ### WHAT IS OPENTELEMENTRY?
- **OpenTelementry(OTel) is a open-source observability framework and set of standards used to generate and collect and export telementry data such as metrics, logs and traces from applications.**
- **Telementry is also mainly used to track requests across multiple services**

2. ### TRACE
- Trace is a complete journey of a single request 
- Trace enables us to:
    - Which service handles the request 
    - How long did each service take 

3. ### SPAN 
- A span is an single operation within the journey of an request

4. ### TRACE ID vs SPAN ID
- **Trace Id**is a **unique identifier for the entire trace**
- **Span ID** is **indiviual span having a unique identifier**  
```
Trace ID: abc123

Frontend
Span ID: 111
    │
    └── API Service
        Span ID: 222
            │
            └── Payment Service
                Span ID: 333
                    │
                    └── Database
                        Span ID: 444
```
5. ### CONTEXT PROPAGATION 
- Context propogation allows OpenTelementry to know which span belongs to which Trace
- **How?** By the below example, if you see when Frontend passes request to API it also passes something called as **trace context** and this allows OTel to know it belongs to a certain Trace 
```
User Request
    ↓
Frontend
Trace ID: abc123
    │
    │ Pass Trace Context
    ▼
API
Trace ID: abc123
    │
    │ Pass Trace Context
    ▼
Payment
Trace ID: abc123
```
- **Trace context** typically contains:
    1. Trace ID
    2. Current Span ID

6. ### INSTRUMENTATION 
- Instrumentation is the process by which we **add OTel to an application so it can generate telementry data**, such as traces, metrics and logs

1. **AUTOMATIC INSTRUMENTATION**
- Sometimes OpenTelementry **automatically instruments supported libraries and framework** 
- Example:
```
Incoming HTTP Request
        ↓
OpenTelemetry automatically creates a span
        ↓
Application calls another HTTP service
        ↓
OpenTelemetry creates another span
        ↓
Database query
        ↓
OpenTelemetry creates another span
```
- Basically, OTel can automatically and easily follow the request and how an application is working

2. **MANUAL INSTRUMENTATION**
- Sometimes OTel cannotm understand the operation in our app and we might want to manually create a span 
- for example:
```
User Checkout Request
        ↓
Validate Cart
        ↓
Process Payment
        ↓
Send Confirmation
```
7. ### OpenTelementry Architecture 
```
Application
    ↓
Instrumentation
    ↓
OpenTelemetry SDK
    ↓
OpenTelemetry Collector
    ↓
Tracing Backend
```
1. **APPLICATION**
- This is basically the target, application or a service

2. **INSTRUMENTATION**
- Instrumentation is the process of OTel connecting to an application/service and it is used to **generate telementry data**

3. **OpenTelementry SDK**
- It is used inside the application to **manage and process the data collected during the process of instrumentation INSIDE THE APPLICATION**
- The SDK can handle things like:
    1. Creating and managing spans 
    2. Sampling 
    3. Process telementry data 
    4. Exporting telementry data 

4. **OpenTelementry Collector**
- It is used to **collect the telementry data from applications and PROCESS THEM OUTSIDE THE APPLICATION them**
- The collector can then:
    1. Receive telementry
    2. Process or transform it
    3. Filter it
    4. Batch it 

8. ### TRACING BACKEND 
- Tracing backend lets you query and visualize the traces that is collected from instrumentation and processed by SDK
- **Tracing backend stores the traces**

## OVERALL FLOW 
```
User Request
      ↓
Application
      ↓
Instrumentation creates spans
      ↓
OpenTelemetry SDK processes telemetry
      ↓
OpenTelemetry Collector receives telemetry
      ↓
Collector processes / batches / exports it
      ↓
Tracing Backend stores traces
      ↓
You visualize and analyze the trace
```

10. ### HOW PROPOGATION HAPPENS
```
Trace ID: abc123

Frontend
Span ID: 111
    │
    └── API Service
        Span ID: 222
        Parent: 111
            │
            └── Payment Service
                Span ID: 333
                Parent: 222
                    │
                    └── Database
                        Span ID: 444
                        Parent: 333
```
> If the API Service creates a new Span ID, why does it still belong to the same trace as the Frontend span?
- Answer: **The API span belongs to the same trace because it receives the same Trace ID through context propagation, while the Parent Span ID connects it to the Frontend span.**

### REAL PRODUCTION FLOW 
```
Slow Request Reported
        ↓
Use Tracing
        ↓
Trace the specific request
        ↓
Identify which service/span is slow
        ↓
Example: Payment Service → 4.5s
        ↓
Investigate the trace further
        ↓
Check which operation inside Payment Service is slow
        ↓
Database?
External API?
Application logic?
        ↓
Use Metrics
        ↓
Check whether this affects many requests
        ↓
P50 / P95 / P99 Latency
Request Rate
Error Rate
CPU / Memory
Pod Replicas / Scaling
        ↓
Identify the root cause
        ↓
Fix the issue
        ↓
Use Metrics + Traces to verify improvement
```
### PROMETHEUS vs GRAFANA vs OPENTELEMENTRY 
> Prometheus collects and stores metrics for monitoring, while Grafana visualizes those metrics using panels and dashboards. OpenTelemetry provides a standardized way to generate and export telemetry such as traces, metrics, and logs. In our case, tracing helps identify where a request is slow or failing, and then Prometheus metrics and Grafana can help investigate the broader cause.