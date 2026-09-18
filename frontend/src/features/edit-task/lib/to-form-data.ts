import { TTaskSchemaUpdate } from '@/entities/task'
import { TTaskWithStudents } from '@/shared/model'

export const toFormData = (task: TTaskWithStudents): TTaskSchemaUpdate => ({
  title: task.task.title,
  description: task.task.description ?? '',
  students: task.students?.map((el) => el.id.toString()) ?? [],
  files: [],
  deletedFileIds: [],
})
