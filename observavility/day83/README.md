# DAY 83 - GRAFANA LOKI

1. ### WHAT IS GRAFANA LOKI?
- Grafana loki is an open source log aggregation system that is used for collecting, storing and querying logs from applications and infrastructure 
- Instead of checking logs on every pods we have a centralized place for analysing the logs and that is grafana loki 

2. #### WHY DO WE NEED CENTRALIZED LOGGING?
- **Without centraliized logging**, we might have to check logs of each pod manually
```
Service 1 → Check its logs manually
Service 2 → Check its logs manually
Service 3 → Check its logs manually
Service 4 → Check its logs manually
```
- and in k8s we might have multiple pods and it is very hard to read and debug each pod during an prod incident and hence we need a centralized log collection 
- **With centralized logging**
```
All Applications / Pods
          ↓
     Log Collection
          ↓
      Grafana Loki
          ↓
Centralized Logs
          ↓
Search and Investigate
```
3. ### HOW LOGS REACH GRAFANA LOKI
```
Application / Pod / Server
        ↓
     Generates Logs
        ↓
Log Collector / Agent
        ↓
    Grafana Loki
        ↓
      Grafana
        ↓
Search and Investigate
```
1. **Application**
- This is the target and it is responsible for generating logs

2. **Log collector/Agent**
- A collector or agent reads the logs and forwards them to loki 

3. **Storage**
- Grafana loki stores the logs along with their labels 

4. **Grafana Queries Loki**
- Grafana loki queries the stored logs according to users desire to get the logs

> The application generates logs, a collector/agent gathers and forwards them, Loki stores them, and Grafana is used to query and investigate them.

4. ### LABELS AND LOG STREAMS 
1. **LABELS**
- Labels are key-value pairs attached to logs that help to identify where the logs came from
- Example:
```
service="payment-service"
environment="production"
namespace="default"
pod="payment-abc123"
```
2. **LOG STREAMS**
- A log stream is a collection of **logs that have the exact same set of labels** 
```
{service="payment-service", environment="production", pod="payment-abc123"}

10:00 INFO Payment started
10:01 INFO Validating card
10:02 ERROR Payment failed
```
> Labels identify and organize logs, while a log stream is the collection of log entries that share the same set of labels.

5. ### HOW LOKI INDEXES LOGS 
- Loki does not index each word inside the log, **it indexes the label inside the logs**
>Loki does not index the full log content. It indexes labels such as service, environment, and namespace, which are used to narrow down the relevant logs. We can then filter the actual log content to find what we need.
