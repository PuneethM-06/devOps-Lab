# Interview.md — Day 68

## Helm Fundamentals & Chart Structure

### 1. What is Helm, and why do we use it with Kubernetes?

**Expected points:**

- Helm is a package manager for Kubernetes.
- Helps package, configure, install, upgrade, and manage Kubernetes applications.
- Reduces duplication when managing Kubernetes manifests.
- Supports reusable templates and environment-specific configuration.
- Helps manage application deployments through Helm Releases.

---

### 2. What is a Helm Chart, and what is the difference between a Chart and a Release?

**Expected points:**

- A Helm Chart is a collection of files that describes a set of Kubernetes resources.
- A Chart is the package/template used to deploy an application.
- A Release is an installed instance of a Helm Chart in a Kubernetes cluster.
- The same Chart can be installed multiple times to create different Releases.

**Mental model:**

```text
Chart
  ↓
helm install
  ↓
Release
  ↓
Kubernetes resources