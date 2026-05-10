export type UserRole = 'employee' | 'admin';

export type AuthUser = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
};
