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

