import {
  IUpdateSolutionByTeacherRequest,
  TSolutionTeacherSchema,
  useUpdateSolutionByTeacherMutation,
} from '@/entities/solution'
import { CHECKED_STATUS_ID } from '@/shared/config'
import { useMutationAction } from '@/shared/lib'

export const useTeacherUpdateSolution = (id: number) => {
  return useMutationAction({
    mutation: useUpdateSolutionByTeacherMutation(),
    prepare: (
      data: TSolutionTeacherSchema,
    ): { id: number; data: IUpdateSolutionByTeacherRequest } => ({
      id,
      data: {
        status_id: data.status,
        grade: data.status === CHECKED_STATUS_ID ? data.grade : undefined,
      },
    }),
    successMessage: 'Решение обновлено',
  })
}
