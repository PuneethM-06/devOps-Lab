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
**WHAT DOES SIGNING AND VERIFICATION MEAN?**
- **Signing** - I, the trusted signer, approved/signed this specific artifact
- **Verification** - Can I cryptographically match what was signed by the trusted signer is the same??

4. ### PUBLIC AND PRIVATE KEY MODEL
- The traditional image signing is a asymmetric cryptography 
```
Private key
    ↓
Signs the image

Public key
    ↓
Verifies the signature
```

**SIGNING**
- A trusted entity such as CI/CD pipeline has access to private key 
```
CI/CD
  ↓
Build image
  ↓
my-app@sha256:111
  ↓
Private key
  ↓
Cosign signs the image
```

**VERIFICATION**
- The verifier uses the corresponding public key 
```
Image
   +
Signature
   +
Public key
        ↓
Cosign verify
        ↓
Valid?
   ┌────┴────┐
   │         │
 Yes         No
   │         │
Trust      Reject
```
-  The public key is safe to distribute unlike private key because it is used for verification 

5. ### KEYLESS SIGNING 
- We know private key signs the image and it is important to keep it safe.
- Instead of manually managing a long lived key, 
```
CI/CD pipeline
      ↓
Proves its identity
      ↓
Cosign performs keyless signing
      ↓
Signature is associated with that identity
      ↓
Later verification
      ↓
Verify artifact + trusted identity
```
6. ### TAMPERING AND TRUST SCENARIOES
**SCENARIO 1**
```
Trusted CI
    ↓
Builds image
    ↓
sha256:111
    ↓
Signs image 
    ↓
Deployment verifies signature
    ↓
Valid + trusted identity
    ↓
Deploy 
```
- The image passes verification because the artifact matches what was signed 
