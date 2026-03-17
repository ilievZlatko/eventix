import { useQuery } from '@tanstack/react-query'
import { getMe } from '@/features/auth/api/auth-api'

export function useMeQuery(token: string | null) {
  return useQuery({
    queryKey: ['me'],
    queryFn: () => getMe(token!),
    enabled: !!token,
    retry: false,
  })
}
