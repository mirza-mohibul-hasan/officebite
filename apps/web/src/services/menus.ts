import { api } from './api';
import type { Menu, MenuPayload } from '../types/menu';

export async function getTodayMenus(date?: string) {
  const { data } = await api.get<{ menus: Menu[] }>('/menus/today', {
    params: date ? { date } : undefined,
  });
  return data.menus;
}

export async function getAdminMenus() {
  const { data } = await api.get<{ menus: Menu[] }>('/admin/menus');
  return data.menus;
}

export async function createMenu(payload: MenuPayload) {
  const { data } = await api.post<{ menu: Menu }>('/admin/menus', payload);
  return data.menu;
}

export async function updateMenu(id: number, payload: MenuPayload) {
  const { data } = await api.put<{ menu: Menu }>(`/admin/menus/${id}`, payload);
  return data.menu;
}

export async function deleteMenu(id: number) {
  await api.delete(`/admin/menus/${id}`);
}
