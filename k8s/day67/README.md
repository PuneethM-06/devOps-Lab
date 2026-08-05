# DAY 67 - KUBERNETES DEBUGGING

## PRODUCTION DEBUGGING FLOW
```
1. Is the Pod running?

↓

2. If not, why?

↓

3. What do the Events say?

↓

4. What do the Logs say?

↓

5. If it's running, can I enter the container?

↓

6. Fix the root cause.

↓

7. Verify the application is healthy.
```

### THREE IMPORTANT COMMANDS
1. **kubectl get pods** - shows the current status of the pods 
2. **kubectl describe pod <pod name>** - Describe the events in the pod in a detailed way 
3. **kubectl logs <pod name>** - Gets the logs of a pod 

> A production pod is failing. What is the first thing you do?
- I first observe the current state instead of restrating the pod 
- I check the status of the pod to see what condition is the pod in using `kubectl get pods`
- I check the application logs inside the failing pod using `kubectl logs <pod name>`

## UNDERSTANDING THE POD STATES
- There can be different states of pods such as:
    1. Running 
    2. Pending 
    3. CrashLoopBackOff
    4. ImagePullBackoff
    5. Completed

1. ### RUNNING
- This is the easiest one 
- It means:
    1. The pod has been scheduled to a worker node 
    2. Container is running inside the pod 
    3. Application is running successfully inside the container 

2. ### PENDING 
- This means that Pod has been accepted by kubernetes but hasnt started running yet 
- common reason for this is:
    1. No worker node has enough CPU
    2. No worker node has enough memory
    3. Node selectors dont match 
    4. Taints prevent scheduling 

3. ### COMPLETED
- This is not an error 
- Think of this as an k8s job
- Example:
```
Database Backup

↓

Runs Once

↓

Finishes

↓

Exits
```
4. ### FAILED
- This is a error, unlike completed 
- This means the pod has stopped 
- Possible reasons could be:
    1. Application exited with an error code 
    2. Container crashed
    3. Fatal startup time 
