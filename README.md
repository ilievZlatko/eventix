# Eventix

Eventix is a modern event booking platform built with **Go** and **React** in a **monorepo architecture**.

The goal of this project is to simulate a **real-world production SaaS platform** similar to Eventbrite or Eventim, supporting event discovery, ticket booking, and organizer dashboards.

This project is also part of a **progressive learning journey**, evolving from a monolith into a full cloud-native architecture with microservices, Kubernetes, and observability.

---

# Tech Stack

## Backend

- Go
- Gin
- PostgreSQL
- pgx
- bcrypt
- JWT (coming next)
- golang-migrate
- Air (live reload)

## Frontend

- React
- TypeScript
- Vite
- TailwindCSS
- shadcn/ui
- Zustand

## Infrastructure

- Docker
- Docker Compose
- Moon monorepo

---

# Monorepo Structure

```
eventix
│
├── apps
│   ├── api        # Go backend
│   └── web        # React frontend
│
├── packages       # shared libraries (future)
│
├── docker-compose.yml
└── README.md
```

---

# Backend Architecture

The backend follows a **modular monolith** architecture.

```
internal
│
├── modules
│   ├── auth
│   ├── users
│   ├── events
│   └── bookings
│
└── platform
    ├── config
    ├── db
    ├── auth
    └── middleware
```

Each module contains its own:

- models
- repositories
- services
- handlers

This approach keeps the code organized while still deploying as a single service.

---

# Running the Project

## Start database

From the repo root:

```
docker compose up -d
```

---

## Run backend

```
cd apps/api
air
```

Server will start at:

```
http://localhost:8080
```

---

# API

Base URL

```
/api/v1
```

## Health check

```
GET /api/v1/health
```

## Register user

```
POST /api/v1/auth/register
```

Example request:

```json
{
  "email": "zlatko@example.com",
  "password": "Password123!",
  "role": "organizer"
}
```

---

# Database

PostgreSQL is used as the primary database.

Migrations are managed with **golang-migrate**.

Run migrations:

```
migrate \
  -path apps/api/migrations \
  -database "postgres://postgres:Password123%21@localhost:5432/eventix?sslmode=disable" \
  up
```

---

# Development Tools

### Air

Live reload for Go development.

```
air
```

### HTTP API testing

Routes are stored in `.http` files and can be executed directly from the editor.

---

# Roadmap

## Phase 1

- User registration
- Login
- JWT authentication
- Event CRUD
- Booking system

## Phase 2

- Organizer dashboard
- Event search and filters
- Pagination
- Validation improvements

## Phase 3

- Redis caching
- RabbitMQ event system

## Phase 4

- Microservices architecture
- gRPC communication

## Phase 5

- Kubernetes deployment
- Helm charts
- Observability (Prometheus + Grafana)

## Phase 6

- Cloud deployment (GCP)
- Terraform infrastructure

---

# Future Vision

Eventix will gradually evolve into a **production-grade distributed system** including:

- microservices
- asynchronous messaging
- scalable cloud infrastructure
- monitoring and observability

---

# Author

Zlatko Iliev
