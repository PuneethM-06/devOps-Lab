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
