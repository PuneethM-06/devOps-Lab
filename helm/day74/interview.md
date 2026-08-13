# Interview.md – Day 77

## Must-Know Fundamentals

### 1. What does Git as the Source of Truth mean?

**Expected topics:**

- Git contains the desired state
- Kubernetes contains the actual state
- Argo CD compares both
- Argo CD reconciles differences

---

### 2. What is Desired State vs Actual State?

**Expected topics:**

- Desired State → Defined in Git
- Actual State → Currently running in Kubernetes
- Match → Synced
- Difference → OutOfSync

---

### 3. What is the difference between Manual Sync and Auto-Sync?

**Expected topics:**

- Manual Sync → Human triggers synchronization
- Auto-Sync → Argo CD triggers synchronization automatically
- Reconciliation brings Kubernetes toward the desired state

---

### 4. How does Argo CD detect Git changes?

**Expected topics:**

- Periodic refresh/checking
- Webhooks
- Webhooks notify Argo CD
- Webhooks do not directly deploy the application

---

### 5. What is reconciliation?

**Expected topics:**

- Compare Desired vs Actual State
- Detect differences
- Apply required changes
- Bring Kubernetes toward the desired state

---

### 6. What is Configuration Drift?

**Expected topics:**

- Kubernetes changed outside Git
- Actual state differs from desired state
- Example: Manual kubectl changes
- Git remains unchanged

---

### 7. What is Self-Healing?

**Expected topics:**

- Fixes configuration drift
- Detects changes made directly in Kubernetes
- Restores the Kubernetes state to match Git
- selfHeal: true

---

### 8. What is Pruning?

**Expected topics:**

- Resource removed from Git
- Resource may still exist in Kubernetes
- prune: true
- Argo CD removes the managed resource from Kubernetes

---

### 9. Explain Auto-Sync, Self-Healing, and Pruning.

**Expected topics:**

- Git changes → Auto-Sync
- Kubernetes manually changed → Self-Healing
- Resource removed from Git → Pruning

---

### 10. What is the role of argocd-repo-server?

**Expected topics:**

- Fetches source from Git
- Reads application configuration
- Generates desired manifests
- Works with YAML, Helm, and Kustomize

---

### 11. What is the role of argocd-application-controller?

**Expected topics:**

- Compares Desired vs Actual State
- Detects OutOfSync resources
- Handles synchronization
- Performs reconciliation

---

## Important Practical Question

### 12. Explain the complete CI/CD and GitOps flow using Argo CD.

**Expected flow:**

Developer changes code  
↓  
Pull Request  
↓  
CI runs tests and validation  
↓  
PR merged  
↓  
Docker image built  
↓  
Image pushed to Container Registry  
↓  
GitOps repository updated with new image tag  
↓  
Git push  
↓  
Argo CD detects the change  
↓  
argocd-repo-server generates desired manifests  
↓  
argocd-application-controller compares Desired vs Actual State  
↓  
OutOfSync  
↓  
Manual Sync / Auto-Sync  
↓  
Argo CD reconciles Kubernetes  
↓  
Kubernetes performs rollout  
↓  
New version running