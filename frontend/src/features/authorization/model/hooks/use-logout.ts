import { useAppDispatch } from '@/shared/lib'
import { useLogoutMutation } from '../../api/auth-api'
import { clearUserData } from '@/entities/user'
import { toast } from 'sonner'
import { getErrorMessage } from '@/shared/api'
import { baseApi } from '@/shared/api'
import { useNavigate } from 'react-router'

export const useLogout = () => {
  const [logoutMutation, { isLoading }] = useLogoutMutation()
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const logout = async (): Promise<void> => {
    try {
      await logoutMutation().unwrap()
      toast.info('Выход из аккаунта')
    } catch (err) {
      toast.error(getErrorMessage(err))
    } finally {
      dispatch(clearUserData())
      dispatch(baseApi.util.resetApiState())
      navigate('/authorization', { replace: true })
    }
  }

  return { logout, isLoading }
}
