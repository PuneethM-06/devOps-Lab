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
