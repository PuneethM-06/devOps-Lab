# DAY 50 - GITHUB ACTIONS CONTINUATION 

## MULTIPLE JOBS
- Every job gets it's own brand new VM
- GitHub runners are ephemeral 
- Multiple jobs give:
    - Isolation 
    - scalability
    - clearer logs
    - easier debugging

### WHY DO WE NEED `actions/checkout` IN MULTIPLE JOBS AND ISNT IT REDUDANT?
- Because every job has it's own new runner which wouldn't have the files cloned or the repo cloned as they start and hence we would need it seperately

## needs
- As we know, without needs GitHub actions starts all the jobs almost at the same time since each job is isolated and each job get their own runner.
- But in some sitations it is needed that a particular job should start after one job is done executing and hence we need needs here 
- needs controls the order only and not share files between jobs 
- A j**ob can wait for more than 1 jobs**