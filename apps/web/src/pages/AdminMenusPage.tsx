import { Plus, Utensils } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { PageHeader } from '../components/PageHeader';

export function AdminMenusPage() {
  return (
    <div className="space-y-6">
      <PageHeader
        title="Menu Management"
        description="Admins will create, update, and remove daily meal menus here."
        action={
          <button
            type="button"
            className="inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700"
          >
            <Plus size={16} aria-hidden="true" />
            New menu
          </button>
        }
      />
      <EmptyState
        icon={Utensils}
        title="No menus yet"
        description="Menu CRUD screens and API integration will land in the menu management branch."
      />
    </div>
  );
}
