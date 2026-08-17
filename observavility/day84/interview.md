# DAY 84 - ALERTMANAGER 
- Prometheus detecs that a problem exists. ALertmanager manages what happen after the alert fires
```
Application
    ↓
Metrics
    ↓
Prometheus
    ↓
Evaluates Alert Rule
    ↓
Alert Fires
    ↓
Alertmanager
    ↓
Slack / Email / Other Receiver
```

### WHAT DOES ALERT MANAGER DO?
- Route the alert to the correct receiver 
- Group realted alerts together 
- Deduplicate repeated alerts 
- Silence alerts temporarily
- Inhibt less important alerts based on other alerts 
- Send notification on destinations such as slack or email 

- **Definition**
Alertmanager is a componnt of the prometheus ecosystem that recieves alerts from prometheus and manages how they are grouped, deduplicated, silenced, routed and sent notification channels

2. ### ALERT LIFECYCLE 
- The alert lifecycle is `Ìnactive -> Pending -> Firing -> Resolved`
1. **INACTIVE**
- Initial step and here the condition itself is false 

2. **PENDING**
- The condition is true, but the `for` duration is not satisfied yet 
```
Condition becomes true
        ↓
      Pending
        ↓
Wait for 5 minutes
```
3. **FIRING**
- Both, the condition and `for` is true now and hence the alert will be fired/triggered
```
Pending
   ↓
Condition true for 5 minutes
   ↓
Firing
   ↓
Sent to Alertmanager
```
4. **RESOLVED**
- Once the alert is triggered, the following engineer will look into the issue and once fixed it will be marked as resolved 
```
Error rate > 5%
        ↓
Firing
        ↓
Issue is fixed
        ↓
Error rate < 5%
        ↓
Resolved
```
