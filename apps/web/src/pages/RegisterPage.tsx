import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { useRegisterMutation } from '@/features/auth/hooks/use-auth-mutations'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { toast } from 'sonner'
import { Link, useNavigate } from 'react-router'

const registerSchema = z
  .object({
    email: z.email('Please enter a valid email.'),
    password: z.string().min(8, 'Password must be at least 8 characters.'),
    confirmPassword: z
      .string()
      .min(8, 'Password must be at least 8 characters.'),
    role: z.enum(['user', 'organizer']),
  })
  .refine(data => data.password === data.confirmPassword, {
    path: ['confirmPassword'],
    message: 'Passwords do not match.',
  })

type RegisterFormValues = z.infer<typeof registerSchema>

export const RegisterPage = () => {
  const navigate = useNavigate()

  const registerMutation = useRegisterMutation({
    onSuccess: () => {
      toast.success('Registration successful. You can now log in.')
      navigate('/login')
    },
    onError: () => {
      toast.error('Registration failed. Please try again.')
    },
  })

  const form = useForm<RegisterFormValues>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: '',
      password: '',
      confirmPassword: '',
      role: 'user',
    },
  })

  const onSubmit = (values: RegisterFormValues) => {
    registerMutation.mutate(values)
  }

  return (
    <div className='flex min-h-screen items-center justify-center px-4'>
      <Card className='w-full max-w-md'>
        <CardHeader>
          <CardTitle>Create account</CardTitle>
        </CardHeader>

        <CardContent>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className='space-y-4'
          >
            <div className='space-y-2'>
              <Label htmlFor='email'>Email</Label>
              <Input
                id='email'
                type='email'
                {...form.register('email')}
              />
              {form.formState.errors.email ? (
                <p className='text-sm text-red-500'>
                  {form.formState.errors.email.message}
                </p>
              ) : null}
            </div>

            <div className='space-y-2'>
              <Label htmlFor='password'>Password</Label>
              <Input
                id='password'
                type='password'
                {...form.register('password')}
              />
              {form.formState.errors.password ? (
                <p className='text-sm text-red-500'>
                  {form.formState.errors.password.message}
                </p>
              ) : null}
            </div>

            <div className='space-y-2'>
              <Label htmlFor='confirmPassword'>Confirm Password</Label>
              <Input
                id='confirmPassword'
                type='password'
                {...form.register('confirmPassword')}
              />
              {form.formState.errors.confirmPassword ? (
                <p className='text-sm text-red-500'>
                  {form.formState.errors.confirmPassword.message}
                </p>
              ) : null}
            </div>

            <div className='space-y-2'>
              <Label htmlFor='role'>Role</Label>
              <select
                id='role'
                {...form.register('role')}
                className='flex h-10 w-full rounded-md border bg-background px-3 py-2 text-sm'
              >
                <option value='user'>User</option>
                <option value='organizer'>Organizer</option>
              </select>
            </div>

            {registerMutation.isError ? (
              <p className='text-sm text-red-500'>
                Registration failed. Please try again.
              </p>
            ) : null}

            <Button
              type='submit'
              className='w-full'
              disabled={registerMutation.isPending}
            >
              {registerMutation.isPending ? 'Creating account...' : 'Register'}
            </Button>

            <div className='text-center text-sm text-muted-foreground'>
              Already have an account?{' '}
              <Link
                to='/login'
                className='text-primary underline'
              >
                Login
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
