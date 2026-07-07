## DOCKER DAY 19

## DOCKER ARCHITECTURE
```
Docker Client (CLI)
        ↓
Docker Daemon (dockerd)
        ↓
containerd
        ↓
runc
        ↓
Containers
```
1. ### DOCKER CLIENT
- Example: ` docker run ngnix`, `docker ps` etc.
- This is executed through CLI, this basically sends API requests to docker daemon \

2. ### DOCKER DAEMON
- This does all the heavy lifting, it is responsible for building images, pulling images, create containers and manage networks as well 

3. ### DOCKER ENGINE
- Docker engine contains 
```
Docker CLI
+
Docker Daemon
+
containerd
+
runc
```

4. ### CONTAINERD
- Container lifecycle manager 
- It is responsible for:
    - pulling images
    - managing storage
    - managing container lifecycle

5. ### RUNC
- It is container runtime
- It is responsible for 
    - Creating namespaces
    - creating cgroups
    - starting container process

## INTERNAL LEVEL FLOW
```
docker run nginx
        ↓
Docker CLI
        ↓
dockerd
        ↓
Pull image if not present
        ↓
containerd
        ↓
runc
        ↓
Create namespaces
Create cgroups
Start process
        ↓
Container Running
```
## DOCKER ENGINE IN ACTION 
```
You type command
        │
        ▼
Docker CLI
        │
(API Request)
        ▼
Docker Daemon
        │
Checks local images
        │
───────────────
Image exists?
───────────────
Yes            No
 │              │
 │              ▼
 │       Pull image
 │              │
 └──────────────┘
        │
Create container
```
