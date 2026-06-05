import type { Menu } from './menu';
import type { AuthUser } from './auth';

export type OrderStatus = 'placed' | 'confirmed' | 'delivered' | 'cancelled';

export type Order = {
  id: number;
  user_id: number;
  menu_id: number;
  status: OrderStatus;
  created_at: string;
  updated_at: string;
  user?: AuthUser;
  menu?: Menu;
};
