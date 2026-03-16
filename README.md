# Eventix

Eventix is a modern event booking platform built with **Go** and **React** in a **monorepo architecture**.

The goal of this project is to simulate a **real-world production SaaS platform** similar to Eventbrite or Eventim, supporting:

- event discovery
- event creation
- event bookings
- user authentication

This project is also part of a **progressive architecture journey**, evolving from a modular monolith into a distributed cloud-native system with:

- messaging
- microservices
- Kubernetes
- observability

---

# Tech Stack

## Backend

- Go
- Gin
- PostgreSQL
- pgx
- JWT authentication
- bcrypt
- golang-migrate
- Air (live reload)
- CORS middleware

---

## Frontend

- React
- Vite
- TypeScript
- React Router
- TanStack Query
- Zustand
- TailwindCSS
- shadcn/ui
- Sonner (toast notifications)

---

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

Each module contains:

- models
- repositories
- services
- handlers

This structure keeps the codebase organized and scalable while still deploying as a single service.

---

# Frontend Architecture

The frontend is built using **feature-based architecture**.

```
src
│
├── app
│   ├── router.tsx
│   └── query-client.ts
│
├── components
│   └── ui            # shadcn/ui components
│
├── features
│   └── auth
│       ├── api
│       └── hooks
│
├── layouts
│
├── pages
│
├── stores
│
└── lib
    └── api.ts
```

Server state is handled with **TanStack Query**, while **Zustand manages session state**.

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
moon run api:dev
```

---

## Run frontend

```bash
moon run web:dev
```

---

## Run both

```bash
moon run :dev
```

Backend:

```
http://localhost:8080
```

Frontend:

```
http://localhost:5173
```

---

# Database Migrations

Migrations are handled with **golang-migrate**.

Run migrations:

```bash
migrate \
-path apps/api/migrations \
-database "postgres://postgres:Password123%21@localhost:5432/eventix?sslmode=disable" \
up
```

---

# Database Seeding

The project includes a **seed script** for local development.

It creates:

- organizers
- users
- events
- bookings
- additional events for pagination testing

Run seed:

```bash
moon run api:seed
```

---

# Default Seed Users

| Email                  | Role      | Password     |
| ---------------------- | --------- | ------------ |
| organizer1@example.com | organizer | Password123! |
| organizer2@example.com | organizer | Password123! |
| attendee1@example.com  | user      | Password123! |
| attendee2@example.com  | user      | Password123! |

---

# API

Base URL

```
/api/v1
```

---

# Authentication

The API uses **JWT authentication**.

Login returns an access token.

Example response:

```json
{
  "accessToken": "jwt_token_here"
}
```

Frontend stores the token and attaches it automatically to requests.

---

# Endpoints

## Health

```
GET /api/v1/health
```

---

# Authentication

## Register

```
POST /api/v1/auth/register
```

Example request:

```json
{
  "email": "user@example.com",
  "password": "Password123!",
  "role": "user"
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
  "accessToken": "..."
}
```

---

## Current User

```
GET /api/v1/me
```

Requires Authorization header.

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

## Get Event

```
GET /api/v1/events/:id
```

---

## Create Event

Organizer only.

```
POST /api/v1/events
```

---

# Bookings

Users can book events.

---

## Create Booking

```
POST /api/v1/events/:id/bookings
```

Rules:

- cannot book same event twice
- event capacity cannot be exceeded

---

## Get My Bookings

```
GET /api/v1/bookings
```

Response includes event details.

---

## Cancel Booking

```
DELETE /api/v1/bookings/:id
```

Users can only cancel their own bookings.

---

# Frontend Features

The React application currently supports:

- Login page
- Register page
- Auth state with Zustand
- API requests via TanStack Query
- Global notifications using Sonner
- shadcn/ui components
- TailwindCSS styling

---

# Development Tools

### Air

Live reload for Go development.

```
air
```

---

### HTTP API Testing

API routes are stored in `.http` files.

---

# Current MVP Features

Eventix currently supports:

- User registration
- User login
- JWT authentication
- Organizer event creation
- Event discovery
- Pagination
- Event booking
- Booking cancellation
- Database seeding
- Frontend authentication UI

---

# Roadmap

## Phase 1 (Current)

- Authentication
- Event management
- Booking system
- Pagination
- Frontend authentication

---

## Phase 2

- Event search
- Filtering
- Protected routes
- Organizer dashboards

---

## Phase 3

- Redis caching
- RabbitMQ

---

## Phase 4

- Microservices
- gRPC communication

---

## Phase 5

- Kubernetes deployment
- Helm
- Prometheus
- Grafana

---

## Phase 6

- Cloud deployment (GCP)
- Terraform

---

# Future Vision

Eventix will evolve into a **production-grade distributed system** including:

- microservices
- async messaging
- scalable infrastructure
- observability

---

# Author

Zlatko Iliev
