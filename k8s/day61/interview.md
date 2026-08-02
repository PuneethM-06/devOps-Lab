# DAY 61 - KUBERNETES

## WHY DO WE NEED CONFIGMAPS AND SECRETS?
- As we know the principle `BUILD ONCE, DEPLOY ANYWHERE`, so if we have secrets for each env and we are deploying to each env. we should build the image again and deploy which is bad engineering and hence we make use of `configMap` and `secrets`
- **configMap - Non-sensitive informations like API UR?L etc,**
- **Secrets -  Are for sensitive information like API KEYS, pwd etc.**
- In short, ConfigMaps and secrets are used for **ONE IMAGE, MULTIPLE CONFIGURATIONS**

### WHERE DOES THE CONFIGURATION COME FROM?
- K8s has 2 way of doing this 
1. **Environment variables**
2. ** Mounted volumes**

### METHOD 1 - ENVIRONMENT VARIABLES 
- Example: ` syste,.getenv("DATABASE_HOST")
- Here, k8s is responsible for injecting the environment variables 
