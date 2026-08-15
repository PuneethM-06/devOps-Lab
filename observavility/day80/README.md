# DAY 80 - PromQL

### 1. rate()
- We know that counter always increases until the reset
- Seeing the raw number doesnt directly indicate the value of the current request or if the traffic is increasing 
- **rate is used to see the counter increased in x minutes of time and converts that into request per second average** 
- Example
```
10:00 → 1000 requests
10:05 → 1600 requests
Increase = 1600 - 1000 = 600
Time = 5 minutes = 300 seconds
So rate(http_requests_total[5m]) = 2 requests per second  
```
