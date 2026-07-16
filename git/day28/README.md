## GIT DAY 28

## WHY DO WE NEED MERGE OR REBASE
- Image this history:
```
A --- B --- C (main)
```
while you're working another developer merges two commits:
```
A --- B --- C --- F --- G (main)
           \
            D --- E (feature)
```
Now our branch is behind main and we have 2 choices here:
1. Merge
2. Rebase

## WHAT IS MERGE
- Merge combines two histories
- Nothing is re-written and history is preserved exactly 
```
A --- B --- C --- F --- G
           \             \
            D --- E ------ M
```
### ADVANTAGES
1. Never rewrites history
2. safe for shared branches
3. Easy to understand 

### DISADVANTAGES
1. History becomes messy 
```
Merge branch 'main'

Merge branch 'main'

Merge branch 'main'
```

## WHAT IS REBASE
- Instead of creating a merge commit
- Git moves your commits on top of the main 
- Before:
```
A --- B --- C --- F --- G
           \
            D --- E
```
After:
```
A --- B --- C --- F --- G --- D' --- E'
```
- History becomes linear

### ADVANTAGES:
1. Very clean history
2. Looks linear
3. Perfect for feature branches

### DISADVANTAGES
1. It writes history 
2. Never rebase a branch other people are working on

| Merge                     | Rebase                           |
| ------------------------- | -------------------------------- |
| Creates a merge commit    | Rewrites commits                 |
| Preserves history         | Creates linear history           |
| Safe for shared branches  | Best for your own feature branch |
| Doesn't change commit IDs | Creates new commit IDs           |

### SQAUSHING COMMITS
- To take multiple commit records and combine them to a single commit is called as squashing commits
