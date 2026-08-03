# DAY 63 - KUHBERNETES

## ROLLING UPDATE
- A rolling update is a deployment stratergy where k8s gradually replaces old pods with new pods, ensuring the application remains available throughout the deployment process with little or no down time at all 

- During this process k8s waits for the new pod to pass the readiness probe, before removing another pod 

## HOW DOES K8S PERFORM A ROLLING UPDATE 
- Every deployment gets its own replicaset.
- Lets say we had a `replicaset v1` serving the users, so once a new deployment is rolled out. K8s does not immeditaley delete the older replicaset instead it creates a new replicaset called `v2`.Gets the number of pods needed rolling in and once that is done it checks for `readiness` of the pods in the new replicaset and eventually and gradually delete the `replicaset v1 pods`
