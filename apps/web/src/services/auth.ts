import { api } from './api';
import type { AuthSession, LoginRequest } from '../types/auth';

export async function login(payload: LoginRequest) {
  const { data } = await api.post<AuthSession>('/auth/login', payload);
  return data;
}

export async function getCurrentUser() {
  const { data } = await api.get<Pick<AuthSession, 'user'>>('/auth/me');
  return data.user;
}
