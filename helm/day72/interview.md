# Interview.md — Day 72

# Helm Upgrade, Rollback & Diff

## 1. What is the difference between `helm install` and `helm upgrade`?

**Expected Answer:**

* `helm install` is used to create a new Helm release for the first time.
* `helm upgrade` is used to update an existing Helm release with new Chart or configuration changes.

```bash
helm install myapp ./myapp
```

Creates the initial release.

```bash
helm upgrade myapp ./myapp
```

Updates the existing release.

---

## 2. What is a Helm revision, and why does Helm maintain revisions?

**Expected Answer:**

A Helm revision represents a version of a Helm release created during installation or an upgrade.

Helm maintains revision history so that we can:

* Inspect previous states of a release
* Understand what changes happened
* Roll back to a previous revision if a newer release causes problems

Example:

```text
Revision 1 → Install
Revision 2 → Upgrade
Revision 3 → Upgrade
```

To view the history:

```bash
helm history myapp
```

---

## 3. What happens when you run `helm rollback myapp 2`?

**Expected Answer:**

It restores the configuration of Revision 2 and applies it to the Kubernetes resources managed by the Helm release.

```bash
helm rollback myapp 2
```

Important:

* Helm does not delete the newer revision.
* The rollback itself creates a **new revision**.

Example:

```text
Revision 1 → Install
Revision 2 → Upgrade
Revision 3 → Upgrade ❌

helm rollback myapp 2

Revision 4 → Rollback to Revision 2
```

The Kubernetes resources are updated to match the configuration from Revision 2.

---

## 4. What is `helm diff`, and how is it similar to `terraform plan`?

**Expected Answer:**

`helm diff` shows the changes that would occur between the currently deployed release and the proposed changes.

It allows us to review changes **before actually performing the upgrade**.

Example:

```text
Current:
replicas: 3
image: v1

New:
replicas: 5
image: v2
```

The diff can show:

```diff
- replicas: 3
+ replicas: 5

- image: v1
+ image: v2
```

It is conceptually similar to:

```bash
terraform plan
```

because both allow us to **preview changes before applying them**.

### Important

`helm diff` is provided by the **Helm Diff plugin** and is not part of the core Helm CLI by default.

---

## 5. What is the difference between `helm status`, `helm get values`, `helm get manifest`, and `helm template`?

**Expected Answer:**

### `helm status`

Shows the current state/status of a Helm release.

```bash
helm status myapp
```

### `helm get values`

Shows the values associated with the Helm release.

```bash
helm get values myapp
```

### `helm get manifest`

Shows the Kubernetes manifests associated with the release — the YAML resources Helm generated/stored for that release.

```bash
helm get manifest myapp
```

### `helm template`

Renders the Helm Chart locally and prints the resulting Kubernetes YAML without deploying anything.

```bash
helm template myapp ./myapp
```

### Comparison

```text
helm status
    ↓
Current release status

helm get values
    ↓
Release configuration values

helm get manifest
    ↓
Kubernetes manifests associated with release

helm template
    ↓
Render Chart locally
    ↓
No deployment
```

---

# Important Commands

```bash
# Install a new release
helm install myapp ./myapp

# Upgrade an existing release
helm upgrade myapp ./myapp

# View release history
helm history myapp

# Roll back to a previous revision
helm rollback myapp 2

# Show release status
helm status myapp

# Show release values
helm get values myapp

# Show deployed manifests
helm get manifest myapp

# Render templates locally
helm template myapp ./myapp

# Preview upgrade changes
helm diff upgrade myapp ./myapp
```

---

# Core Mental Model

```text
                    HELM RELEASE
                         │
              ┌──────────┴──────────┐
              ↓                     ↓
           INSTALL                UPGRADE
              │                     │
              ↓                     ↓
         Revision 1             Revision 2
                                    ↓
                                Revision 3
                                    ↓
                               Something breaks
                                    ↓
                            helm rollback myapp 2
                                    ↓
                                Revision 4
                                    ↓
                         Revision 2 configuration
                              is restored
```

# Upgrade vs Rollback vs Diff

```text
helm upgrade
    ↓
Actually update the existing release

helm diff upgrade
    ↓
Show what would change

helm rollback
    ↓
Restore a previous release configuration
```

# Troubleshooting Flow

```text
Something looks wrong
        ↓
helm status
        ↓
Check release state
        ↓
helm get values
        ↓
Check release values
        ↓
helm get manifest
        ↓
Check manifests associated with release
        ↓
helm template
        ↓
Check what the Chart currently renders
        ↓
helm diff upgrade
        ↓
Review proposed changes
        ↓
helm upgrade / helm rollback
```

# Key Takeaways

* `helm install` creates a new release.
* `helm upgrade` updates an existing release.
* Every installation or upgrade creates a Helm revision.
* `helm history` lets us inspect release revisions.
* `helm rollback` restores an older revision's configuration.
* A rollback creates a new revision rather than deleting the newer revision.
* `helm diff` previews changes before an upgrade.
* `helm diff` is conceptually similar to `terraform plan`.
* `helm status` shows the current release state.
* `helm get values` shows release values.
* `helm get manifest` shows the Kubernetes manifests associated with the release.
* `helm template` renders the Chart locally without deploying it.
