# KUBERNETES - DAY 62

## WHY DO WE NEED PROBES
- Lets image we made an deployment and after 30 mins due to some reason the application stops responding but when we do `kubectl` to understand the application - Pods return `running`.
- Note:  Kubelet checks if the pods are running healthy and checks if application is healthy and running. 
- Hence `Kubelet` checks the application every now and then using something like `GET /login`

### LIVENESS PROBE
- Kubelet checking if **Are you alive?** - Meaning are you up and running is called liveness probe

### READINESS PROBE
- Kubelet checking if ""Are you ready to serve the users* is called readiness probe 

## LIVENESS PROBE
- A liveness probe is a health check that is done by kubelet to check if the pods are alive, if not kubernetes will deploy or restart the container automatically 

```
Application Running
        │
        ▼
Kubelet performs Liveness Probe
        │
        ▼
Healthy?
     │
 ┌───┴────┐
 │        │
Yes       No
 │        │
 ▼        ▼
Do      Restart
Nothing Container
```

## READINESS PROBE 
- A readiness probe is an health check performed by the kubelet to determine if the application is ready to serve the users 
- R
eadiness probe checks for the health of the application 
```
Application Starting
        │
        ▼
Kubelet performs Readiness Probe
        │
        ▼
Ready?
     │
 ┌───┴────┐
 │        │
Yes       No
 │        │
 ▼        ▼
Add Pod   Remove Pod
to Service from Service
```
- Failure - Remove the pod from the container 

| Liveness Probe                            | Readiness Probe                                     |
| ----------------------------------------- | --------------------------------------------------- |
| Checks if the application is alive        | Checks if the application is ready to serve traffic |
| Failure → Restart the container           | Failure → Remove Pod from Service                   |
| Used for deadlocks, hangs, infinite loops | Used for startup delays or dependency issues        |
| Goal: Recover the application             | Goal: Protect users from failed requests            |

1. ### HTTP PROBE
- Kubelet sends an HTTP request to application 
-  if 200 ok, healthy; unhealthy -> remove the pod from the container 

2. ### TCP PROBE
- checks for establishing an TCP connection to a port
- connection succeeds, good; fails 

3. ### Exec Probe
- Instead of HTTP or TCP, kubelet runs a command inside the container 
- if exit code - 0, good else not 

- ## PROBE YAML 
```
livenessProbe:
  httpGet:
    path: /actuator/health
    port: 8080

  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
```
`httpGet` - use HTTP probe 
`path: /actuator/health` - where to check 
`initialDelaySeconds: 30` - start after 30 seconds of pod deployment 
`periodSeconds: 10` - check every 10 second 
`timeoutSeconds: 5` - wait at most 5 seconds for a response 
`failureThreshold: 3` - restart container after 3 consecutive request failure 

### NOTE:
1. **Liveness probe - Restart container**
2. **Readiness probe - remove pod and replace with new one** 