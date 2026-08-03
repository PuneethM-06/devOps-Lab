# DAY 63 - KUHBERNETES

## ROLLING UPDATE
- A rolling update is a deployment stratergy where k8s gradually replaces old pods with new pods, ensuring the application remains available throughout the deployment process with little or no down time at all 

- During this process k8s waits for the new pod to pass the readiness probe, before removing another pod 
