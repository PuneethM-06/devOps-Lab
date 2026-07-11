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
