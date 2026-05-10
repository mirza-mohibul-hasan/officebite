import { BarChart3, ClipboardList, Utensils } from 'lucide-react';
import { PageHeader } from '../components/PageHeader';

const metrics = [
  { label: 'Meals ordered today', value: '0', icon: ClipboardList },
  { label: 'Menus available', value: '0', icon: Utensils },
  { label: 'Cancelled orders', value: '0', icon: BarChart3 },
];

export function AdminDashboardPage() {
  return (
    <div className="space-y-6">
      <PageHeader title="Admin Dashboard" description="Operational meal analytics will be summarized here." />
      <section className="grid gap-4 md:grid-cols-3">
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
    </div>
  );
}
