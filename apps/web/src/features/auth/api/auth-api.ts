import { api } from '@/lib/api'

export type Role = 'user' | 'organizer'

export type LoginRequest = {
  email: string
  password: string
}

export type LoginResponse = {
  access_token: string
}

export type RegisterRequest = {
  email: string
  password: string
  role: Role
}

export type RegisterResponse = {
  message: string
}

export type MeResponse = {
  id: string
  email: string
  role: Role
}

export async function login(payload: LoginRequest): Promise<LoginResponse> {
  const response = await api.post('/auth/login', payload)
  return response.data
}

export async function register(
  payload: RegisterRequest,
): Promise<RegisterResponse> {
  const response = await api.post('/auth/register', payload)
  return response.data
}

export async function getMe(token: string): Promise<MeResponse> {
  const response = await api.get('/me', {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })
  return response.data
}
