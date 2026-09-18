import { useDeleteClassMutation } from '@/entities/class'
import { useMutationAction } from '@/shared/lib'

export const useDeleteClass = () => {
  const { submit: deleteClassById, isLoading } = useMutationAction<number, number>({
    mutation: useDeleteClassMutation(),
    prepare: (id) => id,
    successMessage: 'Группа удалена',
  })

  return { deleteClassById, isLoading }
}
