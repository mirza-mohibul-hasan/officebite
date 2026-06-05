import { api } from './api';
import type { Order, OrderStatus } from '../types/order';

export async function placeOrder(menuId: number) {
  const { data } = await api.post<{ order: Order }>('/orders', { menu_id: menuId });
  return data.order;
}

export async function getMyOrders() {
  const { data } = await api.get<{ orders: Order[] }>('/orders');
  return data.orders;
}

export async function cancelOrder(orderId: number) {
  const { data } = await api.patch<{ order: Order }>(`/orders/${orderId}/cancel`);
  return data.order;
}

export async function getAdminOrders() {
  const { data } = await api.get<{ orders: Order[] }>('/admin/orders');
  return data.orders;
}

export async function updateOrderStatus(orderId: number, status: OrderStatus) {
  const { data } = await api.patch<{ order: Order }>(`/admin/orders/${orderId}/status`, { status });
  return data.order;
}
