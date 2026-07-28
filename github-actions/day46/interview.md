# GitHub Actions - Interview Questions (Day 46)

## CI/CD Fundamentals

### 1. What is Continuous Integration (CI)?

### 2. What is Continuous Delivery (CD)?

### 3. What is the difference between Continuous Delivery and Continuous Deployment?

### 4. Explain the CI/CD lifecycle.

### 5. Why do companies use CI/CD?

### 6. Why is CI considered a quality gate?

### 7. At what point does CI end and CD begin?

### 8. Why should CI run before code is merged into the main branch?

### 9. Explain CI/CD using a real-world example.

### 10. Explain CI/CD using a Git workflow.

---

# GitHub Actions

### 11. What is GitHub Actions?

### 12. What is a Workflow?

### 13. What is an Event in GitHub Actions?

### 14. What is a Job?

### 15. What is a Step?

### 16. What is a Runner?

### 17. What is the difference between a Workflow and a Job?

### 18. What is the difference between a Job and a Step?

### 19. What happens internally when a workflow starts?

### 20. Why is `runs-on` mandatory?

### 21. What is the difference between GitHub-hosted and Self-hosted runners?

---

# Workflow Keywords

### 22. What is the purpose of the `name` keyword?

### 23. What is the purpose of the `on` keyword?

### 24. What are the most commonly used GitHub Actions events?

### 25. What is the difference between `push` and `pull_request` events?

### 26. Why do most companies configure both `push` and `pull_request` triggers?

### 27. What are branch filters?

### 28. Why wouldn't you run every workflow on every branch?

---

# Jobs & Steps

### 29. Why do workflows have multiple jobs?

### 30. Can jobs run in parallel?

### 31. Why is separating jobs considered a good practice?

### 32. Why should every step have a meaningful name?

---

# uses vs run

### 33. What is the difference between `uses` and `run`?

### 34. When would you use `uses`?

### 35. When would you use `run`?

### 36. Why is `actions/checkout` implemented using `uses` instead of `run`?

### 37. What is the purpose of the `with` keyword?

---

# Checkout Action

### 38. Why is `actions/checkout` usually the first step in most workflows?

### 39. What happens if you don't checkout the repository?

### 40. Does a GitHub-hosted runner automatically contain your repository?

### 41. Explain what `actions/checkout@v4` does internally.

---

# Execution Flow

### 42. Explain the complete lifecycle from `git push` until a workflow completes.

### 43. What happens if one step fails?

### 44. Why does GitHub Actions fail fast by default?

### 45. Explain how GitHub Actions executes shell commands.

### 46. Does GitHub Actions execute commands, or does the runner execute them?

### 47. Where are shell commands executed?

---

# Design & Best Practices

### 48. Why should workflows have a single responsibility?

### 49. Why do companies maintain multiple workflows instead of one large workflow?

### 50. How would you structure CI workflows for a large production repository?