# Docker Security - Interview Questions

## 1. Why is Docker security important?

**Answer:**

Containers share the host operating system kernel. If a container is compromised, it may impact the host or other containers. Following security best practices minimizes this risk.

---

## 2. Why shouldn't Docker containers run as the root user?

**Answer:**

Running as root gives the application elevated privileges. If the application is compromised, an attacker gains root access inside the container. Running as a non-root user follows the Principle of Least Privilege.

---

## 3. Which Dockerfile instruction changes the runtime user?

**Answer:**

```dockerfile
USER appuser
```

---

## 4. Why is `chown` commonly used before switching to a non-root user?

**Answer:**

To ensure the application files and directories are owned by the non-root user so the application can read and write them without permission errors.

Example:

```dockerfile
RUN chown -R appuser:appuser /app
```

---

## 5. What is Trivy?

**Answer:**

Trivy is an open-source vulnerability scanner that scans Docker images, filesystems, Git repositories, and Kubernetes manifests for known security vulnerabilities.

---

## 6. How do you scan a Docker image using Trivy?

```bash
trivy image nginx

# or

trivy image my-app:latest
```

---

## 7. Why should Docker images be scanned before deployment?

**Answer:**

To identify known vulnerabilities in the operating system packages and application dependencies before deploying the image to production.

---

## 8. What should you do if Trivy reports Critical vulnerabilities?

**Answer:**

- Update the base image.
- Upgrade vulnerable packages.
- Remove unnecessary software.
- Rebuild and rescan the image before deployment.

---

## 9. What is image hardening?

**Answer:**

Image hardening is the process of reducing the attack surface by using minimal base images, removing unnecessary packages, using multi-stage builds, keeping images updated, and following security best practices.

---

## 10. Name some techniques to harden a Docker image.

**Answer:**

- Use minimal base images (Slim, Alpine, Distroless)
- Use multi-stage builds
- Remove unnecessary packages
- Run as a non-root user
- Scan images regularly
- Keep base images updated

---

## 11. Why shouldn't secrets be stored inside a Docker image?

**Answer:**

Anyone with access to the image may be able to inspect it. Secrets should instead be injected at runtime or managed using a dedicated secrets management solution.

---

## 12. What is the difference between an environment variable and a secret?

**Answer:**

Environment variables are suitable for non-sensitive configuration such as ports or environment names, while secrets should be used for sensitive data like passwords, API keys, and tokens.

---

## 13. Why is `.dockerignore` important?

**Answer:**

It excludes unnecessary files from the Docker build context, making builds faster, reducing image size, and preventing accidental inclusion of sensitive files such as `.env` or `.git`.

---

## 14. What are some common entries in a `.dockerignore` file?

```text
.git
node_modules
.env
coverage
dist
```

---

## 15. Explain a secure Docker image build workflow.

**Answer:**

A secure workflow typically follows these steps:

1. Use a trusted minimal base image.
2. Build using multi-stage builds.
3. Remove unnecessary packages.
4. Switch to a non-root user.
5. Exclude unnecessary files using `.dockerignore`.
6. Scan the image with Trivy.
7. Push the image to a container registry.

---

## 16. What is the Principle of Least Privilege?

**Answer:**

Applications should be given only the minimum permissions required to perform their tasks. In Docker, this usually means avoiding running containers as the root user unless absolutely necessary.

---

## 17. What are some Docker security best practices?

**Answer:**

- Run containers as a non-root user.
- Use trusted and minimal base images.
- Keep images updated.
- Scan images regularly with Trivy.
- Do not store secrets inside images.
- Use multi-stage builds.
- Keep the Docker build context clean using `.dockerignore`.

---

## 18. You are reviewing a Dockerfile that runs as root and uses `ubuntu:latest` with many unnecessary packages installed. What improvements would you recommend?

**Answer:**

- Replace the base image with a minimal image if appropriate (Slim, Alpine, or Distroless).
- Remove unnecessary packages.
- Use a multi-stage build.
- Create and switch to a non-root user using `USER`.
- Scan the final image using Trivy.
- Keep secrets out of the image and use `.dockerignore` to exclude unnecessary files.