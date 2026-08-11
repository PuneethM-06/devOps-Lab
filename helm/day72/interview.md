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
3. ### HELM ROLLBACK
- It is the process to rollback the release of a version for a certain reason 
```
helm upgrade
     ↓
new revision

helm rollback
     ↓
take an older revision's configuration
     ↓
apply it as a new revision
```
- Syntax: `helm rollback myapp 2`
- Making a Helm rollback reconciles the resources in k8s meaning the pods also reconcile to start running their previous version of pods 
> Helm rollback → restores the previous release configuration → Kubernetes reconciles the resources → Pods end up running the previous version if the workload configuration changed.

4. ### HELM DIFF
- It is responsible for showing what will change -> What exactly is helm gonna change 
- This is not default and comes through **helm plugin** 

5. ### DEBUGGING AND INSPECTING A HELM RELEASE
This is useful for understand debugging when a release doesnt work as expected 
- ### STEPS TO FOLLOW       
1. **Helm status** - `helm status myapp` 
- This helps you understand the current status of the release

2. **HELM GET** - `helm get values myapp`
- Helm also lets you inspect the information stored for the release

3. **HELM TEMPLATE** - `helm template myapp ./myapp`
- It renders the helm template locally and thus printing the kubernetes yaml 
- It is like `terraform plan` and it does not do install any resources in k8s 

4. **DRY RUN** - `helm upgrade myapp ./myapp --dry-run`
- Simulate an install or upgrade without actually installing it 
