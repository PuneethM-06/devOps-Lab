# DAY 74 - ADVANCED HELM AND PRODUCTION PATTERNS

1. ### CONCEPT 1 - HELM UPGRAE ATOMIC
- *Command** - `helm upgrade --atomic`
- When we do a upgrade, and if something goes wrong then the release is left in broken or fail state and to over come this we need to make use of `upgrade --atomic`
- This treats as **all or nothing** - If anything goes wrong, the build automatically rols back to previous state 
```
Current Revision 2
       ↓
helm upgrade --atomic
       ↓
Try Revision 3
       ↓
Failure 
       ↓
Automatic rollback
       ↓
Revision 2 restored
```
2. ### `--wait` and `--timeout`
- `--wait` Helm waits for the resources to become ready before considering the operation successful
```
helm upgrade
    ↓
Create/update Deployment
    ↓
Kubernetes starts Pods
    ↓
Pods become Ready
    ↓
Helm considers operation successful
```
- `--timeout` 
- Imagine pods never become ready and be in pending state and we do not want helm to wait forever and hence we make use of --timeout 
- `helm upgrade myapp ./myapp --wait --timeout 5m`

3. ### HELM HOOKS IN PROD
- We know that helm hooks are kuberenetes resources that are executed at a particular point in the release cycle
```
pre-install
    ↓
Install resources
    ↓
post-install

pre-upgrade
    ↓
Upgrade resources
    ↓
post-upgrade
```
- There are also other hooks such as:
    1. pre-delete
    2. post-delete
    3. pre-rollback
    4. post-rollback 
    5. post-upgrade 

1. ### hook weight
- This is used to control the execution of hooks in an order when we have multiple hooks at the same lifecycle event 
- **lower weight runs first**
```
Hook A → weight -5
Hook B → weight 0
Hook C → weight 5

Execution:

A → B → C
```
4. ### HELM TESTS
- Helm lets you **define test for a deployed release**
- It tests the **deployed application/release**
- They are typically implemented as Kubernetes resources marked with the **helm.sh/hook: test** annotation and are executed using **helm test.**