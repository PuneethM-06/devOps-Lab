## DOCKER DAY 23

## WHY IS PERSISTENT STORAGE NEEDED
- We need persistent storage because containers are ephemeral. This means that containers storage is temporary. upon restarting it will lose its previous memory 

## SOLUTION 
- Docker stores important data outside the containders writable layer.
- So even though when we recreate the container. the data is still safe.

## DOCKER VOLUMES
- Docker Volumes are storage area managed by docker that exists independently of a container 
- Docker volumes are stored on the host machine 
- command to create a docker volume 
```
docker volume create mysql-data
```
`
```
docker run -d \
  --name mysql \
  -v mysql-data:/var/lib/mysql \
  mysql
```
- mysql-data -> This is the docker volume
- /var/lib/mysql -> This is the folder path inside the container 

