# Production-Grade Microservices CI/CD Pipeline on AWS EC2

An automated, high-performance Go microservice deployment architecture running on AWS EC2, featuring multi-stage minimal containerization and zero-downtime continuous deployment via GitHub Actions.

## 🚀 Live Production Endpoint
- **Public URL:** [http://54.252.201.203:8080](http://54.252.201.203:8080)
- **Status:** Real-Time Operational Telemetry & Host Metrics Dashboard

---

## 🛠️ Architecture & Tech Stack
- **Cloud Platform:** Amazon Web Services (AWS EC2 - Ubuntu 24.04 LTS)
- **Language/Runtime:** Golang 1.22
- **Containerization:** Docker Multi-Stage Scratch Image (< 15MB total footprint)
- **CI/CD Pipeline:** GitHub Actions (Automated build, test, push, and remote headless SSH rollout)
- **Registry:** Docker Hub Automated Ingestion

---

## 🔄 Automated Deployment Lifecycle
1. **Push Event:** Developer pushes code changes to the `main` branch.
2. **Build Stage:** GitHub Actions triggers multi-stage Docker build, compiling statically linked CGO-disabled Go binary.
3. **Registry Release:** Automated tag push to Docker Hub utilizing encrypted repository secrets.
4. **Remote Deployment:** Pipeline connects to production AWS EC2 via key-based SSH, terminates stale runtime containers, pulls latest image tag, and performs an instant container spin-up.

---

## 👤 Author
**Mohammad Faiz Ansari**  
Cloud & DevOps Engineer | [LinkedIn Profile](https://www.linkedin.com/in/mohammad-faiz-ansari-devops) • [GitHub](https://github.com/faiz7720)
