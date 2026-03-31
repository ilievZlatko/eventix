import { api } from '@/lib/api'

export type EventDetails = {
  id: string
  title: string
  description: string
  location: string
  starts_at: string
  ends_at: string
  capacity: number
  booked_count: number
  is_booked: boolean
}

export async function getEvent(id: string): Promise<EventDetails> {
  const response = await api.get(`/events/${id}`)
  return response.data
}

export async function createBooking(eventId: string) {
  const res = await api.post(`/events/${eventId}/bookings`)
  return res.data
}
