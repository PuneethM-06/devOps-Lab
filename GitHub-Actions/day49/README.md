## DAY 49 - GITHUB ACTION

### WHY IS GITHUB ACTIONS WORKFLOW SLOW?
- The r**unner is brand new every single time**
- After every small change, it will build everything again from first and hence installing all those dependencies again will take a lot of time 

- Since every workflow run gets a new runner, once the job completes, the runner is destroyed and everything inside it (including the cloned repository and downloaded dependencies) is lost.

- To avoid downloading the same dependencies repeatedly, GitHub stores cache data in its managed cache storage outside the runner.

- During a future workflow run, a new runner is created. Before downloading dependencies, GitHub checks whether a matching cache exists. If it does, the cache is restored onto the new runner, avoiding unnecessary downloads and making the workflow faster.

## CACHE KEYS
- cache keys are used to validate if the cache stored is valid or no 
- Suppose today i am using react 18 and tomorrow if I am using 19, It cannot restore yesterday cache and here is where we make use of cache keys 

- The **cache keys should change whenever the dependencies changes**
- The cache should be implemented by the developer. 
- Implement cache on these folders and use these hash keys 

- Example:
```
- name: Cache npm dependencies
  path: ~/.npm
  key: node-${{ runner.os }}-${{ hashFiles('**/package-lock.json;)}}
```
- Here:
    1. path is where cache should be stored
    2. node is just a prefix and it can be anything 
    3. runner.os - If the runner is ubuntu then os becomes linux 
    4. `{{ hashFiles('**/package-lock.json;)}}`- Github reads this and computes a hash 
- So the final hashkey looks like `node-linux-abs123`

## CACHE HIT and CACHE MISS
- **Cache hit** is a scenario where we the cache key wouldnt have changed and remain the same for the next workflow 

- **Cache Miss** is a scenario where the cache key would have changed and we will have to perform the actions again from first 