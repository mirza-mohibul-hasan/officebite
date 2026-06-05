export type Menu = {
  id: number;
  title: string;
  description: string;
  category: string;
  price: number;
  available_date: string;
  cutoff_time: string;
  max_orders: number;
  is_active: boolean;
  created_at: string;
  updated_at: string;
};

export type MenuPayload = {
  title: string;
  description: string;
  category: string;
  price: number;
  available_date: string;
  cutoff_time: string;
  max_orders: number;
  is_active: boolean;
};
