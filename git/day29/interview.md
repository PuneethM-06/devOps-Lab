# Day 29 – Git Pre-commit Hooks

## What is a Git Hook?

A Git hook is a script that Git automatically executes when specific Git events occur, such as before a commit, after a commit, before a push, or after a merge.

---

## What is a Pre-commit Hook?

A pre-commit hook runs automatically **before a commit is created**.

If the hook passes, the commit is created.

If the hook fails, the commit is blocked until the issues are fixed.

---

## Why Use Pre-commit Hooks?

- Catch issues before code is committed.
- Improve code quality.
- Reduce CI failures.
- Enforce coding standards.
- Prevent accidental commits of secrets or sensitive data.

---

## Common Pre-commit Checks

- Linting (ESLint, ShellCheck, golangci-lint)
- Code formatting (Prettier, gofmt, shfmt, Black)
- Unit tests
- Secret scanning
- YAML/JSON validation
- Trailing whitespace and end-of-file checks

---

## Where are Git Hooks Stored?

```text
.git/hooks/
```

Example:

```text
.git/hooks/pre-commit
```

---

## Pre-commit Hook vs CI Pipeline

| Pre-commit Hook | CI Pipeline |
|-----------------|------------|
| Runs before every local commit | Runs after code is pushed |
| Executes on the developer's machine | Executes on the CI server |
| Provides immediate feedback | Final quality gate before merge/deployment |

Most professional teams use both.

---

## Can Pre-commit Hooks be Bypassed?

Yes.

```bash
git commit --no-verify
```

This skips the pre-commit hook.

However, CI pipelines and branch protection rules still validate the code before it can be merged.

---

## Interview Tips

- Pre-commit hooks automate checks before a commit is created.
- Failed hooks block commits until issues are resolved.
- They improve code quality and reduce CI failures.
- Common checks include linting, formatting, tests, and secret scanning.
- `git commit --no-verify` bypasses local hooks but does not bypass CI or branch protection.