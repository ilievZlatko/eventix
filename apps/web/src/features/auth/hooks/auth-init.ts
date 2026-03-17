import { useEffect } from 'react'
import { useAuthStore } from '@/store/auth-store'
import { useMeQuery } from './use-me-query'

export function useAuthInit() {
  const token = useAuthStore(state => state.token)
  const setUser = useAuthStore(state => state.setUser)
  const clearAuth = useAuthStore(state => state.clearAuth)
  const setIsInitialized = useAuthStore(state => state.setIsInitialized)

  const { data, isError, isSuccess } = useMeQuery(token)

  useEffect(() => {
    if (!token) {
      setIsInitialized()
      return
    }

    if (isSuccess && data) {
      setUser(data)
      setIsInitialized()
      return
    }

    if (isError) {
      clearAuth()
      setIsInitialized()
    }
  }, [token, isSuccess, isError, data, clearAuth, setIsInitialized, setUser])

  useEffect(() => {
    function handleStorage(event: StorageEvent) {
      if (event.key === 'token' && event.newValue === null) {
        clearAuth()
        globalThis.location.href = '/login'
      }
    }

    globalThis.addEventListener('storage', handleStorage)
    return () => globalThis.removeEventListener('storage', handleStorage)
  }, [clearAuth])
}
