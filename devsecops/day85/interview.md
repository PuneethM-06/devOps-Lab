# Interview.md — Day 85: Trivy in GitHub Actions

## Must-Know Fundamentals

### 1. What is Trivy?

**Expected answer:**

Trivy is an open-source security scanner that can scan targets such as container images, repositories, filesystems, dependencies, Kubernetes configurations, and Infrastructure as Code for known vulnerabilities and other security issues.

For container security, Trivy can inspect a Docker image, identify the packages and versions inside it, and check whether they are affected by known vulnerabilities.

---

### 2. Why do we need vulnerability scanning?

**Expected answer:**

An application can have secure application code but still contain vulnerable dependencies or underlying OS packages.

Manually checking every package and version is not practical because applications can contain many direct and transitive dependencies, container images can contain additional OS-level packages, and new vulnerabilities can be discovered over time.

A vulnerability scanner such as Trivy automates this process.

---

### 3. What is a CVE?

**Expected answer:**

CVE stands for **Common Vulnerabilities and Exposures**.

A CVE is a standardized identifier used to track a publicly known security vulnerability.

A CVE is not the vulnerability itself. The vulnerability is the actual security weakness, while the CVE is the identifier used to track and reference it.

---

### 4. How does Trivy identify vulnerabilities in a container image?

**Expected answer:**

Trivy scans the target and identifies the installed packages, dependencies, and their versions. It then checks vulnerability databases and security advisories to determine whether known vulnerabilities affect those versions and reports the matching findings.

The report can include information such as the affected package, installed version, CVE ID, severity, and fixed version when available.

---

### 5. Why is the package version important during vulnerability scanning?

**Expected answer:**

A vulnerability usually affects specific versions or version ranges of a package.

Knowing only that an image contains a package such as `openssl` is not enough. Trivy needs the installed version to determine whether that version is affected by a known vulnerability.

---

### 6. What information can a Trivy vulnerability finding contain?

**Expected answer:**

A finding can include:

- Package or library name
- Installed version
- Vulnerability or CVE ID
- Severity
- Fixed version, if available

---

### 7. What are the common vulnerability severity levels?

**Expected answer:**

Common severity levels include:

- LOW
- MEDIUM
- HIGH
- CRITICAL

Higher severity indicates a more serious vulnerability, but severity does not mean the probability that the vulnerability will definitely be exploited.

A CRITICAL vulnerability also does not automatically mean that the application has already been compromised.

---

### 8. Does finding a CRITICAL CVE mean the application is definitely exploitable?

**Expected answer:**

No.

A CRITICAL CVE means that the known vulnerability has a high severity or potential impact. Whether it is exploitable in a specific environment can depend on factors such as application configuration, whether the vulnerable functionality is used, whether an attacker can reach it, and other conditions required for exploitation.

However, HIGH and CRITICAL findings should still be investigated and remediated.

---

### 9. Why should we scan the final Docker image instead of only a dependency file such as `requirements.txt`?

**Expected answer:**

Scanning only `requirements.txt` focuses mainly on application dependencies.

The final Docker image can contain:

- Application code
- Application dependencies
- The base image
- OS-level packages
- Other installed libraries

Therefore, scanning the final image provides a broader view of the deployable artifact and can identify vulnerabilities that come from the underlying base image or operating system packages.

---

### 10. How do you scan a Docker image using Trivy?

**Expected answer:**

The basic command is:

    trivy image <image-name>

For example:

    trivy image my-python-app:latest

Trivy then examines the image, identifies packages and versions, checks known vulnerability information, and reports matching findings.

---

### 11. Why doesn't finding a vulnerability automatically fail a CI pipeline?

**Expected answer:**

A vulnerability scanner can report vulnerabilities while still returning a successful exit code.

For CI enforcement, Trivy must be configured to return a non-zero exit code when vulnerabilities matching the defined policy are found.

GitHub Actions then uses that non-zero exit code to mark the step or job as failed.

---

### 12. What does `--severity HIGH,CRITICAL` do?

**Expected answer:**

It filters the vulnerability findings based on severity.

For example:

    --severity HIGH,CRITICAL

This means the CI policy is focused on HIGH and CRITICAL vulnerability findings.

---

### 13. What does `--exit-code 1` do?

**Expected answer:**

It tells Trivy to return exit code `1` when vulnerabilities matching the configured criteria are found.

Because exit code `1` is a non-zero exit code, GitHub Actions marks the command or step as failed.

---

### 14. Explain the difference between `--severity` and `--exit-code`.

**Expected answer:**

`--severity HIGH,CRITICAL` defines which vulnerability findings the scan policy cares about.

`--exit-code 1` causes Trivy to return a failure exit code when matching findings exist.

In short:

    --severity → What findings do we care about?

    --exit-code → Should the command fail when those findings exist?

---

### 15. Give an example of a Trivy command that fails on HIGH or CRITICAL vulnerabilities.

**Expected answer:**

    trivy image \
      --severity HIGH,CRITICAL \
      --exit-code 1 \
      my-python-app:latest

This scans the image and returns exit code `1` when HIGH or CRITICAL vulnerabilities matching the scan criteria are found.

---

### 16. Why build the Docker image inside GitHub Actions before scanning it?

**Expected answer:**

A GitHub Actions runner is a fresh environment and does not automatically contain the Docker image built on a developer's local machine.

The workflow needs to:

1. Check out the source code.
2. Build the Docker image.
3. Scan that image with Trivy.

This also allows the CI pipeline to scan the deployable artifact produced from the current change.

---

### 17. Why run Trivy automatically in GitHub Actions instead of asking developers to run it manually?

**Expected answer:**

Running Trivy in GitHub Actions provides consistent and automated security checks.

Every change can be evaluated using the same configuration and severity thresholds, reducing human error and avoiding situations where a developer forgets to run the scan.

It also catches known vulnerabilities before the change is merged or deployed.

---

### 18. What is the role of the Trivy GitHub Action?

**Expected answer:**

The Trivy GitHub Action provides a reusable and convenient way to run Trivy inside a GitHub Actions workflow.

It avoids manually installing and configuring Trivy in every workflow and allows scan settings to be provided through the action's configuration.

Trivy remains the actual vulnerability scanner. The GitHub Action is the integration mechanism used to run Trivy conveniently inside GitHub Actions.

---

### 19. What is responsible for finding vulnerabilities, running the scan, and blocking the merge?

**Expected answer:**

    Trivy
      ↓
    Finds vulnerabilities

    GitHub Actions
      ↓
    Runs the scan automatically
    and determines whether the check passes or fails

    Branch Protection Rules
      ↓
    Can require the security check to pass
    before allowing a merge

---

### 20. Does a failed GitHub Actions workflow automatically block a pull request from merging?

**Expected answer:**

No.

A failed workflow by itself does not necessarily prevent a pull request from being merged.

To enforce merge blocking, branch protection or the repository's merge rules must require the relevant GitHub Actions security check to pass.

---

## Scenario-Based Questions

### 21. Trivy finds a CRITICAL vulnerability, but the GitHub Actions workflow passes. What could be wrong?

**Expected answer:**

The workflow may only be reporting vulnerabilities without being configured to fail.

The Trivy scan may be missing an appropriate failure exit code configuration, such as:

    --exit-code 1

The severity filter or scan configuration should also be checked.

---

### 22. A Trivy check fails, but a developer can still merge the PR. Why?

**Expected answer:**

The failed GitHub Actions check may not be configured as a required status check in the branch protection or merge rules.

Trivy finds the vulnerability, GitHub Actions fails the check, but repository rules are responsible for enforcing whether a successful check is required before merging.

---

### 23. Your Python dependencies are clean, but Trivy finds a vulnerability in the Docker image. Where could it come from?

**Expected answer:**

The vulnerability could come from:

- The base image
- An OS-level package
- Another installed library inside the image

This is why scanning the final Docker image can provide a broader security view than scanning only `requirements.txt`.

---

### 24. Why is this statement incorrect: "Trivy blocks the PR from merging"?

**Expected answer:**

Trivy itself scans for vulnerabilities and reports findings.

The complete enforcement chain is:

    Trivy
      ↓
    Finds vulnerability and returns failure

    GitHub Actions
      ↓
    Marks the check as failed

    Branch protection
      ↓
    Blocks the merge if that check is required

Therefore, Trivy does not directly block the merge by itself.

---

### 25. Explain the complete flow for blocking vulnerable changes from merging.

**Expected answer:**

    Developer opens a PR
            ↓
    GitHub Actions starts
            ↓
    Build Docker image
            ↓
    Trivy scans the image
            ↓
    HIGH/CRITICAL vulnerability found
            ↓
    Trivy returns exit code 1
            ↓
    GitHub Actions check fails
            ↓
    Branch protection requires that check to pass
            ↓
    PR cannot merge into main

---

## Final Interview Answer

> **Trivy is an open-source security scanner that can scan container images, dependencies, repositories, and other targets for known vulnerabilities. For container images, it identifies packages and their versions, checks them against vulnerability databases and security advisories, and reports matching findings such as CVE ID, severity, and fixed version. In GitHub Actions, Trivy can be configured to fail the security check when vulnerabilities matching a defined severity threshold are found. GitHub Actions then marks the check as failed, and branch protection can require that check to pass before allowing the pull request to merge into `main`.**