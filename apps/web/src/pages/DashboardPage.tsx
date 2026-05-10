import { CalendarDays, ClipboardList, ShieldCheck } from 'lucide-react';

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
  return (
    <div className="space-y-8">
      <section className="rounded-lg border border-slate-200 bg-white p-6 shadow-sm">
        <p className="text-sm font-medium uppercase tracking-wide text-brand-700">MVP Bootstrap</p>
        <h2 className="mt-2 text-2xl font-semibold text-slate-950">Office meal management, ready to grow.</h2>
        <p className="mt-3 max-w-2xl text-sm leading-6 text-slate-600">
          The monorepo is wired with a React frontend, Go API, and Dockerized PostgreSQL. Feature branches will fill in
          authentication, menus, orders, and admin analytics.
        </p>
      </section>

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
