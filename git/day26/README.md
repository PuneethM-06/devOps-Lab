## GIT DAY 26

### WHAT IS A BRANCH AND WHY DO WE NEED BRANCH
- Branch is another pointer to commits
- Imagine 20 developers working on the main branch, everyone edit's the code and hence it breaks with merge conflicts etc.

### DIFFERENT TYPES OF BRANCHES
1. MAIN - Production ready code goes here
2. FEATURE - Developing new enhancement goes here
3. BUGFIX - For fixing bugs
4. HOTFIX - Code changes or fixes that goes directly to prod
5. RELEASE - Release branch only receieve fixes before production 

## TRUNK BASED DEVELOPMENT
- Used by Meta, Google and Amazon 
- Developers continously create tiny features that live only for hours
- Requires execellent CI

## BRANCH PROTECTION
- This is github's most important security repository settings.
- Without this anyone can do, `git push origin main`
- This rule includes:
    1. No direct pushing to main 
    2. require PR before merging
    3. At least one approval required
    4. CI checks must pass
    5. Branch must be upto date with the base branch 
    6. Restrict force pushes 
    7. Restrict branch deletion

## GITHUB FLOW
```
main

↓

Create feature branch

↓

Commit

↓

Push

↓

PR

↓

Review

↓

Merge

↓

Delete branch
```

## Git Flow
```
main

↓

develop

↓

feature

↓

develop

↓

release

↓

main
```