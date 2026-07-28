# GitHub Actions - Day 50 Interview Questions

## 1. Why would you split a workflow into multiple jobs instead of putting everything into one job?

**Expected Answer:**
Multiple jobs improve readability, maintainability, isolation, and debugging. Each job runs on its own runner and can execute independently or in parallel.

---

## 2. What is the purpose of `needs`?

**Expected Answer:**
`needs` creates dependencies between jobs. It ensures a job waits for one or more other jobs to complete successfully before starting.

---

## 3. How does GitHub Actions execute jobs if `needs` is not specified?

**Expected Answer:**
Jobs run in parallel because each job gets its own GitHub-hosted runner.

---

## 4. What is a Matrix Strategy, and why is it useful?

**Expected Answer:**
A Matrix Strategy allows the same job to run with different combinations of operating systems, language versions, or configurations, reducing duplicate workflow code and improving test coverage.

---

## 5. What is the difference between `include` and `exclude` in a matrix?

**Expected Answer:**
- `exclude` removes unwanted combinations from the automatically generated matrix.
- `include` adds custom combinations that don't naturally exist or adds extra properties to existing combinations.

---

## 6. What is a reusable workflow?

**Expected Answer:**
A reusable workflow is a workflow triggered using `workflow_call` that can be invoked by other workflows. It helps avoid duplication and promotes reuse across repositories or workflows.

---

## 7. What is the difference between Inputs and Secrets in reusable workflows?

**Expected Answer:**
Inputs are used for normal configuration values and are passed using `with:`. Secrets contain sensitive information and are passed using `secrets:`. Secrets are masked and stored securely in GitHub.

---

## 8. Why doesn't a reusable workflow automatically receive repository secrets?

**Expected Answer:**
For security reasons. Secrets must be explicitly passed or inherited to prevent unintended access and follow the Principle of Least Privilege.

---

## 9. What is the difference between a Job-level `if` and a Step-level `if`?

**Expected Answer:**
A Job-level `if` skips the entire job without allocating a runner, while a Step-level `if` only skips the specific step after the job has already started.

---

## 10. Explain the difference between `fail-fast` and `continue-on-error`.

**Expected Answer:**
`fail-fast` applies to matrix jobs and cancels the remaining matrix jobs when one fails. `continue-on-error` applies to a specific step or job and allows the workflow to continue even if that step or job fails.

---

## 11. What is a DAG (Directed Acyclic Graph) in GitHub Actions?

**Expected Answer:**
A DAG represents the dependency graph between jobs using `needs`, allowing independent jobs to run in parallel while ensuring dependent jobs wait for all prerequisites.

---

## 12. When would you use `continue-on-error`?

**Expected Answer:**
For non-critical tasks such as experimental tests, optional security scans, or coverage uploads where collecting results is useful but failures should not block the pipeline.