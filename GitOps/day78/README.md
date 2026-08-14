# DAY 78 - ARGO CD APP OF APPS PATTERN

1. ### WHY DO WE NEED IT?
- In a production environment we have multiple microservices running
- Example:
```
frontend
backend
auth-service
payment-service
notification-service
postgres
redis
prometheus
grafana
```
- If you have 20 microservices, you might need to create and manage 20 separate Argo CD Application resources.