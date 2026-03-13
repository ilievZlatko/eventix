# Eventix

Eventix is a modern event booking platform built with **Go** and **React** in a **monorepo architecture**.

The goal of this project is to simulate a **real-world production SaaS platform** similar to Eventbrite or Eventim, supporting event discovery, event creation, and booking management.

This project is also part of a **progressive backend architecture journey**, evolving from a modular monolith into a cloud-native system with microservices, messaging, Kubernetes, and observability.

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

The backend follows a **modular monolith architecture**.

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

This keeps the codebase clean and scalable while still deploying as a single service.

---

# Running the Project

## Start the database

From the repository root:

```bash
docker compose up -d
```

---

## Run backend

```bash
cd apps/api
air
```

Server starts at:

```
http://localhost:8080
```

---

# Database Migrations

Migrations are handled using **golang-migrate**.

Run migrations:

```bash
migrate \
-path apps/api/migrations \
-database "postgres://postgres:Password123%21@localhost:5432/eventix?sslmode=disable" \
up
```

---

# Database Seeding

For local development the project includes a **seed script** that populates the database with test data.

The seed script creates:

- organizers
- users
- events
- bookings
- additional sample events for pagination testing

### Default Seed Users

| Email                  | Role      | Password     |
| ---------------------- | --------- | ------------ |
| organizer1@example.com | organizer | Password123! |
| organizer2@example.com | organizer | Password123! |
| attendee1@example.com  | user      | Password123! |
| attendee2@example.com  | user      | Password123! |

### Run the seed script

From `apps/api`:

```bash
go run ./cmd/api/seed
```

This will populate the database with sample users, events, and bookings.

---

### Reset local database (optional)

If you want to start from a clean database:

```bash
docker compose down -v
docker compose up -d
```

Run migrations again:

```bash
migrate \
-path apps/api/migrations \
-database "postgres://postgres:Password123%21@localhost:5432/eventix?sslmode=disable" \
up
```

Then seed again:

```bash
go run ./cmd/api/seed
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

Login returns an access token which must be included in protected requests.

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
  "email": "user@example.com",
  "password": "Password123!",
  "role": "organizer"
}
```

---

## Login

```
POST /api/v1/auth/login
```

Example response:

```json
{
  "accessToken": "jwt_token_here"
}
```

---

## Get Current User

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
  "email": "user@example.com",
  "role": "organizer"
}
```

---

# Events

## Get Events

Supports pagination.

```
GET /api/v1/events?page=1&limit=10
```

Example response:

```json
{
  "data": [],
  "meta": {
    "page": 1,
    "limit": 10
  }
}
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
  "description": "A conference about Go",
  "location": "Berlin",
  "startsAt": "2026-04-10T09:00:00Z",
  "endsAt": "2026-04-10T18:00:00Z",
  "capacity": 100
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

Rules enforced:

- user cannot book the same event twice
- event capacity cannot be exceeded

---

## Get My Bookings

```
GET /api/v1/bookings
```

Returns bookings for the authenticated user including event details.

Example response:

```json
[
  {
    "id": "booking-id",
    "createdAt": "2026-03-12T10:00:00Z",
    "event": {
      "id": "event-id",
      "title": "Go Conference 2026",
      "location": "Berlin",
      "startsAt": "2026-04-10T09:00:00Z"
    }
  }
]
```

---

## Cancel Booking

```
DELETE /api/v1/bookings/:id
```

Users can only cancel **their own bookings**.

---

# Development Tools

### Air

Live reload for Go development.

```bash
air
```

---

### HTTP API Testing

API routes are stored in `.http` files and can be executed directly from the editor.

---

# Current MVP Features

Eventix currently supports:

- User registration
- User login
- JWT authentication
- Organizer event creation
- Event discovery
- Pagination for events
- Event booking
- Booking cancellation
- Viewing user bookings
- Database seeding

---

# Roadmap

## Phase 1 (Current)

- Authentication
- Event management
- Booking system
- Pagination

## Phase 2

- Event search and filtering
- Pagination improvements
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
