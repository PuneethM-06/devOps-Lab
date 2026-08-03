# DAY 65 - KUBERNETES

## HORIZONTAL POD AUTOSCALER
- HPA comes into picture when the traffic to the pods increase leading into more memory utilization
- HPA is a Kubernetes controller that automatically increase or decrease the numbe of pod replicas based on observed metrics such as CPU utilization, enabling applications to server users without any lag

1. ###  OPTION 1: MANUAL SCALING 
- `kubectl scale deployment backend --replicas=10`
- An engineer notices that there is high CPU and hence manually autoscales the replica count for backend service
- the cons of this we need someone to always monitor the CPU usage intead we can make use of **Autoscaler** 

###  HORIZONTAL SCALING
- Increase the number of pods

###  VERTICAL SCALING
- Increase the capacity of existing pods

## METRICS SERVER
- Metric server gives the intimation to the HPA saying, pods CPU is more and now we need to do scaling 
- The overall flow is:
```
Pod

↓

Kubelet

↓

Metrics Server

↓

Horizontal Pod Autoscaler
```
- **Metric server periodically collects metrics from kubelet**
```
Application Running

↓

CPU Usage Increases

↓

Linux Kernel

↓

Kubelet Reads Metrics

↓

Metrics Server Collects Metrics

↓

HPA Reads Metrics

↓

Deployment Replica Count Increased

↓

ReplicaSet Creates New Pods
```
**METRIC SERVER** - Metrics Server is a Kubernetes component that collects CPU and memory usage metrics from the Kubelet running on each Worker Node. Horizontal Pod Autoscaler uses these metrics to determine whether Pods should be scaled up or down.\

## WHY HPA NEEDS CPU REQUESTS 
- CPU UTILIZATION = CURRENT CPU USAGE / CPU REQUEST
- Example:
```
Request = 500m

Current = 750m

750 / 500 = 150%
targetCPUtilization = 70%

HPA says pods are overloaded scale immediately 
```
```
Application Running

↓

Current CPU = 750m

↓

CPU Request = 500m

↓

Metrics Server

↓

HPA calculates

750 / 500 = 150%

↓

Target = 70%

↓

Scale Up
```

## HOW HPA ACTUALLY SCALES
- HPA compares the metric server repor to what is the target 
- It says `depoyment` increase your `replica count`
- Then creates new pods 

## END TO END FLOW - IMPORTANT 
```
High CPU
     │
     ▼
Linux Kernel
     │
     ▼
Kubelet reads CPU usage
     │
     ▼
Metrics Server collects metrics
     │
     ▼
HPA compares CPU utilization with target
     │
     ▼
HPA updates Deployment replicas
     │
     ▼
Deployment Controller updates ReplicaSet
     │
     ▼
ReplicaSet creates new Pending Pods
     │
     ▼
Scheduler assigns Worker Nodes
     │
     ▼
Kubelet starts containers
     │
     ▼
Pods become Ready
     │
     ▼
Service begins sending traffic
```
## SCALE DOWN 
- Once the current gets down to the target scaling down will happen
- HPA would not react immediately after the CPU is down because CPU may vary instead it lets CPU average to be down under the target for some time called as **stabilization window**
