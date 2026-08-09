# DAY 68 - HELM FUNDAMENTALS AND CHART STRUCTURE 

### WHAT IS HELM?
- Helm is a **package manager for Kubernetes**
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

### WHAT IS A HELM CHART?
- A helm chart is a **collection of files that describe set of kuberenetes resources** 

### WHAT IS A HELM RELEASE
- It is a **running instance of helm in kuberenetes** 

### WHAT IS A HELM REPOSITORY?
- A helm repository is a location where helm charts are stored 

### Chart.yml
- A `chart.yml` contains the metadata about the helm chart
- Example
```
apiVersion: v2
name: myapp
description: A Helm chart for my application
type: application
version: 0.1.0
appVersion: "1.0.0"
```
### VALUES.YML
- One of the most important files in yml
- It contains the **default configuration values** used by the Helm templates
- Example:
```
replicaCount: 2

image:
  repository: nginx
  tag: "1.27"
```
- Instead of harcoding: `replicas: 2` we can do `replicas: {{ values.replicaCount }}`

### templates/
- It contains **k8s manifest templates** which will be used by **Helm to process and generate final k8s manifest templates**

### charts/
- Contains dependencies also called as subcharts 
- 

### helmignore
- It is similar to `.gitignore`
- It includes files that must not be included at the time of packaging 
