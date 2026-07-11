# Docker Networking - Interview Questions

## 1. What are the different network drivers available in Docker?

**Answer:**

- Bridge
- Host
- None
- Overlay
- Macvlan

Bridge is the default network for standalone containers.

---

## 2. What is the default network in Docker?

**Answer:**

The default network is the **Bridge Network**. Containers connected to it receive private IP addresses and can communicate with each other on the same Docker host.

---

## 3. What is the difference between Bridge Network and Host Network?

| Bridge Network | Host Network |
|---------------|--------------|
| Container gets its own IP | Container shares the host's IP |
| Requires port publishing (`-p`) | No port publishing required |
| Better isolation | Less isolation |
| Default choice | Used for special networking requirements |

---

## 4. Why would you create a Custom Bridge Network?

**Answer:**

A custom bridge network enables automatic Docker DNS-based service discovery, allowing containers to communicate using container names instead of IP addresses. It also provides better isolation for related containers.

---

## 5. What is Docker DNS?

**Answer:**

Docker provides a built-in DNS server for custom bridge networks that resolves container or service names to their current IP addresses, eliminating the need to hardcode IPs.

---

## 6. Why should you avoid using container IP addresses?

**Answer:**

Container IP addresses are dynamic and may change when containers are recreated. Using container names makes communication more reliable because Docker DNS always resolves the current IP.

---

## 7. Explain Docker port publishing.

**Answer:**

Port publishing maps a host port to a container port.

Example:

```bash
docker run -p 8080:80 nginx
```

This maps:

- Host Port: **8080**
- Container Port: **80**

Users connect to the host port, while the application continues listening on the container port.

---

## 8. Why can't you access a container from your browser without using `-p`?

**Answer:**

Because the application is only accessible within the container's network. Without publishing a port, Docker does not expose it to the host machine.

---

## 9. What is the purpose of the None network?

**Answer:**

The None network starts a container without any network connectivity. The container only has the loopback interface (`lo`). It is useful for offline processing or highly isolated workloads.

---

## 10. What is an Overlay Network?

**Answer:**

An Overlay Network allows containers running on different Docker hosts to communicate as if they were on the same network. It is commonly used with Docker Swarm.

---

## 11. What is a Macvlan Network?

**Answer:**

A Macvlan network assigns a container its own IP address on the physical LAN, making it appear as an independent device on the network. It is mainly used for legacy applications or specialized networking scenarios.

---

## 12. Your frontend container cannot communicate with the backend. What would you check?

**Answer:**

- Verify both containers are running.
- Check that both containers are connected to the same Docker network.
- Verify Docker DNS resolves the backend container name.
- Ensure the backend is listening on the expected port.
- Confirm the application is listening on `0.0.0.0` instead of `127.0.0.1`.
- Test connectivity using `curl` from the frontend container.
- Review application logs.

---

## 13. Which Docker networking commands do you frequently use?

**Answer:**

```bash
docker network ls
docker network inspect <network-name>
docker inspect <container-name>
docker exec -it <container-name> sh
docker logs <container-name>
docker ps
```

---

## 14. What is the difference between `127.0.0.1` and `0.0.0.0` inside a container?

**Answer:**

- `127.0.0.1` listens only inside the container itself.
- `0.0.0.0` listens on all network interfaces, allowing other containers and published ports to access the application.

---

## 15. Explain the difference between communication inside Docker and external access.

**Answer:**

- Containers communicate with each other using Docker networks and Docker DNS (e.g., `http://backend:8080`).
- External users access applications through published host ports (e.g., `http://localhost:8080`).
- Docker forwards traffic from the host port to the container port.