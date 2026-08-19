# interview.md — Day 89: Distroless Images

## Must-Know Fundamentals

### 1. What is a distroless image?

**Expected answer:**

A distroless image is a minimal container runtime image that removes unnecessary general-purpose tooling such as shells and package managers. It contains only the runtime components needed to run the application.

The goal is to reduce unnecessary components and therefore reduce the potential attack surface.

---

### 2. Does distroless mean there is no operating system inside the image?

**Expected answer:**

No. Distroless does not mean there are no OS or runtime components.

It means unnecessary general-purpose distribution tooling is removed, while selected runtime components and libraries required by the application can still be present.

---

### 3. How is a traditional image different from a distroless image?

**Expected answer:**

A traditional image such as Ubuntu, Debian, or Alpine generally contains more general-purpose tooling.

```text
Traditional image

Application
+ Shell
+ Package manager
+ System utilities
+ Libraries
```

A distroless image is more focused on running the application.

```text
Distroless image

Application
+ Required runtime components
+ Required libraries
```

Distroless generally provides less convenience but a more minimal production runtime.

---

### 4. What is an attack surface?

**Expected answer:**

An attack surface is the collection of components, interfaces, and potential points that an attacker could potentially exploit.

Additional unnecessary components can introduce:

- Known vulnerabilities
- Misconfigurations
- Additional executables or capabilities that could be abused

Removing unnecessary components helps reduce the potential attack surface.

---

### 5. Why does removing a shell or package manager help reduce the attack surface?

**Expected answer:**

A shell or package manager provides additional functionality that may not be required for the application to run.

If unnecessary components are removed, there are fewer binaries, packages, and capabilities that may contain vulnerabilities or potentially be abused after a compromise.

However:

> Fewer components do not automatically mean the application is completely secure.

---

### 6. What happens if you run `sh` inside a typical distroless container?

**Expected answer:**

It usually fails because distroless images generally do not include a shell such as `sh` or `bash`.

For example:

```bash
docker exec -it my-container sh
```

The command fails because the shell does not exist inside the container.

---

### 7. Why does the lack of a shell create a trade-off?

**Expected answer:**

Removing the shell helps keep the production image minimal, but it also makes direct debugging inside the application container more difficult.

```text
No shell
    ↓
Smaller production runtime
    +
Reduced attack surface

But
    ↓
Harder to debug directly
```

---

### 8. How can you debug a distroless container?

**Expected answer:**

Start with observability:

- Application logs
- Metrics
- Traces
- Kubernetes events and logs

If deeper debugging is required, use a separate debugging mechanism such as an ephemeral/debug container.

The debugging tools do not need to be permanently included in the production application image.

---

### 9. What is an ephemeral/debug container?

**Expected answer:**

An ephemeral or debug container is a temporary container added for troubleshooting.

Conceptually:

```text
Pod
│
├── Application Container
│      └── Distroless
│          └── No shell
│
└── Debug Container
       ├── Shell
       ├── curl
       ├── ping
       └── Other debugging tools
```

The production image remains minimal while debugging tools are available separately when needed.

---

### 10. Why is an ephemeral debug container better than permanently adding debugging tools?

**Expected answer:**

Debugging tools are only added when required and are not permanently shipped inside the production application image.

```text
Need debugging
      ↓
Add temporary debug container
      ↓
Investigate
      ↓
Remove when finished
```

This preserves the minimal nature of the production image.

---

### 11. Why do Go applications work well with distroless images?

**Expected answer:**

A Go application can be compiled into an executable binary.

The production container does not need the Go compiler or source code to run that compiled application.

```text
Go source code
      ↓
go build
      ↓
Compiled binary
      ↓
Distroless runtime
```

This makes Go a strong candidate for minimal production images.

---

### 12. Why use a multi-stage build with Go?

**Expected answer:**

The builder stage contains everything required to build the application:

- Go compiler
- Source code
- Dependencies
- Build tools

Only the compiled binary and required runtime components are copied into the final image.

```text
Builder stage
├── Go compiler
├── Source code
├── Dependencies
└── Build tools
        ↓
   Compiled binary
        ↓
Final stage
├── Minimal base image
└── Application binary
```

The build tools are not included in the final production image.

---

### 13. What is `scratch`?

**Expected answer:**

`scratch` is essentially an empty base image.

You must provide everything required for the application to run.

For a fully self-contained static binary:

```text
scratch
   +
Application binary
   ↓
Extremely minimal image
```

---

### 14. What is the difference between `scratch` and distroless?

**Expected answer:**

`scratch` is essentially empty, while distroless provides a minimal runtime environment.

| `scratch` | Distroless |
|---|---|
| Essentially empty | Minimal runtime environment |
| Everything must be provided | Selected runtime components available |
| Can be extremely small | Usually slightly larger |
| Best for fully self-contained binaries | Useful when minimal runtime support is needed |

---

### 15. When would `scratch` be a good choice for a Go application?

**Expected answer:**

If the Go binary is fully static and self-contained with no additional runtime requirements, `scratch` can be an excellent minimal choice.

---

### 16. Does using distroless automatically make an application secure?

**Expected answer:**

No.

Distroless reduces unnecessary components and can reduce the attack surface, but the application can still contain:

- Vulnerable dependencies
- Vulnerable runtime libraries
- Application code vulnerabilities
- Misconfigurations
- Secrets included in the image

Distroless is one security layer, not a complete security solution.

---

### 17. How does distroless fit with Gitleaks, Snyk, and Trivy?

**Expected answer:**

Each tool addresses a different security area.

```text
Gitleaks
    ↓
Detect exposed secrets

Snyk
    ↓
Detect dependency vulnerabilities

Trivy
    ↓
Scan container images and artifacts

Distroless
    ↓
Reduce unnecessary runtime components
and attack surface
```

These layers complement each other.

---

### 18. Is distroless mainly suited for development or production?

**Expected answer:**

Distroless is generally more production-focused.

During development, a shell and debugging tools can be useful because the application is actively being tested and changed.

```text
Development
    ↓
More debugging convenience

Production
    ↓
Minimal runtime
+ fewer unnecessary components
```

Distroless can still be used during development, but it is often less convenient.

---

## Scenario-Based Questions

### 19. A Go application is currently running from a `golang` base image. How would you make the production image more minimal?

**Expected answer:**

Use a multi-stage build.

Compile the application in a builder stage and use a distroless image as the final runtime stage. Copy only the compiled binary and required runtime components into the final image.

This removes the Go compiler, source code, and build tools from the production image.

---

### 20. `kubectl exec -it <pod> -- sh` fails against a distroless container. Why?

**Expected answer:**

It fails because the distroless application image generally does not contain a shell such as `sh` or `bash`.

Debugging should instead use logs, metrics, traces, Kubernetes events, or an ephemeral debug container when deeper investigation is needed.

---

### 21. A team says, "We use distroless, so our application is secure." What is wrong with this statement?

**Expected answer:**

Distroless reduces unnecessary components and the potential attack surface, but it does not fix vulnerable application code, dependencies, runtime libraries, secrets, or misconfigurations.

A distroless application can still be vulnerable.

---

### 22. You have a fully static Go binary with no runtime dependencies. Which would be more minimal: `scratch` or distroless?

**Expected answer:**

`scratch`.

Because the binary is fully self-contained, additional runtime components are not required.

---

### 23. Why use an ephemeral debug container instead of permanently adding debugging tools to a production distroless image?

**Expected answer:**

The debug tools are only available when needed and are not permanently included in the production application image.

This keeps the production image minimal and avoids shipping unnecessary tools such as shells, `curl`, or `ping` with every production container.

---

## Final Interview Answer

> **Distroless images are minimal, production-focused container images that remove unnecessary general-purpose tooling such as shells and package managers. They contain the runtime components required to run the application, helping reduce unnecessary packages and the potential attack surface. The trade-off is that direct debugging becomes harder because commands such as `sh` or `bash` may not exist. Debugging can instead rely on logs, metrics, tracing, Kubernetes events, and ephemeral debug containers. Distroless works particularly well with multi-stage builds for Go applications, where the application is compiled in a builder stage and only the binary and required runtime components are copied into the final image. If the Go binary is fully static and self-contained, `scratch` can be an even more minimal option. However, distroless does not automatically make an application secure because vulnerable dependencies, application code, and misconfigurations can still exist.**

# Day 89 — Distroless Images: No Shell, Smaller Attack Surface

## What I Learned

- Learned what distroless images are and why they are production-focused.
- Understood that distroless does not mean there are no OS or runtime components; it means unnecessary general-purpose tooling is removed.
- Compared traditional images with distroless images.
- Learned how removing unnecessary components can reduce the potential attack surface.
- Understood why distroless images generally do not include shells such as `sh` or `bash`.
- Learned the trade-off between a minimal runtime and debugging convenience.
- Covered debugging through:
  - Application logs
  - Metrics
  - Tracing
  - Kubernetes events
  - Ephemeral/debug containers
- Learned why temporary debug containers are better than permanently adding debugging tools to production images.
- Connected distroless with Go and multi-stage builds.
- Understood that the builder stage contains the compiler, source code, dependencies, and build tools, while the final stage contains only what is needed at runtime.
- Compared `scratch` and distroless.
- Learned that `scratch` is essentially empty and works well for fully static, self-contained Go binaries.
- Understood that a smaller attack surface does not automatically mean a secure application.
- Connected distroless with other security layers such as Gitleaks, Snyk, and Trivy.

## Final Mental Model

```text
Traditional Image
        ↓
More general-purpose tools
        ↓
More convenience
Easier debugging
Potentially larger attack surface

Distroless
        ↓
Application
+ required runtime components
        ↓
No general-purpose shell
No package manager
        ↓
Minimal production runtime
Smaller attack surface

scratch
        ↓
Essentially empty base
        +
Fully self-contained application
        ↓
Extremely minimal image
```

```text
Go Application
      ↓
Multi-stage build
      ↓
Builder Stage
├── Go compiler
├── Source code
├── Dependencies
└── Build tools
      ↓
Compiled binary
      ↓
Final Stage
      ↓
Distroless or scratch
      ↓
Production image
```

## Key Takeaway

Distroless images remove unnecessary general-purpose tooling and provide a minimal runtime environment for production applications. This helps reduce the potential attack surface, but debugging becomes less convenient and must rely more on observability or temporary debug containers. For Go applications, multi-stage builds work particularly well because only the compiled binary and required runtime components need to be included in the final image. Distroless improves runtime minimization but does not replace vulnerability scanning, secure coding, dependency management, or other security controls.