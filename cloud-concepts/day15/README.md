## CLOUD CONCEPTS DAY 15

## REGION
- It is cluster of AZ
- An AZ can belong to a single region 
- AZ in the same regions can be connected

## EDGE LOCATIONS
- Edge locations are not Regions
- Edge locations are used for 
    - DNS
    - CDN
    - Caching 
Example:
- Suppose a request is coming from brazil to mumbai and this is slow and hence here we use Edge locations
```
Brazil User
      ↓
Nearest Edge Location
      ↓
Mumbai Region
```
- Edge locations are used to cache static content and help in giving faster response 
