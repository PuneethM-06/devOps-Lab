# interview.md — Day 88: Gitleaks — Pre-Commit Hook Catching Leaked Secrets

## Must-Know Fundamentals

### 1. What is Gitleaks?

**Expected answer:**

Gitleaks is a security tool used to detect secrets accidentally exposed in source code or Git repositories.

Examples include:

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

A leaked credential can potentially provide unauthorized access to the associated cloud account, database, API, or service.

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
```

---

### 3. What counts as a secret?

**Expected answer:**

A secret is sensitive information that should not be exposed to unauthorized users.

Examples include:

- Database passwords
- API keys
- AWS access keys
- Access tokens
- GitHub tokens
- Private keys
- Connection strings containing credentials

```text
DATABASE_PASSWORD=my-password
AWS_ACCESS_KEY_ID=AKIA...
GITHUB_TOKEN=ghp_...
```

Normal configuration values such as `APP_PORT=8080` or `LOG_LEVEL=debug` are generally not secrets.

---

### 4. Why is `.gitignore` not enough to prevent secret leaks?

**Expected answer:**

`.gitignore` mainly prevents specified untracked files from being added to Git.

However, it does not protect secrets placed inside already tracked files.

```text
config.yaml ← Already tracked

DATABASE_PASSWORD=my-secret
```

Also, if a file was already tracked before being added to `.gitignore`, Git can continue tracking it.

---

### 5. How is Gitleaks different from `.gitignore`?

**Expected answer:**

`.gitignore` tells Git to ignore specified untracked files.

Gitleaks scans file content and Git changes for potential secrets.

```text
.gitignore
    ↓
Helps prevent specified untracked files
from being added

Gitleaks
    ↓
Scans content for potential secrets
```

They solve different problems.

---

### 6. How does Gitleaks detect potential secrets?

**Expected answer:**

Gitleaks uses detection rules to identify values that look like secrets.

Detection can involve:

- Known credential patterns
- Keywords and context
- Entropy
- Recognizable secret formats

For example:

```text
AWS_ACCESS_KEY_ID=AKIA...
GITHUB_TOKEN=ghp_...
DATABASE_PASSWORD=...
```

A suspicious value may be flagged for review.

---

### 7. What is entropy, and why can it help detect secrets?

**Expected answer:**

High entropy means a value appears highly random or unpredictable.

Generated secrets often look like long random strings, so entropy can help identify suspicious values.

However, high entropy does not prove that something is a secret.

It could also be:

- A hash
- A generated ID
- Test data
- Another random value

Therefore, findings may need review to determine whether they are real secrets or false positives.

---

### 8. What is a pre-commit hook?

**Expected answer:**

A pre-commit hook is a script or check that runs automatically before Git creates a commit.

```text
Developer changes code
        ↓
git add
        ↓
Changes staged
        ↓
git commit
        ↓
Pre-commit hook runs
        ↓
Gitleaks scans changes
        ↓
Secret found?
   ┌────┴────┐
   │         │
  No        Yes
   │         │
Commit      Commit blocked
continues
```

This helps catch secrets before they enter Git history.

---

### 9. Why is catching secrets during pre-commit useful?

**Expected answer:**

Catching a secret before the commit prevents it from entering Git history in the first place.

This avoids later cleanup such as:

- Revoking or rotating credentials
- Investigating exposure
- Removing secrets from code
- Cleaning Git history
- Investigating possible unauthorized access

Prevention is safer and easier than cleaning up an exposed secret.

---

### 10. Why is deleting a secret in a later commit not enough?

**Expected answer:**

Deleting the secret from the latest version of the file does not remove it from previous Git commits.

```text
Commit 1
↓
Secret added

Commit 2
↓
Secret deleted

Commit 3
↓
Clean code
```

The secret can still exist in the history of Commit 1.

Someone with access to the repository history may still be able to retrieve it.

---

### 11. What should you do if a real secret is leaked?

**Expected answer:**

Treat the leaked secret as potentially compromised.

```text
Secret leaked
      ↓
1. Revoke or disable old secret
      ↓
2. Create replacement if needed
      ↓
3. Remove secret from current code
      ↓
4. Investigate exposure
      ↓
5. Clean Git history if necessary
      ↓
6. Add controls to prevent recurrence
```

The most important early action is to revoke or rotate the exposed credential.

---

### 12. Why should a leaked secret be revoked or rotated?

**Expected answer:**

Once a secret has been exposed, it should be treated as potentially compromised.

Removing it from the repository does not make the old credential safe again.

Revoking or rotating it invalidates the exposed credential and reduces the risk of continued unauthorized use.

---

### 13. Why investigate where the secret was exposed?

**Expected answer:**

The secret may have been exposed in more places than just the source file.

For example:

- Remote Git repositories
- Git history
- CI/CD logs
- Forks or clones
- Other systems

Investigating exposure helps determine the potential impact and where additional cleanup may be needed.

---

### 14. Why is a pre-commit hook alone not enough?

**Expected answer:**

Local Git hooks can be skipped, bypassed, or misconfigured.

Therefore, relying only on a developer's local machine is not sufficient for enforcement.

A second independent layer should run in CI/CD.

---

### 15. Why run Gitleaks both as a pre-commit hook and in CI/CD?

**Expected answer:**

This provides defense in depth.

```text
Layer 1
↓
Pre-commit hook
Catch secrets early before commit

Layer 2
↓
CI/CD scan
Independently verify changes
```

The pre-commit hook provides fast feedback, while CI provides centralized and independent enforcement.

---

### 16. How can branch protection help enforce secret scanning?

**Expected answer:**

Gitleaks can run as a required CI check on pull requests.

```text
Pull Request
      ↓
Gitleaks CI check
      ↓
Pass?
 ┌────┴────┐
 │         │
No        Yes
│          │
Required    Merge can proceed
check fails
```

Branch protection can prevent changes from being merged into protected branches when required security checks fail.

---

### 17. What is the difference between Gitleaks, Snyk, and Trivy?

**Expected answer:**

They focus on different security problems.

```text
Gitleaks
   ↓
Detect accidentally exposed secrets

Snyk
   ↓
Analyze application dependencies
for known vulnerabilities

Trivy
   ↓
Scan artifacts such as container images
for known vulnerabilities
```

Trivy can also generate SBOMs.

These tools complement each other as different layers of a security pipeline.

---

### 18. How do Gitleaks, Snyk, Trivy, and SBOM fit together?

**Expected answer:**

They protect different parts of the software delivery process.

```text
Developer writes code
        ↓
Gitleaks
        ↓
Check for exposed secrets
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
   ┌────┴────┐
   ↓         ↓
Vulnerability SBOM
Scanning       Generation
```

An SBOM provides an inventory of components and versions that can later be used to investigate new vulnerabilities.

---

### 19. What is the overall purpose of using these security layers?

**Expected answer:**

The goal is defense in depth.

```text
Secrets
   ↓
Gitleaks

Dependencies
   ↓
Snyk

Final artifact
   ↓
Trivy

Software inventory
   ↓
SBOM
```

No single tool solves every security problem, so combining multiple layers provides better coverage.

---

## Scenario-Based Questions

### 20. A developer accidentally commits an AWS access key and deletes it in the next commit. What problem remains?

**Expected answer:**

The secret can still exist in the previous Git commit and repository history.

Deleting it from the latest version does not automatically remove it from earlier commits.

The exposed credential should be revoked or rotated because it should be treated as potentially compromised.

---

### 21. A `.env` file is listed in `.gitignore`, but a developer copies the database password into `config.yaml`, which is already tracked. Will `.gitignore` protect the secret?

**Expected answer:**

No.

`.gitignore` does not scan the contents of tracked files.

The secret can still be committed because it exists inside a tracked file.

Gitleaks can help detect the suspicious credential value.

---

### 22. Why might Gitleaks flag a value that is not actually a secret?

**Expected answer:**

Gitleaks uses detection rules such as patterns, keywords, context, and random-looking values.

A hash, generated ID, or test value may look similar to a secret and trigger a detection.

Therefore, findings may need to be reviewed for false positives.

---

### 23. Why is CI scanning still required if every developer has a Gitleaks pre-commit hook?

**Expected answer:**

Local hooks are not guaranteed enforcement mechanisms because they can be skipped, bypassed, or misconfigured.

CI provides an independent security check that applies consistently to changes submitted to the repository.

Together they provide defense in depth.

---

### 24. A real secret is detected after being pushed to GitHub. What should the response be?

**Expected answer:**

The exposed credential should first be revoked or disabled and replaced if necessary.

Then:

- Remove it from the current code.
- Investigate where it was exposed.
- Check systems such as CI logs.
- Clean Git history if necessary.
- Add or improve controls such as Gitleaks pre-commit hooks and CI scanning.

Deleting the secret alone is not enough because the old credential may already have been exposed.

---

## Final Interview Answer

> **Gitleaks is a security tool used to detect secrets accidentally exposed in source code or Git repositories, such as API keys, passwords, access tokens, AWS credentials, and private keys. It can run as a pre-commit hook to catch secrets before they enter Git history and can also run in CI/CD as an independent enforcement layer. This provides defense in depth because local hooks can be skipped or bypassed. If a real secret is already exposed, simply deleting it from the repository is not enough because it may still exist in Git history or other systems. The credential should be treated as compromised, revoked or rotated, and the exposure should be investigated. Gitleaks complements Snyk and Trivy by focusing specifically on secret detection, while Snyk focuses on dependency vulnerabilities and Trivy scans artifacts such as container images for vulnerabilities and can generate SBOMs.**

# Day 88 — Gitleaks: Pre-Commit Hook Catching Leaked Secrets

## What I Learned

- Learned what **Gitleaks** is and how it detects accidentally exposed secrets.
- Covered common secrets such as API keys, passwords, AWS credentials, GitHub tokens, access tokens, and private keys.
- Understood why `.gitignore` is not enough because secrets can exist inside tracked files.
- Learned how Gitleaks uses patterns, keywords, context, recognizable formats, and entropy to identify potential secrets.
- Understood that some findings can be false positives and may require review.
- Learned how **pre-commit hooks** can scan changes before Git creates a commit.
- Understood why preventing a secret from entering Git history is better than cleaning it up later.
- Learned that deleting a secret in a later commit does not remove it from earlier Git history.
- Covered the response to a leaked secret:
  - Revoke or disable the exposed credential.
  - Create replacement credentials if needed.
  - Remove the secret from current code.
  - Investigate where it was exposed.
  - Clean Git history if necessary.
  - Add controls to prevent recurrence.
- Learned why Gitleaks should run both locally and in CI/CD for **defense in depth**.
- Connected Gitleaks with GitHub Actions and branch protection.
- Compared the security tools:
  - **Gitleaks** → Detect exposed secrets.
  - **Snyk** → Analyze application dependencies for known vulnerabilities.
  - **Trivy** → Scan container images and other artifacts for vulnerabilities.
  - **SBOM** → Inventory of components and versions.

## Final Mental Model

```text
Developer writes code
        ↓
Gitleaks Pre-Commit Hook
        ↓
Check for accidentally exposed secrets
        ↓
Commit / Push
        ↓
Gitleaks in CI/CD
        ↓
Independent security check
        ↓
Snyk
        ↓
Check application dependency graph
for known vulnerabilities
        ↓
Build Docker Image
        ↓
Trivy
   ┌────┴────┐
   ↓         ↓
Vulnerability SBOM
Scanning       Generation
   ↓             ↓
Known CVEs     Components + Versions
```

## Key Takeaway

Gitleaks focuses on preventing secrets such as API keys, passwords, and tokens from entering Git repositories. Running it as both a pre-commit hook and a CI/CD check provides defense in depth: the local hook catches issues early, while CI independently enforces the security check. If a real secret is exposed, deleting it is not enough—the credential should be treated as compromised and revoked or rotated.