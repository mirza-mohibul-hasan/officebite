# OfficeBite

OfficeBite is a production-style MVP for managing office meals across employees and admins.

## Stack

- Frontend: React, TypeScript, Vite, Tailwind CSS, React Router, Zustand, TanStack Query
- Backend: Go, Gin, GORM, PostgreSQL
- Infrastructure: Docker, Docker Compose
- Authentication: JWT

## Monorepo Layout

```text
officebite/
├── apps/
│   ├── api/
│   └── web/
├── infrastructure/
│   └── docker/
├── docs/
└── README.md
```

## Quick Start

```bash
docker compose up --build
```

Services:

- Web: http://localhost:5173
- API: http://localhost:8080
- PostgreSQL: localhost:5432

Seeded MVP accounts:

- Admin: `admin@officebite.local` / `password123`
- Employee: `employee@officebite.local` / `password123`

## Local Development

API:

```bash
cd apps/api
go run ./cmd/server
```

Web:

```bash
cd apps/web
npm install
npm run dev
```

## Branch Workflow

- `main`: stable release branch
- `develop`: integration branch
- Feature branches: `feature/<scope>`

Current MVP phases live in `docs/architecture.md`.

## MVP Features

- JWT login/logout with employee and admin roles
- Protected frontend routes
- Employee dashboard, daily menu view, order placement, cancellation, and history
- Admin menu CRUD
- Admin user management
- Admin all-orders view
- Admin order lifecycle controls
- Admin analytics summary

## API Overview

All application routes are prefixed with `/api/v1`.

- `POST /auth/login`
- `GET /auth/me`
- `GET /menus/today?date=YYYY-MM-DD`
- `GET /orders`
- `POST /orders`
- `PATCH /orders/:id/cancel`
- `GET /admin/menus`
- `POST /admin/menus`
- `PUT /admin/menus/:id`
- `DELETE /admin/menus/:id`
- `GET /admin/orders`
- `PATCH /admin/orders/:id/status`
- `GET /admin/users`
- `POST /admin/users`
- `PUT /admin/users/:id`
- `GET /admin/dashboard/summary`

More detail lives in `docs/api.md`.

## Production Notes

Production hardening notes live in `docs/production.md`.
