# Interview.md — Day 71

# Helm Hooks

## 1. What is a Helm Hook, and why would you use one?

**Expected Answer:**

A Helm Hook is a Kubernetes resource that Helm executes at a specific point in the Helm lifecycle.

Hooks are used when we need a task to run before or after Helm operations such as installation, upgrade, or deletion.

A common example is running a database migration before deploying an application.

```text
helm install
     ↓
DB migration Hook
     ↓
Migration completes
     ↓
Application resources are installed
```

---

## 2. How do we identify a Kubernetes resource as a Helm Hook, and where do we define what the Hook actually does?

**Expected Answer:**

We identify a Kubernetes resource as a Helm Hook using the `helm.sh/hook` annotation.

Example:

```yaml
metadata:
  annotations:
    "helm.sh/hook": pre-install
```

The actual behavior is defined by the Kubernetes resource itself.

For example:

```yaml
kind: Job
```

with:

```yaml
containers:
  - name: migration
    image: myapp:migrate
    command: ["./migrate"]
```

Here:

```text
helm.sh/hook: pre-install
        ↓
WHEN → before installation

Job + image + command
        ↓
WHAT → database migration
```

> The Kubernetes resource defines **what** should happen, while the `helm.sh/hook` annotation defines **when** it should happen.

---

## 3. What is the difference between `pre-install` and `pre-upgrade` Hooks?

**Expected Answer:**

### `pre-install`

Runs before Helm installs the normal Kubernetes resources for the first time.

```bash
helm install myapp ./myapp
```

### `pre-upgrade`

Runs before Helm upgrades the resources of an existing release.

```bash
helm upgrade myapp ./myapp
```

Example:

```text
First deployment:

helm install
     ↓
pre-install Hook
     ↓
Create resources


Existing release:

helm upgrade
     ↓
pre-upgrade Hook
     ↓
Update resources
```

---

## 4. What is the purpose of `post-install` and `post-upgrade` Hooks?

**Expected Answer:**

These Hooks are similar to the `pre` Hooks, but they execute **after** the corresponding Helm operation.

### `post-install`

Runs after the normal Kubernetes resources have been installed.

```text
helm install
     ↓
Create resources
     ↓
post-install Hook
```

### `post-upgrade`

Runs after the normal Kubernetes resources have been upgraded.

```text
helm upgrade
     ↓
Update resources
     ↓
post-upgrade Hook
```

---

## 5. What happens if a Helm Hook fails?

**Expected Answer:**

If a Hook fails, Helm does not simply ignore the failure and continue with the operation.

The Helm installation or upgrade fails.

Example:

```text
helm install
     ↓
pre-install Hook
     ↓
   FAIL
     ↓
Helm install fails
```

For a database migration:

```text
Database migration
       ↓
    FAILED
       ↓
Application deployment
       ✕
```

This prevents the new application version from being deployed when a required migration has failed.

---

# Important Helm Hook Annotations

## Identify a Hook

```yaml
annotations:
  "helm.sh/hook": pre-install
```

## Multiple Hook Events

A resource can be associated with multiple Hook events:

```yaml
annotations:
  "helm.sh/hook": pre-install,pre-upgrade
```

This allows the same resource to run during both initial installation and upgrades.

---

# Important Hook Types

```text
pre-install
    → Before installing normal resources

post-install
    → After installing normal resources

pre-upgrade
    → Before upgrading normal resources

post-upgrade
    → After upgrading normal resources

pre-delete
    → Before deleting the release

post-delete
    → After deleting the release
```

For interviews, the most important four are:

```text
pre-install
post-install
pre-upgrade
post-upgrade
```

---

# Hook Deletion Policy

Hook resources such as Jobs can remain in the cluster after they complete.

Helm provides `helm.sh/hook-delete-policy` to control when Hook resources should be removed.

Example:

```yaml
annotations:
  "helm.sh/hook": pre-install
  "helm.sh/hook-delete-policy": hook-succeeded
```

### Common policies

```text
hook-succeeded
    → Delete the Hook after successful execution

hook-failed
    → Delete the Hook after failed execution

before-hook-creation
    → Delete the previous Hook before creating a new one
```

Multiple policies can also be combined:

```yaml
"helm.sh/hook-delete-policy": before-hook-creation,hook-succeeded
```

---

# Real-World Example — Database Migration

A database migration can be implemented as a Kubernetes Job and marked as a `pre-install` or `pre-upgrade` Hook.

```yaml
apiVersion: batch/v1
kind: Job

metadata:
  name: db-migration
  annotations:
    "helm.sh/hook": pre-install,pre-upgrade
    "helm.sh/hook-delete-policy": hook-succeeded

spec:
  template:
    spec:
      restartPolicy: Never

      containers:
        - name: migration
          image: myapp:migrate
          command: ["./migrate"]
```

The flow is:

```text
helm install / helm upgrade
          ↓
Migration Hook
          ↓
    ┌─────┴─────┐
    ↓           ↓
 SUCCESS       FAILURE
    ↓             ↓
Continue       Helm operation
Helm           fails
    ↓
Deploy/update
application
```

# Core Mental Model

```text
                  HELM OPERATION
                        │
           ┌────────────┴────────────┐
           ↓                         ↓
       PRE HOOK                  POST HOOK
           ↓                         ↓
   Do something BEFORE        Do something AFTER
           │                         │
           └────────────┬────────────┘
                        ↓
               Normal K8s resources
```

# Key Takeaways

* A Helm Hook is a Kubernetes resource executed at a specific point in the Helm lifecycle.
* `helm.sh/hook` identifies a resource as a Hook.
* The Kubernetes resource itself defines what the Hook actually does.
* `pre-install` runs before the initial installation.
* `pre-upgrade` runs before upgrading an existing release.
* `post-install` runs after installation.
* `post-upgrade` runs after an upgrade.
* Helm also supports delete-related Hooks.
* A failed Hook causes the Helm operation to fail.
* Hook deletion policies control what happens to Hook resources after execution.
* Database migrations are a common real-world use case for `pre-install` and `pre-upgrade` Hooks.
