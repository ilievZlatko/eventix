import { Navigate, useLocation } from 'react-router'

import { AppLayout } from '@/layouts/AppLayout'
import { useAuthStore } from '@/store/auth-store'

export const ProtectedRoute = () => {
  const token = useAuthStore(state => state.token)
  const isInitialized = useAuthStore(state => state.isInitialized)
  const location = useLocation()

  if (!isInitialized) {
    return <div className='p-6'>Loading...</div>
  }

  if (!token) {
    return (
      <Navigate
        to='/login'
        state={{ from: location }}
        replace
      />
    )
  }

  return <AppLayout />
}

