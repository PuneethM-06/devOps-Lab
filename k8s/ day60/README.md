# DAY 60 - KUBERNETES


## WHY DO WE NEED SERVICES?
- Let's suppose in a worked node we have 3 pods and these pods are running `front end`, `back-end` and `database`.
Now if front end needs to communicate or send request to backend it needs to know the IP address of backend. 
- Pods are ephemeral, once a pod restarts its IP restarts and hence it is hard for the frontend ot remember the backend IP all the time and also there are other issue with this method:
    1. Poda are ephemeral
    2. Scaling 
    3. Node failures 