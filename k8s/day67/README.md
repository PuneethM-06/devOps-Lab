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

5. ### UNKNOWN
- This is a rare case scenario.
- This means the control plane has lost communication with the worker node 


| Status                | Meaning                                      |
| --------------------- | -------------------------------------------- |
| Running               | Pod is running successfully.                 |
| Pending               | Waiting to be scheduled or started.          |
| Completed (Succeeded) | Finished successfully (common for Jobs).     |
| Failed                | Pod terminated due to an error.              |
| Unknown               | Kubernetes cannot determine the Pod's state. |

## CRASHLOOPBACKOFF
```
Container Starts

↓

Application Crashes

↓

Kubelet Restarts Container

↓

Application Crashes Again

↓

Kubelet Restarts Again

↓

Keeps Crashing

↓

Wait Longer

↓

Restart Again

↓

CrashLoopBackOff
```
- crashLoopBackOff means the contaiuner starts, crashes repeatedly, and kubernetes delays each restart using an increasing backoff to avoid continuous restart loops

### COMMON CAUSES
1. Application Bug
2. Wrong environment variables
3. Secrets missing 
4. Database not reachable 
5. Port already in use
6. OOMKILLED

## PENDING PODS 
- This is a scenario where kubernetes is not possible to assign a pod to a worker node 
- The possible reasons are:
    1. CPU and memory inside the worker nodes does not have the minimum space that the pod is asking 
    2. Node selector does not match 
   ```
   nodeSelector:
  region: us-east-1
  But every node is on region = us-west-1
  ```
### NOTE:
> Here the question that will arise is cant k8s make use of HPA and scale? 
- HPA cannot create worker nodes it can **only scale the pods**
- **Cluster Auto scaler** is responsible for creating new worker nodes

## IMAGEPULLBACKOFF AND ERRIMAGE PULL
- This error is seen when the container **FAILS TO START BECAUSE IT CANNOT PULL THE IMAGE**
- The different between `ERRIMAGE PULL` and `IMAGEPULLBACKOFF` is that, `ERRIMAGE PULL` occurs when the image is failed fetch for the first time, while `IMAGEPULLBACKOFF` kubelet keeps trying and it doesnt happen 
- The reason is, container tries to pull the image from the docker hub or ghcr where we have mentioned but the image wouldnt be there at all 
- common reason for this to occue is:
    1. Wrong image name 
    2. Wrong image tag
    3. Registry unavailable 
    4. authenticaton failed 

| CrashLoopBackOff                | ImagePullBackOff                            |
| ------------------------------- | ------------------------------------------- |
| Image downloaded successfully   | Image could not be downloaded               |
| Container starts                | Container never starts                      |
| Application crashes             | Image pull fails                            |
| `kubectl logs` is useful        | `kubectl logs` is usually not useful        |
| Root cause often in application | Root cause often in image name/tag/registry |

```
kubectl get pods

↓

Is the container running?

├── YES
│      │
│      ├── Application crashing?
│      │        ↓
│      │   kubectl logs
│      │
│      └── Running but unhealthy?
│               ↓
│         kubectl exec
│
└── NO
       │
       ├── Pending
       │      ↓
       │ kubectl describe
       │
       ├── ImagePullBackOff
       │      ↓
       │ kubectl describe
       │
       └── FailedScheduling
              ↓
         kubectl describe
```

## OOMKILLED - OUT OF MEMORY KILLED
- OOMKilled stands for Out Of Memory Killed. It occurs when a container tries to use more memory than its configured memory limit. To protect the stability of the system and other running containers, the Linux Kernel's OOM Killer terminates that container. Kubernetes (specifically the Kubelet) then notices that the container has exited and restarts it according to the Pod's restart policy.

| Command                             | Purpose                                                      |
| ----------------------------------- | ------------------------------------------------------------ |
| `kubectl logs <pod>`                | View application logs                                        |
| `kubectl logs -f <pod>`             | Stream logs live                                             |
| `kubectl logs --previous <pod>`     | View logs from the previously crashed container              |
| `kubectl logs <pod> -c <container>` | View logs from a specific container in a multi-container Pod |

