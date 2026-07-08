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
