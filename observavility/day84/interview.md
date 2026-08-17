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
````
