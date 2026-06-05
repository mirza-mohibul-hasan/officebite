import { FormEvent, useState } from 'react';
import { useMutation } from '@tanstack/react-query';
import { useLocation, useNavigate } from 'react-router';
import { LogIn } from 'lucide-react';
import { ErrorState } from '../components/ErrorState';
import { getApiErrorMessage } from '../services/errors';
import { login } from '../services/auth';
import { useAuthStore } from '../store/authStore';

export function LoginPage() {
  const navigate = useNavigate();
  const location = useLocation();
  const setSession = useAuthStore((state) => state.setSession);
  const [email, setEmail] = useState('employee@officebite.local');
  const [password, setPassword] = useState('password123');

  const loginMutation = useMutation({
    mutationFn: login,
    onSuccess: (session) => {
      setSession(session);
      const from = location.state?.from?.pathname;
      navigate(from ?? (session.user.role === 'admin' ? '/admin' : '/'), { replace: true });
    },
  });

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    loginMutation.mutate({ email, password });
  }

  return (
    <main className="grid min-h-screen place-items-center bg-slate-50 px-6">
      <section className="w-full max-w-md rounded-lg border border-slate-200 bg-white p-6 shadow-sm">
        <p className="text-sm font-medium uppercase tracking-wide text-brand-700">OfficeBite</p>
        <h1 className="mt-2 text-2xl font-semibold text-slate-950">Sign in</h1>
        <p className="mt-3 text-sm leading-6 text-slate-600">Use your office meal account to continue.</p>

        <form className="mt-6 grid gap-4" onSubmit={handleSubmit}>
          {loginMutation.isError ? (
            <ErrorState message={getApiErrorMessage(loginMutation.error, 'Check your email and password, then try again.')} />
          ) : null}
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Email
            <input
              type="email"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              className="h-11 rounded-md border border-slate-300 px-3 text-slate-950 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              autoComplete="email"
              required
            />
          </label>
          <label className="grid gap-2 text-sm font-medium text-slate-700">
            Password
            <input
              type="password"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              className="h-11 rounded-md border border-slate-300 px-3 text-slate-950 outline-none focus:border-brand-600 focus:ring-2 focus:ring-brand-100"
              autoComplete="current-password"
              required
            />
          </label>
          <button
            type="submit"
            disabled={loginMutation.isPending}
            className="inline-flex h-11 items-center justify-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700 disabled:cursor-not-allowed disabled:opacity-70"
          >
            <LogIn size={16} aria-hidden="true" />
            {loginMutation.isPending ? 'Signing in...' : 'Sign in'}
          </button>
        </form>

        <div className="mt-5 rounded-md bg-slate-50 p-3 text-xs leading-5 text-slate-600">
          Seeded accounts: employee@officebite.local / password123, admin@officebite.local / password123.
        </div>
      </section>
    </main>
  );
}
