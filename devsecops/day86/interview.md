# interview.md — Day 86: Snyk — Scan Go Dependencies & Auto-Fix PRs

## Must-Know Fundamentals

### 1. What is Snyk?

**Expected answer:**

Snyk is a developer security platform that helps identify and remediate known security vulnerabilities in application dependencies and other software components.

For a Go application, Snyk can analyze the project's dependency graph, identify vulnerable dependencies, report the findings, and help suggest or automate remediation when a suitable fix is available.

---

### 2. What is the role of `go.mod`?

**Expected answer:**

`go.mod` defines a Go module and declares its dependency requirements.

For example:

    module github.com/example/my-app

    go 1.25

    require github.com/gin-gonic/gin v1.11.0

Here, `github.com/gin-gonic/gin v1.11.0` is a dependency required by the application.

The `go 1.25` line is not a dependency. It specifies the Go version associated with the module.

---

### 3. What is a direct dependency?

**Expected answer:**

A direct dependency is a dependency that the application directly requires or imports.

For example:

    Your Go Application
            ↓
    Package A

Package A is a direct dependency.

---

### 4. What is an indirect dependency?

**Expected answer:**

An indirect dependency is a dependency required by another dependency.

For example:

    Your Go Application
            ↓
    Package A            ← Direct dependency
            ↓
    Package B            ← Indirect dependency

Even if the developer never explicitly added Package B, it can still become part of the application's dependency graph because Package A requires it.

---

### 5. What is `go.sum`?

**Expected answer:**

`go.sum` records checksums used to verify downloaded Go module content.

A checksum acts like a fingerprint of the module's contents. When Go downloads a module, it can verify that the downloaded content matches the expected checksum.

The simple mental model is:

    go.mod
        ↓
    What modules and dependency versions does my project require?

    go.sum
        ↓
    What checksums are recorded to help verify downloaded module content?

---

### 6. How does Snyk find vulnerabilities in indirect dependencies?

**Expected answer:**

Snyk analyzes the complete dependency graph rather than only looking at dependencies directly added by the developer.

For example:

    Your Go Application
            ↓
    Package A            ← Direct dependency
            ↓
    Package B            ← Indirect dependency
            ↓
    Known CVE ❌

Because Package B exists in the resolved dependency graph, Snyk can identify its version and check whether it is affected by known vulnerabilities.

---

### 7. What information can Snyk provide when it finds a vulnerable dependency?

**Expected answer:**

Snyk can provide information such as:

- The vulnerable package
- The installed version
- Vulnerability or CVE information
- Severity
- How the dependency entered the dependency graph
- Available remediation or fixed versions

Understanding how the dependency entered the application is especially useful when the vulnerable package is an indirect dependency.

---

### 8. Why can blindly upgrading an indirect dependency be risky?

**Expected answer:**

The dependency that introduced the indirect package may rely on a specific API or behavior.

For example:

    Package A
        ↓
    Depends on Package B v1.5

If Package B is blindly upgraded:

    Package B v1.5
            ↓
        Upgrade
            ↓
    Package B v1.6

the new version could introduce compatibility issues, API changes, behavioral changes, or dependency conflicts.

Therefore, we should understand the dependency relationship, choose the appropriate remediation, and test the application.

---

### 9. What is Snyk Auto-Fix?

**Expected answer:**

Snyk Auto-Fix refers to Snyk's ability to help remediate vulnerabilities by identifying suitable dependency changes when a fix is available.

The general flow is:

    Vulnerability found
            ↓
    Fix available?
       ┌────┴────┐
       │         │
      Yes        No
       │         │
    Propose/     Investigate
    apply fix    manually

Snyk may suggest or automate a dependency update that removes the vulnerability, but not every vulnerability has an available automatic fix.

---

### 10. What happens if no fix is available?

**Expected answer:**

If no suitable dependency upgrade is available, the issue may need to be investigated manually.

Possible approaches include:

- Investigating alternative versions
- Applying a workaround
- Changing configuration
- Replacing the dependency
- Applying another appropriate mitigation

The correct approach depends on the vulnerability and the application.

---

### 11. What are automated fix PRs?

**Expected answer:**

When a suitable remediation is available, Snyk can help automate the remediation workflow by proposing or creating a pull request containing dependency changes.

The general flow is:

    Snyk finds vulnerability
            ↓
    Suitable fix available
            ↓
    Dependency update proposed
            ↓
    Fix PR created
            ↓
    Review
            ↓
    Tests
            ↓
    Merge if safe

---

### 12. Why should a Snyk-generated fix PR still be reviewed and tested?

**Expected answer:**

Snyk's goal is to remediate the known vulnerability, but it does not fully understand the application's business logic or guarantee that an updated dependency will not break the application.

A dependency update can introduce:

- Breaking API changes
- Behavioral changes
- Compatibility issues
- Dependency conflicts
- Test failures

Therefore, the fix should be reviewed and the application should be tested before merging.

---

### 13. How can Snyk be integrated into CI/CD?

**Expected answer:**

Snyk can be added to a CI/CD pipeline so that dependency scans run automatically on events such as pushes and pull requests.

The general flow is:

    Push / Pull Request
            ↓
    CI pipeline starts
            ↓
    Snyk scans Go dependencies
            ↓
    Vulnerabilities found?
            ↓
        Pass / Fail

This ensures that security scanning is automated and consistent.

---

### 14. Why is CI/CD integration better than manually running Snyk?

**Expected answer:**

CI/CD integration automates the scan so developers do not have to remember to run it manually.

Benefits include:

- Consistent security checks
- Reduced human error
- Automatic scanning on configured events
- Clear pass/fail results
- Easier enforcement through branch protection

---

### 15. Does Snyk directly block a pull request from merging?

**Expected answer:**

No.

Snyk identifies vulnerabilities and can cause the CI check to fail based on the configured policy.

The complete enforcement flow is:

    Snyk
        ↓
    Finds vulnerability

    CI/CD
        ↓
    Runs the scan and marks the check as pass/fail

    Branch protection
        ↓
    Requires the check to pass

    If the check fails
        ↓
    Merge can be blocked

Therefore, Snyk itself does not directly enforce the merge restriction. Repository or branch protection rules enforce it.

---

### 16. How does Snyk complement Trivy?

**Expected answer:**

Snyk and Trivy have overlapping capabilities, so they should not be simplified as tools that can only scan one type of target.

For our workflow:

    Go Application
            ↓
    Dependency graph
            ↓
    Snyk
            ↓
    Find vulnerable dependencies
    + Help with remediation

Then:

    Build Docker image
            ↓
    Trivy
            ↓
    Scan the final image
    + Application dependencies
    + Base image
    + OS packages

Snyk can help catch dependency vulnerabilities early in the developer workflow, while Trivy can inspect the final packaged Docker image.

---

## Scenario-Based Questions

### 17. You never added Package B directly. How can Snyk still find a CVE in it?

**Expected answer:**

Snyk analyzes the complete dependency graph, including indirect or transitive dependencies.

If Package A depends on Package B, Package B becomes part of the application's dependency graph even though the developer never directly added it.

Snyk can therefore identify Package B, determine its version, and check it for known vulnerabilities.

---

### 18. Snyk finds Package B v1.5 vulnerable, and Package B v1.6 fixes the issue. Why should you not blindly force the upgrade?

**Expected answer:**

Package B may be an indirect dependency required by Package A.

Package A may depend on the API or behavior of Package B v1.5. Forcing an upgrade could introduce compatibility or behavioral issues.

The dependency relationship should be understood, an appropriate remediation should be chosen, and the application should be tested.

---

### 19. Why is automatically merging every Snyk-generated fix PR dangerous?

**Expected answer:**

A dependency update may fix the CVE but still break the application.

The updated package may introduce API changes, behavioral changes, compatibility problems, or dependency conflicts.

Therefore, the fix PR should be reviewed and tested before merging.

---

### 20. Why is running Snyk in CI/CD useful?

**Expected answer:**

Snyk can run automatically on events such as pushes and pull requests.

Developers do not need to remember to run the scan manually, the same security policy can be applied consistently, and the result can be used as a pass/fail CI check.

Branch protection can then require the security check to pass before merging into `main`.

---

### 21. Snyk reports that your Go dependencies are clean, but Trivy finds a HIGH vulnerability in the final Docker image. Why can this happen?

**Expected answer:**

Snyk may have found no vulnerable packages in the Go dependency graph, but the Docker image can contain additional components.

The vulnerability could come from:

- The base image
- The underlying operating system
- OS-level packages
- Other components included in the final image

This is why scanning both the dependency graph and the final container image provides broader coverage.

---

## Final Interview Answer

> **Snyk is a developer security platform that can identify vulnerabilities in a Go application's dependency graph, including both direct and indirect dependencies. It can provide information about vulnerable packages, how they entered the dependency graph, and available remediation options. When integrated into CI/CD, Snyk can automatically scan dependencies on events such as pull requests and fail checks when the configured security policy is violated. Snyk can also help propose or create dependency update PRs, but those changes must still be reviewed and tested because a security fix can introduce compatibility or behavioral issues. Snyk complements tools such as Trivy, which can additionally scan the final Docker image, including base image and OS-level packages.**