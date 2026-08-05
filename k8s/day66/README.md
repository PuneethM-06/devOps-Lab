# DAY 66 - KUBERNETES

## WHY DO NAMESPACES EXIST
- Namespaces are like folders inside a desktop
- Namespaces are used to organize resources inside a k8s cluster 
- Example:
```
Cluster

├── dev
│     ├── API
│     └── Redis
│
├── qa
│     ├── API
│     └── Redis
│
├── prod
│     ├── API
│     └── Redis
│
└── monitoring
      ├── Prometheus
      └── Grafana
```
- A namespace is a logicial partition within a k8s cluster that organizes and isolates resources. It allows multiple teams, applications or environments to share the same cluster without resource naming conflicts

