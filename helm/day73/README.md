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
4. **helm dependency update**
- Once we declare the depedency in the chart, helm still needs to fetch that dependency and hence do to that we use this command 
- `helm dependency update`
```
Chart.yaml
    ↓
Reads dependencies
    ↓
Finds Redis repository
    ↓
Gets requested Redis Chart version
    ↓
Downloads dependency
    ↓
Stores it under charts/
```
5. ### CHART LOCK
- When we run `helm dependency update` Helm resolves the dependency versions and can create a `chart.lock`
- Chart.lock records the **exact dependency versions** **that were resolved when you run helm dependency update, so the same dependency versions can be reproduced later.**