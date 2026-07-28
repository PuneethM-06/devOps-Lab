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

## Reusable Workflows

### Why?

Used when multiple repositories or workflows share the same CI/CD pipeline.

Instead of copying the same workflow YAML everywhere, create one reusable workflow and let other workflows call it.

This follows the DRY (Don't Repeat Yourself) principle.

---

### Reusable Workflow

```yaml
name: Reusable Build

on:
  workflow_call

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - run: echo "Building..."
```

---

### Calling Workflow

```yaml
name: CI

on:
  push

jobs:
  build:
    uses: ./.github/workflows/reusable-build.yml
```

---

### Key Points

- Triggered by `workflow_call`
- Cannot start on its own
- Must be called from another workflow
- Used to reuse complete workflow logic (jobs)
- Reduces duplication and improves maintainability

### PASSING VARIABLES TO REUSABLE WORKFLOWS
- Here instead of hardcoding the value of the reusable workflow we can pass variables to it 

- Inside `resuable-build.yml`
```
name: Reusable workflow
on:
    workflow_call
        inputs:
            enviornment:
                required: true
                type: string 

jobs:
    build:
        runs_on: ubuntu-latest
        
        steps:
            - name: echo" Building {{ inputs.enviornement }}"
```

- Inside `ci.yml`
```
name: Reuising the resuable build
on:
    push

jobs:
    build:
        uses: ./.github/workflow/reusable-build.yml
    with:
        enviornment: dev
```

## SECRETS IN RESUABLE WORKFLOW
- The worst part is **hardcoding** the secrets instead.\
- We can store it in *SECRETS* and what we can do is, `secrets.AWS_ACCESS_ID`

### HOW IT CAN BE DONE USING CODING 
- In `resuable-build.yml`
```
name: reusable build 

on:
    workflow_call:
        inputs:
            environment:
                required: true 
                type: string 
            
            secrets:
                aws-key:
                    required: true
jobs:
    build:
        runs-on: ubunutu-latest
    
    steps:
        - name: Deploying
            run: echo" Deploying ${{ inputs.environment }}
        
        - name: AWS S3
            env:
                AWS_ACCESS_KEY: ${{ secrets.aws-key }}
```
- Inside the `ci.yml`

```
name: Inside ci.yml
on:
    push

jobs:
    deploy:
        uses: ./.github/workflow/resuable-build.yml
        
        with:
            environment: production
         secrets:
            aws-key: ${{secrets.AWS_ACCESS_KEY}}
```
## Job Outputs

- Every job runs on its own isolated runner.
- Because of this, jobs cannot directly share variables or data with each other.
- To pass small pieces of data (such as a version, Docker image tag, commit SHA, or URL) from one job to another, GitHub Actions provides **Job Outputs**.

Example:

Job A
   ↓
Version = v1.2.3
   ↓
Job Output
   ↓
Job B

---

## Step Outputs

- A job can contain multiple steps.
- Before a value can become a Job Output, it is first created as a **Step Output**.
- The job then exposes that Step Output as a Job Output so that other jobs can access it.

Flow:

Step
   ↓
Step Output
   ↓
Job Output
   ↓
Another Job

### HOW DO WE DO IT CODE
- Here:
    1. We have to mark a particular step with `id`
    2. We have to push that to the next step using `$GITHUB_OUTPUT`
- Example:
```
steps:
    - name: Generate version
        id: build-info

        run: |
            echo "docker-tag=v1.2.3" >> $GITHUB_OUTPUT
```
- This is the step outpout and can another job access it yet. NO

- And we can do it by using:
```
jobs:
    builds:
        outputs:
            version: ${{ steps.version.outputs.version}}
```
### TO BE CONTINUES NOT DOING IT NOW 

## CONDITIONAL EXECUTION (if)
- We can use it on `jobs` and we can use it on `steps`
### JOB LEVEL
```
jobs:
    deploy:
        if: github.ref == 'refs/heads/main'
```
- This tells do the deployment if it is `main` branch 

### STEP LEVEL
```
-name: Notify slack
    if: failure()
    run: ,,,
```

## continue-on-error
- Here it is similar to what we learnt on `fail-fast`.
-  But the catch is, `fail-fast` is for matrix while `continue-on-error` is for both steps and jobs 

```
name: Continue On Error Demo

on:
  push:

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
      - name: Build
        run: echo "Building..."

      - name: Experimental Test
        continue-on-error: true
        run: exit 1

      - name: Deploy
        run: echo "Deploying..."
```
