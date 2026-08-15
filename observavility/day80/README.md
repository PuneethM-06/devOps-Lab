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
- having `rate(http_requests_total)` has no time range and hence prometheus has no historical time range for calculating 
> **NOTE**: `rate()`is generally made for counters fo cummulative increase in in value

- ### FOR GAUGE
- Gauge can go in both the ways (increasing and decreasing)
- And hence for gauge - usually query the curent value directly because it can both increase and decrase 
```
Counter + rate() → how fast the cumulative value is changing/increasing per second.

Gauge alone → current value.

Gauge + rate() → rate of change, which is technically possible but usually not the metric you want.
```

### 2. irate vc rate
- **rate** - Calulcalates the increase/difference in the requests in that given time range and calculates the requests per second
- **irate** - Calculates the increase/difference in the requests in the last 2 recent values and calculates the requests per second
```
10:00 → 100
10:01 → 200
10:02 → 300
10:03 → 400
10:04 → 500
10:05 → 700
irate(...[5m]); so 
10:04 → 500
10:05 → 700
700 - 500 = 200 requests 
200/60 = 3.33 requests per second 
```
- #### PRACTICAL USAGE
- `rate()` → dashboards, recording rules, alerts; generally the safer default.
- `irate()` → when you specifically want to inspect short-term, rapidly changing behavior.

3. ### sum
- `rate()` - calculates the per-second rate for each matching time series and sum adds those together 
- Example:
```
Instance A → 2 requests/sec
Instance B → 3 requests/sec
Instance C → 5 requests/sec
sum(rate(http_requests_total[5m])) = 2 + 3 + 5 = 10 requests/sec
```
