# DAY 73 - HELM CHART DEPENDENCIES, REPOSITORIES AND PACKAGING 

1. ### CONCEPT 1 - HELM CHART DEPENDENCIES
- Helm allows one chart to declare another chart as a dependency 
- **A helm dependency is another chart that your chart depends on**
```
Your Chart
    ↓
Declares dependencies
    ↓
Helm obtains those Charts
    ↓
Your application + dependencies
    ↓
Installed together
```
- `chart.yaml` can declare other helm charts that your chart depends on

2. ### CONCEPT 2 - DEPENDENCY FIELD 
- It tells which chart is the dependency 
- Example:
1. **name** - says which dependency chart to use 
```
dependencies:
    - name: redis
```
- **name tells which chart is the depedency**

2. **version** - Which version of the dependency chart should I use 
- Exmaple:
`version: "20.0.0"`

3. **repository** - where can I find the chart specified by `name` and `version`.
```
name
 ↓
Which Chart?

version
 ↓
Which version?

repository
 ↓
Where can Helm download/find it?
```
