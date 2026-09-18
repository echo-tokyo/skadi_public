import {
  IUpdateTaskRequest,
  TTaskSchemaUpdate,
  useUpdateTaskMutation,
} from '@/entities/task'
import { useMutationAction } from '@/shared/lib'

const prepare =
  (id: number) =>
  (data: TTaskSchemaUpdate): FormData => {
    const body: IUpdateTaskRequest = {
      id,
      title: data.title,
      description: data.description,
      students: data.students.map(Number),
      delete_files: data.deletedFileIds,
      file: data.files,
    }

    const fd = new FormData()

    fd.append('id', body.id.toString())
    fd.append('title', body.title)
    fd.append('description', body.description)

    body.students?.forEach((el) => fd.append('students', String(el)))
    body.delete_files?.forEach((el) => fd.append('delete_files', String(el)))
    body.file?.forEach((el) => fd.append('file', el))

    return fd
  }

export const useEditTask = (id: number) => {
  return useMutationAction({
    mutation: useUpdateTaskMutation(),
    prepare: prepare(id),
    successMessage: 'Задание обновлено',
  })
}
