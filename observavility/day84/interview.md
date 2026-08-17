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
### OVERALLL FLOW/MENTAL MODEL
```
Condition False
      ↓
   Inactive

Condition becomes True
      ↓
   Pending
      ↓
True for entire `for` duration?
      ↓
     Yes
      ↓
   Firing
      ↓
Alertmanager
      ↓
Slack / Email

Condition becomes False
      ↓
Resolved / Inactive
```

### Grouping 
- It is the process of grouping similar alerts to ensure that there are no similar errors being flooded 
```
5 Payment Service Alerts
        ↓
   Alertmanager
        ↓
One Grouped Notification
        ↓
 Payment Service: 5 instances affected
```

### routing 
- Routing determines where the alert must be **routed based on the labels** 
```
Payment Service Alert
        ↓
team="payments"
        ↓
Payments Slack Channel
```

### Deduplication 
- Dedupliation prevent alertmanager from repeatedly sending the same alert notification 
```
Prometheus
    ↓
Same alert updates
    ↓
Alertmanager
    ↓
Recognizes existing alert
    ↓
Avoids unnecessary duplicate notifications
```
### Silencing 
- Silencing temporarily suppresses notiifcations for specific alerts 
```
Planned Maintenance
        ↓
Create Silence
        ↓
Matching Alerts
        ↓
Notifications Suppressed
```
### Inhibition 
- Inhibition is when one alert ca automatically supress another related alert
```
DatabaseDown → CRITICAL

DatabaseConnectionError → WARNING
ApplicationError → WARNING
```
- If DB is down other alerts maybe the consequences of these alerts and hence it will be automatically silenced