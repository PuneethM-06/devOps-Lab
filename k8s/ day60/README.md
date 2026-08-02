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
- **Labels acts as identity card for k8s pods, a single pod can have multiple labels**
- Example:
```
labels:
  app: backend
  env: production
  version: v1
```
### SELECTORS
- Labels store information, while selectors search for information 
- Similar to google, selectors are responsible for searching pods with the right labels and give information around it 

- Selectors are used by:
      1. **DEPLOYMENT**: Uses selectors to identify the pods
      2. **SERVICE**: Uses selectors to identify which pods receieve traffic

**LABEL**: A label is a key-value pair attache to a k8s object that is used to identify and organize resources
**SELECTOR**:  A selector is a query that identifies k8s object based on matching labels 

```
Labels → Selectors → Service → kube-proxy
```
> If a Pod's label changes from backend to frontend will Pod be able t receieve request from backend service?
- Nope, because when selector queries to match the pod with label backend; it wont be listed and hence backend service will not say kube-proxy to route traffic there 
- `Labels → Selectors → Service → kube-proxy`

## CLUSTERIP (default service type)
- **Cluster IP is a default kubernetes service that exposes an application only within the kubernetes cluster by assigning it a stable internal IP and DNS name**

```
Frontend Pod
      │
      ▼
backend-service
(ClusterIP)
10.96.20.15
      │
      ▼
kube-proxy
      │
      ▼
Backend Pod 1

Backend Pod 2

Backend Pod 3
```
### A REAL PRODUCTION ARCHITECTURE
```
                    Internet
                        │
                        ▼
                Load Balancer
                        │
                        ▼
                 Frontend Service
                  (LoadBalancer)
                        │
                        ▼
                  Frontend Pods
                        │
                        ▼
                  Backend Service
                   (ClusterIP)
                        │
                        ▼
                  Backend Pods
                        │
                        ▼
                  Database Service
                   (ClusterIP)
                        │
                        ▼
                   Database Pod
```
## NODEPORT
- Nodeport is a k8s service type that exposes an application externally by opening a fixed port on every worker node and forwards the traffic to the corresponding clusterIP services, which then reaches the pods 
```
                Outside the Cluster
                     (Your Laptop)
                           │
                           ▼
               http://NodeIP:30080
                           │
                           ▼
                    Worker Node
                    Port 30080
                           │
                           ▼
                    NodePort Service
                           │
                           ▼
                     ClusterIP Service
                           │
                           ▼
                      kube-proxy
                           │
                    Load Balances
                           │
          ┌────────────────┼────────────────┐
          ▼                ▼                ▼
      Backend Pod 1   Backend Pod 2   Backend Pod 3
```
## LOAD BALANCER
- A Loadbalancer service exposes an application externally by provisioning a cloud providers load balancer, which forwards incoming traffic to the k8s service
```
Internet User
      │
      ▼
AWS Load Balancer
      │
      ▼
NodePort
      │
      ▼
ClusterIP
      │
      ▼
kube-proxy
      │
      ▼
Backend Pod
```
