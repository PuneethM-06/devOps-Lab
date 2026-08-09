# Interview.md — Day 68

# Helm Fundamentals & Chart Structure

## 1. What is Helm, and why do we use it with Kubernetes?

**Expected Answer:**

Helm is a **package manager for Kubernetes**. It helps package, configure, install, upgrade, and manage Kubernetes applications.

Without Helm, maintaining separate raw Kubernetes YAML files for different environments can lead to duplicated and difficult-to-maintain configurations.

Helm solves this by providing **reusable templates and configurable values**, allowing the same Chart to be used across environments such as development, staging, and production.

---

## 2. What is a Helm Chart, and what is a Helm Release? What is the difference between them?

**Expected Answer:**

A **Helm Chart** is a collection of files that describes a set of Kubernetes resources. It contains files such as `Chart.yaml`, `values.yaml`, and Kubernetes manifest templates.

A **Helm Release** is an **installed instance of a Helm Chart** in a Kubernetes cluster. Helm tracks and manages the Release, allowing operations such as upgrades, rollbacks, status checks, and uninstallations.

**Key difference:**

```text
Chart
  ↓
The package / definition

Release
  ↓
An installed instance of that Chart
```

The same Chart can be installed multiple times, creating different Releases.

---

## 3. What are `Chart.yaml`, `values.yaml`, and `templates/` used for?

**Expected Answer:**

### `Chart.yaml`

Contains metadata about the Helm Chart.

Example:

```yaml
apiVersion: v2
name: myapp
description: A Helm chart for my application
type: application
version: 0.1.0
appVersion: "1.0.0"
```

### `values.yaml`

Contains the **default configuration values** used by Helm templates.

Example:

```yaml
replicaCount: 2

image:
  repository: nginx
  tag: "1.27"
```

### `templates/`

Contains **Kubernetes manifest templates** that Helm processes using the supplied values to generate the final Kubernetes manifests.

Example:

```yaml
replicas: {{ .Values.replicaCount }}
```

---

## 4. What happens when you run `helm install myapp ./mychart`?

**Expected Answer:**

The flow is:

```text
Create a Helm Chart
        ↓
Configure values / templates
        ↓
helm install
        ↓
Helm combines:
  Templates + Values
        ↓
Renders Kubernetes manifests
        ↓
Sends them to the Kubernetes API
        ↓
Kubernetes creates the actual resources
        ↓
Helm tracks this installation as a Release
```

`myapp` is the **Release name**, while `./mychart` is the **Chart location**.

The rendered manifests can create resources such as:

* Deployment
* Service
* ConfigMap
* Secret
* Ingress
* HPA
* ServiceAccount

Helm does not replace Kubernetes. Kubernetes is responsible for creating and running the actual resources.

---

## 5. What is a Helm Repository, and how is it different from a Helm Chart and a Helm Release?

**Expected Answer:**

A **Helm Repository** is a location used to store and distribute Helm Charts.

A **Helm Chart** is a collection of files that defines a Kubernetes application and its resources.

A **Helm Release** is an installed instance of a Helm Chart that Helm tracks and manages.

```text
Helm Repository
      ↓
    Chart
      ↓
 helm install
      ↓
   Release
      ↓
Kubernetes resources
```

### Quick comparison

| Term           | Meaning                                        |
| -------------- | ---------------------------------------------- |
| **Helm**       | Package manager for Kubernetes                 |
| **Repository** | Location where Charts are stored/distributed   |
| **Chart**      | Package/definition of a Kubernetes application |
| **Release**    | Installed instance of a Chart                  |

---

# Basic Helm Commands

```bash
helm version
helm create myapp
helm install myapp ./mychart
helm list
helm status myapp
helm uninstall myapp
```

### Command meanings

* `helm version` — Shows the installed Helm version.
* `helm create` — Creates a new Helm Chart.
* `helm install` — Installs a Chart and creates a Release.
* `helm list` — Lists Helm Releases.
* `helm status` — Shows the status/details of a Release.
* `helm uninstall` — Removes a Release.

---

# Core Helm Mental Model

```text
                Helm Chart
                    │
          ┌─────────┴─────────┐
          ↓                   ↓
    values.yaml          templates/
          │                   │
          └─────────┬─────────┘
                    ↓
              Helm rendering
                    ↓
          Kubernetes manifests
                    ↓
            Kubernetes API
                    ↓
          Kubernetes resources
                    ↓
              Helm Release
```

# Key Takeaways

* Helm is a **package manager for Kubernetes**.
* A **Chart** is the package/definition of a Kubernetes application.
* `values.yaml` provides configuration values.
* `templates/` contains reusable Kubernetes manifest templates.
* Helm renders templates using values into Kubernetes manifests.
* Kubernetes creates the actual resources from those manifests.
* A **Release** is an installed instance of a Chart that Helm tracks and manages.
* A **Repository** stores and distributes Charts.
* Helm does not replace Kubernetes; it works on top of Kubernetes to simplify application packaging and lifecycle management.
