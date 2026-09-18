import { useUpdateSolutionByStudentMutation } from '@/entities/solution'
import { useMutationAction } from '@/shared/lib'
import { TStatusId } from '@/shared/model'

// FIXME: status должен быть статусом для ученика: 1 | 2 | 3
type TInput = { id: number; status: TStatusId }

const prepare = ({ id, status }: TInput): { id: number; data: FormData } => {
  const fd = new FormData()
  fd.append('status_id', String(Number(status)))
  return { id, data: fd }
}

export const useStudentUpdateSolutionStatus = () => {
  return useMutationAction({
    mutation: useUpdateSolutionByStudentMutation(),
    prepare,
  })
}
