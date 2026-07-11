## DOCKER DAY 23

## WHY IS PERSISTENT STORAGE NEEDED
- We need persistent storage because containers are ephemeral. This means that containers storage is temporary. upon restarting it will lose its previous memory 

## SOLUTION 
- Docker stores important data outside the containders writable layer.
- So even though when we recreate the container. the data is still safe.

