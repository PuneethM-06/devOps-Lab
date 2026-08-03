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

