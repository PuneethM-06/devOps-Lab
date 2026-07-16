## DAY 29 GIT 
### WHAT ARE GIT HOOKS
- Git hooks are scripts that Git automatically run when certain git event occur
- Example:
    - Before a commit
    - After a commit 
    - Before a push
    - After a merge

## PRE COMMIT HOOK
- A pre-commit hook runs when a user tries to perform a git commit 
- If the pre commit hook event succeds then you commit is recorded else it is not 
- They are used for:
    1. Syntax error
    2. Formatting issue
    3. Secrets
    4. Lint errors

- They are stored in `.git/hooks/`

### GIT HOOKS CAN BE BYPASSED?
- Yes, using `git commit --no-verify`

### LOCAL HOOKS VS CI
- Local hooks runs on our local machine and when a git event occurs
- While CI runs after a push 