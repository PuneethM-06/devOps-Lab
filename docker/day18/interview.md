# Interview.md – Day 18: Docker Fundamentals

## MUST-KNOW FUNDAMENTALS

### 1. What problem was Docker created to solve?

**Answer:**

Docker was created to solve the **"Works on my machine"** problem.

An application may work correctly in the developer's environment but fail in testing or production because of differences in:

* OS
* Library versions
* Runtime versions
* Dependencies
* Configuration

Docker packages the application and its dependencies into a consistent environment so that it can run reliably across different environments.

**Follow-up: What is "Works on my machine"?**

> "Works on my machine" means an application works correctly in one environment but fails in another because the environments have different dependencies, configurations, or versions.

This problem is commonly associated with **dependency hell**.

---

### 2. What is the difference between a Virtual Machine and a Container?

**Answer:**

The major difference is that a VM contains a **complete guest operating system**, while a container is an **isolated process that shares the host kernel**.

| Virtual Machine             | Container                             |
| --------------------------- | ------------------------------------- |
| Contains a guest OS         | Does not contain a separate OS kernel |
| Has its own kernel          | Shares the host kernel                |
| Higher resource consumption | Lower resource consumption            |
| Slower startup              | Faster startup                        |
| Stronger isolation          | Process-level isolation               |

**Follow-up: Why does a VM need more resources?**

A VM needs memory and CPU for the guest operating system in addition to the application and its dependencies.

---

### 3. Why are containers lightweight and why do they start faster than VMs?

**Answer:**

Containers are lightweight because they do not need to boot a complete operating system. They share the host operating system's kernel and package only the application and its dependencies.

Containers start faster because Docker mainly needs to:

1. Create namespaces.
2. Apply cgroups.
3. Set up the filesystem.
4. Configure networking.
5. Start the application process.

A VM, on the other hand, needs to boot an entire guest operating system.

---

### 4. What exactly is a container?

**Answer:**

> A container is an isolated process that packages an application and its dependencies while sharing the host operating system's kernel.

A container is **not a mini virtual machine**.

It is simply a process that is isolated using Linux kernel features such as:

* Namespaces
* Cgroups

---

### 5. Why can a single server run hundreds of containers but only a limited number of virtual machines?

**Answer:**

Containers share the host kernel and do not require a separate guest operating system for every application.

VMs require:

* A guest OS
* Memory for the guest OS
* CPU resources for the guest OS
* Application dependencies

Therefore, containers generally consume significantly fewer resources than VMs, allowing a host to run many more containers.

---

### 6. Why do companies still use Virtual Machines if containers are lightweight?

**Answer:**

VMs are still widely used because they provide **stronger isolation** by giving each VM its own guest OS and kernel.

VMs are also useful when different operating systems or kernels are required.

For example:

* Running Windows workloads on a Linux infrastructure.
* Strong isolation between workloads.
* Running applications that require a specific kernel or OS.

---

# SHARED KERNEL & USER SPACE

### 7. Can you run an Ubuntu container on a CentOS host? Why?

**Answer:**

Yes.

A container can have an Ubuntu user space while the host runs CentOS because the container shares the **host kernel**, but it can have its own user-space files, libraries, binaries, and application dependencies.

The important point is:

> The container's user space can be different from the host's user space, but the container still uses the host kernel.

---

### 8. Can you run a Windows container natively on a Linux host? Why not?

**Answer:**

No, not natively.

A Linux host provides a Linux kernel, while Windows containers require a Windows kernel.

Containers share the host kernel, so a Linux kernel cannot natively provide the Windows kernel functionality required by Windows containers.

A VM can be used when a different operating system and kernel are required.

---

### 9. What does this statement mean?

> Containers share the host kernel but have their own user space.

**Answer:**

The **kernel** is shared between the host and containers, but each container can have its own:

* Filesystem
* Libraries
* Binaries
* Application
* Configuration

For example:

```text
Host
└── Linux Kernel
      ├── Container 1 → Ubuntu user space
      ├── Container 2 → Alpine user space
      └── Container 3 → Debian user space
```

The containers have different user spaces but use the same host kernel.

---

# NAMESPACES

### 10. What are Linux namespaces and why are they needed?

**Answer:**

A namespace is a Linux kernel feature that gives a process its **own isolated view of a resource**.

Namespaces are used to isolate containers from each other.

For example:

* PID namespace → Process isolation
* Network namespace → Network isolation
* Mount namespace → Filesystem isolation
* UTS namespace → Hostname isolation
* IPC namespace → Inter-process communication isolation
* User namespace → User and privilege isolation

---

### 11. Why do containers feel like separate machines even though they share the same kernel?

**Answer:**

Containers feel like separate machines because namespaces isolate what each container can see.

A container can have its own:

* Processes
* Network interfaces
* IP address
* Routing table
* Filesystem
* Hostname
* Users

However, underneath all of this, the containers are still using the same host kernel.

---

### 12. Containers are just processes. Why can't one container see the processes of another container?

**Answer:**

Because of **PID namespaces**.

Each container gets its own process namespace, so the container sees only the processes that belong to its namespace.

For example:

```text
Host
├── Container A
│   ├── PID 1
│   ├── PID 2
│   └── PID 3
│
└── Container B
    ├── PID 1
    ├── PID 2
    └── PID 3
```

Container A cannot normally see Container B's processes because they are isolated by PID namespaces.

---

### 13. What is a PID namespace?

**Answer:**

A PID namespace provides **process isolation**.

It allows a container to have its own process ID space.

For example, a container may see:

```bash
ps -ef
```

and see only a few processes even though the host is running hundreds of processes.

**Practical scenario:**

**Why can't Container A kill processes running in Container B?**

Because the processes belong to different PID namespaces, so Container A does not normally have visibility or access to Container B's processes.

---

### 14. What is a Network namespace?

**Answer:**

A network namespace provides **network isolation**.

Each container can have its own:

1. Network interfaces
2. IP address
3. Routing table
4. Port space

**Practical scenario:**

**Why does every container get its own IP address and routing table?**

Because each container can have its own isolated network namespace.

---

### 15. What is the difference between a PID namespace and a Network namespace?

**Answer:**

The difference is **what resource they isolate**.

```text
PID Namespace
      ↓
Process Isolation

Network Namespace
      ↓
Network Isolation
```

PID namespaces isolate processes, while network namespaces isolate networking resources.

---

# CGROUPS

### 16. What are cgroups and why are they needed?

**Answer:**

Cgroups, or **Control Groups**, are a Linux kernel feature used to limit, measure, and isolate the resource usage of a group of processes.

They prevent one container from consuming an uncontrolled amount of system resources.

---

### 17. What resources can cgroups control?

**Answer:**

Common resources controlled by cgroups include:

* CPU
* Memory
* PIDs
* Disk I/O

For example:

```bash
docker run --cpus="1" nginx
```

This limits the container's CPU usage.

Another example:

```bash
docker run -m 512m nginx
```

This limits the container's memory usage to 512 MB.

---

### 18. What is the difference between namespaces and cgroups?

**Answer:**

The simplest way to remember this is:

> **Namespaces → What a process can see.**

> **Cgroups → How much a process can use.**

For example:

```text
NAMESPACES
    ↓
Isolation
    ↓
What can I see?

CGROUPS
    ↓
Resource Control
    ↓
How much can I use?
```

Namespaces provide isolation, while cgroups control resource usage.

---

### 19. What happens if a container exceeds its memory limit?

**Answer:**

If a container exceeds its configured memory limit, the Linux kernel may invoke the **Out-Of-Memory (OOM) killer** and terminate processes to reclaim memory.

In a containerized environment, this can result in the container being terminated and potentially restarted by the container runtime or Kubernetes.

---

### 20. What is OOMKilled?

**Answer:**

**OOMKilled** means **Out Of Memory Killed**.

It happens when a process exceeds the available or configured memory and the kernel kills it to reclaim memory.

**Follow-up: Why does it happen?**

Common reasons include:

* Application memory leaks
* Memory limit is too low
* Application suddenly consumes more memory
* High workload

---

### 21. What is CPU throttling?

**Answer:**

CPU throttling happens when a process is restricted from using more CPU because of its configured cgroup CPU limits or quota.

Instead of killing the process, the kernel limits how much CPU time it can consume.

**Follow-up: How is it different from OOMKilled?**

```text
CPU Throttling
      ↓
Process is slowed down

OOMKilled
      ↓
Process is killed
```

So CPU throttling affects performance, while OOMKilled terminates the process.

---

# PRODUCTION & TROUBLESHOOTING

### 22. A container is repeatedly getting OOMKilled. How would you troubleshoot it?

**Answer:**

I would troubleshoot it step by step:

1. **Confirm that the container is actually being OOMKilled.**
2. **Check the container or pod logs.**
3. **Check the configured memory limits.**
4. **Check the application's actual memory usage.**
5. **Investigate whether there is a memory leak or sudden memory spike.**
6. **Fix the application if necessary or increase the memory limit if the workload genuinely requires more memory.**

The goal is not simply to increase the memory limit. First, I would determine **why the application is consuming more memory than expected**.

---

### 23. A pod is in `CrashLoopBackOff`. What does it mean?

**Answer:**

`CrashLoopBackOff` means the container inside the pod is repeatedly starting, crashing, and being restarted by Kubernetes.

The basic flow is:

```text
Container starts
      ↓
Application crashes
      ↓
Kubernetes restarts it
      ↓
Application crashes again
      ↓
Kubernetes restarts it
      ↓
Backoff increases
```

`CrashLoopBackOff` does **not** tell us the root cause.

It tells us that the container is repeatedly crashing and Kubernetes is backing off before restarting it again.

---

### 24. What are some common reasons for `CrashLoopBackOff`?

**Answer:**

Common reasons include:

* OOMKilled
* Application crash
* Wrong configuration
* Missing environment variables
* Failed dependency connections
* Incorrect application startup command
* Missing files or configuration

The important point is:

> `CrashLoopBackOff` is a symptom, not the root cause.

---

### 25. A container becomes very slow but does not crash. What could be happening?

**Answer:**

One possible reason is **CPU throttling**.

If the container has a CPU limit and tries to consume more CPU than allowed, the cgroup CPU limit can throttle the process.

The container remains running, but the application becomes slow.

```text
High CPU demand
      ↓
CPU limit reached
      ↓
CPU throttling
      ↓
Application becomes slow
```

---

# MOST IMPORTANT QUESTION OF DAY 18 ⭐

### 26. Walk me through what happens internally when you run:

```bash
docker run nginx
```

**Answer:**

At a high level, Docker goes through the following process:

```text
docker run nginx
        ↓
Image Lookup/Pull
        ↓
Create Writable Container Layer
        ↓
Create Namespaces
        ↓
Create Cgroups
        ↓
Configure Networking
        ↓
Mount Filesystem
        ↓
Start nginx Process (PID 1)
        ↓
Container Running
```

### Step-by-step explanation:

**1. Image Lookup/Pull**

Docker checks whether the `nginx` image is available locally.

If it is not available, Docker pulls it from a container registry such as Docker Hub.

**2. Create Writable Container Layer**

Docker creates the container's writable layer on top of the image's read-only layers.

**3. Create Namespaces**

Docker creates namespaces to isolate the container's:

* Processes
* Networking
* Filesystem
* Hostname
* Users
* IPC

**4. Create Cgroups**

Docker applies cgroups to control resources such as:

* CPU
* Memory
* PIDs
* Disk I/O

**5. Configure Networking**

Docker sets up the container's network interfaces, IP address, routing, and connectivity according to the configured Docker network.

**6. Mount Filesystem**

Docker mounts the container's filesystem so that the application can access its required files, binaries, libraries, and directories.

**7. Start nginx Process**

Docker starts the nginx process inside the configured namespaces and cgroups.

The main process becomes **PID 1 inside the container**.

**8. Container Running**

The container remains running as long as its main process is running.

---

# QUICK REVISION

## NAMESPACES vs CGROUPS

```text
NAMESPACES
    ↓
Isolation
    ↓
What can the process SEE?

CGROUPS
    ↓
Resource Control
    ↓
How much can the process USE?
```

## VM vs CONTAINER

```text
VM
    ↓
Guest OS
    ↓
Application
    ↓
More resources
    ↓
Slower startup
    ↓
Stronger isolation

CONTAINER
    ↓
Application + Dependencies
    ↓
Shared Host Kernel
    ↓
Fewer resources
    ↓
Faster startup
    ↓
Process-level isolation
```

## OOMKilled vs CPU Throttling

```text
OOMKilled
    ↓
Memory limit exceeded
    ↓
Process killed

CPU Throttling
    ↓
CPU limit reached
    ↓
Process slowed down
```

## CRASHLOOPBACKOFF

```text
Container starts
      ↓
Application crashes
      ↓
Kubernetes restarts it
      ↓
Application crashes again
      ↓
Kubernetes backs off
      ↓
CrashLoopBackOff
```

---

# INTERVIEW-READY CHECKLIST

If you can answer these questions confidently **without looking at the answers**, you should understand the core Day 18 Docker fundamentals:

* [ ] Explain the "Works on my machine" problem.
* [ ] Explain Dependency Hell.
* [ ] Explain VM vs Container.
* [ ] Explain why containers are lightweight.
* [ ] Explain why containers start faster.
* [ ] Define a container.
* [ ] Explain the shared kernel.
* [ ] Explain user space.
* [ ] Explain why Ubuntu containers can run on a CentOS host.
* [ ] Explain why Windows containers cannot run natively on a Linux kernel.
* [ ] Explain namespaces.
* [ ] Explain PID namespaces.
* [ ] Explain Network namespaces.
* [ ] Explain cgroups.
* [ ] Explain Namespaces vs Cgroups.
* [ ] Explain OOMKilled.
* [ ] Explain CPU throttling.
* [ ] Troubleshoot OOMKilled.
* [ ] Explain CrashLoopBackOff.
* [ ] Troubleshoot common CrashLoopBackOff causes.
* [ ] Explain what happens when a container is slow due to CPU throttling.
* [ ] Explain the internal flow of `docker run nginx`.

## FINAL GOAL ⭐

> If you can explain these concepts in your own words and walk through `docker run nginx` without looking at your notes, your Day 18 Docker fundamentals are interview-ready.
