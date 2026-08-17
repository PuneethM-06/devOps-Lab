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
