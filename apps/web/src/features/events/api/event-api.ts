import { api } from '@/lib/api'
import type { Event } from './events-api'

export async function getEvent(id: string): Promise<Event> {
  const response = await api.get(`/events/${id}`)
  return response.data
}

export async function createBooking(eventId: string) {
  const res = await api.post(`/events/${eventId}/bookings`)
  return res.data
}
