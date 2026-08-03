# Day 63 – Kubernetes Rolling Updates & Rollbacks

## Topics Covered
- Why Rolling Updates are required
- Rolling Update workflow
- Deployment, ReplicaSet and Pods
- Update Strategies
- maxSurge
- maxUnavailable
- Readiness Probe during deployments
- Rollbacks
- Rollout History
- Best Practices

---

## Interview Questions

### 1. What is a Rolling Update?

A Rolling Update is a deployment strategy in Kubernetes where old Pods are gradually replaced with new Pods while ensuring the application remains available, resulting in zero or minimal downtime.

---

### 2. Why do we need Rolling Updates?

Deleting all existing Pods before deploying a new version causes downtime because there are no Pods available to serve user requests. Rolling Updates gradually replace Pods while keeping the application available.

---

### 3. How does Kubernetes perform a Rolling Update?

When a Deployment is updated, Kubernetes creates a new ReplicaSet instead of modifying the existing one. The new ReplicaSet gradually creates Pods while the old ReplicaSet is scaled down one Pod at a time until the deployment is complete.

---

### 4. Does Kubernetes modify an existing ReplicaSet during an update?

No. Kubernetes creates a new ReplicaSet for the updated Pod template. The old ReplicaSet is retained for rollback purposes.

---

### 5. What is maxSurge?

maxSurge defines the maximum number of additional Pods Kubernetes can temporarily create above the desired replica count during a Rolling Update.

Example:
- Replicas = 5
- maxSurge = 2

Maximum Pods during deployment = 7

---

### 6. What is maxUnavailable?

maxUnavailable defines the maximum number of Pods that can be unavailable during a Rolling Update while maintaining application availability.

Example:
- Replicas = 5
- maxUnavailable = 1

Minimum available Pods = 4

---

### 7. Why are Readiness Probes important during Rolling Updates?

Readiness Probes ensure that a newly created Pod is fully initialized and capable of serving user requests before Kubernetes removes an old Pod. This enables zero-downtime deployments.

---

### 8. Explain the relationship between Deployment, ReplicaSet and Pods.

- Deployment manages application updates and rollbacks.
- Deployment creates and manages ReplicaSets.
- ReplicaSets maintain the desired number of Pods.
- Pods run the actual application containers.

---

### 9. What is a Rollback?

A Rollback is the process of reverting a Deployment to a previously working version when the current deployment introduces issues. Kubernetes performs this by scaling down the current ReplicaSet and scaling up the previous ReplicaSet.

---

### 10. Why does Kubernetes retain old ReplicaSets?

Old ReplicaSets are retained so Kubernetes can quickly roll back to a previous stable version without recreating the entire Deployment configuration.

---

### 11. How do you check the rollout status of a Deployment?

```bash
kubectl rollout status deployment <deployment-name>