# DAY 72 - HELM UPGRADE, ROLLBACK AND DIFF

1. ### HELM UPGRADE
- `helm install myapp ./myapp` creates the first installation and creates a helm release 
- We run helm upgrade after the first installation, because we want to upgrade/update the existing values and not create a new one
```
helm install
    ↓
Create a new release

helm upgrade
    ↓
Modify an existing release
```
2. ### HELM REVISION 
- This is the most important because **rollback only makes sense because Helm keeps revisions of a 
release**
- **Definition** - A Helm revision is a version of Helm release created when the release is installed or upgraded
- Example:
```
helm install myapp ./myapp
        ↓
Revision 1

helm upgrade myapp ./myapp
        ↓
Revision 2

helm upgrade myapp ./myapp
        ↓
Revision 3
```
