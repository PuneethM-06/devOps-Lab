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
- K8s provides the environment variables to the container from the config maps before the application starts and the application reads from the containers
```
ConfigMap
     │
     ▼
Environment Variable
     │
     ▼
Container
     │
     ▼
Application
```
### Definition:
- Environment variables allow k8s to inject configmaps or secret values into a container, which the application can make use at the runtime

### METHOD 2 - MOUNTED VOLUMES - ConfigMaps/secret volumes
- A configvolume/Secret volume is a way to provide file inside a container 
- In this case instead of injecting the configMap/secrets k8s creates the config file in the container which will be made use at the runtime 
- K8s converts/creates the configMaps/secrets into files and provides during the runtime right?

## CREATING A CONFIGMAP USING ENV
- Suppose our application needs:
```
LOG_LEVEL=DEBUG
DATABASE_HOST=postgres.default.svc.cluster.local
```

We create a ConfigMap:
```
apiVersion: v1
kind: ConfigMap

metadata:
  name: backend-config

data:
  LOG_LEVEL: "DEBUG"
  DATABASE_HOST: "postgres.default.svc.cluster.local"\
```
- Once the configMap is created the next thing we'd like to do is, **Create a deployment pod** and inside the **pod spec** we say to use the configmap 

- The next part is to make use of `env`
```
containers:
- name: backend
  image: backend:v1

  env:
```
- we use `env` because we want k8s to create environment variables 
```
containers:
- name: backend
  image: backend:v1

  env:
  - name: LOG_LEVEL
    valueFrom:
      configMapKeyRef:
        name: backend-config
        key: LOG_LEVEL
```
- Here we are giving the name to our environemnt variable and we are saing refer the value from `ConfigMapKeyRef`because we created `kind configMap` and we are using` name: backend-config` because we gave the `configmap meta data name` called as backend-config and then the key it needs it refer tere is `LOG_LEVEL`

## CREATING A CONFIGMAP USING A MOUNTED VALUE
- Here the congif map we are gonna use is:
```
apiVersion: v1
kind: ConfigMap

metadata:
  name: backend-config

data:
  application.properties: |
    LOG_LEVEL=DEBUG
    DATABASE_HOST=postgres.default.svc.cluster.local
```
- Here notice carefully, **WE ARE NOT STORING INDIVIUAL VALUES INSTEAD WE ARE STORING IT IN A FILE CALLED `application.properties`**

- #### Now the deployment
```
containers:
- name: backend
  image: backend:v1

  volumeMounts:
  - name: config-volume
    mountPath: /etc/config

volumes:
- name: config-volume
  configMap:
    name: backend-config
```
- **volumemounts** - define where it should be mounted
- **volumes** - defines what storage exisits for the pod 

- **mountPath** - creates a directory inside the container 
- Then we create volume at the container level (look at the indentation), once that is done. the `name` of `volumeMounts` and `volumes` must match. **Because this connects the container to the volume**
- Now the volume gets the data from the configMap

## CREATING A SECRET 
```
apiVersion: v1
kind: Secret

metadata:
  name: backend-secret

type: Opaque

stringData:
  DATABASE_PASSWORD: MyPassword123
  API_KEY: sk_live_xxxxx
```

- `kind: secret` - says we are storing secrets 
- `type opaque` - default secret type and it means it stores key-value pair
- `stringData` - we defined our secrets as key-value pairs

- ### IN DEPLOYMENT
```
apiVersion: apps/v1
kind: Deployment

...
spec:
  template:
    spec:
      containers:
      - name: backend
        image: backend:v1

        env:
        - name: DATABASE_PASSWORD
          valueFrom:
            secretKeyRef:
              name: backend-secret
              key: DATABASE_PASSWORD
```
