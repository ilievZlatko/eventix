import { useMutation, useQueryClient } from '@tanstack/react-query'
import { createBooking } from '../api/event-api'

export function useCreateBooking(eventId: string) {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => createBooking(eventId),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['event', eventId] })
    },
  })
}
