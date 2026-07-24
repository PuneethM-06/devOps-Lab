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

### WHY IS DEVOPS-LAB TWICE
```
/home/runner/work/
│
└── devOps-Lab/          ← Workspace for your repository
      │
      └── devOps-Lab/    ← Actual cloned repository
            │
            ├── README.md
            ├── .github
            ├── docker
            ├── linux
            └── ...
```

- when a workflow starts it will be not having access to our repo and that is why the first step is to have `- uses: actions/checkout@v4`

- similarly we can do `ls` or `ls -la` to inspect the repo

- **find** - It performs recursive search through directories
