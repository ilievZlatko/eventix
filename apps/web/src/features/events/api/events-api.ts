import { api } from '@/lib/api'
import type { EventDetails } from './event-api'

export type EventsResponse = {
  data: EventDetails[]
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
