# DAY 68 - HELM FUNDAMENTALS AND CHART STRUCTURE 

### WHAT IS HELM?
- Helm is a **package manager for Kubernetes**
- Helm helps you package, configure, install, upgrade and manage kuberenetes application

- Example:
- Suppose we have an application that needs to run in 3 environments say dev, stage and prod. Now, dev needs `replicas: 3` while prod needs `replicas: 5`. Without helm, we end up creating nearly same identical yaml files and here is where helm kicks in.

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

### WHAT IS A HELM CHART?
- A helm chart is a **collection of files that describe set of kuberenetes resources** 

### WHAT IS A HELM RELEASE
- It is an **installed helm chart instance running in k8s** 

### WHAT IS A HELM REPOSITORY?
- A helm repository is a location where helm charts are stored 

### Chart.yaml
- A `chart.yaml` contains the metadata about the helm chart
- Example
```
apiVersion: v2
name: myapp
description: A Helm chart for my application
type: application
version: 0.1.0
appVersion: "1.0.0"
```
### values.yaml
- One of the most important files in yaml
- It contains the **default configuration values** used by the Helm templates
- Example:
```
replicaCount: 2

image:
  repository: nginx
  tag: "1.27"
```
- Instead of harcoding: `replicas: 2` we can do `replicas: {{ .Values.replicaCount }}`

### templates/
- It contains **k8s manifest templates** which will be used by **Helm to process and generate final k8s manifest templates**

### charts/
- Contains dependencies also called as subcharts 

### helmignore
- It is similar to `.gitignore`
- It includes files that must not be included at the time of packaging 

### Everything put together 
```
myapp/
│
├── Chart.yaml
│       ↓
│   Chart metadata
│
├── values.yaml
│       ↓
│   Default configuration
│
├── templates/
│       ↓
│   Kubernetes templates
│
├── charts/
│       ↓
│   Dependencies
│
└── .helmignore
        ↓
    Files to exclude
```

### WHAT HAPPENS WHEN HELM INSTALLS A CHART
- suppose we have 
```
myapp/
├── Chart.yaml
├── values.yaml
└── templates/
    ├── deployment.yaml
    └── service.yaml
```
- And values.yml contains `replicaCount: 3`
- And a template uses something like: `replicas: {{ .Values.replicaCount }}`
- And when we run `helm install myapp ./myapp`
- conceptually this happens
```
                  Helm Chart
                      │
          ┌───────────┴───────────┐
          ↓                       ↓
     values.yaml             templates/
          │                       │
          └───────────┬───────────┘
                      ↓
                Helm renders
                      ↓
          Kubernetes manifests
                      ↓
             Kubernetes API
                      ↓
             Resources created
                      ↓
                Helm Release
```
- Helm takes the `values` and the `templates` together and uses them to produces `Kubernetes manifests`.
- These manifests are sent to k8s to produce the actual resources.
- Helm keeps track of these deployments using Release 
