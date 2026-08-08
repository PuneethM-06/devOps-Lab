## DOCKER DAY 18

### WHY WAS DOCKER CREATED?

* Docker was created to solve the issue of "Works on my machine".
* When an application was developed using certain tech stacks and versions, and was deployed to prod, prod could result in an outage because different environments had different versions. To overcome this, we came up with Docker, VMs, and containers.
* The above process was called **DEPENDENCY HELL**.

## SOLUTION 1 - VIRTUAL MACHINES

* Virtual machines are complete computers that are inside another computer.
* A VM contains:

  * OS
  * Application
  * Dependencies

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

* A hypervisor is software that can be used to run multiple VMs on a single physical machine.

## TYPE 1 - HYPERVISOR

* Runs directly on the hardware.
* Example: Xen

## TYPE 2 - HYPERVISOR

* Runs on top of an OS.
* Example: Oracle VM VirtualBox

## LIMITATIONS OF VM

* Requires lots of memory since it has applications, an OS, and dependencies.
* Slow startup time.

## SOLUTION 2 - CONTAINERS

* A container is an isolated process running on a host operating system.

* A container is not a VM; it is not a mini OS. It is an isolated process that is running on a host OS.

* Containers ensure that each process runs under the same kernel.

* Example: Virtual machines

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

* Example: Containers

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

* VMs make use of separate kernels, which are expensive, while containers, on the other hand, ensure that they are run under the same kernel.

* Containers share the host kernel but have their own user space.

## NAMESPACES

* A namespace is a Linux kernel feature that provides a process with its **own isolated view of a resource**.
* So the process still shares the host OS, but it sees what the namespace allows it to see.
* Without namespaces, all processes can see each other.
* With namespaces, Container A cannot see Container B processes.

## TYPES OF NAMESPACES

1. ### PID NAMESPACES (PROCESS ISOLATION)

* Suppose we run `ps -ef` and we get 10 processes running as output.
* If I run the same command inside a container, then I get 3 processes.
* So in reality, the container thinks it has only 3 processes, but there are actually more which cannot be seen by it, and this is called **PROCESS ISOLATION**.
* This is useful because one container cannot kill or disturb the other container processes.

2. ### NET NAMESPACE - NETWORK ISOLATION

* Every container gets its own

  1. IP address
  2. Network interfaces
  3. Routing tables
  4. Port spaces
* Container A cannot connect to Container B.

3. ### MNT NAMESPACES - MOUNT/FILESYSTEM NAMESPACES

* Every container thinks it has its own filesystem.
* Example: Inside container

```
/
├── app
├── bin
├── etc
└── tmp
```

* Example: Host

```
/home
/var
/opt
/etc
```

* Each container can see the mount of the filesystem of its own.

4. ### UTS NAMESPACE - HOSTNAME ISOLATION

* Each container has its own hostname.

5. ### IPC NAMESPACE - INTER-PROCESS COMMUNICATION

* Processes communicate through:

  * Shared memory
  * Message Queues
  * Semaphores

6. ### USER NAMESPACES

* `root` inside a container is not the same as `root` on a host.
* This is used for providing security to a container.

## C GROUPS

* C group stands for **CONTROL GROUP**.
* C group is a kernel feature where it limits, measures, and isolates resource usage of a group of processes.

## NOTE:

1. **NAMESPACES MAKE SURE WHAT A PROCESS CAN SEE AND COMMUNICATE.**
2. **CGROUPS MAKE SURE WHAT RESOURCES A PROCESS CAN MAKE USE OF.**

## RESOURCES CONTROLLED BY CGROUPS

1. ### CPU

* This limits the CPU usage.

```
docker run --cpus="1" nginx
```

Container can use only 1 CPU core.

2. ### MEMORY

* Limits memory usage.

```
docker run -m 512m nginx
```

3. ### NETWORK

* Can also control network usage.

4. ### PIDs

* Can limit the number of processes.

## OOMKILLED

* OOMKILLED - Out of memory killed.
* This is one of the common production incidents where the kernel kills a container because it was using more than the limited memory.

## CPU THROTTLING

* This is the situation where the kernel doesn't kill a process; instead, it slows it down.

## WHAT IS CRASHLOOPBACKOFF?

* It is a situation where:

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

* This doesn't tell us the root cause, but it tells us there is some problem because the pods are getting killed and Kubernetes is starting them again.

## REASONS WHY COMPANIES STILL USE VMS

1. Stronger isolation since they have their own kernels.
2. We cannot run Windows containers natively on Linux in containers, but we can do it in a VM.

## NOTE:

Containers are lightweight because they do not include a separate guest operating system or kernel. Instead, they package only the application and its dependencies while sharing the host operating system's kernel. This reduces memory and storage consumption significantly.

Containers start faster because they do not need to boot an entire operating system. Starting a container mainly involves creating namespaces, applying cgroups, setting up the filesystem and networking, and starting the application process.

## OVER ALL FLOW

```
docker run nginx
        ↓
Image Lookup/Pull
        ↓
Container Creation
        ↓
Namespaces Creation
        ↓
Cgroups Creation
        ↓
Networking Setup
        ↓
Filesystem Mount
        ↓
Start nginx Process (PID 1)
        ↓
Container Running
```
