# Interview.md – Day 18 (Important Questions Only)

## Must-Know Fundamentals

### 1. What problem was Docker trying to solve?

**Follow-up:** What is the "Works on my machine" problem?

---

### 2. What is the difference between a Virtual Machine and a Container?

**Expected topics:**

* Guest OS and kernel
* Shared kernel
* Resource consumption
* Startup time
* Isolation

---

### 3. Why are containers lightweight and why do they start faster than VMs?

---

### 4. What exactly is a container?

**Expected Answer:**

> A container is an isolated process that packages an application and its dependencies while sharing the host operating system's kernel.

---

### 5. Why can a single server run hundreds of containers but only a limited number of virtual machines?

---

### 6. Why do companies still use Virtual Machines if containers are lightweight?

---

## Shared Kernel & User Space

### 7. Can you run an Ubuntu container on a CentOS host? Why?

---

### 8. Can you run a Windows container natively on a Linux host? Why not?

---

### 9. What does the statement below mean?

> Containers share the host kernel but have their own user space.

---

## Namespaces

### 10. What are Linux namespaces and why are they needed?

---

### 11. Why do containers feel like separate machines even though they share the same kernel?

---

### 12. Containers are just processes. Why can't one container see the processes of another container?

---

### 13. What is a PID namespace?

**Practical scenario:**
Why can't Container A kill processes running in Container B?

---

### 14. What is a Network namespace?

**Practical scenario:**
Why does every container get its own IP address and routing table?

---

### 15. What is the difference between PID namespace and Network namespace?

---

## Cgroups

### 16. What are cgroups and why are they needed?

---

### 17. What resources can cgroups control?

**Expected:**

* CPU
* Memory
* PIDs
* Disk I/O

---

### 18. What is the difference between namespaces and cgroups?

**Expected Answer:**

* Namespaces → What a process can see.
* Cgroups → How much a process can use.

---

### 19. What happens if a container exceeds its memory limit?

---

### 20. What is OOMKilled?

**Follow-up:** Why does it happen?

---

### 21. What is CPU throttling?

**Follow-up:** How is it different from OOMKilled?

---

## Production & Troubleshooting

### 22. A container is repeatedly getting OOMKilled. How would you troubleshoot it?

**Expected approach:**

1. Confirm OOMKilled.
2. Check logs.
3. Check memory limits.
4. Investigate memory usage.
5. Fix the application or increase limits.

---

### 23. A pod is in `CrashLoopBackOff`. What does it mean?

---

### 24. What are some common reasons for `CrashLoopBackOff`?

**Expected:**

* OOMKilled
* Application crash
* Wrong configuration
* Missing environment variables
* Failed dependency connections

---

### 25. A container becomes very slow but does not crash. What could be happening?

**Expected Answer:**
CPU throttling due to cgroup limits.

---

## Most Important Question of Day 18 ⭐

### 26. Walk me through what happens internally when you run:

```bash
docker run nginx
```

**Expected Flow:**

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

---

# If you can answer these 26 questions confidently without notes, your Day 18 fundamentals are interview-ready.
