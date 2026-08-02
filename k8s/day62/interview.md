# KUBERNETES - DAY 62

## WHY DO WE NEED PROBES
- Lets image we made an deployment and after 30 mins due to some reason the application stops responding but when we do `kubectl` to understand the application - Pods return `running`.
- Note:  Kubelet checks if the pods are running healthy and checks if application is healthy and running. 
- Hence `Kubelet` checks the application every now and then using something like `GET /login`

