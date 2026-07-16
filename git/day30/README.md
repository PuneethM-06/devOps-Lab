## GIT DAY 30

## GIT STASH 
- git stash temporarily saves your uncommitted changes without creating a commit.

### GIT STASH COMMANDS
1. `git stash` - Saves work 
2. `git stash list` - view stashes
3. `git stash pop` - Restore latest stash
4. `git stash apply` - Restore without deleting 
5.  `git stash drop` - Delete a stash


## GIT BALME 
- Git blame shows who last modified each line of the file.

## GIT BISECT
- git bisect helps you understand the actual commit that introduced the bug using binary search 
- It checks the middle co mmit 
```
1  2  3  4  5  6  7  8
```
Git checks:
```
1 2 3 4 5 6 7 8
        ↑
```
- If commit 4 is good then the bug must be in 5,6,7,8
- Now get checks the middle again:
```
5 6 7 8
  ↑
```
- If commit 6 is bad then the bug is between:
```
5 6
```
git checks 5
