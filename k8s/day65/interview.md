# DAY 65 - KUBERNETES

## HORIZONTAL POD AUTOSCALER
- HPA comes into picture when the traffic to the pods increase leading into more memory utilization 

1. ###  OPTION 1: MANUAL SCALING 
- `kubectl scale deployment backend --replicas=10`
- An engineer notices that there is high CPU and hence manually autoscales the replica count for backend service
- the cons of this we need someone to always monitor the CPU usage intead we can make use of **Autoscaler** 