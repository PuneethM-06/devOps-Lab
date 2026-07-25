# devops-lab

Hands-on DevOps learning repo — 12 months to Cloud/DevOps roles 
Every concept practised, documented, and committed. No tutorials without code.

**Target:** Cloud / DevOps / Platform Engineer roles 
**Timeline:** 12 months  
**Method:** One evolving project + daily commits + structured notes

---

## How this repo is structured

Every topic has its own folder. Inside each folder:
- `dayX.md` files — what I learned, commands I ran, what confused me, interview angle
- `scripts/` — actual bash/python scripts written while learning
- `README.md` — index of that topic and key takeaways

No copy-pasted theory. Everything here was typed and run by me.

---

## Progress tracker

### Phase 1 — Foundations (Months 1–3)

#### Linux & Bash
- [x] Day 1 — Navigation & filesystem
- [x] Day 2 — File permissions & ownership
- [x] Day 3 — Processes & signals
- [x] Day 4 — Logs & troubleshooting
- [x] Day 5 — Text processing (grep, awk, sed, cut)
- [x] Day 6 — Networking commands (curl, dig, ss, netstat)
- [x] Day 7 — Bash scripting (scripts that actually do something)

#### Networking
- [x] Day 8 — IP addresses, subnets, CIDR notation
- [x] Day 9 — DNS — how a domain becomes an IP
- [x] Day 10 — TCP vs UDP, 3-way handshake, ports
- [x] Day 11 — HTTP vs HTTPS, TLS handshake, status codes
- [x] Day 12 — Firewalls, security groups, stateful vs stateless

#### Cloud Concepts (Vendor-neutral)
- [x] Day 13 — What is cloud, on-prem vs cloud, trade-offs
- [x] Day 14 — IaaS vs PaaS vs SaaS with real examples
- [x] Day 15 — Regions, availability zones, edge locations
- [x] Day 16 — Scalability, high availability, fault tolerance
- [x] Day 17 — Shared responsibility model

#### Docker
- [x] Day 18 — What containers are and why (vs VMs)
- [x] Day 19 — Writing a Dockerfile, image layers, caching
- [x] Day 20 — Multi-stage builds, image size optimisation
- [x] Day 21 — docker-compose — multi-container apps
- [x] Day 22 — Networking between containers
- [x] Day 23 — Volumes — persistent data
- [x] Day 24 — Security — non-root user, Trivy scanning
- [x] Day 25 — Push to GitHub Container Registry (GHCR)

#### Git (professional level)
- [x] Day 26 — Branching strategy, branch protection on main
- [x] Day 27 — Conventional commits, commit discipline
- [x] Day 28 — Rebase vs merge, interactive rebase, squashing
- [x] Day 29 — Pre-commit hooks — lint before every commit
- [x] Day 30 — git bisect, git blame, git stash

#### Python & Bash (automation focus)
- [ ] Day 31 — Variables, loops, functions, error handling
- [ ] Day 32 — File handling, reading/writing, argparse
- [ ] Day 33 — HTTP requests with httpx, parse JSON responses
- [ ] Day 34 — boto3 — interact with AWS from Python
- [ ] Day 35 — CLI tools with click or typer

#### AWS Deep Dive
- [x] Day 36 — EC2 — launch, SSH, security groups, user data
- [x] Day 37 — S3 — buckets, policies, versioning, presigned URLs
- [x] Day 38 — IAM — users, roles, policies, least privilege
- [x] Day 39 — VPC — subnets, route tables, IGW, NAT gateway
- [x] Day 40 — RDS — managed Postgres, Multi-AZ, backups
- [x] Day 41 — ECS / Fargate — deploy Docker container on AWS
- [x] Day 42 — ALB — load balancer, target groups, health checks
- [x] Day 43 — CloudWatch — metrics, logs, alarms, dashboards
- [x] Day 44 — Route53 — DNS, hosted zones, routing policies
- [x] Day 45 — SAA exam prep — practice tests, weak areas

**Phase 1 milestone:** Dockerised Python app running on AWS EC2, pushed via GitHub, CI pipeline green

---

### Phase 2 — Core DevOps + GitOps (Months 4–6)

#### GitHub Actions (CI/CD)
- [x] Day 46 — First workflow — checkout, test, build on push
- [x] Day 47 — Docker build and push to GHCR in pipeline
- [x] Day 48 — Caching pip deps, measuring speed improvement
- [x] Day 49 — Manual approval gates for production deploys
- [x] Day 50 — Matrix builds, reusable workflows, artifacts
- [x] Day 51 — Secrets and environment variables in pipelines

#### Terraform (Infrastructure as Code)
- [ ] Day 52 — First .tf file — provider, resource, plan, apply
- [ ] Day 53 — Variables, outputs, terraform.tfvars
- [ ] Day 54 — S3 backend + DynamoDB state locking
- [ ] Day 55 — Modules — networking and compute modules
- [ ] Day 56 — Terraform in CI — plan on PR, apply on merge
- [ ] Day 57 — Workspaces — dev and staging environments
- [ ] Day 58 — Import existing resources, lifecycle rules

#### Kubernetes
- [ ] Day 59 — Pods, Deployments, ReplicaSets — the why
- [ ] Day 60 — Services — ClusterIP, NodePort, LoadBalancer
- [ ] Day 61 — ConfigMap and Secrets — mount as env or volume
- [ ] Day 62 — Liveness and readiness probes
- [ ] Day 63 — Rolling updates and rollbacks
- [ ] Day 64 — Resource requests and limits, OOMKilled
- [ ] Day 65 — HorizontalPodAutoscaler — scale on CPU
- [ ] Day 66 — Namespaces, RBAC basics
- [ ] Day 67 — Debugging — CrashLoopBackOff, Pending, OOMKilled

#### Helm
- [ ] Day 68 — Chart structure — Chart.yaml, values.yaml, templates
- [ ] Day 69 — values-dev.yaml vs values-prod.yaml
- [ ] Day 70 — Hooks — run DB migration before app deploys
- [ ] Day 71 — Upgrade, rollback, diff

#### ArgoCD / GitOps
- [ ] Day 72 — What GitOps is and why it's different from CI/CD
- [ ] Day 73 — Install ArgoCD, connect to Git repo
- [ ] Day 74 — Auto-sync — push to Git, watch it deploy
- [ ] Day 75 — Self-heal — manual change gets reverted
- [ ] Day 76 — App of Apps pattern for multiple services

#### Observability
- [ ] Day 77 — Prometheus — instrument Flask app with metrics
- [ ] Day 78 — PromQL — rate, histogram_quantile, alerts
- [ ] Day 79 — Grafana — dashboards from Prometheus data
- [ ] Day 80 — OpenTelemetry — trace a request across services
- [ ] Day 81 — Grafana Loki — structured logs, LogQL queries
- [ ] Day 82 — Alertmanager — fire Slack alert on error spike

**Phase 2 milestone:** Green CI pipeline on every push, ArgoCD deploying from Git, Grafana dashboard live

---

### Phase 3 — DevSecOps + Projects (Months 7–9)

#### Container & Supply Chain Security
- [ ] Day 83 — Trivy in GitHub Actions — block CVEs from merging
- [ ] Day 84 — Snyk — scan Python deps, auto-fix PRs
- [ ] Day 85 — SBOM — generate with Trivy, understand why it matters
- [ ] Day 86 — gitleaks — pre-commit hook catching leaked secrets
- [ ] Day 87 — Distroless images — no shell, smaller attack surface
- [ ] Day 88 — Image signing with cosign

#### Secrets Management
- [ ] Day 89 — AWS Secrets Manager — store, fetch at runtime
- [ ] Day 90 — IAM roles everywhere — no static credentials
- [ ] Day 91 — External Secrets Operator — sync AWS secrets to K8s
- [ ] Day 92 — HashiCorp Vault — dynamic secrets demo

#### Policy as Code
- [ ] Day 93 — OPA / Conftest — write Terraform policies
- [ ] Day 94 — Enforce in pipeline — block non-compliant infra

#### Project 1 — cloud-platform-kit
- [ ] Week 1 — Scaffold 3 microservices, Dockerfile each
- [ ] Week 2 — GitHub Actions pipeline for all 3
- [ ] Week 3 — Terraform infra — VPC, ECS, ALB, RDS
- [ ] Week 4 — ArgoCD GitOps deployment
- [ ] Week 5 — OTel traces across all 3 services
- [ ] Week 6 — Grafana dashboard, Alertmanager rules

#### Project 2 — iac-security-scanner
- [ ] Week 1 — Python CLI that parses Terraform files
- [ ] Week 2 — Detect misconfigs — public S3, open port 22
- [ ] Week 3 — LLM integration — Claude API generates fixes
- [ ] Week 4 — Publish as a GitHub Action

**Phase 3 milestone:** Both projects running end-to-end, Trivy gate in CI, applying to jobs

---

### Phase 4 — Interview Prep + Apply (Months 10–12)

#### Interview preparation
- [ ] Walk through CI/CD pipeline end to end out loud
- [ ] Write a Dockerfile from memory in under 5 minutes
- [ ] Debug CrashLoopBackOff live — logs, describe, exec
- [ ] Write bash log parser on a whiteboard
- [ ] Answer "what happens when you type a URL" — full chain

#### System design
- [ ] ALB → K8s → RDS → Cache — draw and explain
- [ ] Blue/green vs canary deploy — tradeoffs
- [ ] Multi-region setup with GDPR considerations
- [ ] Incident response — something broke, what do you do

#### Applications
- [ ] LinkedIn rewritten in English, Open to Work EU turned on
- [ ] Resume updated — TDP story framed as a journey
- [ ] Ireland applications — Critical Skills Employment Permit
- [ ] Netherlands applications — Kennismigrant visa
- [ ] Germany applications — EU Blue Card
- [ ] 30 applications per month target

**Phase 4 milestone:** First offer. Move.

---

## The two main projects

### cloud-platform-kit
3-microservice developer platform. API gateway + vulnerability scanner + notification service.
Full GitOps with ArgoCD, infrastructure in Terraform, distributed tracing with OpenTelemetry.
This is the project that shows I can build and operate a platform, not just run a single app.

### iac-security-scanner
Python CLI that scans Terraform configs for misconfigurations and uses an LLM to generate
human-readable fixes with corrected Terraform code. Published as a GitHub Action.
This is the project that ties my Wiz experience to open-source DevSecOps tooling.

---

## Environment
- **Code editor:** GitHub Codespaces (browser-based, no local install needed)
- **Cloud:** AWS Free Tier — personal account, zero company data
- **Container registry:** GitHub Container Registry (GHCR)
- **Certifications target:** AWS Solutions Architect Associate (Month 3)

---

- *Started: June 2026*
- *Every commit is a step. Start before you feel ready.*