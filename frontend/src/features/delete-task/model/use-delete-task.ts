import { useDeleteTaskMutation } from '@/entities/task'
import { useMutationAction } from '@/shared/lib'

export const useDeleteTask = () => {
  const { isLoading, submit: deleteTaskById } = useMutationAction<
    number,
    number
  >({
    mutation: useDeleteTaskMutation(),
    prepare: (id) => id,
    successMessage: 'Задание удалено',
  })

  return { deleteTaskById, isLoading }
}
