# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability in **Talken**, please report it privately to:


Email: [bitsgenix@gmail.com](mailto:bitsgenix@gmail.com)
Subject: Talken Security Vulnerability Report

Do **not** open a public issue for security vulnerabilities.

---

## Supported Versions

| Version   | Supported          |
|-----------|--------------------|
| `main`    | Yes (development)  |
| `v1.x`    | Yes (release)      |
| `< v1.0`  | No                 |

You should upgrade to the latest `main` or `v1.x` release as soon as fixes are published.

---

## Security Release Process

1. **Report**                     – Vulnerability reported privately via email.  
2. **Acknowledge**               – We will respond within 48 hours.  
3. **Fix**                       – We develop, test, and QA a patch.  
4. **Public Release**            – Create a security release (tagged `vX.Y.Z-security`), update CHANGELOG.  
5. **Notification**              – Notify the community via GitHub release notes.

---

## Reporting Guidelines

When reporting, please include:

- A clear title and description.  
- Steps to reproduce the issue.  
- Impact assessment (data leakage, unauthorized access, etc.).  
- Any suggested fixes or references.

---

## Data Exposure

Talken stores minimal data locally. No central server holds messages. Nevertheless, keep your device secure and encrypted.

---

## External Dependencies

- Keep dependencies up to date.  
- Review changelogs of upstream libraries for security patches.  
- If you find a vulnerability in a dependency, report it upstream and notify us.

---

Thank you for helping keep **Talken** safe!
