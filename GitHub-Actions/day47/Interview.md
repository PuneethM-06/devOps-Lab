# Day 47 & Day 48 Interview Questions

## GitHub Actions Fundamentals

### Q1. Explain the architecture of GitHub Actions.

### Q2. What is the difference between a Workflow, Job and Step?

### Q3. What is the difference between `uses` and `run`?

### Q4. Why is `actions/checkout` required?

### Q5. What happens if `actions/checkout` is not used?

### Q6. Explain the lifecycle of a GitHub Actions workflow.

### Q7. What is a Runner?

### Q8. Does every step run on a new runner?

### Q9. Do multiple jobs share the same filesystem?

### Q10. Explain the difference between Step Logs and Workflow Logs.

### Q11. How do you debug a failed GitHub Actions workflow?

### Q12. Why should you start reading logs from the first failed step?

### Q13. Explain Linux exit codes.

### Q14. Why does a GitHub Actions workflow stop after a failed step?

### Q15. How would you debug a "No such file or directory" error inside a workflow?

### Q16. Why should language runtimes be explicitly configured instead of relying on the runner defaults?

---

## Docker + GitHub Actions

### Q17. Why should Docker images be built in CI instead of on a developer's machine?

### Q18. Explain Docker image tagging.

### Q19. What is Semantic Versioning (SemVer)?

### Q20. Why is `latest` not recommended for production deployments?

### Q21. Why are Git commit SHA tags useful?

### Q22. Can one Docker image have multiple tags? Explain.

### Q23. Explain the difference between an Image ID and an Image Tag.

### Q24. What is GitHub Container Registry (GHCR)?

### Q25. Why would a company choose GHCR over Docker Hub?

### Q26. How does GitHub Actions authenticate with GHCR?

### Q27. Where are credentials stored in GitHub Actions?

### Q28. Why should secrets never be hardcoded in workflows?

### Q29. What does `docker/metadata-action` do?

### Q30. Does `docker/metadata-action` decide semantic versions?

### Q31. How does `docker/metadata-action` know about commit SHAs and Git tags?

### Q32. Explain the complete Docker CI pipeline from push to image upload.

### Q33. How would you design a production Docker workflow for linux-sysmonitor?

### Q34. Explain the difference between:
- latest
- Semantic Version
- Commit SHA

### Q35. Walk me through what happens after you push code to GitHub until the Docker image appears in GHCR.