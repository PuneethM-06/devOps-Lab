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
- **Instead we can have one ArgoCD that manages other Argo CD applications and it is called App of Apps pattern**
```
                    Parent Application
                           │
          ┌────────────────┼────────────────┐
          │                │                │
          ▼                ▼                ▼
     frontend-app      backend-app      database-app
          │                │                │
          ▼                ▼                ▼
      Kubernetes       Kubernetes       Kubernetes
```
- The parent application doesnt directly deploy all your services.
- Instead it manages all the child application resources
- Each child application manages its k8s resources
```
Parent App
    ↓
Creates / manages Child Apps
    ↓
Each Child App deploys its own resources
```