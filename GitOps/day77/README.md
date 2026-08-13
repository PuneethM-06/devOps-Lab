# DAY 77 - GITOPS

### HOW DOES ARGO CD KNOW SOMETHNG CHANGED IN GIT?
1. **POLLING.PERIODIC REFRESH**
- Argo does periodic checks on the git repository
```
Argo CD
   │
   ├── Check Git
   │
   ├── Check Git again
   │
   ├── Check Git again
   │
   └── Detect new commit
```
- The desired state may have changed. Let me fetch and regenerate it and changes in the k8s cluster 
