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
