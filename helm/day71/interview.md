# DAY 71 - HELM HOOKS 

1. ### WHAT IS A HELM HOOK 
- An helm hook specfies run this kubernetes resource at a specific point during a helm operation 
- Example: When we do `helm install myapp ./myapp`, helm rednders the chart and starts creating kubernetes resources and before that we might need to do something and that is where we have Helm hook coming into picture 
```
helm install
     ↓
Helm sees migration Hook
     ↓
Run migration Job
     ↓
Migration completes
     ↓
Deploy application resources
```
- Helm hook is executed before or during or after an helm operation 

2. ### HOOK ANNOTATION 
- Hook annotation lets helm know that a particular k8s resource is a hook 
- Example:
```
metadata:
  annotations:
    "helm.sh/hook": pre-install
```
- When a job is ran, helm recognizes that job asa pre-install hook and hence we should complete this before going with the helm operation 
- syntax is `helm.sh/hook`
- Example:
```
apiVersion: batch/v1
kind: Job

metadata:
  name: db-migration
  annotations:
    "helm.sh/hook": pre-install

spec:
  template:
    spec:
      restartPolicy: Never

      containers:
        - name: migration
          image: myapp:migrate
          command: ["./migrate"]
```
- **The above script has a pre-install, and hence this will be executed before the installation happens**

3. ### Pre-install vs Pre-upgrade
**pre-install** runs when we are doing initial installation 
```
helm install
     ↓
pre-install Hook
     ↓
Normal resources created
```
**pre-upgrade** runs when already an installation is happened and not the first time
```
helm upgrade
     ↓
pre-upgrade Hook
     ↓
Upgrade normal resources
```
4. ### Post install and post-upgrade
- **Post install** runs after installing normal resources
```
helm install
     ↓
Create normal K8s resources
     ↓
post-install Hook
```
- **Post upgrade** runs after the helm has upgraded the normal resources
```
┌──────────────────────────────────────┐
│           HELM OPERATION             │
│                                      │
│ pre-install  → install → post-install
│                                      │
│ pre-upgrade  → upgrade → post-upgrade
└──────────────────────────────────────┘
```
### Execution flow 
- **For an install**
```
helm install
     ↓
pre-install Hook
     ↓
Create normal Kubernetes resources
     ↓
post-install Hook
```
- **For an upgrade**
```
helm upgrade
     ↓
pre-upgrade Hook
     ↓
Update normal Kubernetes resources
     ↓
post-upgrade Hook
```
5. ### HOOK DELETION POLICY 
- Example:
```
annotations:
  "helm.sh/hook": pre-install
  "helm.sh/hook-delete-policy": hook-succeeded
```
- Helm gives us this feature, to ensure that the previous hoook resources dont exist around in the cluster 
- **hook-succeed** means delete this hook resource after the hook completes successfully 
