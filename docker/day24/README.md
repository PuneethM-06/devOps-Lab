## DOCKER DAY 24


### WHY DO WE NEED DOCKER SECURIT
- A container is still running code on your host system 
- Compromising a container can give access to secrets, access mounted volumes etc.

## COMMON SECURITY RISKS
1. ### RUNNING AS A ROOT
2. Vulnerable images 
3. Secrets inside images 
4. LArge images 

## RUNNING CONTAINERS AS NON ROOT USER
- `USER appuser`
- Run the remainng instruction and the application using the user
- Common pattern is:
    1. Run as root during build 
    2. Then switch to app user
```
RUN apt-get update && apt-get install ...
Then 
USER appuser
```
