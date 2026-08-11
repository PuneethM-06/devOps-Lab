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
