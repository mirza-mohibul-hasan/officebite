import { useNavigate } from 'react-router';
import { ShieldCheck, UserRound } from 'lucide-react';
import { useAuthStore } from '../store/authStore';
import type { UserRole } from '../types/auth';

export function LoginPage() {
  const navigate = useNavigate();
  const setSession = useAuthStore((state) => state.setSession);

  function continueAs(role: UserRole) {
    setSession({
      token: `dev-${role}-token`,
      user: {
        id: role === 'admin' ? 1 : 2,
        name: role === 'admin' ? 'Admin User' : 'Employee User',
        email: role === 'admin' ? 'admin@officebite.local' : 'employee@officebite.local',
        role,
      },
    });
    navigate(role === 'admin' ? '/admin' : '/', { replace: true });
  }

  return (
    <main className="grid min-h-screen place-items-center bg-slate-50 px-6">
      <section className="w-full max-w-md rounded-lg border border-slate-200 bg-white p-6 shadow-sm">
        <p className="text-sm font-medium uppercase tracking-wide text-brand-700">OfficeBite</p>
        <h1 className="mt-2 text-2xl font-semibold text-slate-950">Sign in</h1>
        <p className="mt-3 text-sm leading-6 text-slate-600">
          Authentication UI will be connected in the auth branch. These temporary actions keep the protected shell
          testable during frontend foundation work.
        </p>
        <div className="mt-6 grid gap-3">
          <button
            type="button"
            onClick={() => continueAs('employee')}
            className="inline-flex h-11 items-center justify-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700"
          >
            <UserRound size={16} aria-hidden="true" />
            Continue as employee
          </button>
          <button
            type="button"
            onClick={() => continueAs('admin')}
            className="inline-flex h-11 items-center justify-center gap-2 rounded-md border border-slate-300 px-4 text-sm font-semibold text-slate-700 hover:bg-slate-100"
          >
            <ShieldCheck size={16} aria-hidden="true" />
            Continue as admin
          </button>
        </div>
      </section>
    </main>
  );
}
