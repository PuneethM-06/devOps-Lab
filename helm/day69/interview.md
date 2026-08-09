# Interview.md — Day 69

# Helm Values & Environment Configuration

## 1. What are Helm Values, and why do we use them?

**Expected Answer:**

Helm Values are **configuration parameters provided to a Helm Chart and consumed by its templates**.

They allow us to change application configuration without modifying the underlying Kubernetes templates.

For example:

```yaml
replicaCount: 3
```

can be consumed by:

```yaml
replicas: {{ .Values.replicaCount }}
```

This allows the same Chart to be configured differently for different environments.

---

## 2. What is the purpose of `values.yaml`?

**Expected Answer:**

`values.yaml` contains the **default configuration values** for a Helm Chart.

For example:

```yaml
replicaCount: 2

image:
  repository: nginx
  tag: "1.27"
```

Templates can access these values using `.Values`.

`values.yaml` does not directly create Kubernetes resources. It provides configuration that is consumed by the templates.

---

## 3. What is `.Values` in Helm?

**Expected Answer:**

`.Values` is used inside Helm templates to **access the values provided to the Chart**.

For example:

```yaml
replicaCount: 3
```

can be accessed using:

```yaml
{{ .Values.replicaCount }}
```

Nested values can also be accessed:

```yaml
image:
  repository: nginx
  tag: "1.27"
```

Using:

```yaml
{{ .Values.image.repository }}
{{ .Values.image.tag }}
```

---

## 4. What is the difference between `values.yaml`, `values-dev.yaml`, and `values-prod.yaml`?

**Expected Answer:**

* `values.yaml` → contains the **default configuration**.
* `values-dev.yaml` → contains **development-specific configuration**.
* `values-prod.yaml` → contains **production-specific configuration**.

For example:

```yaml
# values-dev.yaml
replicaCount: 1
```

```yaml
# values-prod.yaml
replicaCount: 5
```

The same Chart and templates can be reused for both environments.

---

## 5. How do you tell Helm to use a specific values file?

**Expected Answer:**

Use `-f` or `--values`.

Example:

```bash
helm install myapp-dev ./myapp -f values-dev.yaml
```

The following are equivalent:

```bash
-f values-dev.yaml
```

and:

```bash
--values values-dev.yaml
```

They tell Helm to use the specified values file when rendering the Chart.

---

## 6. What is `--set` in Helm?

**Expected Answer:**

`--set` is used to **override a specific value directly from the command line**.

Example:

```bash
helm install myapp ./myapp --set replicaCount=5
```

If the default value is:

```yaml
replicaCount: 2
```

the final value becomes:

```yaml
replicaCount: 5
```

It is useful for small or temporary overrides.

---

## 7. What is the precedence of Helm values?

**Expected Answer:**

The simplified precedence is:

```text
--set
  ↓
-f / --values
  ↓
values.yaml
```

The higher-precedence value overrides the lower-precedence value when the same key is defined.

For example:

```yaml
# values.yaml
replicaCount: 2
```

```yaml
# values-prod.yaml
replicaCount: 5
```

and:

```bash
--set replicaCount=10
```

The final value is:

```text
replicaCount = 10
```

because `--set` has the highest precedence.

---

## 8. What happens when multiple `-f` values files are provided?

**Expected Answer:**

Multiple values files can be provided:

```bash
helm install myapp ./myapp \
  -f values-prod.yaml \
  -f values-hotfix.yaml
```

If both files define the same value, the **right-most values file takes precedence**.

So:

```text
values.yaml
     ↓
values-prod.yaml
     ↓
values-hotfix.yaml
     ↓
--set
```

The more specific/later value wins.

---

## 9. How can the same Helm Chart be used for different environments?

**Expected Answer:**

The same Chart can contain common Kubernetes templates while different values files provide environment-specific configuration.

Example:

```text
myapp/
├── Chart.yaml
├── values.yaml
├── values-dev.yaml
├── values-prod.yaml
└── templates/
    └── deployment.yaml
```

The same template can use:

```yaml
replicas: {{ .Values.replicaCount }}
```

Development:

```bash
helm install myapp-dev ./myapp -f values-dev.yaml
```

Production:

```bash
helm install myapp-prod ./myapp -f values-prod.yaml
```

This allows us to maintain **one set of templates** instead of duplicating Kubernetes manifests for every environment.

---

## 10. What is the difference between a Helm template and a values file?

**Expected Answer:**

The **template defines how the Kubernetes resource is structured**, while the **values file provides configuration for that template**.

```text
templates/
    ↓
HOW the resource is structured

values-dev.yaml
values-prod.yaml
    ↓
HOW the resource is configured
```

For example:

```yaml
# Template
replicas: {{ .Values.replicaCount }}
```

Development:

```yaml
# values-dev.yaml
replicaCount: 1
```

Production:

```yaml
# values-prod.yaml
replicaCount: 5
```

Same template, different configuration.

---

# Important Helm Commands

```bash
helm install myapp ./myapp -f values-dev.yaml

helm install myapp ./myapp -f values-prod.yaml

helm install myapp ./myapp --set replicaCount=5
```

### Key options

* `-f` / `--values` → Use a specific values file.
* `--set` → Override a specific value directly from the command line.

---

# Core Mental Model

```text
                    ONE HELM CHART
                         │
                     templates/
                         │
            ┌────────────┼────────────┐
            ↓            ↓            ↓
       values-dev   values-stage   values-prod
            │            │            │
            └────────────┼────────────┘
                         ↓
                    Helm renders
                         ↓
              Kubernetes manifests
                         ↓
                   K8s resources
```

### Key takeaway

> **The same Helm Chart can be reused across environments by keeping the Kubernetes templates common and changing the configuration through different values files.**
