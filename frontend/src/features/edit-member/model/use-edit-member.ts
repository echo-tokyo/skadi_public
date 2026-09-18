import { TMemberFormData, useUpdateMemberMutation } from '@/entities/member'
import { transformToUpdateRequest } from '../lib/transform-to-update-request'
import { useMutationAction } from '@/shared/lib'

export const useEditMember = (id: number) => {
  return useMutationAction<TMemberFormData, { id: number; data: ReturnType<typeof transformToUpdateRequest> }>({
    mutation: useUpdateMemberMutation(),
    prepare: (formData) => ({ id, data: transformToUpdateRequest(formData) }),
    successMessage: 'Пользователь обновлён',
  })
}
