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


## WHAT IS A DOCKER IMAGE?
- Docker Image is a blueprint
- A docker image contains:
        - Application code
        - runtime
        - dependencies
        - filesystem 
        - metadata
- Docker Images are immutable and cannot be changed

## WHAT IS A DOCKER CONTAINER 
- Running instance of an image
- unlike an image it is running, writeable

## WHY DOESN'T DOCKER DUPLICATE THE IMAGE?
- Docker shares the same read-only image among multiple containers, while each container gets its own small writable layer.

## READ-ONLY LAYER vs WRITEABLE LAYER
- As mentioned, docker doesn;t duplicate the containers, It creates multiple containers for the same read-only image with each container getting it's own writeable layer
```
              Ubuntu Image
             (Read-only)
                    │
        ┌───────────┴───────────┐
        │                       │
        ▼                       ▼
Container A              Container B
Writable Layer          Writable Layer
```
- This is the reason why containers are lightweight 

## CONTAINER LIFECYCLE
```
             docker create
                   │
                   ▼
               CREATED
                   │
          docker start
                   │
                   ▼
               RUNNING
              ↙       ↘
docker pause         docker stop
      │                  │
      ▼                  ▼
   PAUSED             STOPPED
      │                  │
docker unpause      docker start
      │                  │
      └──────────────┬───┘
                     ▼
                  RUNNING
                     │
              docker rm
                     ▼
                  DELETED
```
1. ### CREATE
- Memory is allocated
- CPU is allocated
- Process begins to run 

2. ### PAUSE
- Process is frozen
- CPU usage stopped
- Memory remains allocated

3. ### STOP
- Memory is released
- Process exists
- CPU is no longer used

## DOCKER FILE
- A docker file is a simple recipe for building a docker image.

## WHAT HAPPENS IN A DOCKER BUILD?

- docker build -t myapp .
So when a user executes this, 
- docker: Mentions docker
- build: docker build; mentions the task
- -t: assigns a tag
- myapp: Image name 
- . - build it in the current directory 
The flow is
```
docker build
       │
       ▼
Docker CLI
       │
       ▼
dockerd
       │
Reads Dockerfile
       │
Executes instructions one by one
       │
Creates image layers
       │
Stores final image
```
## ORDER OF A DOCKER FILE 
1. FROM 
2. WORKDIR
3. COPY
4. RUN 
5. CMD

- Example:
```
FROM python:3.12

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

CMD ["python", "app.py"]
```
### ABBREVIATION
1. `FROM python:3.12` - This mentions the base image
2. `WORKDIR /app` = Creating workdir
3. `COPY . .` - Copy files to 
4. `RUN pip install` - Install dependencies
5. `CMD` -  store the default command - Basically, CMD runs when container starts running 

Note, not when image is created. But when container starts running 

## OVERALL FLOW
```
Dockerfile
      │
docker build
      ▼
Image
      │
docker run
      ▼
Container
```

## NOTE: UNDERSTANDING CMD
When we execute docker build, the Docker CLI sends the build request to dockerd. The Docker daemon reads the Dockerfile and executes each instruction sequentially. Instructions such as FROM, COPY, and RUN create image layers, while instructions like CMD are stored as metadata in the final image—they are not executed during the build. Once all the instructions are processed, Docker creates the final image. Later, when someone runs docker run using that image, Docker creates and starts a container. At that point, Docker reads the stored CMD instruction and executes it as the container's default startup command, which starts the application inside the container.

## what does FROM do?
- FROM specifies the base image from which I want to build on 
- This base image already contains all the needed file system, python, pip and also system libraries 

## WHAT DOES WORKDIR DO?
- It basically tells, from this point onwards it runs all the instruction in mentioned work directory 
- If there are multiple WORKDIR, It doesnt replace it instead it appends it 
- WORKDIR is preferred over RUN cd because every RUN instruction starts in a new shell

## WHAT DOES COPY DO?
- COPY files from your host machine to the docker image 
- General syntax is COPY <source> <destination>
- COPY ..  means copy everything from the build context from the host machine to current workdir in the image 
- COPY ..  can be expensive and hence it is a better approach if we can COPY needed files 
- Build context is a set of local files and directories that DOCKER CLI packs and sends to Docker daemonm

## Docker Build Workflow

```text
                     User
                      │
                      ▼
             docker build -t myapp .
                      │
                      ▼
                Docker CLI
                      │
        Sends Build Context + Dockerfile
          to dockerd (REST API)
                      │
                      ▼
             Docker Daemon (dockerd)
                      │
         Reads Dockerfile Line by Line
                      │
                      ▼
        ┌─────────────────────────────────┐
        │ 1. FROM python:3.12             │
        │    • Pull base image if needed  │
        │    • Use as foundation          │
        └─────────────────────────────────┘
                      │
                      ▼
        ┌─────────────────────────────────┐
        │ 2. WORKDIR /app                 │
        │    • Set current directory      │
        │    • Create /app if missing     │
        └─────────────────────────────────┘
                      │
                      ▼
        ┌─────────────────────────────────┐
        │ 3. COPY . .                     │
        │    • Copy build context         │
        │      (host → image)             │
        │    • Into current WORKDIR       │
        └─────────────────────────────────┘
                      │
                      ▼
          Create Image Layers (Cache Aware)
                      │
                      ▼
          Store Final Docker Image Locally
                      │
                      ▼
            Image Ready for docker run
```
## DOCKER INSTRUCTION ADD 
- COPY and ADD are mostly the same but they have 2 differences
- ADD = COPY + 2 EXTRA FEATURES
- Those 2 features are:
        - Automatically extract local compressed archives
        - Download files from a URL

### QUESTION:When should you use ADD instead of COPY?
- ANSWER: Use `COPY` for copying files and directories. Use `ADD` when you want to make use of one of its 2 features that is extract local archives or download files from a URL

## DOCKER INSTRUCTION RUN 
- Docker runs this while building an image and the result becomes the part of the image forever
- Every RUN creates a layer 

## DOCKER INSTRUCTION ENTRY POINT
- Here ENTRYPOINT acts as the The executable while CMD acts as the default arguements

## DOCKERFILE INSTRUCTION - ENV
- Sets enviornment variable inside the image 
- ENV PORT=8080 - so when the application starts by default it will be heard from port 8080
- This will be stored as metadata and will be executed when the container starts running and not when the image is created
- ENV or metadata will not create a layer in the docker image 

## DOCKER INSTRUCTION ARGS
- ARG defines the build time variable 
- we need ARG to make the DOCKERFILE resuable and configurable 

## DOCKER INSTRUCTION EXPOSE
- It tells where the application is listening 
- Like ENV and ARG it is also a metadata and does not create a layer
- It does not expose a port, it says where it is listening 

## DOCKER INSTRUCTION LABEL 
- Stores metadata about the image 
- Label cannot affect the application or also wont create any layers

## IMAGE LAYERS
- An immutable read-only file that represents a specific set of silesystem changes is called as a Image Layer
- Read-only ensures that there is consistency, Reproducibility, safe sharing and efficient stoarage and importantly nothing breaks 


### UNION FILESYSTEM 
- Merges multiple directories into a single logical view
```
Layer A
/bin

Layer B
/usr

Layer C
/app

becomes:
/
├── bin
├── usr
└── app
```
## WHAT IS A BUILD CACHE 
- It is used by docker to provide speed and save storage while making use of the image layers that are unchanged

## DOCKER IGNRORE FILE
- The issue is that when we do a COPY .., all the files lets say node_modules which is 855mb will also be copied and to prevent that we need docker igore file
- Basically, to send the needed files 
- It is similar to gitignore
- for most projects
```
.git
.gitignore
node_modules
coverage
logs
*.log
*.tmp
.env
.vscode
.idea
```

### EXAMPLE DOCKER FILE 
```
FROM node:22-alpine
WORKDIR app/
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
ENV NODE_ENV=production
EXPOSE 3000
CMD ["npm", "start"]
```

