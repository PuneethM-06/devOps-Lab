# Day 26 – Git Branching Strategy & Branch Protection (Interview Notes)

## What is a Git branch?
A Git branch is a movable pointer to a sequence of commits. It allows developers to work independently without affecting other branches until the changes are merged.

---

## Why use feature branches instead of committing directly to `main`?
- Keeps the `main` branch stable and production-ready.
- Allows isolated development.
- Enables code reviews through Pull Requests.
- Prevents incomplete or broken code from reaching production.
- Allows CI/CD checks to run before merging.

---

## Common Branch Types

- **main** – Production-ready code.
- **feature/** – New features.
- **bugfix/** – Non-critical bug fixes.
- **hotfix/** – Urgent production fixes.
- **release/** – Prepare a release before deployment.

---

## GitHub Flow
```
main
   ↓
Create feature branch
   ↓
Develop & Commit
   ↓
Push
   ↓
Open Pull Request
   ↓
Code Review + CI
   ↓
Merge to main
   ↓
Delete feature branch
```

---

## Branch Protection
Branch protection prevents unsafe changes to important branches like `main`.

Common rules:
- No direct pushes to `main`
- Pull Request required
- Required approvals
- CI checks must pass
- Branch must be up to date before merging
- Restrict force pushes and branch deletion

---

## Feature Branch vs Hotfix

| Feature Branch | Hotfix Branch |
|---------------|---------------|
| Develops new functionality | Fixes urgent production issues |
| Normal development cycle | High-priority production fix |
| Merged after review | Reviewed and deployed quickly |

---

## GitHub Flow vs Git Flow vs Trunk-Based Development

- **GitHub Flow:** Simple workflow using `main` and feature branches. Most common for modern development.
- **Git Flow:** Uses `develop`, `release`, and `hotfix` branches. Suitable for scheduled releases.
- **Trunk-Based Development:** Small, short-lived branches merged frequently into `main`. Common in organizations with strong CI/CD.

---

## Stacked Pull Requests
When one feature depends on another, create the next branch from the previous feature branch instead of `main`.

Example:
```
main
 └── feat/dockerfile
      └── feat/docker-compose
           └── feat/docker-security
```

After the earlier PR merges, update the base branch of the dependent PR to `main`.

---

## Interview Tips
- Never commit directly to `main`.
- Keep feature branches focused on a single change.
- Open Pull Requests early for review.
- Keep your feature branch updated with the latest `main`.
- Understand when to use feature, bugfix, and hotfix branches.