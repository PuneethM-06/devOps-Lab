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