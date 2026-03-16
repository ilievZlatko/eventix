import { createBrowserRouter } from 'react-router'

import { AppLayout } from '@/layouts/AppLayout'
import { AuthLayout } from '@/layouts/AuthLayout'

import { LoginPage } from '@/pages/LoginPage'
import { RegisterPage } from '@/pages/RegisterPage'
import { EventsPage } from '@/pages/EventsPage'
import { EventPage } from '@/pages/EventPage'

export const router = createBrowserRouter([
  {
    element: <AuthLayout />,
    children: [
      {
        path: '/login',
        element: <LoginPage />,
      },
      {
        path: '/register',
        element: <RegisterPage />,
      },
    ],
  },
  {
    element: <AppLayout />,
    children: [
      {
        path: '/',
        element: <EventsPage />,
      },
      {
        path: '/events/:id',
        element: <EventPage />,
      },
    ],
  },
])
