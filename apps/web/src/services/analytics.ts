import { api } from './api';
import type { DashboardSummary } from '../types/analytics';

export async function getDashboardSummary() {
  const { data } = await api.get<{ summary: DashboardSummary }>('/admin/dashboard/summary');
  return data.summary;
}
