# Docker Workflow

The root `docker-compose.yml` defines the MVP runtime:

- PostgreSQL database
- Go API
- React web app

The checked-in `docker-compose.override.yml` is for local development and adds source mounts for quicker iteration.

## Commands

```bash
docker compose up --build
docker compose down
docker compose down -v
```

## Production Web Image

The web Dockerfile includes a `production` target that builds static assets and serves them with Nginx:

```bash
docker build -f infrastructure/docker/web.Dockerfile --target production -t officebite-web:prod .
```
