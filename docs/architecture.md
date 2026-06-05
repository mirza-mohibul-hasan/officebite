# OfficeBite MVP Architecture

OfficeBite is organized as a focused monorepo with separate deployable apps and shared infrastructure.

## Apps

- `apps/api`: Go API using Gin for HTTP routing, GORM for persistence, PostgreSQL as the database, and JWT for authentication.
- `apps/web`: React SPA using Vite, TypeScript, Tailwind CSS, React Router, Zustand, and TanStack Query.

## Backend Boundaries

- `handlers`: HTTP request/response adapters.
- `services`: business use cases such as login, menu management, order placement, and analytics.
- `repository`: database access through GORM.
- `models`: persistence/domain structs.
- `middleware`: authentication, role checks, CORS, and request concerns.
- `routes`: route grouping and dependency wiring.
- `config`: environment loading and runtime settings.
- `utils`: cross-cutting helpers.

## Implemented MVP Scope

- Authentication: JWT login, persisted frontend session, protected routes, employee/admin roles.
- Menus: admin CRUD APIs and UI, employee daily menu view.
- Orders: employee order placement, cancellation, history, and admin order review.
- Dashboard: admin summary metrics for orders, menus, cancellations, and estimated revenue.

## Scope Guardrails

The MVP intentionally keeps one API service and one web app. It avoids microservices, queue infrastructure, Kubernetes, and enterprise-grade observability until product needs justify them.

## Branch Sequence

1. `feature/project-bootstrap`
2. `feature/docker-setup`
3. `feature/backend-foundation`
4. `feature/frontend-foundation`
5. `feature/authentication`
6. `feature/menu-management`
7. `feature/order-management`
8. `feature/admin-dashboard`
9. `feature/polishing`
