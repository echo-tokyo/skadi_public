import {
  ICreateTaskRequest,
  TTaskSchemaCreate,
  useCreateTaskMutation,
} from '@/entities/task'
import { useMutationAction } from '@/shared/lib'

const prepare = (data: TTaskSchemaCreate): FormData => {
  const body: ICreateTaskRequest = {
    title: data.title,
    description: data.description,
    students: data.students.map(Number),
    classes: data.classes.map(Number),
    file: data.files,
  }

  const formData = new FormData()

  formData.append('title', body.title)
  formData.append('description', body.description)

  body.students?.forEach((id) => formData.append('students', String(id)))
  body.classes?.forEach((id) => formData.append('classes', String(id)))
  body.file?.forEach((file) => formData.append('file', file))

  return formData
}

export const useCreateTask = () => {
  return useMutationAction({
    mutation: useCreateTaskMutation(),
    prepare,
    successMessage: 'Задание создано',
  })
}
