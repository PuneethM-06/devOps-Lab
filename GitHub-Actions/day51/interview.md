# GitHub Actions - Day 51 Interview Questions

## 1. What are Environment Variables (`env`) in GitHub Actions?

**Expected Answer:**
Environment variables are values defined at the workflow, job, or step level that can be reused throughout the workflow. They help avoid hardcoding values and improve maintainability.

---

## 2. What is the precedence order of environment variables?

**Expected Answer:**
Step `env` > Job `env` > Workflow `env`.

The most specific scope overrides the broader scope.

---

## 3. What is the difference between `env`, `vars`, and `secrets`?

**Expected Answer:**

- **env** → Workflow-specific variables defined inside YAML.
- **vars** → Non-sensitive configuration stored in GitHub Repository/Organization Variables.
- **secrets** → Sensitive information such as passwords, API keys, and tokens stored securely in GitHub Secrets.

---

## 4. When should you use Repository Variables instead of Secrets?

**Expected Answer:**
Repository Variables should be used for non-sensitive configuration such as application names, AWS regions, ports, and Docker image names. Secrets should only be used for confidential information.

---

## 5. What is the `GITHUB_TOKEN`?

**Expected Answer:**
`GITHUB_TOKEN` is a temporary token automatically created by GitHub for every workflow run. It allows workflows to authenticate with GitHub APIs based on the permissions granted to it.

---

## 6. Why should workflow permissions be explicitly defined?

**Expected Answer:**
Explicit permissions follow the Principle of Least Privilege by granting workflows only the permissions they require, reducing the impact of compromised workflows or third-party actions.

---

## 7. What is the Principle of Least Privilege?

**Expected Answer:**
It is the security principle of granting users, applications, or workflows only the minimum permissions required to perform their tasks and nothing more.

---

## 8. What are GitHub Environments?

**Expected Answer:**
Environments represent deployment targets such as Development, Staging, and Production. They can contain environment-specific secrets, variables, and deployment protection rules.

---

## 9. What are Environment Protection Rules?

**Expected Answer:**
Environment Protection Rules secure deployments by enforcing policies such as required reviewers, branch restrictions, and wait timers before deployments can proceed.

---

## 10. What is a Manual Approval Gate?

**Expected Answer:**
A Manual Approval Gate pauses a deployment until an authorized reviewer approves it. It is commonly used before production deployments.

---

## 11. What is the difference between Branch Protection Rules and Environment Protection Rules?

**Expected Answer:**

- Branch Protection Rules protect source code before merging into protected branches.
- Environment Protection Rules protect deployments after code has been merged by enforcing deployment approvals and restrictions.

---

## 12. Name some GitHub Actions security best practices.

**Expected Answer:**

- Never hardcode secrets.
- Use Repository Variables for non-sensitive configuration.
- Store sensitive data in GitHub Secrets.
- Follow the Principle of Least Privilege.
- Explicitly define workflow permissions.
- Pin GitHub Action versions.
- Use trusted third-party actions.
- Protect production environments with approval gates.
- Separate development and production credentials.

---

## 13. How would you design a production-grade CI/CD pipeline?

**Expected Answer:**
Checkout → Build → Lint/Test/Security Scan (parallel) → Package → Publish Artifact/Container → Deploy to Staging → Smoke Tests → Manual Approval → Deploy to Production → Notifications.