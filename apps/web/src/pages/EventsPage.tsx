import { useState } from 'react'
import { Link } from 'react-router'
import { useEventsQuery } from '@/features/events/hooks/use-events-query'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'

export const EventsPage = () => {
  const [page, setPage] = useState(1)
  const { data, isLoading, error } = useEventsQuery(page)

  if (isLoading) return <div className='p-6'>Loading Events...</div>

  if (error || !data) return <div className='p-6'>Failed to load events</div>

  return (
    <div className='space-y-6'>
      <div className='grid gap-4 md:grid-cols-2'>
        {data.data.map(event => (
          <Link
            key={event.id}
            to={`/events/${event.id}`}
          >
            <Card>
              <CardHeader>
                <CardTitle>{event.title}</CardTitle>
              </CardHeader>

              <CardContent>
                <p className='text-sm text-muted-foreground'>
                  {event.location}
                </p>
                <p className='text-sm'>
                  {new Date(event.starts_at).toLocaleDateString()}
                </p>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>

      <div className='flex items-center gap-2'>
        <Button
          disabled={page === 1}
          onClick={() => setPage(page - 1)}
        >
          Previous
        </Button>
        <Button
          disabled={page === data.meta.total_pages}
          onClick={() => setPage(page + 1)}
        >
          Next
        </Button>

        <span className='text-sm text-muted-foreground ml-auto'>
          Page {data.meta.page} of {data.meta.total_pages} ({data.meta.total}{' '}
          events)
        </span>
      </div>
    </div>
  )
}
