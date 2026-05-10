import { ClipboardList } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { PageHeader } from '../components/PageHeader';

export function OrderHistoryPage() {
  return (
    <div className="space-y-6">
      <PageHeader title="My Orders" description="Employees will review placed and cancelled meal orders here." />
      <EmptyState
        icon={ClipboardList}
        title="No order history yet"
        description="The order management branch will add order placement, cancellation, and history data."
      />
    </div>
  );
}
