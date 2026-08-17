# interview.md — Day 87: SBOM — Generate with Trivy & Why It Matters

## Must-Know Fundamentals

### 1. What is an SBOM?

**Expected answer:**

SBOM stands for **Software Bill of Materials**.

It is an inventory of the software components and their versions that make up a specific application or software artifact.

For example:

    Docker Image
    │
    ├── Go dependency A v1.2
    ├── Go dependency B v2.0
    ├── OS package X v3.1
    └── OS package Y v4.5

An SBOM helps answer:

> What components and versions are inside this software artifact?

---

### 2. Why do we need an SBOM?

**Expected answer:**

An SBOM provides visibility into the components and versions used inside software.

When a new vulnerability is discovered, instead of manually checking every application or asking every development team which package versions they use, we can search the collected SBOM inventory to identify potentially affected applications or artifacts.

The flow is:

    New CVE announced
            ↓
    Identify affected component
            ↓
    Search SBOM inventory
            ↓
    Find applications/artifacts containing it
            ↓
    Check versions
            ↓
    Identify potentially affected software
            ↓
    Prioritize remediation

---

### 3. What is the difference between an SBOM and vulnerability scanning?

**Expected answer:**

An SBOM answers:

> What components and versions do we have?

A vulnerability scanner answers:

> Are any of those components known to be vulnerable?

For example:

    SBOM
        ↓
    openssl v3.0.2

    Vulnerability Scanner
        ↓
    Check openssl v3.0.2 against known
    vulnerability information
        ↓
    CVE findings / severity / remediation information

An SBOM is primarily an inventory. It is not itself a vulnerability report.

---

### 4. How can SBOMs and vulnerability scanning work together?

**Expected answer:**

An SBOM provides the inventory of components and versions, while vulnerability analysis checks those components against known vulnerability information.

The flow can be:

    Software Artifact
            ↓
    Generate SBOM
            ↓
    Components + Versions
            ↓
    Vulnerability Analysis
            ↓
    Known CVEs / Findings

This makes it easier to understand what software exists and whether known vulnerabilities affect it.

---

### 5. What information can an SBOM contain?

**Expected answer:**

An SBOM can contain information such as:

- Component or package name
- Version
- Identifiers
- Supplier or source information
- Dependency relationships
- Other metadata depending on the SBOM format

The exact fields can vary depending on the standard and tool being used.

---

### 6. Why is version information important in an SBOM?

**Expected answer:**

A package name alone is not enough to determine whether software is potentially affected by a vulnerability.

For example:

    Package A
    Affected versions: v1.0 to v1.5
    Fixed version: v1.6

If the SBOM contains:

    Package A v1.3

we can immediately see that the version falls within the affected range.

Therefore, the component name and version together are critical for vulnerability investigation.

---

### 7. What can Trivy do with an SBOM?

**Expected answer:**

Trivy can inspect artifacts such as container images and generate an SBOM containing the components and versions it discovers.

Conceptually:

    Docker Image
            ↓
    Trivy
            ↓
    Inspect components
            ↓
    Generate SBOM
            ↓
    Components + versions

Trivy can also perform vulnerability scanning, which is a separate function from generating the SBOM.

---

### 8. What is the difference between Trivy SBOM generation and Trivy vulnerability scanning?

**Expected answer:**

When generating an SBOM, Trivy answers:

> What components and versions are inside this artifact?

When performing vulnerability scanning, Trivy answers:

> Are any of these components affected by known vulnerabilities?

Therefore:

    SBOM Generation
        ↓
    Inventory

    Vulnerability Scanning
        ↓
    Security analysis

The same artifact can be used for both operations.

---

### 9. Why is generating an SBOM from the final Docker image useful?

**Expected answer:**

A final Docker image can provide a more complete picture of the actual artifact than looking only at the application's dependency file.

The final image can contain:

- Application dependencies
- Base image components
- Underlying operating system packages
- Other installed components

Therefore, generating an SBOM from the final image can capture components that are not visible from `go.mod` alone.

---

### 10. What are CycloneDX and SPDX?

**Expected answer:**

CycloneDX and SPDX are standardized formats for representing SBOM information.

They provide a common way for tools and systems to describe, exchange, and process software component inventory information.

Conceptually:

    Tool A
        ↓
    CycloneDX / SPDX
        ↓
    Standardized SBOM
        ↓
    Tool B can understand it

Without standards, every tool could generate its own custom format, making interoperability more difficult.

---

### 11. Why are standard SBOM formats important?

**Expected answer:**

Standard formats provide a common structure for representing and sharing SBOM information.

This allows different security tools, platforms, and systems to exchange and understand the same software inventory without requiring every integration to support a unique custom format.

---

### 12. How can Trivy generate an SBOM from a Docker image?

**Expected answer:**

A Trivy workflow can inspect a Docker image and generate an SBOM in a supported format.

For example:

    trivy image --format cyclonedx --output sbom.json my-go-app:latest

Conceptually:

    my-go-app:latest
            ↓
    Trivy
            ↓
    Inspect components
            ↓
    Generate CycloneDX SBOM
            ↓
    sbom.json

The SBOM contains an inventory of components and versions discovered inside the target artifact.

---

### 13. Why should SBOM generation be automated in CI/CD?

**Expected answer:**

Automating SBOM generation ensures that the inventory is generated consistently for the artifacts being built.

A typical flow is:

    Push / Pull Request
            ↓
    CI/CD Pipeline
            ↓
    Build Docker Image
            ↓
    Generate SBOM
            ↓
    Store or publish SBOM
            ↓
    Deploy artifact

This avoids relying on developers to manually generate the SBOM.

---

### 14. Why should the SBOM describe the artifact that was actually built?

**Expected answer:**

The final artifact can contain additional components beyond what is directly visible in the source repository.

For a Docker image, this can include:

- Application dependencies
- Base image components
- OS packages
- Other installed components

Generating the SBOM from the built artifact provides a more accurate inventory of what is actually packaged and potentially deployed.

---

### 15. How do SBOMs help when a new CVE is discovered?

**Expected answer:**

When a new CVE affects a component, stored SBOMs can be searched to identify artifacts containing that component and determine which versions they use.

For example:

    New CVE affects openssl
            ↓
    Search stored SBOMs
            ↓
    Find artifacts containing openssl
            ↓
    Check versions
            ↓
    Identify potentially affected applications
            ↓
    Prioritize remediation

This reduces the need to manually inspect every repository or application.

---

### 16. Does an SBOM automatically prove that an application is vulnerable?

**Expected answer:**

No.

An SBOM tells us that a component and version exist inside an artifact. Vulnerability analysis is still required to determine whether that component version is affected by a known vulnerability.

An SBOM helps identify **potentially affected** software and supports faster investigation.

---

### 17. What is the relationship between `go.mod` and an SBOM?

**Expected answer:**

`go.mod` declares the Go module and dependency requirements for the application.

An SBOM is an inventory of the components and versions inside a specific software artifact.

For example:

    go.mod
        ↓
    Application dependency requirements

    Final Docker Image
        ↓
    Application dependencies
    + Base image
    + OS packages
    + Other components
        ↓
    SBOM

Therefore, an SBOM generated from the final image can provide a broader inventory than `go.mod` alone.

---

### 18. What is the role of `go.sum`?

**Expected answer:**

`go.sum` records checksums used to verify downloaded Go module content.

It helps verify that downloaded module content matches the expected checksum information.

The simple distinction is:

    go.mod
        ↓
    What dependencies and versions does
    the application require?

    go.sum
        ↓
    What checksum information is used to
    verify downloaded module content?

---

### 19. What is the role of Snyk in this workflow?

**Expected answer:**

Snyk can analyze the application's dependency graph, including direct and indirect dependencies, and identify known vulnerabilities.

It can also help with remediation by identifying or proposing dependency changes when suitable fixes are available.

Conceptually:

    Go Application
            ↓
    Dependency Graph
            ↓
    Snyk
            ↓
    Known vulnerabilities?
            ↓
    Report / Remediation

---

### 20. What is the role of Trivy in this workflow?

**Expected answer:**

Trivy can inspect artifacts such as container images.

It can perform:

- Vulnerability scanning
- SBOM generation

Conceptually:

    Docker Image
            ↓
    Trivy
       ┌────┴────┐
       │         │
       ↓         ↓
    Vulnerability  SBOM
    Scanning       Generation
       ↓             ↓
    Known CVEs     Components
    + Severity     + Versions

---

## Scenario-Based Questions

### 21. A critical CVE affects `openssl` versions below 3.0.5. Your company has 100 containerized applications. How can stored SBOMs help?

**Expected answer:**

The security team can search the collected SBOM inventory to identify which applications or artifacts contain `openssl` and determine which versions they use.

This helps identify potentially affected applications without manually checking every application or contacting every team first.

The affected teams can then investigate and remediate the issue based on priority and impact.

---

### 22. An SBOM contains `Package A v1.3`, and a CVE affects Package A versions v1.0 through v1.5. Why is the SBOM useful?

**Expected answer:**

The SBOM provides both the package name and version.

Since `Package A v1.3` falls within the affected version range, the application can be identified as potentially affected and investigated further.

Without version information, we would not be able to make that comparison accurately.

---

### 23. Why is an SBOM not the same as a vulnerability report?

**Expected answer:**

An SBOM is an inventory of components and versions.

A vulnerability report requires vulnerability analysis to determine whether those components are affected by known CVEs.

The distinction is:

    SBOM
        ↓
    What do we have?

    Vulnerability Scan
        ↓
    Is what we have known to be vulnerable?

---

### 24. Why are CycloneDX and SPDX useful?

**Expected answer:**

They provide standardized formats for representing and exchanging SBOM information.

Because the structure is standardized, different tools and systems can understand and process the same SBOM more easily.

---

### 25. Why is generating an SBOM from a Docker image generally more useful than only reading `go.mod`?

**Expected answer:**

`go.mod` describes the application's Go dependency requirements.

The final Docker image can additionally contain the base image, underlying OS packages, and other installed components.

Therefore, an SBOM generated from the final image can provide a more complete inventory of the actual packaged artifact.

---

## Final Interview Answer

> **An SBOM, or Software Bill of Materials, is a standardized inventory of the software components and versions that make up an application or artifact. It helps organizations understand what is inside their software and respond faster when new vulnerabilities are discovered. For example, when a new CVE affects a package, stored SBOMs can be searched to identify which applications or artifacts contain that package and which versions they use. An SBOM itself is not a vulnerability report; vulnerability scanners such as Trivy or Snyk perform the analysis needed to identify known vulnerabilities. Trivy can also generate SBOMs from artifacts such as Docker images, providing visibility into application dependencies, base image components, OS packages, and other components included in the final artifact. Standards such as CycloneDX and SPDX make SBOM information easier to exchange and process across different tools and systems.**