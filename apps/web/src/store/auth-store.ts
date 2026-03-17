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
  isInitialized: boolean

  setAuth: (token: string, user: AuthUser) => void
  setUser: (user: AuthUser) => void
  setIsInitialized: () => void
  clearAuth: () => void
}

export const useAuthStore = create<AuthState>(set => ({
  token: localStorage.getItem('token'),
  user: null,
  isInitialized: false,

  setAuth: (token, user) => {
    localStorage.setItem('token', token)
    set({ token, user })
  },

  setUser: user => {
    set({ user })
  },

  setIsInitialized: () => {
    set({ isInitialized: true })
  },

  clearAuth: () => {
    localStorage.removeItem('token')
    set({ user: null, token: null })
  },
}))
