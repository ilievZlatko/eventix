import { Outlet, useNavigate } from 'react-router'
import { useAuthStore } from '@/store/auth-store'
import { Button } from '@/components/ui/button'
import { toast } from 'sonner'

export const AppLayout = () => {
  const user = useAuthStore(state => state.user)
  const clearAuth = useAuthStore(state => state.clearAuth)

  const navigate = useNavigate()

  function handleLogout() {
    clearAuth()
    toast.success('Logged out successfully')
    navigate('/login')
  }

  return (
    <div className='min-h-screen bg-gray-50'>
      <header className='border-b bg-white'>
        <div className='mx-auto flex max-w-6xl items-center justify-between px-6 py-4'>
          <div className='font-semibold'>Eventix</div>

          <div className='flex items-center gap-4'>
            <span className='text-sm text-muted-foreground'>{user?.email}</span>

            <Button
              variant='outline'
              onClick={handleLogout}
            >
              Logout
            </Button>
          </div>
        </div>
      </header>

      <main className='mx-auto max-w-6xl px-6 py-8'>
        <Outlet />
      </main>
    </div>
  )
}
