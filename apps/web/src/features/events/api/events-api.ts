import { api } from '@/lib/api'

export type Event = {
  id: string
  title: string
  description: string
  location: string
  starts_at: string
}

export type EventsResponse = {
  data: Event[]
  meta: {
    page: number
    limit: number
    total: number
    total_pages: number
  }
}

export async function getEvents(page = 1, limit = 10): Promise<EventsResponse> {
  const response = await api.get(`/events?page=${page}&limit=${limit}`)
  return response.data
}
