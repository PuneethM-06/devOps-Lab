# DAY 69 - HELM VALUES & ENVIRONMENT CONFIGURATION 

1. ### WHAT ARE HELM VALUES 
- Helm values are **configuration parameters that are supplied to helm chart and used by helm templates**
- Example:
```
replicaCount: 2

image:
  repository: nginx
  tag: "1.27"
```
so if the template says:
`replicas: {{ .Values.replicaCount }}`
- helm understand the value as `replicas: 2`

### NESTED HELM VALUE
- Example:
```
image:
  repository: nginx
  tag: "1.27"
  replicaCount: 2
```
- `replicas: {{ .Values.image.replicaCount}}`