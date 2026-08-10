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
