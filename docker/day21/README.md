## DOCKER DAY 21

## WHAT PROBLEM WILL DOCKER COMPOSE TRY TO SOLVE?
- Suppose an application has 3 containers, now trying to start each container might be difficult and time consuming and inorder to overcome this issue we need docker-compose
- Problems without docker compose are:
    - Too many commands
    - Easy to make mistakes
    - Hard to share
    - No version control 

- Docker compose resolves this issue with a single yaml file, so creating a yaml file and then with a single command everything comes up and does that job 

- ### DOCKER IS IMPERATIVE - We Tell Docker Exactly What To Do
- ### DOCKER COMPOSE IS DECLARATIVE - We describe the desired state and everything comes up and works 

## FUNDAMENTALS OF DOCKER-COMPOSE FILE
- Instead of writing multiple commands, we describe
    - Which containers to run 
    - which image to use
    - which port to expose
    - which volumes to mount 

### WHY YAML?
- YAML isn't a markup language
    - Easy to edit 
    - Human readable 
    - Widely used 
    - Supports nested structures

1. ### SERVICES
- Everything that we want to run in a docker compose file is ran under this.
```
services:

  frontend:
    image: nginx

  backend:
    image: node

  database:
    image: postgres
```
2. ### SERVICES
- Tells docker to create a container using particular image 

3. ### BUILD
- Sometimes we dont have an image but we have an docker file in the current directory before creating the container

4. ### CONTAINER_NAME
- Generally, docker compose creates container name automatically. But we can overwrite using container_name 

5. ### PORTS
- Interview question 8080:80, Here 8080 is Host port and 80 is container port

5. ### RESTART
- Tells what docker should do if a container stops

6. ### ENVIRONMENT 
- Used to pass enviornment variables to the container 

7. ### COMMAND
- Docker has a default CMD but docker compose can overwrite it 

8. ### WORKING_DIR
- specifies the working directory inside the container before the command execution starts

| `image`                        | `build`                              |
| ------------------------------ | ------------------------------------ |
| Uses an existing image         | Builds an image from a Dockerfile    |
| Pulls if missing               | Builds locally before running        |
| Faster if image already exists | Needed for your own application code |

```
docker compose up
        │
        ▼
Read compose.yaml
        │
        ▼
Validate YAML
        │
        ▼
Build images (if required)
        │
        ▼
Pull missing images
        │
        ▼
Create networks
        │
        ▼
Create volumes
        │
        ▼
Create containers
        │
        ▼
Start containers
```
