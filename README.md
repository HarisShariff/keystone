# Keystone

<p align="center">
  <a href="https://github.com/HarisShariff/keystone/actions/workflows/build.yml">
    <img alt="Build Status" src="https://github.com/HarisShariff/keystone/actions/workflows/build.yml/badge.svg">
  </a>
  <a href="https://goreportcard.com/report/github.com/HarisShariff/keystone">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/HarisShariff/keystone">
  </a>
  <a href="https://coveralls.io/github/HarisShariff/keystone">
    <img alt="Coverage Status" src="https://coveralls.io/repos/github/HarisShariff/keystone/badge.svg?branch=main">
  </a>
  <a href="https://github.com/HarisShariff/keystone/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/HarisShariff/keystone.svg">
  </a>
  <a href="https://github.com/HarisShariff/keystone/stargazers">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/HarisShariff/keystone.svg">
  </a>
  <a href="https://github.com/HarisShariff/keystone/network/members">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/HarisShariff/keystone.svg">
  </a>
</p>

Small. Strong. Foundational.

Keystone is a lightweight, open-source Identity and Access Management (IAM) system for Go.
It provides a clean, embeddable alternative to bulky solutions like Keycloak, perfect for internal systems such as hospitals, manufacturing setups, and SMEs.

Keystone supports:

- HTTP server mode (expose IAM APIs).
- Go library mode (embed directly into your services).

Built with clean architecture principles, secure token handling, and modular design.

## Features

- Dual Operation Modes: HTTP server or Go library.
- Clean Architecture: Transport and storage separated via interfaces.
- Secure Authentication: Password hashing (bcrypt), JWT access/identity tokens.
- Role-Based Access Control (RBAC): Simple and extensible role support.
- Pluggable Storage: Oracle database backend (PostgreSQL and others in future).
- Dockerized Deployment: Lightweight container image.
- Minimal External Dependencies: Standard library preferred.

## Why Keystone?

Keystone is designed for developers and teams who need lightweight, secure, and embeddable Identity and Access Management (IAM) — without the heavy baggage of traditional solutions like Keycloak or Auth0.

Our focus is simple:

- **Lightweight and embeddable**: Keystone can run as a standalone HTTP server or be embedded directly as a Go library inside your services.
- **Security-first foundation**: We provide secure password hashing, JWT token issuance, and role-based access control without external dependencies.
- **Designed for SMEs and Internal Systems**: Perfect for back-office medical groups, manufacturing units, and SMEs where simplicity, speed, and security matter more than complex external SSO integrations.
- **Extendable for the future**: While our core remains lightweight, Keystone is designed to evolve — future modules like OAuth2, SSO, and MFA can be added without sacrificing the clean, modular architecture.

Keystone exists for those who need an identity system that is **small, strong, and foundational** — not one that demands you bring a battleship to a fishing trip.

## Getting Started

Run as HTTP Server

```bash
docker pull ghcr.io/HarisShariff/keystone:latest
docker run -p 8080:8080 ghcr.io/HarisShariff/keystone:latest
```

Use as Go Library

```bash
go get github.com/HarisShariff/keystone
```

```go
import "github.com/HarisShariff/keystone"

func main() {
    ks := keystone.New()
    userID, err := ks.RegisterUser("username", "password")
    // handle err, use userID
}
```

## Project Structure

```
/cmd           # Entry points (server, tools)
/internal
  /core        # Business logic (auth, user, token, role)
  /infrastructure
    /db        # Storage adapters
    /crypto    # Password hashing
    /jwt       # JWT utilities
  /transport
    /http      # HTTP handlers
  /config      # Configuration loading
/pkg
  /keystone    # Public Go library API
/scripts
Dockerfile
```

## Contributing

Contributing

We welcome contributions!

Please follow these steps: 1. Check Existing Issues: See if your feature, task, or bug is already tracked. 2. Open the Right Issue Type:

- 🛠️ Feature – New capabilities or improvements.
- 🧹 Task – Housekeeping, setup, refactoring.
- 🐞 Bug – Something broken that needs fixing. 3. Fork the Repository and create a new branch. 4. Submit a Pull Request once ready — link it to the issue!

👉 We have dedicated templates for Features, Tasks, and Bugs to make it easy.

## License

This project is licensed under the MIT License.

## About

Author [Haris Shariff](https://github.com/HarisShariff).
Created for developers and teams who want a strong, simple, and reliable IAM foundation without the heavy enterprise overhead.

For contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).
For community behavior expectations, see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
