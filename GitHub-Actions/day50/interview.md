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
