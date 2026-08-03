# DAY 65 - KUBERNETES

## HORIZONTAL POD AUTOSCALER
- HPA comes into picture when the traffic to the pods increase leading into more memory utilization
- HPA is a Kubernetes controller that automatically increase or decrease the numbe of pod replicas based on observed metrics such as CPU utilization, enabling applications to server users without any lag

1. ###  OPTION 1: MANUAL SCALING 
- `kubectl scale deployment backend --replicas=10`
- An engineer notices that there is high CPU and hence manually autoscales the replica count for backend service
- the cons of this we need someone to always monitor the CPU usage intead we can make use of **Autoscaler** 

## HORIZONTAL SCALING
- Increase the number of pods
