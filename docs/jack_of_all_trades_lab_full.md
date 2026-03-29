# Project: Jack of All Trades Lab

## Overview

This project is a hands-on infrastructure and IAM showcase designed to demonstrate real, practical understanding of modern systems.

The core idea is not just to *build something that works*, but to:

> Make invisible infrastructure processes visible and understandable in real time.

---

## Core Vision

A live system where you can **watch authentication and request flow happen step-by-step**:

anonymous request  
→ redirect to OIDC provider  
→ login  
→ JWT issued  
→ proxy validates token  
→ identity headers injected  
→ backend processes authenticated request  

All of this is visualized in real time in the frontend.

---

## Goals

- Deep understanding of networking (TCP, ports, connection states)
- Reverse proxy and request routing
- Authentication (OIDC / OAuth2)
- JWT validation and identity propagation
- Containerization with Docker
- Observability and system insight
- Real-world architecture design

---

## Infrastructure Stack

- **DNS**: Cloudflare (DNS only initially)
- **Compute**: Hetzner VPS (4 CPU / 8GB RAM)
- **OS**: Debian/Ubuntu Linux
- **Firewall**: UFW (ports 22, 80, 443)
- **SSH Security**:
  - Key-based authentication only
  - Password authentication disabled
  - Fail2ban enabled

---

## Architecture

Internet  
→ Nginx (TLS termination, reverse proxy)  
→ oauth2-proxy (OIDC authentication)  
→ Backend service  
→ Event stream  

All services run via Docker Compose.

---

## Development Approach

The system is built incrementally:

1. Base server setup (SSH, firewall, security)
2. Docker + simple backend service
3. Nginx reverse proxy with manual TLS (Certbot)
4. OIDC authentication using oauth2-proxy
5. Event system for real-time updates
6. Frontend visualization
7. Observability layer (Prometheus / Grafana)

---

## Key Feature: Live Auth Flow Visualization

The defining feature of this project.

A frontend dashboard that displays:

- Incoming requests (anonymous vs authenticated)
- Authentication steps in real time
- JWT contents (decoded)
- Identity headers passed to backend
- Request lifecycle timeline

Example timeline:

[12:01:03] request received  
[12:01:04] redirect to OIDC  
[12:01:07] login successful  
[12:01:07] JWT validated  
[12:01:07] user identified  

---

## Event System

Backend emits structured events:

{ "type": "request_received" }  
{ "type": "redirect_to_oidc" }  
{ "type": "token_received", "email": "user@example.com" }  
{ "type": "request_authenticated" }  

These events are streamed to the frontend using:

- Server-Sent Events (SSE)

---

## Frontend (Showcase UI)

The UI is designed to feel alive and reactive:

- Real-time event feed (scrolling)
- Authentication timeline per request
- JWT viewer
- Identity headers display
- Request metadata (IP, latency)

Inspired by real-time analytics tools like Databuddy.

---

## Future Extensions

- IAM API (CRUD):
  - Users
  - Roles
  - Permissions
- Visualization of authorization decisions
- Replay authentication flows
- Failure simulations (expired token, invalid JWT)
- Metrics integration (Prometheus)
- Grafana dashboards

---

## Security Strategy

Current:

- Public SSH (for learning)
- Fail2ban protection
- Key-only authentication

Future:

- Restrict SSH to Tailscale network
- Remove public SSH exposure entirely

---

## Key Learning Areas

- TCP/IP behavior (SYN, RST, connection states)
- Port exposure (open vs closed vs filtered)
- Reverse proxy mechanics
- TLS and certificate management
- OIDC authentication flow
- JWT handling and validation
- Identity propagation via headers
- Container networking
- Basic observability principles

---

## Philosophy

This project prioritizes understanding over convenience.

Instead of hiding complexity, it exposes it:

> The system should explain itself while running.

---

## Why This Project Matters

This project demonstrates:

- Systems thinking across multiple layers
- Practical IAM implementation
- Ability to design and explain architecture
- Understanding of modern infrastructure patterns

---

## Current Status

- VPS provisioned
- SSH configured and hardened
- Firewall configured
- Docker Compose tested locally

---

## Next Steps

1. Set up Docker Compose on VPS
2. Deploy Nginx container with TLS
3. Deploy backend service container
4. Configure reverse proxy routing
5. Implement SSE event streaming
6. Build live frontend dashboard

---

## End Goal

A fully working system where:

- Authentication flow is visible in real time
- Infrastructure behavior is observable
- Identity propagation is transparent

A project that not only works, but teaches.
