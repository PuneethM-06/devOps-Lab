# Day 30 – Git Stash, Git Blame & Git Bisect

## Git Stash

Temporarily saves **uncommitted changes** without creating a commit, allowing you to switch branches or work on another task.

### Common Commands

```bash
git stash
```
Save current changes.

```bash
git stash list
```
View saved stashes.

```bash
git stash pop
```
Restore and remove the latest stash.

```bash
git stash apply
```
Restore the latest stash without removing it.

```bash
git stash drop
```
Delete a stash.

---

## Git Blame

Shows **who last modified each line of a file**, along with the commit hash and modification date.

### Why Use It?

- Identify who last changed a line.
- Find the commit that modified it.
- Understand the context of a code change.
- Useful for debugging and code investigation.

> **Note:** `git blame` does **not** identify who introduced a bug—it only shows the last person who modified that line.

---

## Git Bisect

A debugging tool that uses **binary search** to identify the commit that introduced a bug.

### Workflow

```bash
git bisect start
```

Mark the current commit as bad:

```bash
git bisect bad
```

Mark a known good commit:

```bash
git bisect good <commit-hash>
```

Test each commit Git checks out and mark it as:

```bash
git bisect good
```

or

```bash
git bisect bad
```

Finish:

```bash
git bisect reset
```

---

## Useful Git Bisect Commands

Automatically run a test script:

```bash
git bisect run ./test.sh
```

Skip an untestable commit:

```bash
git bisect skip
```

---

## Interview Tips

### Git Stash
- Temporarily stores uncommitted changes.
- Useful when switching branches without creating temporary commits.
- `pop` restores and removes the stash.
- `apply` restores the stash without removing it.

### Git Blame
- Shows who last modified each line of a file.
- Helps identify the commit associated with a line.
- Used for debugging and understanding code history.

### Git Bisect
- Uses binary search to efficiently locate the commit that introduced a bug.
- Requires one known good commit and one known bad commit.
- `git bisect run` automates testing.
- `git bisect skip` skips commits that cannot be tested.
- Does **not** fix bugs; it only identifies the offending commit.