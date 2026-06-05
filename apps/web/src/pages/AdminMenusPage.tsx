import { FormEvent, useState } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Pencil, Plus, Trash2, Utensils } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { createMenu, deleteMenu, getAdminMenus, updateMenu } from '../services/menus';
import type { Menu, MenuPayload } from '../types/menu';
import { formatCurrency } from '../utils/formatCurrency';
import { formatDate, todayISODate } from '../utils/formatDate';

const emptyForm: MenuPayload = {
  title: '',
  description: '',
  category: 'lunch',
  price: 1200,
  available_date: todayISODate(),
  cutoff_time: `${todayISODate()}T10:00`,
  max_orders: 0,
  is_active: true,
};

export function AdminMenusPage() {
  const queryClient = useQueryClient();
  const [form, setForm] = useState<MenuPayload>(emptyForm);
  const [editingId, setEditingId] = useState<number | null>(null);

  const menusQuery = useQuery({
    queryKey: ['admin', 'menus'],
    queryFn: getAdminMenus,
  });

  const saveMutation = useMutation({
    mutationFn: (payload: MenuPayload) => (editingId ? updateMenu(editingId, payload) : createMenu(payload)),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'menus'] });
      queryClient.invalidateQueries({ queryKey: ['menus'] });
      setForm(emptyForm);
      setEditingId(null);
    },
  });

  const deleteMutation = useMutation({
    mutationFn: deleteMenu,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'menus'] });
      queryClient.invalidateQueries({ queryKey: ['menus'] });
    },
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    saveMutation.mutate({ ...form, cutoff_time: new Date(form.cutoff_time).toISOString() });
  }

  function startEdit(menu: Menu) {
    setEditingId(menu.id);
    setForm({
      title: menu.title,
      description: menu.description,
      category: menu.category,
      price: menu.price,
      available_date: menu.available_date.slice(0, 10),
      cutoff_time: menu.cutoff_time.slice(0, 16),
      max_orders: menu.max_orders,
      is_active: menu.is_active,
    });
  }

  return (
    <div className="space-y-6">
      <PageHeader
        title="Menu Management"
        description="Create, update, and remove daily meal menus."
        action={
          <button
            type="button"
            onClick={() => {
              setEditingId(null);
              setForm(emptyForm);
            }}
            className="inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700"
          >
            <Plus size={16} aria-hidden="true" />
            New menu
          </button>
        }
      />

      <form className="grid gap-4 rounded-lg border border-slate-200 bg-white p-5 shadow-sm" onSubmit={handleSubmit}>
        {saveMutation.isError ? <ErrorState message="Could not save this menu." /> : null}
        <div className="grid gap-4 md:grid-cols-2">
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Title
            <input
              value={form.title}
              onChange={(event) => setForm({ ...form, title: event.target.value })}
              className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              required
            />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Available date
            <input
              type="date"
              value={form.available_date}
              onChange={(event) => setForm({ ...form, available_date: event.target.value })}
              className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              required
            />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Category
            <input
              value={form.category}
              onChange={(event) => setForm({ ...form, category: event.target.value })}
              className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              required
            />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Cutoff time
            <input
              type="datetime-local"
              value={form.cutoff_time}
              onChange={(event) => setForm({ ...form, cutoff_time: event.target.value })}
              className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              required
            />
          </label>
        </div>
        <label className="grid gap-2 text-sm font-medium text-slate-700">
          Description
          <textarea
            value={form.description}
            onChange={(event) => setForm({ ...form, description: event.target.value })}
            className="min-h-24 rounded-md border border-slate-300 px-3 py-2 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
            required
          />
        </label>
        <label className="grid gap-2 text-sm font-medium text-slate-700 md:max-w-xs">
          Price in cents
          <input
            type="number"
            min={1}
            value={form.price}
            onChange={(event) => setForm({ ...form, price: Number(event.target.value) })}
            className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
            required
          />
        </label>
        <div className="grid gap-4 md:grid-cols-2">
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Max orders
            <input
              type="number"
              min={0}
              value={form.max_orders}
              onChange={(event) => setForm({ ...form, max_orders: Number(event.target.value) })}
              className="h-10 rounded-md border border-slate-300 px-3 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
            />
          </label>
          <label className="flex items-center gap-2 pt-7 text-sm font-medium text-slate-700">
            <input
              type="checkbox"
              checked={form.is_active}
              onChange={(event) => setForm({ ...form, is_active: event.target.checked })}
            />
            Active menu
          </label>
        </div>
        <div className="flex gap-3">
          <button
            type="submit"
            disabled={saveMutation.isPending}
            className="inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700 disabled:opacity-70"
          >
            {editingId ? 'Update menu' : 'Create menu'}
          </button>
          {editingId ? (
            <button
              type="button"
              onClick={() => {
                setEditingId(null);
                setForm(emptyForm);
              }}
              className="h-10 rounded-md border border-slate-300 px-4 text-sm font-semibold text-slate-700 hover:bg-slate-100"
            >
              Cancel
            </button>
          ) : null}
        </div>
      </form>

      {menusQuery.isLoading ? <LoadingState label="Loading menus" /> : null}
      {menusQuery.isError ? <ErrorState message="Could not load menus." /> : null}
      {menusQuery.data?.length === 0 ? (
        <EmptyState icon={Utensils} title="No menus yet" description="Create the first menu to publish meal options." />
      ) : null}

      <section className="grid gap-4">
        {menusQuery.data?.map((menu) => (
          <article key={menu.id} className="rounded-lg border border-slate-200 bg-white p-5 shadow-sm">
            <div className="flex flex-col gap-4 md:flex-row md:items-start md:justify-between">
              <div>
                <div className="flex flex-wrap items-center gap-3">
                  <h3 className="text-lg font-semibold text-slate-950">{menu.title}</h3>
                  <span className="rounded-md bg-slate-100 px-2 py-1 text-xs font-medium text-slate-600">
                    {formatDate(menu.available_date)}
                  </span>
                  <span className="rounded-md bg-brand-50 px-2 py-1 text-xs font-semibold text-brand-700">
                    {formatCurrency(menu.price)}
                  </span>
                  <span className="rounded-md bg-slate-100 px-2 py-1 text-xs capitalize text-slate-600">
                    {menu.category}
                  </span>
                  <span
                    className={
                      menu.is_active
                        ? 'rounded-md bg-emerald-50 px-2 py-1 text-xs text-emerald-700'
                        : 'rounded-md bg-red-50 px-2 py-1 text-xs text-red-700'
                    }
                  >
                    {menu.is_active ? 'Active' : 'Inactive'}
                  </span>
                </div>
                <p className="mt-2 text-sm leading-6 text-slate-600">{menu.description}</p>
                <p className="mt-1 text-xs text-slate-500">
                  Cutoff {formatDate(menu.cutoff_time)} {menu.max_orders > 0 ? `- Capacity ${menu.max_orders}` : '- Unlimited capacity'}
                </p>
              </div>
              <div className="flex gap-2">
                <button
                  type="button"
                  onClick={() => startEdit(menu)}
                  className="grid h-10 w-10 place-items-center rounded-md border border-slate-200 text-slate-600 hover:bg-slate-100"
                  title="Edit menu"
                >
                  <Pencil size={16} aria-hidden="true" />
                </button>
                <button
                  type="button"
                  onClick={() => deleteMutation.mutate(menu.id)}
                  className="grid h-10 w-10 place-items-center rounded-md border border-red-200 text-red-600 hover:bg-red-50"
                  title="Delete menu"
                >
                  <Trash2 size={16} aria-hidden="true" />
                </button>
              </div>
            </div>
          </article>
        ))}
      </section>
    </div>
  );
}
