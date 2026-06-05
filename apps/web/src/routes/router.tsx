import { createBrowserRouter } from 'react-router';
import { AppLayout } from '../layouts/AppLayout';
import { AdminDashboardPage } from '../pages/AdminDashboardPage';
import { AdminMenusPage } from '../pages/AdminMenusPage';
import { AdminOrdersPage } from '../pages/AdminOrdersPage';
import { DashboardPage } from '../pages/DashboardPage';
import { MenuPage } from '../pages/MenuPage';
import { OrderHistoryPage } from '../pages/OrderHistoryPage';
import { LoginPage } from '../pages/LoginPage';
import { GuestRoute } from './GuestRoute';
import { ProtectedRoute } from './ProtectedRoute';

export const router = createBrowserRouter([
  {
    element: <ProtectedRoute />,
    children: [
      {
        path: '/',
        element: <AppLayout />,
        children: [
          {
            index: true,
            element: <DashboardPage />,
          },
          {
            path: 'menu',
            element: <MenuPage />,
          },
          {
            path: 'orders',
            element: <OrderHistoryPage />,
          },
        ],
      },
    ],
  },
  {
    element: <ProtectedRoute allowedRoles={['admin']} />,
    children: [
      {
        path: '/admin',
        element: <AppLayout />,
        children: [
          {
            index: true,
            element: <AdminDashboardPage />,
          },
          {
            path: 'menus',
            element: <AdminMenusPage />,
          },
          {
            path: 'orders',
            element: <AdminOrdersPage />,
          },
        ],
      },
    ],
  },
  {
    element: <GuestRoute />,
    children: [
      {
        path: '/login',
        element: <LoginPage />,
      },
    ],
  },
]);
