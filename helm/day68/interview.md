# DAY 68 - HELM FUNDAMENTALS AND CHART STRUCTURE 

### WHAT IS HELM?
- Helm is a package manager for Kubernetes
- Helm helps you package, configure, install, upgrade and manage kuberenetes application

- Example:
- Suppose we have an application that needs to run in 3 environments say dev, stage and prod. Now, dev needs `replicas: 3` while prod needs `replicas: 5`. Without helm, we end up creating nearly same identical yml files and here is where helm kicks in.

```
              HELM
                │
                ▼
          Helm Chart
                │
        ┌───────┴───────┐
        │               │
   values-dev       values-prod
        │               │
        ▼               ▼
  Kubernetes YAML   Kubernetes YAML
        │               │
        └───────┬───────┘
                ▼
           Kubernetes
```
