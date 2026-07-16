# Day 28 – Merge vs Rebase, Interactive Rebase & Squashing

## Merge vs Rebase

### Merge
- Combines two branch histories.
- Creates a **merge commit**.
- Preserves commit history.
- Safe for shared/public branches.
- Does not rewrite commit history.

### Rebase
- Replays your commits on top of the latest base branch.
- Creates a clean, linear commit history.
- Rewrites commit history by creating new commits with new commit hashes.
- Best suited for your own feature branch.

---

## When to Use Merge vs Rebase

**Use Rebase**
- Update your own feature branch with the latest `main`.
- Before opening or updating a Pull Request.
- When you want a clean, linear Git history.

**Use Merge**
- On shared branches.
- When you want to preserve the exact commit history.
- When rewriting history could affect other developers.

> **Golden Rule:** Never rebase a shared branch that other developers are actively using.

---

## Interactive Rebase

Interactive Rebase (`git rebase -i`) allows you to edit commit history before merging.

Common operations:
- `pick` – Keep the commit.
- `reword` – Change the commit message.
- `squash` – Combine multiple commits into one.
- `drop` – Remove a commit.
- Reorder commits.

---

## Commit Squashing

Squashing combines multiple related commits into a single meaningful commit.

Benefits:
- Cleaner Git history.
- Easier code reviews.
- Easier to revert changes.
- Removes unnecessary commits like `fix typo` or `forgot semicolon`.

Example:

Instead of:
```text
fix typo
forgot chmod
fix lint
```

Use:
```text
feat: add Docker Compose configuration
```

---

## Force Push After Rebase

After a rebase, Git creates new commits with new commit hashes.

A normal `git push` fails because the remote branch still has the old commit history.

Use:
```bash
git push --force-with-lease
```

Instead of:
```bash
git push --force
```

### Why `--force-with-lease`?
- Checks whether the remote branch has changed since your last fetch.
- Prevents accidentally overwriting another developer's work.
- Safer than `git push --force`.

---

## Interview Tips

- Merge preserves history; Rebase rewrites history.
- Rebase creates a clean, linear commit history.
- Never rebase shared branches.
- Interactive Rebase is used to clean commit history before merging.
- Squash related commits into one meaningful commit.
- After rebasing a pushed branch, use `git push --force-with-lease`, not `git push --force`.