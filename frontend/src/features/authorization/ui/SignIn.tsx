import { Button, Input } from '@/shared/ui'
import { FC, ReactNode } from 'react'
import styles from './styles.module.scss'
import { useSignIn } from '../model/hooks/use-sign-in'
import { Controller, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { authSchema, TAuthSchema } from '../model/schema'

const { wrapper, auth } = styles
export const SignIn: FC = (): ReactNode => {
  const { signIn, isLoading } = useSignIn()

  const {
    control,
    handleSubmit,
    getValues,
    formState: { isDirty },
  } = useForm<TAuthSchema>({
    defaultValues: {
      login: '',
      password: '',
    },
    resolver: zodResolver(authSchema),
  })

  const onSubmit = (): void => {
    signIn({
      username: getValues('login').trim(),
      password: getValues('password'),
    })
  }

  return (
    <div className={wrapper}>
      <form className={auth} onSubmit={handleSubmit(onSubmit)}>
        <Controller
          name='login'
          control={control}
          render={({ fieldState, field }) => (
            <Input
              title='Логин'
              fluid
              value={field.value}
              onChange={field.onChange}
              isValid={!fieldState.error}
              description={fieldState.error?.message}
            />
          )}
        />
        <Controller
          name='password'
          control={control}
          render={({ fieldState, field }) => (
            <Input
              title='Пароль'
              fluid
              type='password'
              value={field.value}
              onChange={field.onChange}
              isValid={!fieldState.error}
              description={fieldState.error?.message}
            />
          )}
        />
        <Button fluid type='submit' disabled={!isDirty} isLoading={isLoading}>
          Войти
        </Button>
      </form>
    </div>
  )
}
