# DAY 88 - gitLeaks 

1. ### WHAT IS GIT LEAKS?
- Gitleaks is a tool that is used to identify secrets that are accidentally in Git repositories and code 
- The core purpose is **Catch secrets before they become exposed in a Git repository, Git history, or deployed code.**

2. ### WHAT COUNTS AS A SECRET?
- A secret is a sensitive information that should not be exposed to unauthorized people 
- common example:
```
Secrets
│
├── API keys
├── Passwords
├── Database credentials
├── Access tokens
├── GitHub tokens
├── AWS access keys
├── Private keys
└── Connection strings containing credentials
```
> Config values such as ports etc, are not really secrets 

3. ### WHY .gitignore is not enough
- .gitignore can help secrets not going to github for untracked files (files inside .gitignore). But secrets can go when they are places inside a tracked file 

4. ### HOW GITLEAKS DETECTS SECRETS 
- It uses **detection rules to identiy a secret** 
1. **Known patterns**
- Some credentials follow known patterns 
- Example: 
```
AWS_ACCESS_KEY_ID=AKIA...
GITHUB_TOKEN=ghp_...
```

2.**Keywords and context**
- Variable names can provide context
```
password
secret
token
api_key
access_key
```
3. **ENTROPY**
- some secrets are long strings, which can be indicated as a secret 
- Example: `skdjf83jdKLSJDF9sdfKJ23klsdf...`

5. ### PRE-COMMIT HOOKS
- **A script or a check that automatically runs before Git creates a commit** 
```
Developer changes code
        ↓
git add
        ↓
Changes are staged
        ↓
git commit
        ↓
Pre-commit hook runs
        ↓
Gitleaks scans staged changes
        ↓
Secret found?
   ┌────┴────┐
   │         │
  No        Yes
   │         │
Commit      Commit blocked
continues
```
6. ### WHY DELETING A SECRET IS NOT ENOUGH 
- Because once it is committed, anyone access to the repo can see the secret in git history 

7. ### GIT LEAKS IN CI/CD
- A pre-commit hook is usefl but it cannot be the only defense because **pre-commit hooks can ebe bypassed**
```
Developer
    ↓
Pre-commit hook
    ↓
Gitleaks scans before commit
    ↓
Commit / Push
    ↓
GitHub Actions / CI
    ↓
Gitleaks scans again
    ↓
Secret found?
   ┌────┴────┐
   │         │
  No        Yes
   │         │
Pipeline    Pipeline fails
continues
```
8. ### WHAT HAPPENS IF A SECRET IS ALREADY LEAKED?
```
Secret leaked
      ↓
1. Revoke or rotate the secret
      ↓
2. Investigate where it was exposed
      ↓
3. Remove it from the code/repository
      ↓
4. Clean Git history if necessary
      ↓
5. Add controls to prevent it again
```
1. **Revoke or Rotate the secret**
- Revoke the old token and create a new one 

2. **Investigate where it was exposed**
- Ask for questions such as:
    - Public repo?
    - Who had access?
    - Secret was pushed to a remote repo
    - Is it copied to logs?

3. **Remove it from current code**
- remove the token from the code if it is accidentaly pushed 

4. **Clean git history**
- A better practise would be to clean repo history, but we still need to revoke the old token and create new one 
5. **Prevent Recurrence**
- Add layers such as
```
Gitleaks pre-commit hook
        +
Gitleaks CI scan
        +
Better secret management
```
9. ### FINAL SECURITY PIPELINE MENTAL MODEL
```
Developer writes code
        ↓
Gitleaks
        ↓
Check for accidentally exposed secrets
        ↓
Commit / Push
        ↓
Snyk
        ↓
Check application dependencies
for known vulnerabilities
        ↓
Build Docker Image
        ↓
Trivy
        ├── Vulnerability Scan
        │       ↓
        │   Check application dependencies,
        │   base image and OS packages
        │
        └── SBOM Generation
                ↓
            Components + Versions
```
| Tool         | Main purpose                                                |
| ------------ | ----------------------------------------------------------- |
| **Gitleaks** | Detect accidentally exposed secrets                         |
| **Snyk**     | Analyze application dependency vulnerabilities              |
| **Trivy**    | Scan artifacts such as container images for vulnerabilities |
| **SBOM**     | Inventory of components and versions inside an artifact     |
