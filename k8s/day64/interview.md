# Day 64 – Resource Requests, Limits & OOMKilled

## Topics Covered

- Why Resource Requests & Limits are required
- CPU & Memory Units
- Resource Requests
- Resource Limits
- Scheduler and Resource Requests
- CPU Throttling
- Memory Limits
- OOMKilled
- Linux Kernel vs Kubernetes Responsibilities
- Best Practices

---

## Interview Questions

### 1. Why do we need Resource Requests and Limits?

Resource Requests and Limits prevent a single Pod from consuming excessive CPU or memory, ensuring fair resource allocation and maintaining node stability for all applications running on the same worker node.

---

### 2. What is a Resource Request?

A Resource Request specifies the minimum CPU and memory required by a Pod. The Kubernetes Scheduler uses these values to determine whether a worker node has enough available resources before scheduling the Pod.

---

### 3. What is a Resource Limit?

A Resource Limit specifies the maximum amount of CPU and memory a container is allowed to consume while it is running.

---

### 4. What is the difference between Requests and Limits?

| Requests | Limits |
|----------|--------|
| Minimum guaranteed resources | Maximum allowed resources |
| Used by the Scheduler | Enforced at runtime |
| Determines Pod placement | Restricts resource consumption |

---

### 5. Can a Pod use more CPU than its Request?

Yes. A Resource Request only guarantees the minimum amount of CPU or memory for scheduling. A Pod can consume more than its Request until it reaches its configured Limit.

---

### 6. What happens if a Worker Node does not have enough resources to satisfy a Pod's Request?

The Scheduler will not place the Pod on that node. It searches for another suitable node with sufficient available resources.

---

### 7. What does `500m` CPU mean?

CPU is measured in millicores.

- 1000m = 1 CPU Core
- 500m = 0.5 CPU Core
- 250m = 0.25 CPU Core

---

### 8. What do `Mi` and `Gi` represent?

They represent memory units.

- Mi = Mebibyte
- Gi = Gibibyte

Examples:
- 512Mi ≈ 512 MB
- 1Gi ≈ 1 GB

---

### 9. What happens when a container exceeds its CPU Limit?

The Linux kernel throttles the container's CPU usage. The container is not terminated but may run more slowly.

---

### 10. What happens when a container exceeds its Memory Limit?

The Linux kernel's Out Of Memory (OOM) Killer terminates the container. Kubernetes then restarts the container if it is managed by a controller such as a Deployment.

---

### 11. What is OOMKilled?

OOMKilled occurs when a container exceeds its configured memory limit. The Linux kernel's OOM Killer terminates the container to protect the worker node from running out of memory.

---

### 12. Who terminates a container during an OOMKilled event?

The Linux kernel's OOM Killer terminates the container.

---

### 13. Who restarts an OOMKilled container?

Kubelet detects that the container has exited. If the Pod is managed by a Deployment/ReplicaSet, Kubernetes starts a new container to maintain the desired state.

---

### 14. What is CPU Throttling?

CPU Throttling occurs when a container attempts to consume more CPU than its configured limit. Instead of terminating the container, the Linux kernel restricts its CPU usage, causing the application to run more slowly.

---

### 15. Explain the complete Resource Management lifecycle.

1. The developer defines Resource Requests and Limits.
2. The Scheduler uses Requests to choose an appropriate worker node.
3. Kubelet starts the container.
4. The application begins consuming CPU and memory.
5. CPU usage beyond the configured Limit is throttled by the Linux kernel.
6. Memory usage beyond the configured Limit triggers the Linux kernel's OOM Killer.
7. Kubelet detects the container exit and restarts it if required by the Pod's restart policy and its managing controller.

---

### 16. Which Kubernetes component is responsible for each task?

| Component | Responsibility |
|-----------|----------------|
| Scheduler | Selects the worker node based on Resource Requests |
| Kubelet | Starts, monitors and restarts containers |
| Linux Kernel | Enforces CPU and Memory limits |
| OOM Killer | Terminates containers that exceed memory limits |
| Deployment / ReplicaSet | Maintains the desired number of Pods |