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
