import { TTaskSchemaCreate } from '../model/schema'

export const defaultValues: TTaskSchemaCreate = {
  title: '',
  description: '',
  students: [],
  classes: [],
  files: [],
  deletedFileIds: [],
}
