import { Navigate, Outlet } from 'react-router';
import { useAuthStore } from '../store/authStore';

export function GuestRoute() {
  const isAuthenticated = useAuthStore((state) => state.isAuthenticated);

  if (isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
}
