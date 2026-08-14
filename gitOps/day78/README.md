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
```
frontend Child Application
        ↓
Points to frontend manifests
        ↓
Deployment
Service
ConfigMap
```

4. ### REPOSITORY STRUCTURE
```
gitops-repo/
│
├── root-app.yaml
│
├── apps/
│   ├── frontend-app.yaml
│   ├── backend-app.yaml
│   └── monitoring-app.yaml
│
├── frontend/
│   ├── deployment.yaml
│   └── service.yaml
│
├── backend/
│   ├── deployment.yaml
│   └── service.yaml
│
└── monitoring/
    └── ...
```
- This can be the structure to work on and the flow is like 
````
Git Repository
│
├── root-app.yaml
│       │
│       ▼
│     apps/
│       │
│       ├── frontend-app.yaml
│       │        │
│       │        ▼
│       │     frontend/
│       │        ├── deployment.yaml
│       │        └── service.yaml
│       │
│       └── backend-app.yaml
│                │
│                ▼
│             backend/
│                ├── deployment.yaml
│                └── service.yaml
````
- root points to apps(frontend-app.yaml, backend-app.yaml) and this frontend in turn might point to `frontend/`

## IMPORTANT 
```
1. You apply root-app.yaml
        ↓
2. Kubernetes creates the Parent Application
        ↓
3. Parent App looks at:
   repoURL + path: apps/
        ↓
4. It reads:
   frontend-app.yaml
   backend-app.yaml
   monitoring-app.yaml
        ↓
5. These YAMLs are applied to Kubernetes
   as Child Application resources
        ↓
6. Each Child Application has its own:
   repoURL + path
        ↓
7. Each Child App reads its application manifests
        ↓
8. Creates/manages Kubernetes resources
   like Deployment, Service, ConfigMap, etc.
```
5. ### WHY USE APP OF APPS IN PRODUCTION 
- **Scalability and management**

6. ### LIMITATION
- More complex
- Parent application can have larger blast radius 
- App of Apps is not always necessary

> When a new Child Application YAML is added, the Parent Application syncs it and creates/manages it as a new Argo CD Application. That Child Application then deploys and manages the Kubernetes resources defined in its configured source path.