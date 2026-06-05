# OfficeBite API

Base URL: `/api/v1`

## Auth

### Login

`POST /auth/login`

```json
{
  "email": "employee@officebite.local",
  "password": "password123"
}
```

Returns:

```json
{
  "token": "jwt",
  "user": {
    "id": 1,
    "name": "Employee User",
    "email": "employee@officebite.local",
    "role": "employee"
  }
}
```

### Current User

`GET /auth/me`

Requires `Authorization: Bearer <token>`.

## Menus

### Today's Menus

`GET /menus/today?date=YYYY-MM-DD`

Requires authentication.

### Admin Menu CRUD

Requires admin role.

- `GET /admin/menus`
- `POST /admin/menus`
- `PUT /admin/menus/:id`
- `DELETE /admin/menus/:id`

Menu payload:

```json
{
  "title": "Chicken rice bowl",
  "description": "Grilled chicken with rice and vegetables",
  "category": "lunch",
  "price": 1200,
  "available_date": "2026-06-05",
  "cutoff_time": "2026-06-05T10:00:00Z",
  "max_orders": 40,
  "is_active": true
}
```

Prices are stored in cents.

## Orders

### Employee Orders

Requires authentication.

- `GET /orders`
- `POST /orders`
- `PATCH /orders/:id/cancel`

Place order payload:

```json
{
  "menu_id": 1
}
```

### Admin Orders

Requires admin role.

- `GET /admin/orders`
- `PATCH /admin/orders/:id/status`

Status payload:

```json
{
  "status": "confirmed"
}
```

Allowed statuses: `placed`, `confirmed`, `delivered`, `cancelled`.

## Users

Requires admin role.

- `GET /admin/users`
- `POST /admin/users`
- `PUT /admin/users/:id`

User payload:

```json
{
  "name": "Ayesha Rahman",
  "email": "ayesha@example.com",
  "password": "password123",
  "role": "employee",
  "department": "Finance",
  "is_active": true
}
```

## Analytics

Requires admin role.

- `GET /admin/dashboard/summary`
