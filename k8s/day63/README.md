# DAY 63 - KUHBERNETES

## ROLLING UPDATE
- A rolling update is a deployment stratergy where k8s gradually replaces old pods with new pods, ensuring the application remains available throughout the deployment process with little or no down time at all 

- During this process k8s waits for the new pod to pass the readiness probe, before removing another pod 

## HOW DOES K8S PERFORM A ROLLING UPDATE 
- Every deployment gets its own replicaset.
- Lets say we had a `replicaset v1` serving the users, so once a new deployment is rolled out. K8s does not immeditaley delete the older replicaset instead it creates a new replicaset called `v2`.Gets the number of pods needed rolling in and once that is done it checks for `readiness` of the pods in the new replicaset.
- Please note that the **previous replicaset is not deleted**, because tomorrow if you want to go back to the previous replicaset because the current has a bug then it might be handy 

## maxSurge and maxUnavailable 

### PROBLEM
- when we create a new deployment, K8s has to keep two important goals in mind:
    1. DONT CREATE TOO MANY PODS, BECAUSE PODS CONSUME MEMORY
    2. DONT REMOVE TOO MANY OLD PODS BECAUSE USER NEEDS APPLICATION 

> How many max pods can I remove?
- maxUnavailable 

> How many extra pods can I temporarily create?
- maxSurge 

- Example:
```
maxUnavailable: 1
maxSurge: 1
```
- This says I can remove maximum of one pod at a time; and create one pod at a time in the new replicaset
- Creates a new pod at the new replicaset; waits or readiness probe to pass -> Removes one pod from the prev replicaset 

- **maxSurge**: defines the **maximum number of new pods K8s can temporarily create above the desired replica count during an update rollout**
- **maxUnavilable** - defines the maximum number of pods thaty can be unavailabe during a rolling update while maintaing application availability

- If `REPLICAS  = R`;
- Maximum pods during an update = `R + maxSurge`
- Minimum available pods = `R - maxunavailable`

## ROLLBACK
- If there is a issue with the new deployment, K8s can immediately do a rollback to the older replicaset because it would have not deleted it in the first place
```
v1 Deployment
        │
        ▼
ReplicaSet-v1
      (3 Pods)

        │
Rolling Update
        ▼

ReplicaSet-v1
      (0 Pods)

ReplicaSet-v2
      (3 Pods)

        │
Bug Found
        ▼

Rollback

        │
        ▼

ReplicaSet-v1
      (3 Pods)

ReplicaSet-v2
      (0 Pods)
```

