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
