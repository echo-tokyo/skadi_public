import { useSignInMutation } from '../../api/auth-api'
import { useNavigate } from 'react-router'
import { setUserData } from '@/entities/user'
import { useAppDispatch } from '@/shared/lib'
import { ISignInRequest } from '../types'
import { toast } from 'sonner'
import { getErrorMessage } from '@/shared/api'

export const useSignIn = () => {
  const [signInMutation, { isLoading }] = useSignInMutation()
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const signIn = async (formData: ISignInRequest): Promise<void> => {
    try {
      const result = await signInMutation(formData).unwrap()

      dispatch(setUserData(result))

      navigate('/personal-area')
    } catch (err) {
      toast.error(getErrorMessage(err))
    }
  }

  return { signIn, isLoading }
}
