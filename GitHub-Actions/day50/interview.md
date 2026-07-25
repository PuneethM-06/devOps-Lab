# DAY 50 - GITHUB ACTIONS CONTINUATION 

## MULTIPLE JOBS
- Every job gets it's own brand new VM
- GitHub runners are ephemeral 
- Multiple jobs give:
    - Isolation 
    - scalability
    - clearer logs
    - easier debugging

### WHY DO WE NEED `actions/checkout` IN MULTIPLE JOBS AND ISNT IT REDUDANT?
- Because every job has it's own new runner which wouldn't have the files cloned or the repo cloned as they start and hence we would need it seperately

## needs
- **needs create DAG - Dependency Graph**
- As we know, without needs GitHub actions starts all the jobs almost at the same time since each job is isolated and each job get their own runner.
- But in some sitations it is needed that a particular job should start after one job is done executing and hence we need needs here 
- needs controls the order only and not share files between jobs 
- A j**ob can wait for more than 1 jobs**

## MATRIX STRATERGY
- **A matrix stratergy lets you run your job against multiple combinations**
- Lets say we have a customer who makes use of windows, ubuntu and also macos. And then they also want to run tests on Python 3.10, Python 3.11 and Python 3.12
- So here we need to create a seperate job and seperate python version for each and that would make a mess of jobs and that is where we can use **matrix**
- **We write one job making use of matrix startergy and it creates all possible combinations**
- Example:
```
stratergy:
    matrix:
        python-version:
            - "3.10"
            - "3.12"
```
- Adding multiple OS:
```
stratergy:
    matrix:
        os:
            - ubuntu-latest
            - windows-latest
        python-version:
            - "3.10"
            - 3.12"
```
- Each combination gets its own runner 

#### exclude
- Remove specific combinations from the automatically generated matrix
```
stratergy:
    matrix:
        os:
          - ubuntu-latest
          - windows-latest
        node-version:
            - 18
            - 20
        exclude:
            - os: windows-latest
            - node-version: 20
```
- The flow is:
    1. Generate ALL combinations
    2. Apply exclude
    3. Give back the final results

#### include
- **Add special combinations that arent a part of the normal cartesian product** 
- This solves the opposite problem of exclude
```
stratergy:
    matrix:
        - windows-latest
        - ubuntu-latest
    node-version:
        - 18
        - 20 
    include:
        - os: windows-latest
        - node-version: 22
```

#### fail-fast
- decides whether github should cancel the remaining matrix jobs if one fails 
- By default:
```
fail-fast: true
```
- By this, **GitHub cancels the remaining matrix jobs once a failure occurs, provided they are still running or waiting**
- with `fail-fast: false`, it ensures that once a job fails, the other jobs continue to execute and the results are given out 

#### max-parallel
- we can restrict the number of jobs running maximum.
```
stratgery:
    matrix:
        - windows-latest
        - ubuntu-latest
    max-parallel: 2

    node-version:
        - 20
        - 18
```
### SUMMARY 
| Feature        | Purpose                                                                |
| -------------- | ---------------------------------------------------------------------- |
| `exclude`      | Remove specific combinations from the generated matrix.                |
| `include`      | Add extra or custom combinations (and optional extra variables).       |
| `fail-fast`    | Stop remaining matrix jobs after the first failure (default behavior). |
| `max-parallel` | Limit how many matrix jobs execute concurrently.                       |

## RESUSABLE WORKFLOWS
- Lets say we have a common task for all the workflows we are doing and it might be cloning a repo.
- This looks okay for few repositories but what we are working in a an organization, a small change will lead to a change in all the exisiting repo adn hence it is not feasible and that is where we get the concept of Resuable workflow 
```
Reusable Workflow

        ▲
        │
Repo A──┤

Repo B──┤

Repo C──┘
```
> A reusable workflow is a GitHub Actions workflow that can be called by other workflows using workflow_call. It allows multiple repositories or workflows to share a common CI/CD pipeline, reducing duplication and making maintenance easier. Instead of updating the same workflow in many repositories, you update the reusable workflow once, and every workflow that calls it automatically benefits. This follows the DRY (Don't Repeat Yourself) principle.

### RESUABLE WORKFLOW IN CODING 
- **WITHOUT RESUABLE WORKFLOW**
```
name: without resusable workflow

on:
    push:

jobs:
    - name: Job 1
```

- **WITH REUSABLE WORKFLOW**
```
name: with workflow

on:
    -workflow_call
jobs:
    - name: Job 2
```
- ***INSIDE THE FOLDER**
```
.github/
└── workflows/
    ├── ci.yml
    └── reusable-build.yml
```

