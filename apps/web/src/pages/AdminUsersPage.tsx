import { FormEvent, useState } from 'react';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { Pencil, UserPlus, Users } from 'lucide-react';
import { EmptyState } from '../components/EmptyState';
import { ErrorState } from '../components/ErrorState';
import { LoadingState } from '../components/LoadingState';
import { PageHeader } from '../components/PageHeader';
import { createUser, getAdminUsers, updateUser, type UserPayload } from '../services/users';
import type { AuthUser } from '../types/auth';

const emptyForm: UserPayload = {
  name: '',
  email: '',
  password: '',
  role: 'employee',
  department: '',
  is_active: true,
};

export function AdminUsersPage() {
  const queryClient = useQueryClient();
  const [form, setForm] = useState<UserPayload>(emptyForm);
  const [editingId, setEditingId] = useState<number | null>(null);
  const usersQuery = useQuery({ queryKey: ['admin', 'users'], queryFn: getAdminUsers });
  const saveMutation = useMutation({
    mutationFn: (payload: UserPayload) => (editingId ? updateUser(editingId, payload) : createUser(payload)),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['admin', 'users'] });
      setForm(emptyForm);
      setEditingId(null);
    },
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    saveMutation.mutate(form);
  }

  function startEdit(user: AuthUser) {
    setEditingId(user.id);
    setForm({
      name: user.name,
      email: user.email,
      password: '',
      role: user.role,
      department: user.department,
      is_active: user.is_active,
    });
  }

  return (
    <div className="space-y-6">
      <PageHeader title="User Management" description="Manage employees, admins, departments, and account access." />

      <form className="grid gap-4 rounded-lg border border-slate-200 bg-white p-5 shadow-sm" onSubmit={handleSubmit}>
        {saveMutation.isError ? <ErrorState message="Could not save this user." /> : null}
        <div className="grid gap-4 md:grid-cols-2">
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Name
            <input className="h-10 rounded-md border border-slate-300 px-3" value={form.name} onChange={(event) => setForm({ ...form, name: event.target.value })} required />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Email
            <input type="email" className="h-10 rounded-md border border-slate-300 px-3" value={form.email} onChange={(event) => setForm({ ...form, email: event.target.value })} required />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Password {editingId ? <span className="text-xs font-normal text-slate-500">Leave blank to keep current</span> : null}
            <input type="password" className="h-10 rounded-md border border-slate-300 px-3" value={form.password ?? ''} onChange={(event) => setForm({ ...form, password: event.target.value })} required={!editingId} />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Department
            <input className="h-10 rounded-md border border-slate-300 px-3" value={form.department} onChange={(event) => setForm({ ...form, department: event.target.value })} />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Role
            <select className="h-10 rounded-md border border-slate-300 px-3" value={form.role} onChange={(event) => setForm({ ...form, role: event.target.value as UserPayload['role'] })}>
              <option value="employee">Employee</option>
              <option value="admin">Admin</option>
            </select>
          </label>
          <label className="flex items-center gap-2 pt-7 text-sm font-medium text-slate-700">
            <input type="checkbox" checked={form.is_active} onChange={(event) => setForm({ ...form, is_active: event.target.checked })} />
            Active account
          </label>
        </div>
        <div className="flex gap-3">
          <button type="submit" className="inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700">
            <UserPlus size={16} aria-hidden="true" />
            {editingId ? 'Update user' : 'Create user'}
          </button>
          {editingId ? (
            <button type="button" onClick={() => { setEditingId(null); setForm(emptyForm); }} className="h-10 rounded-md border border-slate-300 px-4 text-sm font-semibold text-slate-700">
              Cancel
            </button>
          ) : null}
        </div>
      </form>

      {usersQuery.isLoading ? <LoadingState label="Loading users" /> : null}
      {usersQuery.isError ? <ErrorState message="Could not load users." /> : null}
      {usersQuery.data?.length === 0 ? <EmptyState icon={Users} title="No users yet" description="Create employees and admins to use OfficeBite." /> : null}

      <section className="grid gap-3">
        {usersQuery.data?.map((user) => (
          <article key={user.id} className="flex flex-col gap-3 rounded-lg border border-slate-200 bg-white p-4 shadow-sm md:flex-row md:items-center md:justify-between">
            <div>
              <div className="flex flex-wrap items-center gap-2">
                <h3 className="font-semibold text-slate-950">{user.name}</h3>
                <span className="rounded-md bg-slate-100 px-2 py-1 text-xs capitalize text-slate-600">{user.role}</span>
                <span className={user.is_active ? 'rounded-md bg-emerald-50 px-2 py-1 text-xs text-emerald-700' : 'rounded-md bg-red-50 px-2 py-1 text-xs text-red-700'}>
                  {user.is_active ? 'Active' : 'Inactive'}
                </span>
              </div>
              <p className="mt-1 text-sm text-slate-600">{user.email} {user.department ? `- ${user.department}` : ''}</p>
            </div>
            <button type="button" onClick={() => startEdit(user)} className="grid h-10 w-10 place-items-center rounded-md border border-slate-200 text-slate-600 hover:bg-slate-100" title="Edit user">
              <Pencil size={16} aria-hidden="true" />
            </button>
          </article>
        ))}
      </section>
    </div>
  );
}
