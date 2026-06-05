import { NavLink, Outlet, useNavigate } from 'react-router';
import { BarChart3, ClipboardList, LogOut, Menu as MenuIcon, Utensils } from 'lucide-react';
import { useAuthStore } from '../store/authStore';

const employeeLinks = [
  { label: 'Dashboard', to: '/', icon: Utensils },
  { label: 'Menu', to: '/menu', icon: MenuIcon },
  { label: 'Orders', to: '/orders', icon: ClipboardList },
];

const adminLinks = [
  { label: 'Admin', to: '/admin', icon: BarChart3 },
  { label: 'Menus', to: '/admin/menus', icon: MenuIcon },
  { label: 'All Orders', to: '/admin/orders', icon: ClipboardList },
];

export function AppLayout() {
  const navigate = useNavigate();
  const { user, clearSession } = useAuthStore();
  const links = user?.role === 'admin' ? [...employeeLinks, ...adminLinks] : employeeLinks;

  function handleLogout() {
    clearSession();
    navigate('/login', { replace: true });
  }

  return (
    <div className="min-h-screen bg-slate-50">
      <header className="border-b border-slate-200 bg-white">
        <div className="mx-auto flex max-w-6xl flex-col gap-4 px-6 py-4 lg:flex-row lg:items-center lg:justify-between">
          <div className="flex items-center gap-3">
            <span className="grid h-10 w-10 place-items-center rounded-md bg-brand-600 text-white">
              <Utensils size={20} aria-hidden="true" />
            </span>
            <div>
              <p className="text-sm font-semibold uppercase tracking-wide text-brand-700">OfficeBite</p>
              <h1 className="text-lg font-semibold text-slate-950">Meal Management</h1>
            </div>
          </div>
          <div className="flex flex-col gap-3 sm:flex-row sm:items-center">
            <nav className="flex flex-wrap gap-1" aria-label="Primary navigation">
              {links.map((link) => {
                const Icon = link.icon;

                return (
                  <NavLink
                    key={link.to}
                    to={link.to}
                    end={link.to === '/' || link.to === '/admin'}
                    className={({ isActive }) =>
                      [
                        'inline-flex h-10 items-center gap-2 rounded-md px-3 text-sm font-medium transition',
                        isActive ? 'bg-brand-50 text-brand-700' : 'text-slate-600 hover:bg-slate-100 hover:text-slate-950',
                      ].join(' ')
                    }
                  >
                    <Icon size={16} aria-hidden="true" />
                    {link.label}
                  </NavLink>
                );
              })}
            </nav>
            <div className="flex items-center gap-3 border-t border-slate-200 pt-3 sm:border-l sm:border-t-0 sm:pl-3 sm:pt-0">
              <div className="text-right">
                <p className="text-sm font-medium text-slate-950">{user?.name}</p>
                <p className="text-xs capitalize text-slate-500">{user?.role}</p>
              </div>
              <button
                type="button"
                onClick={handleLogout}
                className="grid h-10 w-10 place-items-center rounded-md border border-slate-200 text-slate-600 hover:bg-slate-100 hover:text-slate-950"
                title="Log out"
              >
                <LogOut size={16} aria-hidden="true" />
              </button>
            </div>
          </div>
        </div>
      </header>
      <main className="mx-auto max-w-6xl px-6 py-8">
        <Outlet />
      </main>
    </div>
  );
}
