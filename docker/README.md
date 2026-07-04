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

## SOLUTION 2 -  CONTAINERS
- A container is an isolated process running on a host operating system 
- A container is not a VM, it is not a mini OS. It is a isolated process that is running on a host OS
- Containers ensure that each processes are ran under the same kernel.

- Example: Virtual machines
```
VM 1
├── Application
└── Ubuntu Kernel

VM 2
├── Application
└── Ubuntu Kernel

VM 3
├── Application
└── Ubuntu Kernel
```
- Example: Containers
```
Container 1
├── Application

Container 2
├── Application

Container 3
├── Application

↓
Single Host Kernel
```
- VM make use of seperate kernel which are expensive while containers on the other hand ensure that they are run under the same kernel

- Containers share the host kernel but have their own user space.

## NAMESPACES

- A namespace is a linux kernel feature that provides a process it's **own isolated view of a resource**
- So the process still shares the host OD, but it sees what the namespace allows it to see 
- without namespaces, all process can see each other.
- with namespaces, Container A cannot see Container B process

## TYPER OF NAMESPACES

1. ### PID NAME SPACES (PROCESS ISOLATION)
- Suppose we run `ps -ef`and we get 10 process running as output and then 
- If i run the same command inside a container, then i get 3 process.
- So in reality, container thinks it has only 3 processes but there are actually more which cannot be seen by it and is called as a PROCESS ISOLATION.
- This is useful because 1 contanier cannot kill or disturb the other container processes.


2. ### NET NAMESPACE - NETWORK ISOLATION 
- Every container gets its own
    1. IP address
    2. Network interfaces
    3. Routing tables
    4. Port spaces
- Container A cannot connect to Container B

3. ### MNT NAMESPACES - MOUNT/FILESYSTEM NAMESPACES
- Every container thinks it has its own filesystem
- Example: Inside container
```
/
├── app
├── bin
├── etc
└── tmp
```
- Example: Host
```
/home
/var
/opt
/etc
```
Each container can see the mount of file system of its own 

4. ### UTS NAMESPACE - HOSTNAME ISOLATION
- Each container has its own host name 

5. ### IPC NAMESPACE - INTER-PROCESS COMMUNICATION 
- Processes communicate through:
    - Shared memory
    - Messages Queues
    - Semaphores
6. ### USER NAME SPACES
- `root` inside container is not the same of `root` in a host
- This is used for providing security to a container 


## C GROUPS
- C group stands for **CONTROL GROUP**
- C group is a kernel feature where it limits, measures and isolates resource usage of a group of processes.

## NOTE:
1. **NAMESPACE MAKE SURE WHAT A PROCESS CAN SEE AND COMMUNICATE.**
2. **CGROUPS MAKE SURE IN WHAT A PROCESS CAN MAKE USE OF**

## RESOURCES CONTROLLED BY CGROUPS
1. ### CPU
- This limits the CPU usag
```
docker run --cpus="1" nginx
```
Container can use only 1 CPU core

2. ### MEMORY
- Limit memory usage
```
docker run -m 512m nginx
```
3. ### NETWORK
- Can also control network usage 

4. ### PIDs
- Can limit number of proceeses

## OOMKILLED
- OOMKILLED -  Out of memory killed
- This is one of the common production incident where kernel kills a container becuase it was using more than the limited memory 

## CPU THROTTLING 
- This is the situation where kernel doesnt kill a process instead it slows it down 

## WHAT IS CRASHLOOPBACKOFF?
- It is a situation where /
```
Start Container
      ↓
Application crashes
      ↓
Kubernetes restarts it
      ↓
Application crashes again
      ↓
Kubernetes restarts it
      ↓
Application crashes again
```
- This doesnt tell us the root cause but it tells us there is some problem because of which the pods are getting killed and k8s is starting it again 

