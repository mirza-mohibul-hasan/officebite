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
  "price": 1200,
  "available_date": "2026-06-05"
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

## Analytics

Requires admin role.

- `GET /admin/dashboard/summary`
