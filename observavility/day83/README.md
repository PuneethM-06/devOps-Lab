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
