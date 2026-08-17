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
