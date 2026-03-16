import type { Role } from '@/features/auth/api/auth-api'
import { create } from 'zustand'

type AuthUser = {
  id: string
  email: string
  role: Role
}

type AuthState = {
  token: string | null
  user: AuthUser | null
  setAuth: (token: string, user: AuthUser) => void
  clearAuth: () => void
}

export const useAuthStore = create<AuthState>(set => ({
  token: localStorage.getItem('token'),
  user: null,

  setAuth: (token, user) => {
    localStorage.setItem('token', token)
    set({ token, user })
  },

  clearAuth: () => {
    localStorage.removeItem('token')
    set({ user: null, token: null })
  },
}))
