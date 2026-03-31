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
  const isFull = event.booked_count >= event.capacity
  const isBooked = event.is_booked

  function getButtonLabel() {
    if (bookingMutation.isPending) return 'Booking...'
    if (isBooked) return 'Already booked'
    if (isFull) return 'Event is full'
    return 'Book Event'
  }

  function handleBook() {
    if (isFull) {
      toast.error('Event is full')
      return
    }

    if (isBooked) {
      toast.error('You have already booked this event')
      return
    }

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
    <div className='max-w-2xl space-y-6'>
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

          <div className='text-sm font-medium'>
            {event.booked_count} / {event.capacity} spots taken
          </div>

          {isBooked && (
            <div className='text-sm text-green-600'>
              You have already booked this event
            </div>
          )}

          {isFull && <div className='text-sm text-red-600'>Event is full</div>}

          <Button
            onClick={handleBook}
            disabled={bookingMutation.isPending}
          >
            {getButtonLabel()}
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}
