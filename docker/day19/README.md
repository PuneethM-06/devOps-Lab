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
### CONTAINERD AND RUNC
- Docker orchaestrates the process while the actual container is started by runc

- CONTAINERD is container run time manager. It is reponsible for:
        - Pulling images
        - Unpacking images
        - Managing images
        - starting and terminating containers

### RUNC
- Runc is a low-level container runtime 
- It's job is to create a linux container and start the process
- It is responsible for creating namespaces, Cgroups

## DOCKER REGISTRY
- Docker registry is like a collection of container images
- Example: dockerhub, ECR etc

### REPOSITORY
- A repositroy are groups of different versions of the same images 

### TAGS
- Represent the particular version of an image 

## NOTE: Never use latest as default because as versions change and we pull the latest everytimne, there can be a possiblity that our application might behave differently 

- A single image can create multiple containers
- Multiple containers can be created from a single image because the image is read-only and shared.
- Images are immutable and are READ-ONLY
## QUESTION

1.  If dockerd crashes while 10 containers are already running, what happens to those containers?
- Answer: These containers that are already running will work fine since it is managed by containerd. But New containers cannot be created because the Docker CLI cannot communicate with the Docker daemon (dockerd), which is responsible for orchestrating container creation and management.