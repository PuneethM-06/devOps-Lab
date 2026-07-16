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
