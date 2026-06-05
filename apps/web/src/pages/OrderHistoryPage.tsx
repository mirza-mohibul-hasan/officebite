import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { ClipboardList, XCircle } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { cancelOrder, getMyOrders } from '../services/orders';
import { formatCurrency } from '../utils/formatCurrency';
import { formatDate } from '../utils/formatDate';

export function OrderHistoryPage() {
  const queryClient = useQueryClient();
  const ordersQuery = useQuery({
    queryKey: ['orders', 'mine'],
    queryFn: getMyOrders,
  });
  const cancelMutation = useMutation({
    mutationFn: cancelOrder,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ['orders'] }),
  });

  return (
    <div className="space-y-6">
      <PageHeader title="My Orders" description="Employees will review placed and cancelled meal orders here." />
      {ordersQuery.isLoading ? <LoadingState label="Loading orders" /> : null}
      {ordersQuery.isError ? <ErrorState message="Could not load your orders." /> : null}
      {cancelMutation.isError ? <ErrorState message="Could not cancel this order." /> : null}
      {ordersQuery.data?.length === 0 ? (
        <EmptyState icon={ClipboardList} title="No order history yet" description="Place an order from today's menu." />
      ) : null}

      <section className="grid gap-4">
        {ordersQuery.data?.map((order) => (
          <article key={order.id} className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
            <div className="flex flex-col gap-4 md:flex-row md:items-start md:justify-between">
              <div>
                <div className="flex flex-wrap items-center gap-3">
                  <h3 className="text-lg font-semibold text-slate-950">{order.menu?.title ?? 'Meal order'}</h3>
                  <span className="rounded-md bg-slate-100 px-2 py-1 text-xs font-medium capitalize text-slate-600">
                    {order.status}
                  </span>
                  {order.menu ? (
                    <span className="rounded-md bg-brand-50 px-2 py-1 text-xs font-semibold text-brand-700">
                      {formatCurrency(order.menu.price)}
                    </span>
                  ) : null}
                </div>
                <p className="mt-2 text-sm text-slate-600">Ordered {formatDate(order.created_at)}</p>
              </div>
              {order.status === 'placed' ? (
                <button
                  type="button"
                  onClick={() => cancelMutation.mutate(order.id)}
                  disabled={cancelMutation.isPending}
                  className="inline-flex h-10 items-center gap-2 rounded-md border border-red-200 px-4 text-sm font-semibold text-red-600 hover:bg-red-50 disabled:opacity-70"
                >
                  <XCircle size={16} aria-hidden="true" />
                  Cancel
                </button>
              ) : null}
            </div>
          </article>
        ))}
      </section>
    </div>
  );
}
