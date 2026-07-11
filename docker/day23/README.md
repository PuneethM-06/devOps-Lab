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

## WHAT IS A BIND MOUNT 
- A bind mount maps a specific folder from your host machine to a container 
- Unlike a docker volume, Docker does not manage this storage 

- Syntax
```
docker run -v <host-path>:<container-path> image
```
```
Host
/home/puneeth/project
        │
        ▼
Container
/app
```
### BIND MOUNTS ARE USUALLY PREFFERED FOR LARGE APPLICATIONS BECAUSE OF LIVE SYNCHRONISATION OF FILES 

Development → Bind Mounts
Production → Docker Volumes

### WHY AREN'T BIND MOUNTS PREFERRED FOR PRODUCTION
- Because they depend on host-specific directory paths, reducing portability and these volumes are not managed by docker

## NAMED VOLUMES vs BIND VOLUMES

- Named volumes are under docker and it decides its storage path 
- Named volumes are used for production 
- Bind volumes are used for development 
- Bind volumes storage is not decided by docker and it is stored in our local system

## ANONYMOUS VOLUMES
- A docker managed volume without a user-defined name 
```
Example of named vol: docker run -v mysql-data:/var/lib/mysql mysql
Example of Anonymous volume: docker run -v /var/lib/mysql mysql
```

## VOLUME LIFECYCLE
1. CREATE A VOLUME 
2. ATTACH IT TO A CONTAINER
3. STOP THE CONTAINER
4. REMOVE THE CONTAINER
5. START A NEW CONTAINER 

```
Create Volume
      │
      ▼
Attach to Container
      │
      ▼
Stop Container
      │
      ▼
Remove Container
      │
      ▼
Volume Still Exists 
      │
      ▼
Attach to New Container
      │
      ▼
Data Still Available 
```

## VOLUME PLUGINS 
- A volume plugin lets docker store volume data somewhere other than the local disks 
- If a server crashes, the volumes inside the server also die and we might lose the data but with help of plugins.
If server A crashes, we can use plugins give the stored data to server B and use 

- Docker supports multiple plugins but the default is local 
- Syntax
```
docker volume create \
  --driver <driver-name> \
  my-volume
```
## COMMANDS TO KNOW
```
docker volume ls

docker volume inspect <volume>

docker volume rm <volume>

docker volume prune

docker inspect <container>
```
