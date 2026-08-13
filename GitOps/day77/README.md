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

2. ### WEBHOOKS
- Instead of waiting for argo CD's next check, we can notify Argo CD by web hooks
```
Developer
    ↓
git push
    ↓
GitHub
    ↓
Webhook sent to Argo CD
    ↓
Argo CD refreshes application
    ↓
Detects new desired state
    ↓
Compares with Kubernetes
```
- *Webhooks doesnt deploy the application, it is responsible only for notifying the argoCD about it*

2. ### SELF HEALING 
- **AUTO-SYNC IS TRIGGED WHEN DESIRED STATE CHANGES**
- **SELF HEALING IS TRIGGERED WHEN THE ACTUAL STATE CHANGES AUTOMATICALLY**
- Example: Someone manually changed the actual state using command `kubectl scale deployment my-app --replicas=1`, now self healing kicks in and then fixes it.
```
Someone manually changes Kubernetes
            ↓
Actual state changes
            ↓
Argo CD detects drift
            ↓
Actual ≠ Desired
            ↓
Self-Healing kicks in
            ↓
Argo CD reconciles Kubernetes
            ↓
Kubernetes returns to Git's desired state
```
