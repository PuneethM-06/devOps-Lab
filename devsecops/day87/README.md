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