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
