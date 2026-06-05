CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(120) NOT NULL,
  email VARCHAR(160) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(24) NOT NULL DEFAULT 'employee' CHECK (role IN ('employee', 'admin')),
  department VARCHAR(120) NOT NULL DEFAULT '',
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS menus (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(160) NOT NULL,
  description TEXT NOT NULL,
  category VARCHAR(80) NOT NULL DEFAULT 'lunch',
  price BIGINT NOT NULL,
  available_date DATE NOT NULL,
  cutoff_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  max_orders INTEGER NOT NULL DEFAULT 0,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_menus_available_date ON menus (available_date);

CREATE TABLE IF NOT EXISTS orders (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  menu_id BIGINT NOT NULL REFERENCES menus(id) ON UPDATE CASCADE ON DELETE RESTRICT,
  status VARCHAR(24) NOT NULL DEFAULT 'placed' CHECK (status IN ('placed', 'confirmed', 'delivered', 'cancelled')),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders (user_id);
CREATE INDEX IF NOT EXISTS idx_orders_menu_id ON orders (menu_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders (status);
CREATE UNIQUE INDEX IF NOT EXISTS idx_orders_active_user_menu ON orders (user_id, menu_id) WHERE status = 'placed';
