# DAY 90 - IMAGE SIGNING WITH COSIGN

1. ### WHAT PROBLEM DOES IMAGE SIGNING SOLVE 
- Signing an image is adding a cryptoraphic proof that is associated with the image 
```
Trusted CI builds image
        ↓
Specific image artifact
        ↓
Cosign signs it
        ↓
Later:
Verify signature
        ↓
Valid and from trusted identity?
   ┌────────┴────────┐
   │                 │
  Yes               No
   │                 │
Trust              Reject
```
