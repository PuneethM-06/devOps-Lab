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
