import { useAuthInit } from '@/features/auth/hooks/auth-init'
import { useAuthStore } from '@/store/auth-store'

export default function AppInitializer({
  children,
}: {
  children: React.ReactNode
}) {
  useAuthInit()
  const isInitialized = useAuthStore(state => state.isInitialized)

  if (!isInitialized) {
    return <div className='p-6'>Loading...</div>
  }

  return <>{children}</>
}
