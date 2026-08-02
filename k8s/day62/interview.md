# Day 62 – Kubernetes Liveness & Readiness Probes

## Topics Covered
- Why Kubernetes Probes are required
- Pod Health vs Application Health
- Kubelet and Health Checks
- Liveness Probe
- Readiness Probe
- Liveness vs Readiness
- HTTP, TCP and Exec Probes
- Probe Configuration
- Production Scenarios
- Best Practices

---

## Interview Questions

### 1. Why do we need Kubernetes Probes?

A Pod being in the `Running` state only means the container process is running. It does not guarantee that the application inside the container is healthy or capable of serving requests. Kubernetes uses probes to monitor the application's health and readiness.

---

### 2. Which Kubernetes component performs health checks?

The **Kubelet** running on each worker node performs health checks by periodically executing the configured probes.

---

### 3. What is a Liveness Probe?

A Liveness Probe is a health check performed by the Kubelet to determine whether an application is still alive. If the probe fails repeatedly, Kubernetes restarts the container automatically.

---

### 4. When should a Liveness Probe be used?

Liveness Probes should detect application failures that can be resolved by restarting the container, such as:
- Deadlocks
- Infinite loops
- Hung or unresponsive applications
- Memory-related issues causing the application to stop responding

---

### 5. What is a Readiness Probe?

A Readiness Probe checks whether an application is ready to accept user traffic. If the probe fails, Kubernetes removes the Pod from the Service endpoints but does not restart the container.

---

### 6. When should a Readiness Probe fail?

Examples include:
- Application startup is still in progress
- Database is unavailable
- Redis or external dependency is unavailable
- Cache initialization is still running

---

### 7. Difference between Liveness and Readiness Probes

| Liveness Probe | Readiness Probe |
|----------------|-----------------|
| Checks if the application is alive | Checks if the application is ready to serve traffic |
| Failure restarts the container | Failure removes the Pod from Service endpoints |
| Used for hung or crashed applications | Used for startup delays or temporary dependency failures |

---

### 8. Can both Liveness and Readiness fail at the same time?

Yes. If an application becomes completely unresponsive, the Readiness Probe removes the Pod from the Service while the Liveness Probe eventually restarts the container.

---

### 9. What are the different types of probes?

- HTTP Probe – Checks an HTTP endpoint (most common for web applications).
- TCP Probe – Verifies whether a TCP port accepts connections.
- Exec Probe – Executes a command inside the container.

---

### 10. Which probe would you use for a Spring Boot application?

HTTP Probe using an endpoint such as:

`/actuator/health`

---

### 11. Which probe would you use for PostgreSQL?

TCP Probe because PostgreSQL exposes a TCP port but not an HTTP health endpoint.

---

### 12. Which probe would you use for a legacy application that creates a health file?

Exec Probe, since Kubernetes can execute a command to verify the existence or contents of the file.

---

### 13. What is the purpose of initialDelaySeconds?

It delays the first health check after the container starts, giving the application enough time to initialize. This prevents slow-starting applications from failing health checks prematurely.

---

### 14. What happens if initialDelaySeconds is configured incorrectly?

If the application requires more startup time than configured, Kubernetes may repeatedly fail the health checks and continuously restart the container, potentially causing a CrashLoopBackOff.

---

### 15. Why shouldn't Liveness and Readiness always perform the same function?

Because they answer different questions:
- Liveness determines whether the application should be restarted.
- Readiness determines whether the application should receive user traffic.