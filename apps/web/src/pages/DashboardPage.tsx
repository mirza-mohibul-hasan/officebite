import { CalendarDays, ClipboardList, ShieldCheck } from 'lucide-react';
import { PageHeader } from '../components/PageHeader';
import { useCurrentUser } from '../hooks/useCurrentUser';

const cards = [
  {
    title: "Today's menu",
    description: 'Daily meal discovery and ordering will land in the menu branch.',
    icon: CalendarDays,
  },
  {
    title: 'My orders',
    description: 'Employees will place, cancel, and review meal orders here.',
    icon: ClipboardList,
  },
  {
    title: 'Admin controls',
    description: 'Managers will maintain menus, view orders, and monitor analytics.',
    icon: ShieldCheck,
  },
];

export function DashboardPage() {
  const user = useCurrentUser();

  return (
    <div className="space-y-8">
      <PageHeader
        title={`Welcome, ${user.name}`}
        description="Review today's menu, place a meal order, and keep track of your meal history."
      />

      <section className="grid gap-4 md:grid-cols-3">
        {cards.map((card) => {
          const Icon = card.icon;

          return (
            <article key={card.title} className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
              <Icon className="text-brand-600" size={24} aria-hidden="true" />
              <h3 className="mt-4 text-base font-semibold text-slate-950">{card.title}</h3>
              <p className="mt-2 text-sm leading-6 text-slate-600">{card.description}</p>
            </article>
          );
        })}
      </section>
    </div>
  );
}
