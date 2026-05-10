import { ClipboardList } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { PageHeader } from '../components/PageHeader';

export function AdminOrdersPage() {
  return (
    <div className="space-y-6">
      <PageHeader title="Order Management" description="Admins will monitor all meal orders and statuses here." />
      <EmptyState
        icon={ClipboardList}
        title="No orders to review"
        description="The order management branch will add all-order views, status data, and admin workflow actions."
      />
    </div>
  );
}
