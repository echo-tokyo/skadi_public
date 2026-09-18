import {
  IUpdateSolutionByStudentRequest,
  TSolutionStudentSchema,
  useUpdateSolutionByStudentMutation,
} from '@/entities/solution'
import { useMutationAction } from '@/shared/lib'
import { TStatusId } from '@/shared/model'

const prepare =
  (id: number) =>
  (data: TSolutionStudentSchema): { id: number; data: FormData } => {
    const body: IUpdateSolutionByStudentRequest = {
      status_id: Number(data.status) as TStatusId,
      answer: data.answer,
      delete_files: data.deleted_file_ids,
      file: data.file_answer,
    }

    const fd = new FormData()
    fd.append('status_id', String(body.status_id))
    fd.append('answer', String(body.answer))

    body.delete_files?.forEach((el) => fd.append('delete_files', String(el)))
    body.file?.forEach((el) => fd.append('file', el))

    return { id, data: fd }
  }

export const useStudentUpdateSolution = (id: number) => {
  return useMutationAction({
    mutation: useUpdateSolutionByStudentMutation(),
    prepare: prepare(id),
    successMessage: 'Решение обновлено',
  })
}
