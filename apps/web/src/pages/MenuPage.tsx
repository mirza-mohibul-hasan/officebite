import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { CalendarDays, ShoppingCart } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { getTodayMenus } from '../services/menus';
import { placeOrder } from '../services/orders';
import { formatCurrency } from '../utils/formatCurrency';
import { formatDate, todayISODate } from '../utils/formatDate';

export function MenuPage() {
  const queryClient = useQueryClient();
  const today = todayISODate();
  const menusQuery = useQuery({
    queryKey: ['menus', 'today', today],
    queryFn: () => getTodayMenus(today),
  });
  const orderMutation = useMutation({
    mutationFn: placeOrder,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['orders'] });
    },
  });

  return (
    <div className="space-y-6">
      <PageHeader title="Today's Menu" description={`Available meals for ${formatDate(today)}.`} />

      {menusQuery.isLoading ? <LoadingState label="Loading menus" /> : null}
      {menusQuery.isError ? <ErrorState message="Could not load today's menus." /> : null}
      {orderMutation.isError ? <ErrorState message="Could not place this order. You may already have an active order for this meal." /> : null}
      {menusQuery.data?.length === 0 ? (
        <EmptyState
          icon={CalendarDays}
          title="No menu published yet"
          description="Check back later or ask an admin to publish today's meal options."
        />
      ) : null}

      <section className="grid gap-4 md:grid-cols-2">
        {menusQuery.data?.map((menu) => (
          <article key={menu.id} className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
            <div className="flex items-start justify-between gap-4">
              <div>
                <h3 className="text-lg font-semibold text-slate-950">{menu.title}</h3>
                <p className="mt-2 text-sm leading-6 text-slate-600">{menu.description}</p>
              </div>
              <p className="shrink-0 rounded-md bg-slate-100 px-3 py-1 text-sm font-semibold text-slate-900">
                {formatCurrency(menu.price)}
              </p>
            </div>
            <button
              type="button"
              onClick={() => orderMutation.mutate(menu.id)}
              disabled={orderMutation.isPending}
              className="mt-5 inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700 disabled:opacity-70"
            >
              <ShoppingCart size={16} aria-hidden="true" />
              Order meal
            </button>
          </article>
        ))}
      </section>
    </div>
  );
}
