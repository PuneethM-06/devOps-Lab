# GITHUB ACTIONS - CI/CD

## CONTINIOUS INTEGRATIONS
- It is the process of **validating every code change by building, testing and verifying before it is merged or deployed**

## CONTINOUS DELIVERY
- It is the **process of automatically preparing a validate code change for deployment/release.**

- Example: **CI** is responsible for **validating** the code changes and deciding if this code is good enough to be merged to the main branch. While **CD** is responsible for **DELIVERING** to the users by building artifacts, building Docker images etc.

## WHAT IS GITHUB ACTIONS?
- Every `yml` file is a workflow when we talk wrt to github actions 

### WHAT IS A WORKFLOW
- A workflow is an automated process that GitHub executes whenever there is an event occured

1. **WORKFLOW** - It is the entire automation script
2. **EVENT** - When should the workflow start to run and execute
3. **RUNNER** - Once the event is recognised where will the workflow script be executed 
4. **JOB** - A workflow can have one or more jobs in a workflow script 
5. **STEPS** - These are "steps" inside every job 
6. **ACTIONS** - Actions are reusable code written by GitHub or the community that can be reused

## OVERALL WORKFLOW
```
You push code
        │
        ▼
Event occurs
(push)
        │
        ▼
Workflow starts
(ci.yml)
        │
        ▼
GitHub creates Runner
(Ubuntu VM)
        │
        ▼
Job starts
(Build)
        │
        ▼
Step 1
Checkout Code
        │
        ▼
Step 2
Install Python
        │
        ▼
Step 3
Run Tests
        │
        ▼
Step 4
Build Application
        │
        ▼
Workflow finishes
```

## GITHUB ACTIONS VOCABULARY 

1. **name** - gives your workflow human readable name 
2. **on** - says when should the workflow start
3. **job** - These are tasks that will be executed once the workflow stars executing
4. **runs-on** - Tells where the jobs should run Example: ubuntu-latest, windows-latest, macos-latest
5. **steps** - organized way of executing commands or following the workflow in a job 
6. **uses** - Execute a resuable action created by GitHub or the community
7. **run** - Used to execute commands 
8. **with** - When we use an action, we need input with the action and that's when we use with 

## COMMON EVENTS

| Event               | Meaning                    | Typical Use                 |
| ------------------- | -------------------------- | --------------------------- |
| `push`              | Code is pushed             | CI                          |
| `pull_request`      | PR opened/updated          | CI before merge             |
| `workflow_dispatch` | Manual trigger             | Manual deployments          |
| `schedule`          | Cron schedule              | Nightly jobs                |
| `release`           | GitHub Release created     | Publish artifacts           |
| `workflow_call`     | Called by another workflow | Reusable workflows (Day 50) |

## PRODUCTION READY WORKFLOW 

```
Feature Branch
       │
       ▼
Push
       │
       ▼
CI Runs ✅

──────────────

Open PR
       │
       ▼
CI Runs Again ✅

──────────────

Merge
       │
       ▼
Main Branch
       │
       ▼
CD Starts 🚀
```
