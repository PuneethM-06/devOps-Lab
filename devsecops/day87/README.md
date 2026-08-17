# DAY 87 - SBOM 

1. ### WHAT IS SBOM
- **SBOM - SOFTWARE BILL OF MATERIALS**
- It is basically an **inventory of components that make up an application or artifact** 

2. ### WHY DO WE NEED SBOM 
- Suppose a new CVE is out for package X and we have hundreds of applications and containers.
- Checking for each of these applications and containers can be a tedious task and hence we here we can make use of SBOM
- SBOM helps us understand which application is using those packages/versions and we go directly and make the fix 
```
New CVE announced
        ↓
Affected package identified
        ↓
Check SBOM inventory
        ↓
Which applications contain it?
        ↓
Identify affected versions
        ↓
Prioritize remediation
```
3. ### SBOM vs VULNERABILITY SCANNING 
- SBOM answers what versions and packages we have in our applications and containers running 
- Vulnerability scanner helps us understand which package/version has a CVE reported 

4. ### WHAT DOES AN SBOM CONTAIN 
- As we know SBOM is a inventory of components and versions/packages that applications contain
```
Component / Package Name
        ↓
Version
        ↓
Identifier
        ↓
Supplier / Source
        ↓
Dependency relationship
```
- SBOM can also contain **dependency relationship** to see how each package is related to each other 

5. ### GENERATING AN SBOM WITH TRIVY
- Suppose we have 
```
Docker Image
│
├── Your Go application
├── Go dependencies
├── Base image components
├── OS packages
└── Other installed components
```
- Trivy scans for these and creates an inventory of list of packages and versions an application is using as an SBOM
```
Docker Image
      ↓
Trivy inspects components
      ↓
Identifies packages + versions
      ↓
Generates SBOM
```
 > command is: `trivy image --format cyclonedx --output sbom.json my-go-app:latest`
```
trivy image
    ↓
Scan/inspect the container image

--format cyclonedx
    ↓
Generate the SBOM in CycloneDX format

--output sbom.json
    ↓
Save the generated SBOM to a file

my-go-app:latest
    ↓
Target image
```
6. ### SBOM FORMATS:CycloneDX and SPDX
- SBOM is an inventory and there must be a standard format that the is needed for understanding and exchanging and hence we have:
    1. CycloneDX
    2. SPDX
- CycloneDX is commonly used in software supply-chain and security workflows
```
Component
Version
Dependencies
Identifiers
Metadata
```
- SPDX is also widely use for describing components and related information 
7. ### WHY SBOMs MATTER WHEN NEW CVE IS DISCOVERED?
```
New CVE affects Package X
        ↓
Check SBOM inventory
        ↓
Which applications/artifacts contain Package X?
        ↓
Which versions are they using?
        ↓
Identify potentially affected applications
        ↓
Prioritize remediation
```
8. ### SBOM IN CI/CD
```
Go application
      ↓
Build Docker image
      ↓
my-go-app:latest
      ↓
Trivy generates SBOM
      ↓
sbom.json
```
- The core reason for doing this is **automation and consistency**
- **Actual workflow**
```
Push / Pull Request
        ↓
GitHub Actions
        ↓
Build Docker image
        ↓
Trivy generates SBOM
        ↓
Save SBOM as artifact / publish it
        ↓
Continue pipeline
```

