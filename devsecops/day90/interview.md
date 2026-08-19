# interview.md — Day 90: Image Signing with Cosign

## Must-Know Fundamentals

### 1. What problem does image signing solve?

**Expected answer:**

Vulnerability scanning alone does not prove that an image came from a trusted source or that the specific artifact being deployed is the expected one.

Image signing provides cryptographic proof that can be used to verify:

- The identity of the trusted signer
- The integrity of the signed artifact

---

### 2. Why is vulnerability scanning alone not enough?

**Expected answer:**

A vulnerability scanner such as Trivy can identify known vulnerabilities, but it does not prove:

- Who built the image
- Whether the image came from a trusted CI/CD pipeline
- Whether the artifact being deployed is the approved artifact

An image can have zero critical CVEs and still be untrusted.

---

### 3. What is the difference between an image tag and an image digest?

**Expected answer:**

An image tag is a mutable label.

```text
my-app:v1
```

The same tag can later point to a different image.

An image digest identifies specific image content.

```text
my-app@sha256:abc123...
```

If the image content changes, the digest changes.

---

### 4. Why is signing an image by its digest more reliable than trusting a tag?

**Expected answer:**

A tag can be overwritten or moved to point to a different image.

A digest is tied to specific image content.

```text
Image A
↓
sha256:111

Image content changes
↓
Image B
↓
sha256:222
```

Signing the digest allows us to verify the specific artifact that was signed rather than relying on a mutable tag.

---

### 5. What does image signing actually mean?

**Expected answer:**

Image signing creates cryptographic proof associated with a specific artifact.

Conceptually:

```text
Trusted CI
    ↓
Build image
    ↓
Specific image digest
    ↓
Cosign signs the artifact
    ↓
Later verification
    ↓
Valid trusted signature?
    ↓
Allow / Reject
```

---

### 6. What does image signing help verify?

**Expected answer:**

Image signing primarily helps verify:

- **Authenticity** — Was the artifact signed by a trusted identity or key?
- **Integrity** — Does the artifact being verified match the artifact that was signed?

It does not prove that the image is vulnerability-free.

---

### 7. Why should a signature for one digest not validate another digest?

**Expected answer:**

The signature is associated with a specific artifact.

For example:

```text
Signed artifact
sha256:111
```

If the image content changes:

```text
New artifact
sha256:222
```

The new artifact is different, so the original signature should not validate it.

---

### 8. What is the private/public key model for image signing?

**Expected answer:**

In traditional asymmetric signing:

```text
Private key
    ↓
Signs the image

Public key
    ↓
Verifies the signature
```

The private key creates the signature, while the corresponding public key is used to verify it.

---

### 9. Why must the private signing key be protected?

**Expected answer:**

If the private key is compromised, an attacker could sign a malicious or altered image and make it appear as though it came from the trusted signer.

Therefore, the private key must be securely stored, protected, and rotated if compromised.

---

### 10. Why can the public key be shared?

**Expected answer:**

The public key is used for verification only.

Sharing it allows systems to verify signatures, but possession of the public key does not allow someone to create a valid signature.

```text
Private key → Sign
Public key  → Verify
```

---

### 11. What is keyless signing?

**Expected answer:**

Keyless signing allows artifacts to be signed without the user manually managing a long-lived private signing key in the traditional way.

Trust can instead be associated with an authenticated identity, which makes it particularly useful for CI/CD systems.

---

### 12. Why is keyless signing useful in CI/CD?

**Expected answer:**

It reduces the operational burden of manually managing a long-lived private signing key.

Traditional signing requires managing:

- Key storage
- Secret protection
- CI/CD access
- Key rotation

With keyless signing, the CI/CD identity can be used to establish trust during the signing process.

---

### 13. What happens if a signed image is modified?

**Expected answer:**

If the image content changes, its digest changes.

The original signature is tied to the original artifact, so it should not validate the modified artifact.

```text
Original image
sha256:111
↓
Signed

Modified image
sha256:222
↓
Original signature fails verification
```

---

### 14. What should happen to an unsigned image if policy requires trusted signatures?

**Expected answer:**

The image should be rejected.

Even if the image passes vulnerability scanning, the absence of a valid signature from the trusted identity means its origin and integrity cannot be verified according to the configured trust policy.

---

### 15. Is a signed image automatically secure?

**Expected answer:**

No.

A valid signature does not mean the image has:

- No vulnerabilities
- No insecure application code
- No misconfigurations
- No malicious behavior

Signing solves a different problem: artifact authenticity and integrity.

---

### 16. Where can image verification happen?

**Expected answer:**

Verification can happen at different points, including:

- CI/CD pipelines
- Deployment stages
- Kubernetes admission or policy layers

The goal is to verify the artifact before allowing it to run.

---

### 17. Why is signing alone not enough?

**Expected answer:**

Signing creates cryptographic proof, but that proof has no practical enforcement value if nobody verifies it.

The complete model is:

```text
Sign
  ↓
Verify
  ↓
Check trusted identity
  ↓
Policy enforcement
  ↓
Allow / Reject
```

Without verification and enforcement, an unsigned or untrusted image could still be deployed.

---

## Scenario-Based Questions

### 18. Trivy reports zero critical CVEs. Is that enough to trust and deploy the image?

**Expected answer:**

No.

Trivy checks for known vulnerabilities, but it does not prove the image came from a trusted source or that the deployed artifact is the expected one.

Image signing and verification can provide additional trust through authenticity and integrity checks.

---

### 19. Why is a digest more reliable than a tag for signing?

**Expected answer:**

A tag can be overwritten or moved to point to another image.

A digest identifies specific image content, so changing the image produces a different digest.

Therefore, signing a digest ties the signature to a specific artifact.

---

### 20. A trusted CI signs `my-app@sha256:111`. Later the image becomes `sha256:222`. Why should verification fail?

**Expected answer:**

The signature was created for the specific artifact identified by `sha256:111`.

Since `sha256:222` represents different image content, it is a different artifact and should not validate against the original signature.

---

### 21. What happens if a private signing key is compromised?

**Expected answer:**

An attacker could potentially create valid signatures for malicious or altered images, making them appear to come from the trusted signer.

The compromised key should therefore be revoked or rotated, and trust policies may need to be updated.

---

### 22. Your CI signs an image, but the deployment process never verifies the signature. What is the problem?

**Expected answer:**

Signing alone does not enforce trust.

Without verification and policy enforcement, the deployment process could still deploy an unsigned, tampered, or untrusted image.

The signature must be verified before deployment, and the result must be enforced through an allow/reject decision.

---

## Final Security Pipeline Mental Model

```text
Developer
    ↓
Gitleaks
    ↓
Check for leaked secrets
    ↓
Build application/image
    ↓
Snyk
    ↓
Check dependencies
    ↓
Trivy
    ↓
Scan image for known vulnerabilities
    +
Generate/inspect SBOM
    ↓
Cosign
    ↓
Sign specific artifact
    ↓
Push to registry
    ↓
Deployment
    ↓
Verify signature
    ↓
Check trusted identity
    ↓
Policy enforcement
    ↓
Allow / Reject
```

## Tool Comparison

| Tool / Concept | Main Question |
|---|---|
| **Gitleaks** | Are secrets leaked? |
| **Snyk** | Are dependencies vulnerable? |
| **SBOM** | What components and versions are inside? |
| **Trivy** | Does the image/artifact contain known vulnerabilities? |
| **Distroless** | Can unnecessary runtime components be removed? |
| **Cosign** | Can we verify the artifact's integrity and trusted signer? |

## Final Interview Answer

> **Cosign is used for software supply chain security by signing container images and other artifacts. Vulnerability scanning alone cannot prove who produced an image or whether the specific artifact being deployed is the expected one. Cosign can sign a specific artifact, typically identified by its digest, and later systems can verify the signature against a trusted key or identity. This helps verify artifact authenticity and integrity. However, signing alone is not enough—the signature must be verified and enforced before deployment. A signed image is also not automatically secure, so vulnerability scanning, secret detection, dependency management, and other security controls are still required.**