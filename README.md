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
- JWT authentication
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

This keeps the codebase organized while still deploying as a single service.

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

---

# Authentication

The API uses **JWT authentication**.

Login returns an access token which must be included in protected requests:

```
Authorization: Bearer <token>
```

---

# Endpoints

## Health Check

```
GET /api/v1/health
```

---

# Authentication

## Register User

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

## Login

```
POST /api/v1/auth/login
```

Response:

```json
{
  "accessToken": "jwt_token_here"
}
```

---

## Get Current User

Protected endpoint.

```
GET /api/v1/me
```

Headers:

```
Authorization: Bearer <token>
```

Example response:

```json
{
  "id": "uuid",
  "email": "zlatko@example.com",
  "role": "organizer"
}
```

---

# Events

## Get All Events

```
GET /api/v1/events
```

---

## Get Event By ID

```
GET /api/v1/events/:id
```

---

## Create Event

Only users with role **organizer** can create events.

```
POST /api/v1/events
```

Headers:

```
Authorization: Bearer <token>
```

Example request:

```json
{
  "title": "Go Conference 2026",
  "description": "A conference for Go developers",
  "location": "Sofia, Bulgaria",
  "startsAt": "2026-04-10T09:00:00Z",
  "endsAt": "2026-04-10T18:00:00Z",
  "capacity": 150
}
```

---

# Bookings

Users can book events and manage their bookings.

---

## Create Booking

```
POST /api/v1/events/:id/bookings
```

Headers:

```
Authorization: Bearer <token>
```

Rules enforced:

- user cannot book the same event twice
- event capacity cannot be exceeded

---

## Get My Bookings

```
GET /api/v1/me/bookings
```

Headers:

```
Authorization: Bearer <token>
```

Returns all bookings for the authenticated user.

---

## Cancel Booking

```
DELETE /api/v1/bookings/:id
```

Headers:

```
Authorization: Bearer <token>
```

Users can only cancel **their own bookings**.

---

# Database

PostgreSQL is used as the primary database.

Tables currently implemented:

```
users
events
bookings
```

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

---

### HTTP API Testing

API routes are stored in `.http` files and can be executed directly from the editor.

---

# Current MVP Features

Eventix now supports:

- User registration
- User login
- JWT authentication
- Organizer event creation
- Event discovery
- Event booking
- Booking cancellation
- Viewing user bookings

---

# Roadmap

## Phase 1 (Current)

- Authentication
- Event management
- Booking system

## Phase 2

- Event search and filtering
- Pagination
- Validation improvements
- Organizer dashboards

## Phase 3

- Redis caching
- RabbitMQ event processing

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
