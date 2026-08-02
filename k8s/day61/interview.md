# DAY 61 - KUBERNETES

## WHY DO WE NEED CONFIGMAPS AND SECRETS?
- As we know the principle `BUILD ONCE, DEPLOY ANYWHERE`, so if we have secrets for each env and we are deploying to each env. we should build the image again and deploy which is bad engineering and hence we make use of `configMap` and `secrets`
- **configMap - Non-sensitive informations like API UR?L etc,**
- **Secrets -  Are for sensitive information like API KEYS, pwd etc.**
- In short, ConfigMaps and secrets are used for **ONE IMAGE, MULTIPLE CONFIGURATIONS**