# Production Readiness

OfficeBite is still a focused monolith-style MVP, but this pass adds the baseline production controls expected before a serious deployment.

## Required Environment

Set these values outside the repository:

- `APP_ENV=production`
- `DATABASE_URL`
- `JWT_SECRET`, at least 32 characters and unique per environment
- `JWT_ISSUER`
- `WEB_ORIGIN`, comma-separated if multiple origins are needed
- `AUTO_MIGRATE=false`
- `SEED_DATA=false`
- `VITE_API_URL`

## Deployment Shape

Use the production compose override as a reference:

```bash
docker compose -f docker-compose.prod.yml config
docker compose -f docker-compose.prod.yml up --build
```

The production web target serves static Vite assets with Nginx. The API runs as a non-root user and validates unsafe production defaults on startup.

## Database Migrations

The MVP includes SQL migration files under `apps/api/migrations`. Production disables GORM automigration by default. Apply reviewed migrations as part of deployment before starting the new API version.

## CI

GitHub Actions runs:

- API tests
- Web lint and production build
- Compose config validation

## Remaining Hardening Before Public Internet

- Use managed PostgreSQL backups and restore drills.
- Put the API and web app behind TLS.
- Add centralized logs and metrics.
- Add rate limiting at an edge proxy.
- Replace seeded passwords in all non-local environments.
- Add integration tests against a real PostgreSQL service.
