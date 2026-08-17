# DAY 85 - TRIVY SCAN

### WHAT IS TRIVY SCAN 
- Trivy is a security scan that can scan container images, filesystem contents, repositories, dependencies and other targets for security issues such as known vulnerabilities.

2. ### WHY DO WE NEED TRIVY 
- **Application can be secure, but the components inside may not be and hence we need trivy to perform the scan**
-

3. ### WHAT DOES TRIVY ACTUALLY DO?
- When trivy scans for a container image, it automatically/magically does not know about the vulnerabilities inide. It first needs to **read the software inside the image**