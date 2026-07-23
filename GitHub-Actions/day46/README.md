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
