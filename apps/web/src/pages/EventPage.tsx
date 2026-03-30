import { useParams } from 'react-router'
import { toast } from 'sonner'

import { useEventQuery } from '@/features/events/hooks/use-event-query'
import { useCreateBooking } from '@/features/events/hooks/use-create-booking'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export const EventPage = () => {
  const { id } = useParams()
  const { data, isLoading, isError } = useEventQuery(id!)
  const bookingMutation = useCreateBooking(id!)

  if (isLoading) return <div className='p-6'>Loading...</div>
  if (isError || !data) return <div className='p-6'>Failed to load event</div>

  const event = data

  function handleBook() {
    bookingMutation.mutate(undefined, {
      onSuccess: () => {
        toast.success('Booking successful')
      },
      onError: () => {
        toast.error('Failed to book event')
      },
    })
  }

  return (
    <div className='max-w-2xl'>
      <Card>
        <CardHeader>
          <CardTitle>{event.title}</CardTitle>
        </CardHeader>

        <CardContent className='space-y-4'>
          <p>{event.description}</p>

          <div className='text-sm text-muted-foreground'>
            📍 {event.location}
          </div>

          <div className='text-sm'>
            🕓 {new Date(event.starts_at).toLocaleDateString()}
          </div>

          <Button
            onClick={handleBook}
            disabled={bookingMutation.isPending}
          >
            {bookingMutation.isPending ? 'Booking...' : 'Book Event'}
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}
