# Docker Compose — Interview Questions (Day 21)

## Fundamentals

### 1. What is Docker Compose?

### 2. Why was Docker Compose introduced?

### 3. What problems does Docker Compose solve compared to multiple `docker run` commands?

### 4. Explain the difference between imperative and declarative approaches using Docker.

### 5. Why is Docker Compose commonly used for multi-container applications?

---

## compose.yaml

### 6. What is the purpose of a `compose.yaml` file?

### 7. Why does Docker Compose use YAML?

### 8. Explain the purpose of the `services` section.

### 9. What is the difference between `image` and `build`?

### 10. When would you use `build` instead of `image`?

### 11. What is the purpose of `container_name`?

### 12. Why is `container_name` generally not recommended in large projects?

### 13. Explain port mapping in Docker Compose.

### 14. What is the purpose of the `restart` policy?

### 15. Explain the different restart policies.

### 16. What is the purpose of the `environment` section?

### 17. What is the purpose of the `command` field?

### 18. What is `working_dir`?

### 19. Why is the `version` field usually omitted in modern Compose files?

---

## Docker Compose Lifecycle

### 20. What happens internally when you run `docker compose up`?

### 21. What is the difference between:

- `docker compose up`
- `docker compose up -d`

### 22. What is the difference between:

- `docker compose stop`
- `docker compose down`

### 23. What happens when you execute:

```bash
docker compose down -v
```

### 24. Why does `docker compose start` not work after `docker compose down`?

### 25. What is the purpose of:

```bash
docker compose logs
```

### 26. What is the purpose of:

```bash
docker compose ps
```

### 27. What is the purpose of:

```bash
docker compose exec
```

---

## Environment Variables

### 28. Why should configuration be stored in environment variables?

### 29. What is a `.env` file?

### 30. Explain variable substitution in Docker Compose.

### 31. Why shouldn't production secrets be stored as environment variables?

### 32. What is the difference between Docker Secrets and environment variables?

---

## depends_on & Health Checks

### 33. What does `depends_on` do?

### 34. What does `depends_on` NOT guarantee?

### 35. Why are health checks important?

### 36. How do health checks improve service startup reliability?

---

## Multi-container Applications

### 37. Explain a typical Frontend → Backend → Database architecture using Docker Compose.

### 38. How do different services communicate in a Compose application?

### 39. Why is separating an application into multiple containers beneficial?

### 40. Explain the complete workflow of a request from a browser to the database in a multi-container application.

---

# Scenario-Based Questions

### 41.

Your backend cannot connect to PostgreSQL immediately after `docker compose up`.

What could be the reason?

---

### 42.

A developer committed database credentials inside `compose.yaml`.

Why is this a bad practice?

---

### 43.

You accidentally executed:

```bash
docker compose down -v
```

What happened?

---

### 44.

Your teammate runs:

```bash
docker compose start
```

and receives an error because no containers exist.

Why?

---

### 45.

You need to start an entire application consisting of:

- React
- Spring Boot
- PostgreSQL

using one command.

How would Docker Compose help?