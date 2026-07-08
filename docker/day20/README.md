## DAY 20 - DOCKER

## WHY DO WE NEED DOCKER IMAGE OPTIMIZATION
- Docker image optimization for ann application is important when we are deploying it to k8s.
- It ensures that the application is deployed faster
- Suppose k8s has to replace a node and it has to pull the docker image for that and if the application is 2GB. Recovery time will be slow and hence we need optimal docker images 
- Lesser storage usage 
- Better security 
- Better CI/CD performance 

## NOTE: 
- Smaller image generally doesnt mean it consumes less RAM. Obviously, it consumes less storage but RAM consumption depends on application and its workload

## MULTI-STAGE IMAGE
```
             Stage 1

Source Code
      │
      ▼
npm install
      │
      ▼
npm run build
      │
      ▼
dist/
      │
      │ Copy only dist/
      ▼

             Stage 2

Fresh Image
      │
      ▼
Copy dist/
      │
      ▼
Production Image
```
Here, stage 1 makes used of source code to produce a dist / folder.In stage 2 it is the Fresh image plus dist /

#### "A multi-stage build uses multiple FROM instructions in a single Dockerfile. The first stage builds the application using the required build tools, and the final stage starts from a fresh base image and copies only the build artifacts needed to run the application. This reduces image size, improves security, and removes unnecessary build dependencies."

- After stage 2, when we get the final image. Stage 1 will be discarded
```
# ---------- Stage 1: Build ----------
FROM node:22-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY . .

RUN npm run build

# ---------- Stage 2: Production ----------
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

## CHOOSING THE RIGHT BASE IMAGE 
1. UBUNTU
2. ALPINE
3. DEBIAN
4. SLIM
5. DISTROLESS
6. SCRATCH

1. ### UBUNTU
- It is a full Linux distribution 
- Large image 

2. ### DEBIAN
- Debian is the parent distribution of ubuntu
- Debian is usually smaller than Ubuntu

3. ### ALPINE LINUX
- Tiny image 
- Fast downloads
- Lower storage

4. ### SLIM IMAGES
- Its like UBUNTU -> REMOVE UNWANTED PACKAGES -> SLIM
- Smaller than full images
- Easier debugging

| Base Image |       Size | Compatibility          | Debugging | Typical Use                  |
| ---------- | ---------: | ---------------------- | --------- | ---------------------------- |
| Ubuntu     |      Large | Excellent              | Easy      | Development, general-purpose |
| Debian     |     Medium | Excellent              | Easy      | Production servers           |
| Slim       |      Small | Excellent              | Easy      | Most production apps         |
| Alpine     | Very Small | Can require extra work | Harder    | Lightweight services         |

## NOTE:
- Alpine is not always the best choice because if our application is making use of packages that are compiled against glibc, moving to alpine can break the application 

## IMAGE OPTIMIZATION 

1. ### EVERY LAYER MATTERS
- Here combining RUN commands can optimize the docker image and also ensure that we have a better image history
```
FROM debian:bookworm-slim

RUN apt update && \
    apt install -y curl git
```

### QUESTION: Does having fewer layers always make an image smaller?
- No, If both docker files install the same software, the size maybe almost the same but the advantage is cleaner image history, Better maintainability, Easier cleanup

2. ### CLEAN PACKAGE CACHE
- After the image is created, docker may still have the APT package lists and which are not needed and hence we can do `rm -rf /var/lib/apt/lists/*`
```
RUN apt update && \
    apt install -y curl && \
    rm -rf /var/lib/apt/lists/*
```

3. ### INSTALL ONLY WHAT IS NEEDED
-  Instead of installing git, curl, vim etc. Which is not needed we can install only what is needed

4. ### ORDER MATTERS
- As we already know we need to follow the right order

5. ### USE .dockerignore
- Using dockerignore can make the image smaller and deployments faster and also save from secrets getting leaked

6. ### USE MULTI-STAGE BUILDS
- We've covered this, usiong multi-stage builds make cleaner history and also faster deployments, No build tools in runtime and all 

