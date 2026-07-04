## DOCKER DAY 18

### WHY WAS DOCKER CREATED?
- Docker was created to solve the issue of "Works on my machine".
- When an application was developed using certain tech stacks and versions, and was deployed to prod. Prod resulted in outage because different env had different versions and to over come this we came up with docker, Vm's and containers
- The above process was called as DEPENDENCY HELL

## SOLUTION 1 - VIRTUAL MACHINES
- Virtual machines are complete computers that are inside of an another computer
- An VM contains:
    - OS
    - Application
    - Dependencies
```
Hardware
    ↓
Hypervisor
    ↓
VM1 (Guest OS)
VM2 (Guest OS)
VM3 (Guest OS)
```
## HYPERVISOR
- Hypervisor is a software that can be used to run multiple VM's on a single physical machine

## TYPE 1 - HYPERVISOR
- Runs directly on the hardware
- Example: Xen

## TYPE 2 - HYPERVISOR
- Runs on top of a OS
- Example: Oracle VM VirtualBox

## Limitations of VM
- Lots of memory since it has applications, OS and also dependencies
- Slow start up time 

