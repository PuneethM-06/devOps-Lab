# KUBERNETES DAY 64

```
Worker Node
│
├── Pod A (Spring Boot)
├── Pod B (React)
├── Pod C (Redis)
└── Pod D (Prometheus)
```
- All the four pods share the same worker node, meaningf same CPU, MEMORY, DISK and Network
- Considering a situation, if a pod has a bug and starts consuming all the CPU. Other pods cannot serve the users and hence K8s  needs rule 
- Every pod says:
    - This is the **minimum amount of CPU** and memory I need
    - This is the **maximum amount** I am **allowed to use**

### RESOURCE REQUEST
- Please allocate at least this much CPU and memory for me 

### RESOURCE LIMIT
- Do not let me use more than this 

## CPU and MEMORY UNITS
- In K8s, CPU is measured in:
    1. Cores
    2. Milli cores
- 1 CPU = 1000 millicores
```
100m  = 0.1 CPU

250m  = 0.25 CPU

500m  = 0.5 CPU

1000m = 1 CPU

2000m = 2 CPUs
```
