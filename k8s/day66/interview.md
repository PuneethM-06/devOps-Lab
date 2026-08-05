
# Day 66 – Namespaces & RBAC Basics

## Topics Covered

- Why Namespaces exist
- Namespace Architecture
- Creating Namespaces
- Built-in Namespaces
- RBAC (Role-Based Access Control)
- Role
- ClusterRole
- RoleBinding
- ClusterRoleBinding
- ServiceAccount
- Principle of Least Privilege
- Production Best Practices

---

# Interview Questions

## 1. What is a Namespace in Kubernetes?

A Namespace is a logical partition within a Kubernetes cluster used to organize and isolate resources. It allows multiple teams, applications, or environments to share the same cluster without resource naming conflicts.

---

## 2. Why do we need Namespaces?

Namespaces help:

- Organize resources
- Avoid naming conflicts
- Separate environments (Dev, QA, Stage, Prod)
- Improve resource management
- Provide logical isolation within a cluster

---

## 3. Does a Namespace create a new Kubernetes cluster?

No.

A Namespace does not create a new cluster or worker node. It simply provides logical isolation inside the same Kubernetes cluster.

---

## 4. Can two Deployments have the same name?

Yes.

They can have the same name if they exist in different Namespaces.

Example:

- dev/backend
- prod/backend

These are treated as different resources.

---

## 5. How do you create a Namespace?

Imperative:

```bash
kubectl create namespace dev
```

Declarative:

```yaml
apiVersion: v1
kind: Namespace

metadata:
  name: dev
```

---

## 6. Where are resources created if no Namespace is specified?

Resources are created in the **default** Namespace unless another Namespace is specified using:

```bash
kubectl apply -f deployment.yaml -n dev
```

or

```yaml
metadata:
  namespace: dev
```

---

## 7. What are the built-in Kubernetes Namespaces?

### default

Default Namespace for user workloads.

### kube-system

Contains Kubernetes system components such as:

- CoreDNS
- kube-proxy
- Metrics Server
- CNI plugins

### kube-public

Stores publicly readable cluster information.

### kube-node-lease

Stores Node Lease objects used for worker node heartbeats.

---

## 8. What is RBAC?

RBAC (Role-Based Access Control) is Kubernetes' authorization mechanism that controls who or what can perform specific actions on Kubernetes resources.

---

## 9. Why is RBAC required?

RBAC prevents unauthorized users or applications from modifying Kubernetes resources. It follows the Principle of Least Privilege by granting only the permissions required to perform a specific task.

---

## 10. What is a Role?

A Role defines a set of permissions within a single Namespace.

Example:

- Get Pods
- List Pods
- Create Deployments

These permissions apply only within that Namespace.

---

## 11. What is a ClusterRole?

A ClusterRole defines permissions across the entire Kubernetes cluster.

It is commonly used for:

- Cluster administrators
- Accessing Nodes
- Accessing Namespaces
- Accessing PersistentVolumes
- Cluster-wide permissions

---

## 12. What is the difference between Role and ClusterRole?

| Role | ClusterRole |
|------|-------------|
| Namespace scoped | Cluster scoped |
| Applies to one Namespace | Applies across the cluster |
| Used for application teams | Used for cluster-wide permissions |

---

## 13. What is a RoleBinding?

A RoleBinding assigns a Role (or a ClusterRole) to a User, Group, or ServiceAccount within a Namespace.

A Role alone does not grant permissions until it is bound.

---

## 14. What is a ClusterRoleBinding?

A ClusterRoleBinding assigns a ClusterRole to a User, Group, or ServiceAccount across the entire cluster.

---

## 15. Can a RoleBinding reference a ClusterRole?

Yes.

A RoleBinding can reference a ClusterRole. However, the permissions are still limited to the Namespace where the RoleBinding exists.

---

## 16. What is a ServiceAccount?

A ServiceAccount is an identity used by Pods to authenticate with the Kubernetes API.

Applications running inside Pods use ServiceAccounts instead of human users.

---

## 17. Does every Worker Node have a ServiceAccount?

No.

ServiceAccounts are associated with Pods, not Worker Nodes.

Multiple Pods from the same Deployment commonly share the same ServiceAccount.

---

## 18. What happens if a Pod does not specify a ServiceAccount?

It automatically uses the **default ServiceAccount** in its Namespace.

---

## 19. Explain the complete RBAC flow.

1. A User or Pod sends a request to the Kubernetes API.
2. Kubernetes identifies the User or ServiceAccount.
3. RoleBinding or ClusterRoleBinding is evaluated.
4. The associated Role or ClusterRole is checked.
5. If the required permission exists, the request is allowed.
6. Otherwise, Kubernetes returns a **Forbidden** error.

---

## 20. Explain the Principle of Least Privilege.

The Principle of Least Privilege means granting users and applications only the permissions they require to perform their tasks, reducing the impact of accidental mistakes or security breaches.

---

## 21. Explain the relationship between ServiceAccount and RBAC.

Applications running inside Pods authenticate using a ServiceAccount.

Permissions are granted by:

ServiceAccount

↓

RoleBinding / ClusterRoleBinding

↓

Role / ClusterRole

↓

Kubernetes API Authorization

---

## 22. Production Best Practices

- Avoid using the default Namespace for production applications.
- Separate environments using Namespaces.
- Grant the minimum permissions required.
- Avoid giving ClusterRole or cluster-admin unless necessary.
- Create dedicated ServiceAccounts for applications instead of relying on the default ServiceAccount.
- Regularly audit Roles and RoleBindings.