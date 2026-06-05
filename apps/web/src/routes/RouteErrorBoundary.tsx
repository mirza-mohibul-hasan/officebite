import { isRouteErrorResponse, Link, useRouteError } from 'react-router';
import { AlertTriangle, Home } from 'lucide-react';

export function RouteErrorBoundary() {
  const error = useRouteError();
  const message = isRouteErrorResponse(error)
    ? `${error.status} ${error.statusText}`
    : 'The page could not be loaded.';

  return (
    <main className="grid min-h-screen place-items-center bg-slate-50 px-6">
      <section className="w-full max-w-lg rounded-lg border border-slate-200 bg-white p-6 text-center shadow-sm">
        <AlertTriangle className="mx-auto text-red-500" size={32} aria-hidden="true" />
        <h1 className="mt-4 text-2xl font-semibold text-slate-950">Something went wrong</h1>
        <p className="mt-3 text-sm leading-6 text-slate-600">{message}</p>
        <Link
          to="/"
          className="mt-6 inline-flex h-10 items-center gap-2 rounded-md bg-brand-600 px-4 text-sm font-semibold text-white hover:bg-brand-700"
        >
          <Home size={16} aria-hidden="true" />
          Back to dashboard
        </Link>
      </section>
    </main>
  );
}
