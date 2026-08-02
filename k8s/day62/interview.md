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

