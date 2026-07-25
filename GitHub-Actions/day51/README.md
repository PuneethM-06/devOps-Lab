## DAY 51: GITHUB ACTIONS 

## ENVIRONMENT VARIABLES 
- we are setting environment variables 

### WORKFLOW LEVEL ENV
```
name: CI
on: 
    push
env:
    APP_NAME: linux_sysmonitor
```
- APP_NAME can be accessed at the workflow level 

### JOB LEVEL ENV
```
name: CI
on:
    push
jobs:
    buid:
        env:
            NODE_ENV: Production 
```

### STEP LEVEL ENV
```
name: CI

on:
    push

jobs:
    build:
        steps:
            - name: App name 
                env: 
                    APP_NAME = linux
                run: echo $APP_NAME
```

> What happens if the same environment variable is defined at all three levels?
```
Workflow:
APP_NAME=App1

Job:
APP_NAME=App2

Step:
APP_NAME=App3
```
> Which value does the step actually use?
- Answer: App3

## REPOSITORY VARS
- Repository-wide configurations managed in GitHub settings
- `secrets` are encrypted while `variables` aren't
- This is the only difference, Vars arent encrypted while secrets are. The way of accessing `vars.APP_NAME` and all remains the same 

## PERMISSIONS IN GITHUB ACTIONS
- **Permissions are given to the entire workflow and can also be done to the jobs level**
```
name: CI
on:
    push
permissions:
    contents:read
```

> Why should you explicitly define workflow permissions?
- Answer: To ensure the workflow receives only the **permissions it actually needs**. This reduces the impact of compromised workflows or third-party actions and follows the **Principle of Least Privilege.**

> Should permissions be defined at the workflow level or the job level?
- Answer: Use workflow-level permissions when the same permissions apply to all jobs. Use job-level permissions when specific jobs require additional or different access. This follows the Principle of Least Privilege by granting elevated permissions only to the jobs that need them.

> Use workflow-level permissions when the same permissions apply to all jobs. Use job-level permissions when specific jobs require additional or different access. This follows the Principle of Least Privilege by granting elevated permissions only to the jobs that need them./
- Answer: 
- Branch Protection Rules protect the **source code** by requiring reviews, status checks, and other conditions before code can be merged into a protected branch like main.
- Environment Protection Rules protect **deployments** by requiring approvals, branch restrictions, wait timers, or other checks before code is deployed to environments like Staging or Production.