import axios from 'axios';
import { useAuthStore } from '../store/authStore';

const apiURL = import.meta.env.VITE_API_URL ?? (import.meta.env.DEV ? 'http://localhost:8080/api/v1' : undefined);

if (!apiURL) {
  throw new Error('VITE_API_URL is required');
}

export const api = axios.create({
  baseURL: apiURL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  const token = useAuthStore.getState().token;

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      useAuthStore.getState().clearSession();
    }

    return Promise.reject(error);
  },
);
