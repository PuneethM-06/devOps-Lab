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
- The reason why ArgoCD runs inside the k8s clustyer:
    1. To read the current state of the k8s cluster 
    2. To perform reconcile 

2. ### WHAT ARE WE ACTUALLY INSTALLING 
- When we install Argo CD, k8s creates multiple Argo CD components, usually inside an `argocd` namespace
```
Kubernetes Cluster
│
└── argocd namespace
    │
    ├── argocd-server
    ├── argocd-repo-server
    ├── argocd-application-controller
    └── other supporting components
```
