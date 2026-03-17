import { useMutation } from '@tanstack/react-query'
import {
  getMe,
  login,
  register,
  type LoginRequest,
  type RegisterRequest,
} from '@/features/auth/api/auth-api'
import { useAuthStore } from '@/store/auth-store'
import { useNavigate } from 'react-router'

export function useLoginMutation({
  onError = () => {},
}: {
  onError?: () => void
}) {
  const navigate = useNavigate()
  const setAuth = useAuthStore(state => state.setAuth)
  return useMutation({
    mutationFn: async (payload: LoginRequest) => {
      const loginResponse = await login(payload)
      const user = await getMe(loginResponse.access_token)
      return {
        access_token: loginResponse.access_token,
        user,
      }
    },
    onSuccess: ({ access_token, user }) => {
      setAuth(access_token, user)
      navigate('/')
    },
    onError,
  })
}

export function useRegisterMutation({
  onSuccess = () => {},
  onError = () => {},
}: {
  onSuccess?: () => void
  onError?: () => void
}) {
  return useMutation({
    mutationFn: async (payload: RegisterRequest) => register(payload),
    onSuccess,
    onError,
  })
}
