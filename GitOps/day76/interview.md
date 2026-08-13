# Interview.md – Day 76

## Must-Know Fundamentals

### 1. Where does Argo CD run, and why?

**Expected topics:**

- Runs inside the Kubernetes cluster
- Usually installed in the `argocd` namespace
- Communicates with the Kubernetes API
- Observes the actual state
- Performs synchronization and reconciliation
- Operates continuously without depending on a local machine

---

### 2. What are the main Argo CD components?

**Expected topics:**

- `argocd-server`
  - Web UI
  - API
  - CLI access

- `argocd-repo-server`
  - Accesses Git repositories
  - Reads application configuration
  - Renders Helm charts/manifests

- `argocd-application-controller`
  - Compares desired vs actual state
  - Detects OutOfSync resources
  - Handles synchronization/reconciliation

---

### 3. How is Argo CD installed?

**Expected topics:**

- Create a dedicated `argocd` namespace
- Apply Argo CD installation manifests
- Kubernetes creates the required resources
- Argo CD components run inside the cluster
- `kubectl` sends requests to the Kubernetes API
- Argo CD is not installed directly as a normal application on the local machine

---

### 4. How do you access Argo CD after installation?

**Expected topics:**

- Argo CD Web UI
- Argo CD CLI
- API
- `kubectl port-forward`
- Example: `8080:443`
- Local port → Argo CD service port inside Kubernetes

---

### 5. Where is the initial Argo CD admin password stored?

**Expected topics:**

- Kubernetes Secret
- `argocd-initial-admin-secret`
- Stored in the `argocd` namespace
- Password is Base64 encoded
- Access controlled using Kubernetes authentication and RBAC

---

### 6. What is an Argo CD `Application`?

**Expected topics:**

- Kubernetes Custom Resource
- Defines what Argo CD should manage
- `source` → where the desired configuration is
- `destination` → where it should be deployed
- Used by Argo CD to manage synchronization

---

### 7. What is the difference between `source` and `destination` in an Argo CD Application?

**Expected topics:**

**Source:**

- Git repository URL
- Target revision/branch
- Path to the application configuration

**Destination:**

- Target Kubernetes cluster
- Target namespace

---

### 8. How does Helm fit into an Argo CD workflow?

**Expected topics:**

- Argo CD can work with Helm charts stored in Git
- Uses `Chart.yaml`, `values.yaml`, and `templates/`
- Helm renders the chart into Kubernetes manifests
- Rendered manifests represent the desired state
- Argo CD compares desired state with actual Kubernetes state

---

### 9. What is the difference between manual sync and automated sync?

**Expected topics:**

- Manual sync → Argo CD detects OutOfSync, human triggers sync
- Automated sync → Argo CD automatically triggers synchronization
- Argo CD performs the actual deployment/synchronization

---

### 10. How does Argo CD access a private Git repository?

**Expected topics:**

- Requires credentials
- HTTPS token or SSH key
- Authentication verifies identity
- Authorization determines repository access
- Argo CD needs permission to read the desired configuration

---

## Important Practical Question

### 11. Explain the complete flow from installing Argo CD to deploying an application.

**Expected flow:**

```text
Install Argo CD
        ↓
Argo CD components run inside Kubernetes
        ↓
Access argocd-server
        ↓
Create Argo CD Application
        ↓
Source → Git repository + revision + path
Destination → Kubernetes cluster + namespace
        ↓
argocd-repo-server reads Git
        ↓
Helm renders manifests (if using Helm)
        ↓
Desired State
        ↓
argocd-application-controller
        ↓
Compare Desired vs Actual State
        ↓
OutOfSync?
        ↓
Manual Sync / Automated Sync
        ↓
Kubernetes Cluster Updated