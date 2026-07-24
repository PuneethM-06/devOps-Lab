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

## NOTE:
1. Steps share files in a job 
2. Jobs never share files between them 
- Meaning, once a job is runner is destroyed after the job execution, file changes made in that runner or files from that runner cannot be accessed but the second job 

## STEP LOGS vs WORKFLOW LOGS 

1. ### WORKFLOW LOGS
- A workflow log is a complete log of what happened in the entire workflow 
- It includes;
    - Job creation 
    - Runner setup
    - Every step
    - Every command etc.

2. ### STEP LOGS
- A step log is the output of one step 

## EXIT CODES
- After executing every linux command, it returns a number which decides if the step is green or errored
- `EXIT CODE 0 `-  Sucess
- `ANY NON ZERO VALUE` - FAILURE

```
Run command
      │
      ▼
Exit Code?
      │
 ┌────┴────┐
 │         │
0       Non-zero
 │         │
 ▼         ▼
Continue  Fail Step
```

## OTHER USES
- sometimes the runner image that we use might not have all the dependencies we need and that is where we get other uses.
- Suppose I want to install node version 

```
- name: Install node version
  uses: actions/setup-node@v4
  with:
    node-version: 22
```
