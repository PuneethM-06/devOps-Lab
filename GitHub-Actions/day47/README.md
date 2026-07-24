# DAY 47 OF GITHUB ACTIONS

## RUNNER FILESYSTEM 
- The flow is like this 
```
Your Laptop
      │
      ▼
GitHub receives push
      │
      ▼
Creates a fresh Ubuntu machine
      │
      ▼
Clones your repository
      │
      ▼
Runs every step
      │
      ▼
Deletes the machine
```

- So now here we are going to see where out repo will be cloned and will be executed 

- once the `ci` was updated to read the `pwd`, I can see the path coming to `/home/runner/work/devOps-Lab/devOps-Lab`

- `/home` - home directory for linux
- `/home/runner` - creates a user named runner
- `/home/runner/work` - github workspace where repo is checked out 
- `/home/runner/work/devOps-Lab/devOps-Lab` - a folder created where our repo has been cloned 