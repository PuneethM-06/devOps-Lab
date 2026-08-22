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
- [x] Day 52 — Terraform Fundamentals — provider, resource, init, plan, apply, destroy
- [x] Day 53 — Variables, locals, outputs, expressions, terraform.tfvars
- [x] Day 54 — Terraform State — state file, remote backend (S3), DynamoDB locking
- [x] Day 55 — Data Sources & Modules — reusable networking, compute and storage modules
- [x] Day 56 — Advanced Terraform — count, for_each, dynamic blocks, lifecycle, depends_on
- [x] Day 57 — Multi-Environment Infrastructure — workspaces, dev/staging, backend configuration
- [x] Day 58 — Terraform in CI/CD — fmt, validate, plan on PR, apply on merge, GitHub Actions
- [x] Day 59 — Import & Production Practices — import existing resources, drift detection, lifecycle rules, state commands

#### Kubernetes
- [x] Day 59 — Pods, Deployments, ReplicaSets — the why
- [x] Day 60 — Services — ClusterIP, NodePort, LoadBalancer
- [x] Day 61 — ConfigMap and Secrets — mount as env or volume
- [x] Day 62 — Liveness and readiness probes
- [x] Day 63 — Rolling updates and rollbacks
- [x] Day 64 — Resource requests and limits, OOMKilled
- [x] Day 65 — HorizontalPodAutoscaler — scale on CPU
- [x] Day 66 — Namespaces, RBAC basics
- [x] Day 67 — Debugging — CrashLoopBackOff, Pending, OOMKilled

#### Helm
- [x] Day 68 → Chart Structure
- [x] Day 69 → Values & Environment Configuration
- [x] Day 70 → Hooks
- [x] Day 71 → Release Lifecycle & Core Commands
- [x] Day 72 → Upgrade, Revisions, Rollback & Diff
- [x] Day 73 → Dependencies, Repositories & Packaging
- [x] Day 74 → Advanced Helm & Production Patterns

#### ArgoCD / GitOps
- [x] Day 75 — What GitOps is and why it's different from CI/CD
- [x] Day 76 — Install ArgoCD, connect to Git repo
- [x] Day 77 — Auto-sync — push to Git, watch it deploy
- [x] Day 78 — App of Apps pattern for multiple services

#### Observability
- [x] Day 79 — Prometheus — instrument Flask app with metrics
- [x] Day 80 — PromQL — rate, histogram_quantile, alerts
- [x] Day 81 — Grafana — dashboards from Prometheus data
- [x] Day 82 — OpenTelemetry — trace a request across services
- [x] Day 83 — Grafana Loki — structured logs, LogQL queries
- [x] Day 84 — Alertmanager — fire Slack alert on error spike

**Phase 2 milestone:** Green CI pipeline on every push, ArgoCD deploying from Git, Grafana dashboard live

---

### Phase 3 — DevSecOps + Projects (Months 7–9)

#### Container & Supply Chain Security
- [x] Day 85 — Trivy in GitHub Actions — block CVEs from merging
- [x] Day 86 — Snyk — scan Python deps, auto-fix PRs
- [x] Day 87 — SBOM — generate with Trivy, understand why it matters
- [x] Day 88 — gitleaks — pre-commit hook catching leaked secrets
- [x] Day 89 — Distroless images — no shell, smaller attack surface
- [x] Day 90 — Image signing with cosign

#### Secrets Management
- [x] Day 91 — AWS Secrets Manager — store, fetch at runtime
- [x] Day 92 — IAM roles everywhere — no static credentials
- [x] Day 93 — External Secrets Operator — sync AWS secrets to K8s
- [x] Day 94 — HashiCorp Vault — dynamic secrets demo

#### Policy as Code
- [ ] Day 95 — OPA / Conftest — write Terraform policies
- [ ] Day 96 — Enforce in pipeline — block non-compliant infra

#### GoLang (Cloud Platform Kit Focus)
- [ ] Day 97 — Variables, data types, loops, functions, structs, interfaces, error handling
- [ ] Day 98 — File handling, JSON, environment variables, packages, modules (`go mod`)
- [ ] Day 99 — HTTP client/server (`net/http`), REST APIs, routing, middleware
- [ ] Day 100 — AWS SDK for Go (v2), IAM authentication, S3, STS, CloudWatch basics
- [ ] Day 101 — CLI tools (Cobra), logging, configuration management, project structure

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
