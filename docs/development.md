# Development

## Local Apps

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

Lint and build:

```bash
cd apps/web
npm run lint
npm run build
```

## Docker

```bash
docker compose up --build
```

The API waits for PostgreSQL, runs GORM automigrations when `AUTO_MIGRATE=true`, and seeds two MVP users when `SEED_DATA=true`.

## Verification

```bash
cd apps/api
go test ./...
```

```bash
cd apps/web
npm run build
```

```bash
docker compose config
```
