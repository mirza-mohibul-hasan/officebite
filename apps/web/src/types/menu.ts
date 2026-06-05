export type Menu = {
  id: number;
  title: string;
  description: string;
  price: number;
  available_date: string;
  created_at: string;
  updated_at: string;
};

export type MenuPayload = {
  title: string;
  description: string;
  price: number;
  available_date: string;
};
