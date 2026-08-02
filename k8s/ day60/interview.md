# Day 60 – Kubernetes Services

## Q1. Why do we need Kubernetes Services?

Pods are ephemeral. Whenever a Pod crashes or is recreated, its IP address changes. If applications communicate directly using Pod IPs, communication breaks after Pod recreation.

A Kubernetes Service provides a stable network endpoint that applications can use regardless of Pod lifecycle changes.

---

## Q2. What is a Kubernetes Service?

A Kubernetes Service is an abstraction that provides a stable network endpoint for a group of Pods.

It enables:
- Stable communication
- Service discovery
- Load balancing
- Decoupling applications from Pod IPs

---

## Q3. Why can't applications communicate directly using Pod IP addresses?

Pod IPs are not permanent.

When a Pod is recreated:
- A new Pod is created.
- It receives a different IP address.
- Any application using the old IP will fail.

Services solve this problem by exposing a stable IP and DNS name.

---

## Q4. What are Labels?

Labels are key-value pairs attached to Kubernetes objects to identify, organize, and categorize them.

Example:

```yaml
labels:
  app: backend
  env: production
```

---

## Q5. What are Selectors?

Selectors are queries used to identify Kubernetes objects based on matching Labels.

Example:

```yaml
selector:
  app: backend
```

This Service selects every Pod having:

```yaml
labels:
  app: backend
```

---

## Q6. How does a Service discover Pods?

A Service discovers Pods using Labels and Selectors.

The Service does not store Pod names or Pod IPs.

Instead, it selects every Pod whose labels match its selector.

---

## Q7. What is the relationship between a Service and kube-proxy?

Responsibilities are different.

### Service

- Defines which Pods belong together
- Provides a stable IP and DNS name
- Uses Labels and Selectors

### kube-proxy

- Reads Service information
- Maintains networking rules
- Routes traffic
- Load balances requests across matching Pods

---

## Q8. What is ClusterIP?

ClusterIP is the default Kubernetes Service type.

It exposes an application only inside the Kubernetes cluster by assigning a stable internal IP address and DNS name.

Use Cases:
- Frontend → Backend
- Backend → Database
- Internal microservice communication

---

## Q9. What is NodePort?

NodePort exposes an application outside the Kubernetes cluster by opening the same port on every Worker Node.

Traffic Flow:

Internet
→ Worker Node : NodePort
→ ClusterIP
→ kube-proxy
→ Pod

Primarily used for:
- Development
- Testing
- Small Kubernetes clusters

---

## Q10. What is LoadBalancer?

LoadBalancer is a Kubernetes Service type that exposes an application externally by automatically provisioning a cloud provider's load balancer.

Traffic Flow:

Internet
→ Cloud Load Balancer
→ NodePort
→ ClusterIP
→ kube-proxy
→ Pod

Commonly used in production environments.

---

## Q11. Compare ClusterIP, NodePort and LoadBalancer.

| Feature | ClusterIP | NodePort | LoadBalancer |
|----------|-----------|----------|--------------|
| Internal Communication | Yes | Yes | Yes |
| External Communication | No | Yes | Yes |
| Stable Service IP | Yes | Yes | Yes |
| Opens Port on Worker Nodes | No | Yes | Yes (internally) |
| Cloud Load Balancer | No | No | Yes |

---

## Q12. Explain the complete request flow from the Internet to a Kubernetes Pod.

1. User sends a request over the Internet.
2. The request reaches the cloud provider's Load Balancer.
3. The Load Balancer forwards traffic to a healthy Worker Node's NodePort.
4. NodePort forwards the request to the Service's ClusterIP.
5. The Service identifies matching Pods using Labels and Selectors.
6. kube-proxy routes the request to one of the matching Pods.
7. The request reaches the container inside the Pod.
8. The application processes the request and returns the response.

---

## Q13. Which Service types would you use for an e-commerce application?

Frontend:
- LoadBalancer
- Exposed to Internet users.

Backend:
- ClusterIP
- Accessible only inside the cluster.

Database:
- ClusterIP
- Accessible only by the Backend.

---

## Key Takeaways

- Services provide stable communication endpoints.
- Pods are discovered using Labels and Selectors.
- Services do not know Pod names or Pod IPs.
- kube-proxy routes traffic to matching Pods.
- ClusterIP is used for internal communication.
- NodePort exposes applications through Worker Node ports.
- LoadBalancer provisions a cloud load balancer for external access.
- Service networking is independent of Pod lifecycle.