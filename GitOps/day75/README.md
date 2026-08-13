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
