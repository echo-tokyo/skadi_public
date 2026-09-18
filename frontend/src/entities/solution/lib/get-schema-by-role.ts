import { TRole } from '@/shared/model'
import { solutionTeacherSchema, solutionStudentSchema } from '../model/schemas'

export const getSchemaByRole = (role: TRole) => {
  if (role === 'teacher') return solutionTeacherSchema
  if (role === 'student') return solutionStudentSchema

  throw new Error(`Unexpected role ${role}`)
}
