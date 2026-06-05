import { api } from './api';
import type { AuthUser, UserRole } from '../types/auth';

export type UserPayload = {
  name: string;
  email: string;
  password?: string;
  role: UserRole;
  department: string;
  is_active: boolean;
};

export async function getAdminUsers() {
  const { data } = await api.get<{ users: AuthUser[] }>('/admin/users');
  return data.users;
}

export async function createUser(payload: UserPayload) {
  const { data } = await api.post<{ user: AuthUser }>('/admin/users', payload);
  return data.user;
}

export async function updateUser(id: number, payload: UserPayload) {
  const { data } = await api.put<{ user: AuthUser }>(`/admin/users/${id}`, payload);
  return data.user;
}
