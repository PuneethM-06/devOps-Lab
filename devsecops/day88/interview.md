# interview.md — Day 88: Gitleaks — Pre-Commit Hook Catching Leaked Secrets

## Must-Know Fundamentals

### 1. What is Gitleaks?

**Expected answer:**

Gitleaks is a security tool used to detect secrets that are accidentally exposed in source code or Git repositories.

It can help detect credentials such as:

- API keys
- Passwords
- Access tokens
- AWS credentials
- GitHub tokens
- Private keys

The goal is to catch secrets before they become exposed and create a security risk.

---

### 2. What problem does Gitleaks solve?

**Expected answer:**

Gitleaks helps prevent secrets from being accidentally committed and exposed in Git repositories.

For example:

```text
Developer adds secret
        ↓
git add
        ↓
git commit
        ↓
Gitleaks scans
        ↓
Potential secret detected
        ↓
Commit can be blocked