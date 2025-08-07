# Contributing to Talken

Welcome, and thank you for your interest in contributing to **Talken** – a peer-to-peer, offline-first messaging core designed to work without internet, servers, or even infrastructure. This guide provides all the information you need to contribute effectively.

Talken is **not a chat app**. It is a **core messaging engine** that client applications (mobile, desktop, CLI, etc.) can embed and use. It supports **local communication**, **Bluetooth relays**, **store-and-forward messaging**, and **multi-hop routing**, all in a decentralized, privacy-respecting manner.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Talken Works (Short Recap)](#how-talken-works-short-recap)
- [Development Philosophy](#development-philosophy)
- [How to Contribute](#how-to-contribute)
  - [Issues](#issues)
  - [Feature Suggestions](#feature-suggestions)
  - [Pull Requests](#pull-requests)
- [Semantic Branching](#semantic-branching)
- [Semantic Commit Messages](#semantic-commit-messages)
- [Project Structure](#project-structure)
- [Environment Setup](#environment-setup)
- [Code Style](#code-style)
- [Testing](#testing)
- [Security](#security)
- [License](#license)

---

## Code of Conduct

This project follows the [Contributor Covenant](https://www.contributor-covenant.org/) Code of Conduct. Please be respectful, inclusive, and supportive in your interactions.

---

## How Talken Works (Short Recap)

- Every device generates a **public/private key pair**.
- Devices discover nearby peers via **Bluetooth**, **LAN**, or **Hotspot**.
- Messages are encrypted using the **recipient's public key**.
- If a direct connection isn't available, the message can **hop through other devices** until it reaches its destination.
- Offline devices can receive messages later via **store-and-forward**.
- Talken does not rely on IPs or ports — just public keys and peer discovery.

---

## Development Philosophy

- **Offline-first:** Works without internet.
- **Decentralized:** No central server.
- **Resilient:** Messages survive connectivity loss.
- **Minimal:** No bloat. Lightweight core with pluggable extensions.
- **Universal:** Can run on Linux, Android, CLI, etc.

---

## How to Contribute

### Issues

- Use [GitHub Issues](https://github.com/m-mdy-m/talken/issues) to:
  - Report bugs
  - Request features
  - Discuss design questions

Use labels like `bug`, `enhancement`, `question`, or `design`.

### Feature Suggestions

Before implementing new features, please open an issue to **discuss** and align with the core design principles. Talken is still in early stages and needs a **coherent direction**.

### Pull Requests

Pull requests (PRs) are welcome, **after the core is stabilized**. Until then, PRs should:
- Be scoped to **one logical change**.
- Include documentation and tests where applicable.
- Respect the [branching](#semantic-branching) and [commit](#semantic-commit-messages) conventions.

---

## Semantic Branching

Please use **semantic branch names** that indicate the purpose and type of change:

```

<type>/<short-description>

```

Examples:
- `feat/multi-hop-routing`
- `fix/discovery-crash`
- `docs/architecture-notes`
- `refactor/peer-db`

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `test`: Test-related
- `refactor`: Code reorganization/refactoring
- `perf`: Performance improvements
- `chore`: CI, build scripts, tooling

---

## Semantic Commit Messages

Follow **Conventional Commits**:

```

<type>(scope): short message

```

Examples:
- `feat(router): implement message forwarding logic`
- `fix(bluetooth): resolve discovery loop issue`
- `docs(readme): add usage diagram`

Use these types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Formatting, whitespace
- `refactor`: Code change without affecting behavior
- `perf`: Performance improvement
- `test`: Adding or updating tests
- `chore`: Tooling or infrastructure work

---

## License

By contributing, you agree that your contributions will be licensed under the **MIT License**.

See `LICENSE` file for details.
