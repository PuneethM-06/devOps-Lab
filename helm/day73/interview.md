# Interview.md — Day 73

# Helm Dependencies, Repositories & Packaging

## 1. What is a Helm Chart dependency, and where is it declared?

**Expected Answer:**

A Helm Chart dependency is another Helm Chart that our Chart depends on.

Dependencies are declared in `Chart.yaml` under the `dependencies` section.

Example:

```yaml
dependencies:
  - name: redis
    version: "20.0.0"
    repository: "https://example.com/helm-charts"
```

**Key point:**

* `Chart.yaml` → declares dependencies
* `values.yaml` → provides configuration values

---

## 2. What do `name`, `version`, and `repository` mean in a dependency declaration?

**Expected Answer:**

```yaml
dependencies:
  - name: redis
    version: "20.0.0"
    repository: "https://example.com/helm-charts"
```

* `name` → name of the dependency Chart
* `version` → specific version of the dependency Chart
* `repository` → location where Helm can find the Chart

---

## 3. What is the difference between `helm dependency update` and `helm dependency build`?

**Expected Answer:**

### `helm dependency update`

Resolves the dependencies declared in `Chart.yaml`, updates `Chart.lock`, and downloads the dependencies.

```text
Chart.yaml
    ↓
Resolve dependencies
    ↓
Update Chart.lock
    ↓
Download dependencies
    ↓
charts/
```

### `helm dependency build`

Uses the existing `Chart.lock` to fetch/rebuild the dependency set using the locked dependency information.

```text
Chart.lock
    ↓
helm dependency build
    ↓
Fetch locked dependencies
    ↓
charts/
```

**Key difference:**

> `update` resolves and updates dependency versions, while `build` uses the existing lock information to reproduce the dependency set.

---

## 4. What does `helm package` do, and why would you package a Helm Chart?

**Expected Answer:**

`helm package` packages a Helm Chart directory into a single `.tgz` archive.

```bash
helm package myapp/
```

Result:

```text
myapp-1.0.0.tgz
```

The packaged Chart can then be:

* Published to a Helm repository
* Stored in an OCI registry
* Pulled or installed by other users

---

## 5. What is the difference between a traditional Helm repository and an OCI registry?

**Expected Answer:**

A traditional Helm repository is specifically designed to store and distribute Helm Charts using the Helm repository format and metadata.

An OCI registry can store Helm Charts as **OCI artifacts**.

```text
Traditional Helm Repository
        ↓
Helm-specific repository
        ↓
Chart packages + Helm metadata


OCI Registry
        ↓
OCI-compatible registry
        ↓
Helm Charts stored as OCI artifacts
```

Both support distributing Charts through push/pull operations.

**Important:**

The difference is **not** that OCI eliminates push/pull.

The key difference is:

> **Traditional repository → Helm-specific Chart repository format.**

> **OCI registry → Helm Chart stored as an OCI artifact.**

---

# Important Commands

```bash
# Add a Helm repository
helm repo add myrepo https://example.com/helm-charts

# List configured repositories
helm repo list

# Refresh repository metadata
helm repo update

# Search for Charts
helm search repo myrepo

# Update Chart dependencies
helm dependency update

# Build dependencies from Chart.lock
helm dependency build

# Package a Chart
helm package myapp/

# Pull a Chart from a traditional repository
helm pull myrepo/myapp

# Install from a traditional repository
helm install myapp myrepo/myapp

# Push a Chart to an OCI registry
helm push myapp-1.0.0.tgz oci://registry.example.com/helm

# Pull from an OCI registry
helm pull oci://registry.example.com/helm/myapp --version 1.0.0

# Install from an OCI registry
helm install myapp oci://registry.example.com/helm/myapp --version 1.0.0
```

---

# Core Mental Model

```text
                  Chart.yaml
                      │
                      ↓
               Declare dependencies
                      │
                      ↓
          helm dependency update
                      │
             ┌────────┴────────┐
             ↓                 ↓
        Chart.lock           charts/
             │
             ↓
     helm dependency build
             │
             ↓
      Rebuild dependencies
             │
             ↓
        helm package
             │
             ↓
       myapp-1.0.0.tgz
             │
       ┌─────┴─────┐
       ↓           ↓
 Helm Repository  OCI Registry
       │           │
       └─────┬─────┘
             ↓
       helm install
             ↓
       Helm Chart
             ↓
    Kubernetes resources
```

# Key Takeaways

* Dependencies are declared in `Chart.yaml`.
* `values.yaml` is for configuration, not declaring Chart dependencies.
* `helm dependency update` resolves and downloads dependencies.
* `Chart.lock` records the resolved dependency information.
* `helm dependency build` uses the existing lock information to rebuild the dependency set.
* `helm package` creates a `.tgz` packaged Chart.
* A packaged Chart can be distributed through a traditional Helm repository or an OCI registry.
* Traditional Helm repositories use Helm-specific repository metadata.
* OCI registries store Helm Charts as OCI artifacts.
* Both traditional repositories and OCI registries support distributing Charts through push/pull workflows.
