# Interview.md – Day 75

## Must-Know Fundamentals

### 1. What is GitOps?

**Expected topics:**

- Git as the source of truth
- Desired state
- Actual state
- Argo CD / GitOps controller
- Reconciliation

---

### 2. What problem does GitOps solve?

**Expected topics:**

- Configuration drift
- Manual changes to Kubernetes
- Git and cluster becoming inconsistent
- Auditability and traceability
- Maintaining the desired state

---

### 3. What is the difference between desired state and actual state?

**Expected topics:**

- Desired state → defined in Git
- Actual state → currently running in Kubernetes
- Git acts as the source of truth

---

### 4. What is configuration drift?

**Expected topics:**

- Actual state differs from desired state
- Manual `kubectl` changes
- Application becomes OutOfSync
- Argo CD detects the difference

---

### 5. What is reconciliation in GitOps?

**Expected topics:**

- Compare desired vs actual state
- Detect differences
- Synchronize Kubernetes resources
- Actual state should match desired state

---

### 6. How is traditional CI/CD different from GitOps?

**Expected topics:**

- Traditional CI/CD pipeline directly deploys to Kubernetes
- `kubectl apply` / `helm upgrade`
- Pipeline finishes after deployment
- GitOps continuously compares desired and actual state
- Argo CD manages synchronization

---

### 7. What is the difference between push-based and pull-based deployment?

**Expected topics:**

- Push → CI/CD pipeline pushes changes to Kubernetes
- Pull → Argo CD reads/pulls desired configuration from Git
- Who initiates the deployment
- Separation of CI and deployment responsibilities

---

### 8. What is the role of Argo CD in GitOps?

**Expected topics:**

- GitOps controller
- Reads desired state from Git
- Compares desired vs actual state
- Detects OutOfSync resources
- Sync and reconciliation

---

### 9. Where does Helm fit into a GitOps workflow?

**Expected topics:**

- Helm as a packaging and templating tool
- Helm chart and values stored in Git
- Argo CD can render Helm charts
- Generated manifests represent desired resources
- Argo CD compares and synchronizes them with Kubernetes

---

## Important Practical Question

### 10. Explain the complete CI + GitOps + Argo CD workflow.

**Expected flow:**

```text
Developer
    ↓
Push application code
    ↓
CI Pipeline
    ↓
Test → Build → Push Docker Image
    ↓
Update Helm/Kubernetes configuration in Git
    ↓
Argo CD
    ↓
Compare Desired State vs Actual State
    ↓
Sync / Reconciliation
    ↓
Kubernetes Cluster