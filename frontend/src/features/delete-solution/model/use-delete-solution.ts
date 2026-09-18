import { useDeleteSolutionMutation } from '@/entities/solution'
import { useMutationAction } from '@/shared/lib'

export const useDeleteSolution = () => {
  const { isLoading, submit: deleteSolutionById } = useMutationAction<
    number,
    number
  >({
    mutation: useDeleteSolutionMutation(),
    prepare: (id) => id,
    successMessage: 'Решение удалено',
  })

  return { deleteSolutionById, isLoading }
}
