import { keepPreviousData, useQuery } from '@tanstack/react-query'
import { getEvents } from '../api/events-api'

export function useEventsQuery(page: number) {
  return useQuery({
    queryKey: ['events', page],
    queryFn: () => getEvents(page, 10),
    placeholderData: keepPreviousData,
  })
}
