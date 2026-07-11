# Docker Volumes & Persistent Storage - Interview Questions

## 1. Why do Docker Volumes exist?

**Answer:**

Containers are ephemeral. Any data stored inside the container's writable layer is lost when the container is removed. Docker Volumes provide persistent storage that exists independently of the container lifecycle.

---

## 2. What is a Docker Volume?

**Answer:**

A Docker Volume is Docker-managed persistent storage that allows data to survive container deletion and be reused by new containers.

---

## 3. What is the difference between a Docker Volume and a Bind Mount?

| Docker Volume | Bind Mount |
|---------------|------------|
| Managed by Docker | Managed by the Host |
| Docker chooses storage location | User specifies host path |
| Best for production | Best for development |
| More portable | Host path dependent |
| Ideal for databases | Ideal for source code |

---

## 4. When would you use a Bind Mount instead of a Docker Volume?

**Answer:**

Bind mounts are mainly used during development because changes made on the host are immediately reflected inside the container without rebuilding the image.

---

## 5. Why are Docker Volumes preferred for databases?

**Answer:**

Database data must persist across container restarts, upgrades, and recreation. Docker Volumes separate data from the container lifecycle.

---

## 6. What is an Anonymous Volume?

**Answer:**

An Anonymous Volume is a Docker-managed volume created automatically without a user-defined name. It is persistent but harder to identify and manage than a named volume.

---

## 7. Does deleting a container delete its Docker Volume?

**Answer:**

No. Named volumes have an independent lifecycle and remain until they are explicitly removed.

---

## 8. Can multiple containers use the same Docker Volume?

**Answer:**

Yes. Multiple containers can mount the same volume, although whether this is appropriate depends on the application. Many databases expect exclusive access.

---

## 9. What is the difference between a Named Volume and an Anonymous Volume?

**Answer:**

A Named Volume has a meaningful user-defined name, making it easy to manage and reuse. An Anonymous Volume is automatically named by Docker and is more difficult to identify.

---

## 10. What are Docker Volume Plugins?

**Answer:**

Docker Volume Plugins allow Docker to use external storage systems such as NFS, Amazon EFS, Azure Files, or other storage providers instead of only the local disk.

---

## 11. What is Docker's default volume driver?

**Answer:**

`local`

---

## 12. How would you persist PostgreSQL or MySQL data in Docker?

**Answer:**

Mount a named Docker Volume to the database's data directory.

Example:

```bash
docker run -v postgres-data:/var/lib/postgresql/data postgres
```

---

## 13. Your database lost all its data after recreating the container. What could have gone wrong?

**Answer:**

Possible reasons include:

- No volume was used.
- The wrong volume was attached.
- The application wrote data outside the mounted directory.
- The volume was accidentally deleted.

---

## 14. How do you inspect a Docker Volume?

```bash
docker volume inspect <volume-name>
```

---

## 15. Which Docker volume commands should every DevOps engineer know?

```bash
docker volume create <volume-name>

docker volume ls

docker volume inspect <volume-name>

docker volume rm <volume-name>

docker volume prune
```

---

## 16. Give a real-world example of when you would use a Bind Mount and when you would use a Named Volume.

**Answer:**

- **Bind Mount:** React or Node.js development to synchronize local source code with the container.
- **Named Volume:** PostgreSQL, MySQL, MongoDB, Jenkins, or any application where data must persist after the container is recreated.

---

## 17. Explain the lifecycle of a Docker Volume.

**Answer:**

A Docker Volume is created independently of containers. It can be attached to one or more containers, remains after containers are stopped or removed, and is deleted only when explicitly removed.

---

## 18. How would you troubleshoot a container that lost its persistent data?

**Answer:**

- Verify that a volume was used.
- Check whether the correct volume is attached.
- Inspect the volume using `docker volume inspect`.
- Verify the application's data directory matches the mounted path.
- Ensure the volume wasn't accidentally removed.