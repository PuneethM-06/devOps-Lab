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
```
One Argo CD
│
└── Parent Application
       │
       ├── Child Application → frontend → Kubernetes resources
       ├── Child Application → backend → Kubernetes resources
       ├── Child Application → payment → Kubernetes resources
       └── Child Application → monitoring → Kubernetes resources

```
> In production, we can have multiple microservices, and managing many individual Argo CD Application resources can become difficult. The App of Apps pattern uses a parent Application to manage multiple child Applications. Each child Application then manages and deploys its own Kubernetes resources.

2. ### PARENT APPLICATION vs CHILD APPLICATION
- **Parent application** - It is Argo CD `application resource`
- Its job is to create and manage child applications 
```
Parent Application
        ↓
frontend-app.yaml
backend-app.yaml
monitoring-app.yaml
```
- **Child application** - It is also Argo CD `application resource`
- Each child points to its own manifests
| Parent Application                       | Child Application                       |
| ---------------------------------------- | --------------------------------------- |
| Manages child `Application` resources    | Manages actual Kubernetes resources     |
| Points to YAML files defining child apps | Points to app manifests/Helm charts     |
| Used for grouping and bootstrapping      | Used to deploy individual services/apps |

3. ### HOW DOES THE PARENT APPLICATION ACTUALLY CREATE/MANAGE THOSE CHILD APPLICATION 
- **The parent application points to the git path containing the yaml files of the child applications**
- Example
**Inside parent application**
```
source:
  repoURL: <your-git-repo>
  path: apps
```
- **Inside frontend.yaml (child application)**
```
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: frontend
  namespace: argocd
```
### THE OVERALL FLOW IS 
```
Parent Application
        ↓
Reads Git repository
        ↓
Looks inside apps/
        ↓
Finds frontend.yaml
Finds backend.yaml
Finds monitoring.yaml
        ↓
Creates/updates those Child Applications
```
- then the child takes over
````
frontend Child Application
        ↓
Points to frontend manifests
        ↓
Deployment
Service
ConfigMap
```
