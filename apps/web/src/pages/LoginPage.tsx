import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { useLoginMutation } from '@/features/auth/hooks/use-auth-mutations'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { toast } from 'sonner'
import { Link } from 'react-router'

const loginSchema = z.object({
  email: z.email('Please enter a valid email.'),
  password: z.string().min(8, 'Password must be at least 8 characters.'),
})

type LoginFormValues = z.infer<typeof loginSchema>

export const LoginPage = () => {
  const loginMutation = useLoginMutation({
    onError: () => {
      toast.error('Login failed. Please try again.')
    },
  })

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const onSubmit = (values: LoginFormValues) => {
    loginMutation.mutate(values)
  }

  return (
    <div className='flex min-h-screen items-center justify-center px-4'>
      <Card className='w-full max-w-md'>
        <CardHeader>
          <CardTitle>Login</CardTitle>
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

            {loginMutation.isError ? (
              <p className='text-sm text-red-500'>
                Login failed. Please try again.
              </p>
            ) : null}

            <Button
              type='submit'
              className='w-full'
              disabled={loginMutation.isPending}
            >
              {loginMutation.isPending ? 'Logging in...' : 'Login'}
            </Button>

            <div className='text-center text-sm text-muted-foreground'>
              Don't have an account?{' '}
              <Link
                to='/register'
                className='text-primary underline'
              >
                Register
              </Link>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  )
}
