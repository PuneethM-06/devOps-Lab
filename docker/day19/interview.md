# Day 19 – Docker Fundamentals Interview Questions

## Docker Architecture

### 1. Explain the Docker architecture.

* Docker CLI
* Docker Engine
* dockerd
* containerd
* runc
* Docker Registry

---

### 2. Walk me through what happens internally when you run:

```bash
docker run nginx
```

---

### 3. What is the difference between a Docker image and a Docker container?

---

### 4. Why are Docker images immutable?

---

### 5. What is the writable layer in a Docker container?

---

## Dockerfile

### 6. Explain the purpose of the following Dockerfile instructions:

* FROM
* WORKDIR
* COPY
* ADD
* RUN
* CMD
* ENTRYPOINT
* ENV
* ARG
* EXPOSE
* LABEL

---

### 7. What is the difference between COPY and ADD?

---

### 8. What is the difference between CMD and ENTRYPOINT?

---

### 9. When would you use CMD and ENTRYPOINT together?

---

### 10. What happens if a Dockerfile contains multiple CMD instructions?

---

### 11. What is the difference between ARG and ENV?

---

## Docker Image Layers

### 12. What is a Docker image layer?

---

### 13. Why are Docker image layers read-only?

---

### 14. What is a layered filesystem?

---

### 15. What is a Union Filesystem?

---

### 16. Explain Copy-on-Write with an example.

---

### 17. Why does Docker copy only the modified file instead of the entire layer?

---

### 18. How can multiple Docker images share the same disk space?

---

### 19. What does `docker history` show?

---

### 20. Which Dockerfile instructions create filesystem layers and which only modify metadata?

---

## Docker Build Cache

### 21. What is Docker Build Cache?

---

### 22. How does Docker determine whether a cached layer can be reused?

---

### 23. What is cache invalidation?

---

### 24. Why does changing one COPY instruction cause later RUN instructions to rebuild?

---

### 25. Why should dependency installation happen before copying application source code?

Explain using either a Node.js or Python example.

---

### 26. What does the following command do?

```bash
docker build --no-cache
```

---

## Build Context & .dockerignore

### 27. What is the Docker build context?

---

### 28. When is the build context created?

---

### 29. What is the purpose of a `.dockerignore` file?

---

### 30. How is `.dockerignore` different from `.gitignore`?

---

### 31. Why should `node_modules` be excluded from the build context?

---

### 32. What are the risks of accidentally copying a `.env` file into a Docker image?

---

### 33. What happens if `requirements.txt` is ignored in `.dockerignore` but your Dockerfile contains:

```dockerfile
COPY requirements.txt .
```

---

## Practical & Scenario-Based Questions

### 34. Your Docker build has suddenly become much slower. How would you troubleshoot it?

---

### 35. Your Docker image size is much larger than expected. What could be the possible reasons?

---

### 36. You changed only `README.md`, but Docker rebuilt multiple layers. Why?

---

### 37. Why is the following Dockerfile inefficient?

```dockerfile
FROM node:22

COPY . .

RUN npm install

CMD ["npm", "start"]
```

How would you optimize it?

---

### 38. Compare the following Dockerfiles. Which one is better and why?

**Dockerfile A**

```dockerfile
COPY . .
RUN npm install
```

**Dockerfile B**

```dockerfile
COPY package*.json ./
RUN npm install
COPY . .
```

---

### 39. Why is `npm ci` preferred over `npm install` in CI/CD pipelines?

---

### 40. Design a production-ready Dockerfile for a Node.js or Python application.

Your explanation should cover:

* Layer ordering
* Build cache optimization
* `.dockerignore`
* Environment variables
* Port exposure
* Startup command
* Production best practices
