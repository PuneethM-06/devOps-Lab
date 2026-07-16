# Day 27 – Conventional Commits & Commit Discipline

## What are Conventional Commits?
A standard format for writing commit messages that makes Git history consistent, readable, and easier to automate.

**Format:**
```text
<type>: <description>
```

Optional:
```text
<type>(scope): <description>
```

Example:
```text
feat(auth): add JWT authentication
```

---

## Common Commit Types

| Type | Purpose |
|------|---------|
| feat | New feature |
| fix | Bug fix |
| docs | Documentation changes |
| style | Formatting only (no logic changes) |
| refactor | Improve code without changing behavior |
| test | Add or update tests |
| chore | Maintenance tasks |
| ci | CI/CD pipeline changes |
| build | Build system or dependency changes |
| perf | Performance improvements |

---

## Commit Message Best Practices

- Use the imperative mood.
- Keep the subject concise.
- One logical change per commit.
- Avoid vague messages like `update`, `fix`, or `changes`.

Good:
```text
fix: handle network timeout
```

Bad:
```text
fixed bug
```

---

## Atomic Commits

An atomic commit contains **one logical change** only.

Example:
```text
feat: add Docker Compose configuration
```

Not:
```text
feat: add Docker Compose, update README and fix tests
```

---

## Why Companies Use Conventional Commits

- Easier code reviews
- Cleaner Git history
- Better release notes
- Simpler debugging
- Easier to revert changes
- Enables automated versioning and changelog generation

---

## Interview Tips

- Use meaningful commit messages.
- Commit frequently in small logical units.
- Every commit should represent one purpose.
- Prefer `feat`, `fix`, `docs`, `refactor`, etc., over generic messages like `update` or `misc`.