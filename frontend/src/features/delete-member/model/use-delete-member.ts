import { useDeleteMemberMutation } from '@/entities/member'
import { useMutationAction } from '@/shared/lib'

export const useDeleteMember = () => {
  const { submit: deleteMemberById, isLoading } = useMutationAction<number, number>({
    mutation: useDeleteMemberMutation(),
    prepare: (id) => id,
    successMessage: 'Пользователь удалён',
  })

  return { deleteMemberById, isLoading }
}
