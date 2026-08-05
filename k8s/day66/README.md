# DAY 66 - KUBERNETES

## WHY DO NAMESPACES EXIST
- Namespaces are like folders inside a desktop
- Namespaces are used to organize resources inside a k8s cluster 
- Example:
```
Cluster

├── dev
│     ├── API
│     └── Redis
│
├── qa
│     ├── API
│     └── Redis
│
├── prod
│     ├── API
│     └── Redis
│
└── monitoring
      ├── Prometheus
      └── Grafana
```
- A namespace is a logicial partition within a k8s cluster that organizes and isolates resources. It allows multiple teams, applications or environments to share the same cluster without resource naming conflicts

## WHAT EXACTLY IS A NAMESPACE?
1. **HOW DO WE CREATE NAMESPACES**
- `kubectl create namespace dev`

2. **TO VIEW ALL NAMESPACES**
- `kubectl get namespaces`

3. **DEPLOYMENT INTO DEFAULT NAMESPACE**
- `kubectl apply -f deployment.yaml`

4. **DEPLOYMENT INTO DEV NAMESPACE**
- `kubectl apply -f deployment.yaml -n dev`

> CAN TWO DEPLOYMENTS HAVE THE SAME NAMESPACES?
- Yes, provided they are in different namesapces. Resource names must be unique within a namespace but the same name can exist in multiple namespaces

> CAN KUBERNETES NAMESPACE BE NESTED?
- No, Kubernetes namespaces are flat and cannot be nested. 

##  BUILT IN NAMESPACES
- When a brand new cluster is creayed we can see:
    1. default 
    2. Kube-system 
    3. kube-public
    4. kube-node-lease

1. ### default
- This is the easiest one
- when a deployment is done without explicitly defining a namespace then it is created under `default`

2. ### kube-system 
- This is one of the most important built-in namespace
- This can be thought as the **operating system folder of the kubernetes cluster**
- such as:
    1. coreDNS
    2. kube-proxy
    3. metric server
    4. CNI plugins 

3. ### Kube-public
- This namespace is intendede for resources that are publicly readable across the cluster
- It is rarely used by application developers

4. ### kube-node-lease
- Every worker node periodically says control plane that it is alive and it called as **heartbeat**
- Kubernetes stores these hearbeat **lease object in kube-node-lease**

## RBAC - ROLE BASED ACCESS CONTROL 
- Example:
| User            | Permission                    |
| --------------- | ----------------------------- |
| Developer       | Create Pods in `dev`          |
| QA Engineer     | View Pods in `qa`             |
| DevOps Engineer | Deploy applications to `prod` |
| Cluster Admin   | Full access to the cluster    |

- **HIGH LEVEL FLOW**
```
User
   │
   ▼
Requests an action

↓

RBAC checks

↓

Does this user have permission?

↓

YES → Allow

NO → Deny
```
1. ### PERMISSION
- What actions are allowed
- Example:
    1. Get pods
    2. List pods
    3. Create deployments
    4. Delete services

2. ### ASSIGNMENT
- Who gets these permissions
- Example:
    1. Puneeth
    2. DevOps team 
    3. Service account 

## ROLE vs CLUSTER ROLE 

1. ### ROLE
- A role is namespace-scoped
```
Role

↓

One Namespace

↓

Specific Permissions
```
- Example:
```
Role

Namespace = dev

Permissions:

✓ Get Pods
✓ List Pods
✓ Create Deployments
```
2. ### CLUSTERROLE
- A cluster role unlike is not restricted to a particular namespace but is given permission across the cluster 
```
ClusterRole

↓

Entire Cluster

↓

Permissions
```
- Example:
```
ClusterRole

Permissions

✓ View Nodes
✓ View Namespaces
✓ View Pods everywhere
```
| Role                               | ClusterRole                                               |
| ---------------------------------- | --------------------------------------------------------- |
| Namespace scoped                   | Cluster scoped                                            |
| Permissions apply to one Namespace | Permissions apply across the entire cluster               |
| Used for application teams         | Used for cluster administrators or cluster-wide resources |

## ROLE HIGHLEVEL 
```
Role
   │
   ▼
RoleBinding
   │
   ▼
User - Developer
```

## CLUSTERROLE HIGH LEVEL 
```
ClusterRole
      │
      ▼
ClusterRoleBinding
      │
      ▼
User
```
## SERVICE ACCOUNTS
- For a human to talk to pods, authorize and authenticate himself we are gonna make use of `Roles and clusterRoles`
- For Applications to talk to pod we need **Service accounts**
- Once an service account is created we are gonna do role binding for service accounts as well do have access to applications to roles rolebinding 
- **DEFINITION** - A Service account is a k8s identity used by poids to authenticate with k8s API. It allows application inside a pod to access k8s resources according to permissions granted through BAC
- Each Pod inside a worker node gets a service account; meaning it doesnt mean each pod gets a unique account
- Example:
```
suppose we have 
Deployment

Replicas: 3

ServiceAccount: backend-sa

k8s creates
Pod 1

↓

backend-sa

Pod 2

↓

backend-sa

Pod 3

↓

backend-sa
```
- All pods use the same service account 

> Does every pod have its own service account?
- Every pod runs as a serviceAccount. If no serviceAccount is specified then it makes use of the default service account in its namespace.
Multiple pods from the same deployemnt make use of the same serviceAccount because they represent the same application and there require the same permission 
