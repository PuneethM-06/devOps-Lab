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
2. ### IMAGE TAGS vs IMAGAE DIGESTS
- **Image Tags** are human readble name/version and they can be over written 
- **Image Digests** - They are cryptographic values that can be attached to an image, but they change if the image change too 

3. ### WHAT DOES IMAGE SIGNING ACTUALLY MEAN
```
Trusted CI builds image
        ↓
Image has a specific digest
        ↓
Cosign signs that artifact
        ↓
Signature is stored/associated with it
        ↓
Later, before deployment
        ↓
Verify signature
        ↓
Valid?
   ┌────┴────┐
   │         │
  Yes        No
   │         │
Trust       Reject 
```
