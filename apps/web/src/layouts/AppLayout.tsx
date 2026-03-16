import { Outlet } from 'react-router'

export const AppLayout = () => {
  return (
    <div className='min-h-screen bg-gray-50'>
      <header className='border-b bg-white'>
        <div className='mx-auto max-w-6xl px-6 py-4 font-semibold'>Eventix</div>
      </header>

      <main className='mx-auto max-w-6xl px-6 py-8'>
        <Outlet />
      </main>
    </div>
  )
}
