import { TMemberFormData, useCreateMemberMutation } from '@/entities/member'
import { transformToRequest } from '../lib/transform-to-request'
import { useMutationAction } from '@/shared/lib'

export const useCreateMember = () => {
  return useMutationAction<
    TMemberFormData,
    ReturnType<typeof transformToRequest>
  >({
    mutation: useCreateMemberMutation(),
    prepare: transformToRequest,
    successMessage: 'Пользователь создан',
  })
}
