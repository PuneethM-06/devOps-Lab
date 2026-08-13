# DAY 76 - ARGOCD 

1. ### WHERE DOES ARGOCD RUN?
-  ArgoCD system runs inside your k8s cluster
```
Your Machine
│
├── kubectl
├── argocd CLI
│
└──────────────► Kubernetes Cluster
                       │
                       └── argocd namespace
                              │
                              ├── Argo CD components
                              └── Argo CD manages applications
```
