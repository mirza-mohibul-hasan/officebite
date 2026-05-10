import { CalendarDays } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { PageHeader } from '../components/PageHeader';

export function MenuPage() {
  return (
    <div className="space-y-6">
      <PageHeader title="Today's Menu" description="Daily menu cards and ordering actions will connect here." />
      <EmptyState
        icon={CalendarDays}
        title="No menu loaded yet"
        description="The menu management branch will add daily menu data, pricing, and employee ordering actions."
      />
    </div>
  );
}
