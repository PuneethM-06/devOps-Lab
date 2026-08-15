# DAY 81 - GRAFANA - DASHBOARD FROM PROMETHEUS DATA 

1. ### WHAT IS GRAFANA 
- **Grafana is a visualization and observability platform that queries data sources such as Prometheus and represents the data through dashboard, panels, graphs, tables etc.**
- Grafana acts as a centralized visualization platform where we can connect multiple data sources, such as Prometheus, and represent the data using dashboards, graphs, gauges, tables, and other panels.

### GRANFANA ARCHITECTURE 
```
Grafana Dashboard
        │
        ▼
Grafana Panel
        │
        │ PromQL Query
        ▼
     Prometheus
        │
        │ Queries historical metric data
        ▼
       TSDB
        │
        │ Returns result
        ▼
     Prometheus
        │
        │ Sends query response
        ▼
      Grafana
        │
        ▼
Visualizes the result as a graph
```
> The application exposes metrics, Prometheus scrapes and stores them in its TSDB. Grafana uses Prometheus as a data source and sends PromQL queries to Prometheus. Prometheus queries the TSDB, returns the results to Grafana, and Grafana visualizes those results using dashboards and panels.

### DASHBOARD vs PANEL 
- **DASHBOARD** - A dashboard is a collection of multiple visualizations that give you an overall view of a system or application 
```
Application Dashboard
│
├── Request Rate
├── Error Rate
├── P95 Latency
├── CPU Usage
└── Memory Usage
```
> **DASHBOARD IS A COLLECTION OF PANELS**
