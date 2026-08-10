# DAY 70 - HELM TEMPLATING 

## HELM TEMPLATING
- It is the process of using template expressin and logic inside kubernetes manifests so helm can dynamically generate the final kubernetes manifests 
```
Static Kubernetes YAML
        ↓
Add Helm template expressions
        ↓
Helm processes the template
        ↓
Final Kubernetes YAML\
```
- Example:
- without helm we do something like `replicas: 3` which is fixed and does not change 
- With helm we can do; `replicas: {{ .Values.replicaCount }}`
- The above way is dynamic and Helm will read through those logic or templates and create the final k8s manifest 
- `{{ }}` is template syntax expression 
- Helm process **these expressions before the k8s sees the final manifests**

1. ### USING .values INSIDE TEMPLATES
- Suppose we have a `values.yaml` file
```
replicaCount: 3

image:
  repository: nginx
  tag: "1.27"
```
- Then our templates can read values like 
```
spec:
    replicas: {{ .Values.replicaCount }}

containers:
    - name: myApp
      image: "{{.Values.image.repository}}:{{ .Values.image.tag }}
```
2. ### if/else
- This is used to make the Helm more dyanmic
- Example:
- Suppose we only want to create something when a vale is enabled 
- In `values.yaml`:
```
ingress:
    enabled: true 
```
- In the template:
```
{{ if .Values.ingress.enabled }}
apiVersion: networking.k8s.io/v1
```
- We can also do\
```
{{ if CONDITION }}
    ↓
Condition true → execute this
    ↓
{{ else }}
    ↓
Condition false → execute this
    ↓
{{ end }}
```
- Example:
```
{{- if eq .Values.app_env "dev" }}
docker_image: abc
{{- else }}
docker_image: xyz
```

3. ### with 
- It is used to manage and organize the working of a file 
- Suppose we have a `values.yaml` file:
```
image:
  repository: nginx
  tag: "1.27"
  pullPolicy: IfNotPresent
  name: myapp
  registry: docker.io
```
- with this we have to repeatedly do 
```
.Values.image.repository
.Values.image.tag
.Values.image.pullPolicy
.Values.image.name
.Values.image.registry
```
**with** says, I will be referring to this for the next working block 
```
{{ with .Values.image }}

repository: {{.repository}}
tag: {{ .tag }}
{{ end }}
```
4. ### range 
- This is used for looping through a list of values
- suppose values.yaml
```
ports:
    - 8080
    - 9090
    - 3000
```
- We can do it as:
```
{{ range .Values.ports }}
- port: {{ . }}
{{ end }}
```
5. ### default 
- it lets you fallback to the default value when a value hasnt been supplied
- Example:
```
replicas: {{ default 3 .Values.replicaCount }}
```
