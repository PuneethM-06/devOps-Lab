# DAY 75 - GITOPS AND ITS BASICS 

1. ### WHAT PROBLEM DOES GITOPS SOLVE TODAY 
- Suppose we have a helm based deployment where we have set the replicas to be 2 ( `replicas:2` ), and then someone overrides the git/helm based configuration from 2 to 10 using `kubectl scale deployment my-app --replicas=10`, now the problem is, no one is sure what is the actual deployment that must be going in and there is a drift from the git/helm based congfiguration and to over come this we are making use of GitOps
- **GitOps detects such drifts from the actual state to desired state and gets the cluster back to git based configuration**
```
Git defines desired state
        ↓
GitOps tool checks the cluster
        ↓
Does actual state match desired state?
        │
   Yes ─┴─ No
   │        │
   │        ↓
   │    Reconcile/fix it
   │
Everything is in sync
```

2. ### WHAT IS GITOPS?
- **GitOps is an operational model where git is used as the source of truth for the desired state of appplications and infratsructure**
- By this it enables:
    1. Version history
    2. Traceability
    3. Code review
    4. Rollback 
- Here is how it needs to be done using GitOps/git based changes:
```
Change Git configuration
        ↓
Commit the change
        ↓
Push / create a PR
        ↓
PR gets reviewed and merged
        ↓
Git now contains the new desired state
        ↓
GitOps tool detects the change
        ↓
Kubernetes is updated
```
3. ### DESIRED STATE vs ACTUAL STATE
- **Desired state** - Git based config is the desired state 
- **Actual state** - The resources that are currently in the k8s are actual state 
- ArgoCD always ensure that Desired and actual state are the same 

4. ### CONFIGURATION DRIFT 
- **Configuration drift occurs when the actual state of the k8s resources becomes different from the desried state in git**
- How can drift happen:
    1. Manual changes through commands - `kubectl scale deployment my-app --replicas=5`
    2. Editing resources directly - `kubectl edit deployment my-app`
    3. Emergency fixes - Issue to fix in prod

5. ### RECONCILIATION 
- The process that occurs after finding a congif drift is called reconciliation
- **Taking action to match the  actual state to desired state and eliminate the config drift is called reconciliation**
```
Git
 │
 │ Desired State
 ▼
Argo CD
 │
 │ Compare
 ▼
Kubernetes
 │
 │ Actual State
 ▼

Match?
 │
 ├── Yes → Synced
 │
 └── No  → OutOfSync
              │
              ▼
        Reconciliation / Sync
              │
              ▼
       Actual matches Desired
```
6. ### TRADITIONAL CI/CD vs GitOps
- **TRADITIONAL CI/CD**
```
Developer
    ↓
Pushes code/config to Git
    ↓
CI/CD Pipeline starts
    ↓
Pipeline deploys to Kubernetes
    ↓
kubectl apply / helm upgrade
```
- CI/CD directly pushes the code to k8s
- **GitOps**
```
Developer changes configuration
          ↓
        Git
          ↓
Argo CD detects desired state
          ↓
Argo CD compares with Kubernetes
          ↓
Synchronizes the cluster
```
- **The only difference here is that CI/CD continously doesnt match/check for actual state and desired state are in sync, while ArgoCD does**

7. ### PUSH vs PULL BASED DEPLOYMENT 
- **TRADITIONAL APPROACH - PUSH**
- In this the CI/CD might literally run commands like `kubectl apply -f deployment.yaml` and it might need access to k8s cluster and hence needs creds 
```
GitHub Actions / Jenkins
          │
          │ PUSH deployment
          ▼
     Kubernetes
```
- **GitOps approach - Pull**
```
Git Repository
      ▲
      │ Argo CD reads/pulls desired state
      │
   Argo CD
      │
      ▼
 Kubernetes
```
- Unlike, CI/CD here ArgoCD reads the actual and desired state and reconciles when they are not in sync, with this CI can manage on `Build → Test → Push Image → Update Git configuration` and ArgoCD can manage `Git configuration → Kubernetes deployment`
8. ### OVERALL WHERE GITOPS FIT 
```
Developer
    │
    ▼
Application Code Repository
    │
    ▼
CI Pipeline
 ┌───────────────┐
 │ Test          │
 │ Build Image   │
 │ Push Image    │
 └───────────────┘
    │
    ▼
Container Registry
    │
    │
    └──── Image reference updated in GitOps configuration
                              │
                              ▼
                        GitOps Repository
                              │
                              ▼
                            Argo CD
                              │
                              ▼
                         Kubernetes
```
