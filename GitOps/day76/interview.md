# DAY 76 - ARGOCD 

1. ### WHERE DOES ARGOCD RUN?
-  ArgoCD system runs inside your k8s cluster
```
Your Machine
│
├── kubectl
├── argocd CLI
│
└──────────────► Kubernetes Cluster
                       │
                       └── argocd namespace
                              │
                              ├── Argo CD components
                              └── Argo CD manages applications
```
- The reason why ArgoCD runs inside the k8s clustyer:
    1. To read the current state of the k8s cluster 
    2. To perform reconcile 

2. ### WHAT ARE WE ACTUALLY INSTALLING 
- When we install Argo CD, k8s creates multiple Argo CD components, usually inside an `argocd` namespace
```
Kubernetes Cluster
│
└── argocd namespace
    │
    ├── argocd-server
    ├── argocd-repo-server
    ├── argocd-application-controller
    └── other supporting components
```
- **argocd-server**
- This provides:
    1. Web UI
    2. API
    3. CLI access

- **argocd-repo-server**
- This is responsible for interacting with git repositories for rendering and preparing the application manifests

- **argocd-application-controller**
- This is responsible for checking if the actual and the desired state are in sync and initiating the reconcilation steps

3. ### BEFORE INSTALLING ARGOCD WHAT DO WE NEED 
1. K8s cluster
2. kubectl
3. Git repo 

3. ### HOW DO WE INSTALL ARGOCD
- ArgoCD is deployed in the k8s resources/cluster
```
Create namespace
       ↓
Apply Argo CD installation manifests
       ↓
Kubernetes creates Argo CD resources
       ↓
Pods start running
```
**COMMANDS**
1. `kubectl create namespace argocd`
2. `kubectl apply -n argocd -f <argocd-install-manifest>`

4. ### HOW DO WE ACCESS ARGO CD?
- There are 3 ways in which we can access Argo CD
    1. Argo CD web UI
    2. Argo CD CLI
    3. Argo CD API

- **PORT FORWARDING**
- since ArgoCD is running in k8s cluster it is not directly accesssible as a web ui and hence we need to make use of port forwarding 
- `kubectl port-forward svc/argocd-server -n argocd 8080:443` this enables web ui access locally 
- `8080:443` - **This means traffic sent to port 8080 in your local machine is forwarded to port 443 of the `argocd-server` inside the k8s cluster
```
Your Browser / Local Machine
        │
        │ Request to localhost:8080
        ▼
kubectl port-forward
        │
        │ forwards traffic
        ▼
argocd-server:443
        │
        ▼
Argo CD processes the request
        │
        ▼
Response comes back through the same connection
        ▼
Your Browser
```
> A request sent to port 8080 on your local machine is forwarded to port 443 on the argocd-server inside the Kubernetes cluster. Argo CD processes the request, and the response is sent back to your local machine through that port-forward connection.

5. ### INITIAL ARGO CD LOGIN 
- **The initial password is automatically created and stored in k8s as a secret**
```
Argo CD installation
        ↓
Kubernetes creates resources
        ↓
Initial admin password is stored in a Secret
```
- The password can be retrieved as 
```
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d
```
- Then login 
> The initial password is stored in a Kubernetes Secret so that access can be controlled through Kubernetes authentication and authorization, such as RBAC. Only users or service accounts with the required permissions should be able to read that Secret.

6. ### THE ARGO CD CLI 
 - argoCD CLI - communicate with Argo CD
 - We login to it using `argocd login localhost:8080` amd then we can interact with it 
 ```
 Your Local Machine
        │
        │ argocd CLI
        ▼
localhost:8080
        │
        │ port-forward
        ▼
argocd-server:443
        │
        ▼
Argo CD
```
Here we have two clients:
```
kubectl CLI → Kubernetes API

argocd CLI  → Argo CD API/server
```
