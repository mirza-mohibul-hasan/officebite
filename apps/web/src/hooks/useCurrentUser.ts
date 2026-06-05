import { useAuthStore } from '../store/authStore';

export function useCurrentUser() {
  const user = useAuthStore((state) => state.user);

  if (!user) {
    throw new Error('useCurrentUser must be used inside an authenticated route');
  }

  return user;
}
