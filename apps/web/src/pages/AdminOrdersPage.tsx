import { useQuery } from '@tanstack/react-query';
import { ClipboardList } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { getAdminOrders } from '../services/orders';
import { formatCurrency } from '../utils/formatCurrency';
import { formatDate } from '../utils/formatDate';

export function AdminOrdersPage() {
  const ordersQuery = useQuery({
    queryKey: ['admin', 'orders'],
    queryFn: getAdminOrders,
  });

  return (
    <div className="space-y-6">
      <PageHeader title="Order Management" description="Admins will monitor all meal orders and statuses here." />
      {ordersQuery.isLoading ? <LoadingState label="Loading all orders" /> : null}
      {ordersQuery.isError ? <ErrorState message="Could not load all orders." /> : null}
      {ordersQuery.data?.length === 0 ? (
        <EmptyState icon={ClipboardList} title="No orders to review" description="Orders will appear here as employees place meals." />
      ) : null}

      <section className="overflow-hidden rounded-lg border border-slate-200 bg-white shadow-sm">
        <div className="grid grid-cols-[1.2fr_1.2fr_.7fr_.7fr] gap-4 border-b border-slate-200 bg-slate-50 px-4 py-3 text-xs font-semibold uppercase tracking-wide text-slate-500">
          <span>Employee</span>
          <span>Meal</span>
          <span>Status</span>
          <span>Ordered</span>
        </div>
        {ordersQuery.data?.map((order) => (
          <div
            key={order.id}
            className="grid grid-cols-[1.2fr_1.2fr_.7fr_.7fr] gap-4 border-b border-slate-100 px-4 py-4 text-sm last:border-b-0"
          >
            <span className="font-medium text-slate-950">{order.user?.name ?? 'Employee'}</span>
            <span className="text-slate-600">
              {order.menu?.title ?? 'Meal'} {order.menu ? `(${formatCurrency(order.menu.price)})` : ''}
            </span>
            <span className="capitalize text-slate-600">{order.status}</span>
            <span className="text-slate-600">{formatDate(order.created_at)}</span>
          </div>
        ))}
      </section>
    </div>
  );
}
