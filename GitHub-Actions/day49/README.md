## DAY 49 - GITHUB ACTION

### WHY IS GITHUB ACTIONS WORKFLOW SLOW?
- The r**unner is brand new every single time**
- After every small change, it will build everything again from first and hence installing all those dependencies again will take a lot of time 

- Since every workflow run gets a new runner, once the job completes, the runner is destroyed and everything inside it (including the cloned repository and downloaded dependencies) is lost.

- To avoid downloading the same dependencies repeatedly, GitHub stores cache data in its managed cache storage outside the runner.

- During a future workflow run, a new runner is created. Before downloading dependencies, GitHub checks whether a matching cache exists. If it does, the cache is restored onto the new runner, avoiding unnecessary downloads and making the workflow faster.