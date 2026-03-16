import { Outlet } from 'react-router'

export const AuthLayout = () => {
  return (
    <div className='p-10'>
      <Outlet />
    </div>
  )
}
