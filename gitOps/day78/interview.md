# Interview.md – Day 79

## Argo CD App of Apps Pattern

### 1. What is the App of Apps pattern in Argo CD?

The App of Apps pattern is used to manage multiple Argo CD Applications in a centralized way.

A Parent Application manages multiple Child Applications, and each Child Application manages its own Kubernetes resources.

Argo CD
↓
Parent Application
↓
Child Applications
↓
Kubernetes Resources

---

### 2. Why do we use the App of Apps pattern?

In production environments, we may have multiple microservices and therefore multiple Argo CD Applications.

Managing every Application separately can become difficult.

The App of Apps pattern improves scalability and manageability by allowing multiple Argo CD Applications to be managed centrally through Git.

---

### 3. What is a Parent Application?

A Parent Application is a normal Argo CD Application resource.

It points to a Git repository path containing Child Application YAML definitions.

The Parent Application manages those Child Application resources.

Parent Application
↓
Child Application YAMLs
↓
Child Applications

---

### 4. What is a Child Application?

A Child Application is also an Argo CD Application resource.

Each Child Application points to its own source path containing Kubernetes manifests or a Helm chart.

The Child Application then manages Kubernetes resources such as:

- Deployment
- Service
- ConfigMap
- Secret
- Ingress

---

### 5. What is root-app.yaml?

root-app.yaml is the YAML file that defines the Parent Application.

When it is applied to Kubernetes, it creates the Parent Application.

root-app.yaml
↓
Creates Parent Application
↓
Parent Application manages Child Applications

---

### 6. Explain the complete App of Apps flow.

The complete flow is:

root-app.yaml
↓
Creates Parent Application
↓
Parent Application points to a Git repository
↓
Reads Child Application YAML files
↓
Child Application resources are created in Kubernetes
↓
Each Child Application points to its own manifests
↓
Kubernetes resources are created and managed

---

### 7. What does the Parent Application manage?

The Parent Application manages Child Argo CD Application resources.

It does not directly manage application resources such as Deployments or Services.

Parent Application
↓
Child Applications

---

### 8. What does a Child Application manage?

Each Child Application manages the Kubernetes resources defined in its configured source path.

For example:

Child Application
↓
Deployment
Service
ConfigMap

---

### 9. What happens when a new Child Application YAML is added?

When a new Child Application YAML is added to the Git path managed by the Parent Application, Argo CD can detect and sync the new definition.

The new Child Application is created, and it then manages the Kubernetes resources defined in its own source configuration.

---

### 10. What is the difference between root-app.yaml and the Parent Application?

root-app.yaml is a YAML file that defines the Parent Application.

The Parent Application is the actual Argo CD Application resource created in Kubernetes.

root-app.yaml
↓
Defines and creates
↓
Parent Application
↓
Manages
↓
Child Applications

---

### 11. What are the advantages of the App of Apps pattern?

- Centralized management of multiple Argo CD Applications
- Better scalability for microservices
- Application definitions are managed through Git
- Easier bootstrapping of multiple applications
- New applications can be added through Git

---

### 12. What are the limitations of the App of Apps pattern?

- Adds an additional management layer
- Can make debugging more complex
- Incorrect parent-level changes can affect multiple applications
- May be unnecessary for small environments with only a few applications

---

## Key Interview Answer

### Explain the Argo CD App of Apps pattern.

The Argo CD App of Apps pattern is used to manage multiple Argo CD Applications in a centralized way. A root-app.yaml defines and creates a Parent Application. The Parent Application points to a Git repository path containing Child Application definitions. These Child Applications are created in Kubernetes, and each Child Application points to and manages its own Kubernetes resources such as Deployments, Services, ConfigMaps, and other resources.