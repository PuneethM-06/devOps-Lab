# Day 59 – Kubernetes Architecture Fundamentals

## Q1. Why do we need Kubernetes when we already have Docker?

Docker is responsible for building and running containers. It works well for managing a small number of containers on a single machine. However, in production environments, organizations run hundreds or thousands of containers across multiple servers. Managing them manually becomes difficult.

Kubernetes is a container orchestration platform that automates:

- Scaling
- Self-healing
- Scheduling
- Load balancing
- Rolling updates
- Rollbacks
- High availability
- Service discovery

In short:
- Docker creates and runs containers.
- Kubernetes manages containers at scale.

---

## Q2. What is a Kubernetes Cluster?

A Kubernetes cluster is a collection of physical or virtual machines that work together as a single platform to run containerized applications.

It consists of:

- Control Plane
- Worker Nodes

Benefits:
- High Availability
- Scalability
- Load Distribution

---

## Q3. Explain the Control Plane and Worker Nodes.

### Control Plane

The Control Plane is the brain of the Kubernetes cluster. It manages the entire cluster and makes decisions.

Components:
- API Server
- etcd
- Scheduler
- Controller Manager

### Worker Node

A Worker Node runs the application workloads.

Components:
- kubelet
- containerd
- kube-proxy
- Pods

---

## Q4. What is the Kubernetes API Server?

The API Server is the front door of Kubernetes. Every request from users and internal components goes through the API Server.

Responsibilities:

- Authentication
- Authorization
- Validation
- Exposes Kubernetes APIs
- Stores cluster state in etcd
- Acts as the communication hub between Kubernetes components

---

## Q5. What is etcd?

etcd is a distributed key-value database that stores the complete desired state and configuration of the Kubernetes cluster.

It is the single source of truth for Kubernetes.

It stores:

- Nodes
- Pods
- Deployments
- Services
- Secrets
- ConfigMaps
- Namespaces
- RBAC configurations

If etcd fails:
- Existing Pods continue running.
- New cluster operations cannot be performed.

---

## Q6. What is the Kubernetes Scheduler?

The Scheduler is responsible for selecting the most appropriate Worker Node for newly created Pods.

Scheduling decisions are based on:

- CPU availability
- Memory availability
- Node health
- Scheduling constraints

The Scheduler does not create Pods or start containers.

---

## Q7. What is the Controller Manager?

The Controller Manager continuously compares the desired state stored in etcd with the actual state of the cluster.

If there is a difference, it takes corrective action.

Responsibilities:

- Self-healing
- Maintaining desired replicas
- Reconciliation loop
- Creating missing Pod objects

---

## Q8. What is kubelet?

kubelet is the primary node agent that runs on every Worker Node.

Responsibilities:

- Watches for Pods assigned to its node
- Communicates with the container runtime
- Starts and manages containers
- Monitors Pod health
- Reports node status back to the Control Plane

---

## Q9. What is the Container Runtime?

The Container Runtime (commonly containerd) is responsible for:

- Pulling container images
- Creating containers
- Starting containers
- Stopping containers
- Removing containers

kubelet communicates with the runtime through the Container Runtime Interface (CRI).

---

## Q10. What is kube-proxy?

kube-proxy is a network agent that runs on every Worker Node.

Responsibilities:

- Routes traffic from Services to Pods
- Maintains network rules
- Performs load balancing across Pod replicas
- Provides stable networking for ephemeral Pods

---

## Q11. Explain the complete flow of `kubectl apply -f deployment.yaml`.

1. kubectl sends the request to the API Server.
2. API Server authenticates, authorizes, and validates the request.
3. The desired state is stored in etcd.
4. Controller Manager detects the desired state and creates Pod objects.
5. Scheduler selects the appropriate Worker Node.
6. kubelet on that node detects the assigned Pod.
7. kubelet asks containerd to pull the image and start the container.
8. The Pod becomes Running.
9. kube-proxy configures networking so Services can reach the Pod.

---

## Q12. Explain the hierarchy of Kubernetes objects.

Kubernetes Cluster
→ Worker Node
→ Pod
→ Container
→ Application

The application runs inside a container, the container runs inside a Pod, and the Pod runs on a Worker Node within the Kubernetes cluster.

---

## Key Takeaways

- Docker runs containers.
- Kubernetes orchestrates containers.
- API Server is the front door.
- etcd is the single source of truth.
- Controller Manager maintains the desired state.
- Scheduler selects the Worker Node.
- kubelet starts and manages containers on a node.
- containerd is responsible for container lifecycle.
- kube-proxy enables Service networking and load balancing.
- Pods are the smallest deployable unit in Kubernetes.