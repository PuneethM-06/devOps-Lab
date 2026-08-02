# DAY 60 - KUBERNETES


## WHY DO WE NEED SERVICES?
- Let's suppose in a worked node we have 3 pods and these pods are running `front end`, `back-end` and `database`.
Now if front end needs to communicate or send request to backend it needs to know the IP address of backend. 
- Pods are ephemeral, once a pod restarts its IP restarts and hence it is hard for the frontend ot remember the backend IP all the time and also there are other issue with this method:
    1. Poda are ephemeral
    2. Scaling 
    3. Node failures 

## SOLUTION
- The solution for above problems were service
- no matter:
    1. Which pod is running 
    2. which node its running on 
    3. how many pods exist
- **Application always calls services**
- A Kubernetes Service is an abstraction that provides a stable network endpoint for a group of Pods for communicating .

```
Frontend Pod
      │
      ▼
Backend Service
      │
      ▼
kube-proxy
      │
      ▼
Backend Pod
```
- **Kube-proxy** here is responsible for service networking by routing traffic from a service to a appropriate pod and also doing load balancing 
- **K8s service provides information about what each pods has running on it making kube-proxy routing easy**

## CHAPTER 2 - LABELS AND SELECTORS
- Without labels and selector, k8s would not work they work now
- It is used in:
    1. Deployments
    2. Replicasets 
    3. services
    4. network policies 
