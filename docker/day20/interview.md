# Day 20 – Docker Image Optimization Interview Questions

## Multi-stage Builds

### 1. What is a multi-stage build?

---

### 2. Why do we use multi-stage builds?

---

### 3. Explain the difference between a Builder stage and a Runtime stage.

---

### 4. Can a Dockerfile contain multiple `FROM` instructions? Why?

---

### 5. What does the following instruction do?

```dockerfile
COPY --from=builder /app/dist /usr/share/nginx/html
```

Explain:

* `--from=builder`
* `/app/dist`
* `/usr/share/nginx/html`

---

### 6. Does Docker create multiple final images during a multi-stage build?

---

### 7. Why is only the final stage included in the production image?

---

### 8. What are the advantages of multi-stage builds?

---

### 9. Give a real-world example where multi-stage builds are useful.

---

## Base Images

### 10. What is a Docker base image?

---

### 11. Why does every Dockerfile usually start with `FROM`?

---

### 12. Compare Ubuntu, Debian, Slim, Alpine, and Distroless.

Discuss:

* Image size
* Compatibility
* Debugging
* Security
* Typical use cases

---

### 13. Why isn't Alpine always the best choice?

---

### 14. Why are Slim images commonly used in production?

---

### 15. Which base image would you choose for:

* React application
* Python backend
* Java service
* Linux monitoring tool
* Learning Docker

Explain your reasoning.

---

## Docker Image Optimization

### 16. Why should Docker images be optimized?

---

### 17. What problems do large Docker images cause?

---

### 18. Does a smaller Docker image always use less RAM? Explain.

---

### 19. How does Docker image size affect Kubernetes deployments?

---

### 20. Why do smaller Docker images improve security?

---

### 21. Why should related `RUN` commands be combined?

---

### 22. Why should package cache be removed in the same `RUN` instruction?

---

### 23. Why is the following Dockerfile inefficient?

```dockerfile
RUN apt update

RUN apt install -y curl

RUN rm -rf /var/lib/apt/lists/*
```

How would you improve it?

---

### 24. Why should `COPY package*.json` come before `COPY . .`?

---

### 25. Why should production images install only production dependencies?

---

### 26. Besides multi-stage builds, what are other techniques to reduce Docker image size?

---

## Distroless Images

### 27. What is a Distroless image?

---

### 28. Why are Distroless images considered more secure?

---

### 29. Why can't you run `docker exec -it <container> bash` inside a Distroless container?

---

### 30. How do you debug a Distroless container?

---

### 31. Compare Alpine and Distroless.

---

### 32. When would you choose Distroless over Slim?

---

## Scenario-Based Questions

### 33. Your Docker image is over 2 GB. How would you reduce its size?

---

### 34. A company wants faster Kubernetes deployments. What Docker image optimizations would you recommend?

---

### 35. Why should source code and build tools not be present in the final production image?

---

### 36. Explain the complete lifecycle of a React application's multi-stage Docker build from source code to the final production image.

---

### 37. Your application builds successfully in the Builder stage but fails in the Runtime stage. What could be the possible reasons?

---

### 38. A production container needs frequent interactive debugging. Would you choose Distroless? Why or why not?

---

### 39. Design a production-ready Dockerfile for a React or Node.js application. Explain every optimization you apply.

---

### 40. If you were reviewing a teammate's Dockerfile, what are the first five things you would check before approving the PR?
