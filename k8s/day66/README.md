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

## WHAT EXACTLY IS A NAMESPACE?
1. **HOW DO WE CREATE NAMESPACES**
- `kubectl create namespace dev`

2. **TO VIEW ALL NAMESPACES**
- `kubectl get namespaces`

3. **DEPLOYMENT INTO DEFAULT NAMESPACE**
- `kubectl apply -f deployment.yaml`

4. **DEPLOYMENT INTO DEV NAMESPACE**
- `kubectl apply -f deployment.yaml -n dev`

> CAN TWO DEPLOYMENTS HAVE THE SAME NAMESPACES?
- Yes, provided they are in different namesapces. Resource names must be unique within a namespace but the same name can exist in multiple namespaces

> CAN KUBERNETES NAMESPACE BE NESTED?
- No, Kubernetes namespaces are flat and cannot be nested. 

##  BUILT IN NAMESPACES
- When a brand new cluster is creayed we can see:
    1. default 
    2. Kube-system 
    3. kube-public
    4. kube-node-lease

1. ### default
- This is the easiest one
- when a deployment is done without explicitly defining a namespace then it is created under `default`

2. ### kube-system 
- This is one of the most important built-in namespace
- This can be thought as the **operating system folder of the kubernetes cluster**
- such as:
    1. coreDNS
    2. kube-proxy
    3. metric server
    4. CNI plugins 

3. ### Kube-public
- This namespace is intendede for resources that are publicly readable across the cluster
- It is rarely used by application developers

4. ### kube-node-lease
- Every worker node periodically says control plane that it is alive and it called as **heartbeat**
- Kubernetes stores these hearbeat **lease object in kube-node-lease**