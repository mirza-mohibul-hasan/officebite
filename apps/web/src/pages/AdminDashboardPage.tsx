import { useQuery } from '@tanstack/react-query';
import { BarChart3, ClipboardList, Utensils } from 'lucide-react';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { getDashboardSummary } from '../services/analytics';
import { formatCurrency } from '../utils/formatCurrency';

export function AdminDashboardPage() {
  const summaryQuery = useQuery({
    queryKey: ['admin', 'dashboard', 'summary'],
    queryFn: getDashboardSummary,
  });
  const summary = summaryQuery.data;
  const metrics = [
    { label: 'Meals ordered today', value: summary?.today_orders ?? 0, icon: ClipboardList },
    { label: 'Active orders', value: summary?.placed_orders ?? 0, icon: Utensils },
    { label: 'Cancelled orders', value: summary?.cancelled_orders ?? 0, icon: BarChart3 },
    { label: 'Menus today', value: summary?.today_menus ?? 0, icon: Utensils },
  ];

  return (
    <div className="space-y-6">
      <PageHeader title="Admin Dashboard" description="Operational meal analytics will be summarized here." />
      {summaryQuery.isLoading ? <LoadingState label="Loading dashboard" /> : null}
      {summaryQuery.isError ? <ErrorState message="Could not load dashboard summary." /> : null}

      <section className="grid gap-4 md:grid-cols-4">
        {metrics.map((metric) => {
          const Icon = metric.icon;

          return (
            <article key={metric.label} className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
              <div className="flex items-center justify-between">
                <p className="text-sm font-medium text-slate-600">{metric.label}</p>
                <Icon className="text-brand-600" size={20} aria-hidden="true" />
              </div>
              <p className="mt-4 text-3xl font-semibold text-slate-950">{metric.value}</p>
            </article>
          );
        })}
      </section>
      <section className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
        <p className="text-sm font-medium text-slate-600">Estimated revenue today</p>
        <p className="mt-3 text-3xl font-semibold text-slate-950">{formatCurrency(summary?.today_revenue ?? 0)}</p>
      </section>
    </div>
  );
}
