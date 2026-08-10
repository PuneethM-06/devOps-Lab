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
