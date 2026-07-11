## DOCKER DAY 22

## BRIDGE NETWORK
- whenever we install docker, by default it creates 3 network bridges
    - BRIDGE
    - HOST 
    - NONE

## WHAT IS A BRIDGE NETWORK
- This is an internal virtual switch to which any container that is communicated can talk to each other 
```             Docker Host

      +-----------------------+
      |                       |
      |   Docker Bridge       |
      |      (docker0)        |
      |                       |
      +-----------------------+
         |      |       |
         |      |       |
      Container Container Container
         A         B         C
```

## INTERNALLY,
```
docker run

↓

Create container

↓

Create network namespace

↓

Create virtual ethernet pair (veth)

↓

One end goes inside container

↓

Other end connects to docker0 bridge

↓

Assign container IP

↓

Update routing table

↓

Container starts
```

### WHAT IS VETH PAIR?
- Docker connects containers making use of veth pair
- Every container that is connected to veth pair gets a private IP

### COMMUNICATION FLOW
- Any communication that happens between containers are never gone out of the system 
```
A

↓

docker0 bridge

↓

B
```

### ADVANTAGE OF BRIDGE NETWORK
- Default network -  no setup needed
- Provides isolation between containers and host
- containers can access via NAT

### LIMITATIONS
- Does not provide automatic DNS
- communication is limited to containers on the same docker host
## COMMUNICATION FLOW
```
Container A
    │
   eth0
    │
  veth
    │
+-------------------+
| docker0 bridge    |   ← Virtual switch
+-------------------+
    │
  veth
    │
   eth0
    │
Container B
```

## HOST NETWORK
- By default, Docker created network for each host. Without HOST NETWORK, Docker does not create a seperate network
- Instead, it shares the host machines network

## ADVANTAGES
- No seperate configurations needed
- slightly better performance in network

## DISADVANTAGES
- No network isolation 
- Port conflicts

### WHEN DO WE USE HOST NETWORK OVER BRIDGE NETWORK?
- We use when application needs Host's network directly instead of its own.
- example: Running agents on host to monitor etc. 

## CUSTOM BRIDGE NETWORK
- We need CUSTOM BRIDGE NETWORK to solve the problem of "Containers IP are not stable"
- Example: Lets say, backend is running on 172.17.0.2 and frontend is listening on 172.17.0.2:80 and for some reason if docker rm backend and start again, now backend might be listening on 172.17.0.5 and hence the connecttion fails and hence we need custom brifdge  network 

- So custom bridge network helps you create your own network, and run applications on it so here we do not have to use `http://172.17.0.2:8000` instead we can use `http://backend:8000` and docker resolves backend automatically in the url 

## Advantgae
1. Can be used for Prod
2. Easy to manage
3. Dont have to worry about changing IP
4. Better isolation 

### command to create
- docker network create app-network

```
Interview Questions
1. Why use a custom bridge network instead of the default bridge?

Answer:
Because it provides automatic DNS-based service discovery, making containers communicate using names instead of changing IP addresses.

2. How do containers communicate on a custom bridge network?

By using container names, which Docker resolves to the correct IP through its built-in DNS.

Example:

http://backend:8080
3. What problem does Docker DNS solve?

It removes the need to hardcode container IP addresses, which can change when containers are recreated.

4. What command creates a custom bridge network?
docker network create app-network
```

## DOCKER DNS
- Docker has its own DNS server
- Which behaves as the DNS server and helps in providing IP address to each container for communication 

### WHY DO WE NEED DOCKER DNS
- Suppose we have a backend container that resolves to `172.18.0.3`. If the backend container restarts for some reason. The container may now get `172.18.0.5` because of which the application might not work as expected and hence we make use of DOCKER DNS 

- DOCKER DNS resolves the container name automatically and ensures that there are IP corrections made accordingly for seamless communication and up running of application 
