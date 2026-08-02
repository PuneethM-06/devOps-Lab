# Day 61 – Kubernetes ConfigMaps & Secrets

## Topics Covered
- Why ConfigMaps and Secrets are required
- Build Once, Deploy Anywhere
- ConfigMap vs Secret
- Environment Variables
- Mounted Volumes
- Creating and consuming ConfigMaps
- Creating and consuming Secrets
- Enterprise secret management using Vault
- Captain deployment flow
- ConfigMap/Secret update behavior
- Best practices

## Interview Questions

### 1. Why do we need ConfigMaps?

ConfigMaps store non-sensitive configuration separately from the application, allowing the same Docker image to be deployed across multiple environments without rebuilding it.

---

### 2. Why do we need Secrets?

Secrets securely store sensitive information such as passwords, API keys, and certificates instead of embedding them inside Docker images or application code.

---

### 3. Difference between ConfigMap and Secret?

ConfigMaps store non-sensitive configuration, whereas Secrets store sensitive data.

---

### 4. What are the two ways to consume ConfigMaps or Secrets?

- Environment Variables
- Mounted Volumes

---

### 5. When should you use Environment Variables?

For simple key-value configurations such as database host, usernames, passwords, API URLs, and feature flags.

---

### 6. When should you use Mounted Volumes?

When applications expect configuration as files, such as application.properties, nginx.conf, TLS certificates, or SSH keys.

---

### 7. Explain how a ConfigMap is injected as an environment variable.

The Deployment references a ConfigMap using `configMapKeyRef`. During Pod startup, Kubernetes reads the specified key from the ConfigMap and injects it as an environment variable into the container.

---

### 8. Explain how a Secret is injected into a Pod.

The Deployment references a Secret using `secretKeyRef`. Kubernetes injects the Secret as environment variables or mounted files when the Pod starts.

---

### 9. What happens if a ConfigMap is updated?

If consumed as environment variables, running Pods do not receive the updated values. A restart or rollout is required. If consumed as mounted volumes, Kubernetes updates the files, but whether the application picks up the changes depends on the application.

---

### 10. How are secrets managed in enterprise environments?

Secrets are typically stored in Vault or another secret manager. Deployment platforms such as Captain retrieve secrets from Vault, create Kubernetes Secrets, and the Deployment consumes those Kubernetes Secrets.

---

### 11. Does Kubernetes communicate directly with Vault?

No. Kubernetes only consumes Kubernetes Secrets. A deployment platform or secret management integration (such as Captain, Vault Agent, or External Secrets Operator) retrieves secrets from Vault and creates or updates Kubernetes Secrets.

---

### 12. Why shouldn't Secrets be committed to Git?

Committing secrets exposes sensitive information. Instead, secrets should be stored in a secure secret management system and injected during deployment.