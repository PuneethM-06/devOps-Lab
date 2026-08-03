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
