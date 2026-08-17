# DAY 85 - TRIVY SCAN

### WHAT IS TRIVY SCAN 
- Trivy is a security scan that can scan container images, filesystem contents, repositories, dependencies and other targets for security issues such as known vulnerabilities.

2. ### WHY DO WE NEED TRIVY 
- **Application can be secure, but the components inside may not be and hence we need trivy to perform the scan**
-

3. ### WHAT DOES TRIVY ACTUALLY DO?
- When trivy scans for a container image, it automatically/magically does not know about the vulnerabilities inide. It first needs to **read the software inside the image**
```
Docker Image
     ↓
Trivy examines the image
     ↓
Finds installed packages and their versions
     ↓
Example:
- openssl 1.x.x
- curl 7.x.x
- libc x.x
     ↓
Checks whether those specific versions have known vulnerabilities
     ↓
Reports the findings
```
- It reads the software, and packages inside then checks for CVE's against them 

### WHERE DOES TRIVY GET THIS VULNERABILITY INFORMATION FROM?
- Trivy gets vulnerability information from **vulnerability databases and security advisories**
```
Package + Version
      ↓
openssl 1.2.3
      ↓
Trivy checks vulnerability databases
      ↓
Does a known CVE affect this version?
      ↓
YES → Report the vulnerability
NO  → No matching finding
```
### UNDERSTANDING TRIVY SCAN RESULT 
- A trivy scanned result may look like 
```
Library: openssl
Installed Version: 1.2.3
Vulnerability: CVE-2024-12345
Severity: HIGH
Fixed Version: 1.2.4
```
1. **Library/package**
- Which component has been affected

2. **Installed version**
- Current version present 

3. **Vulnerability**
- This is the identifier for identifying the CVE

4. **Severity**
- This indicates the severity associated with that vulnerability 

5. **Fixed Version**
- This tells which version has a fix for this vulnerability 

### SEVERITY LEVELS - LOW, MEDIUM, HIGH, CRITICAL
```
LOW       → Lower severity
MEDIUM    → Moderate severity
HIGH      → Serious security concern
CRITICAL  → Very serious security concern
```
> Scanning only requirements.txt mainly checks application dependencies, while scanning the final Docker image can also detect vulnerabilities in the base image and OS-level packages included in the deployable artifact.
> Scanning the final Docker image gives us a more complete view because it includes both application dependencies and the underlying packages shipped in the image.

### RUNNING TRIVY SCAN
```
trivy image my-app:1.0
        ↓
Trivy accesses the image
        ↓
Examines its layers and installed components
        ↓
Identifies packages + versions
        ↓
Checks known vulnerability information
        ↓
Displays the findings
```

### EXIT CODES
- **Exit code 0 -> Success**
- **Non-Zero exit code -> Failure**
```
trivy image \
  --severity HIGH,CRITICAL \
  --exit-code 1 \
  my-python-app:latest
```
- If you find vulnerabilities with severity matching high, critical then fail the pipeline using exit code 1
```
Scan image
    ↓
Only consider HIGH and CRITICAL findings
    ↓
Any matching vulnerability found?
      │
   ┌──┴──┐
   │     │
  Yes    No
   │      │
exit 1  exit 0
   │      │
 Fail    Pass
 ```
 ### TRIVY IN GITHUB ACTIONS 
 ```
 PR / Push
    ↓
GitHub Actions runner
    ↓
Checkout repository
    ↓
Build Docker image
    ↓
Scan image with Trivy
    ↓
HIGH/CRITICAL found?
    ├── Yes → exit code 1 → Job fails
    └── No  → exit code 0 → Job passes
```
### BRANCH PROTECTION AND MERGE BLOCKING 
- Failing the flow/actions wont stop someone from merging the PR and we wll need **Branch protection rules for main**
```
Branch protection for main
        ↓
Require status checks to pass
        ↓
Trivy Security Scan must pass
Then 
Trivy scan fails
      ↓
Required GitHub Actions check fails
      ↓
❌ Merge blocked
```