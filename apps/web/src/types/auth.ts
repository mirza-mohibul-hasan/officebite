export type UserRole = 'employee' | 'admin';

export type AuthUser = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  department: string;
  is_active: boolean;
};

export type AuthSession = {
  token: string;
  user: AuthUser;
};

export type LoginRequest = {
  email: string;
  password: string;
};
