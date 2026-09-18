export {
  useGetSolutionByIdQuery,
  useGetSolutionsInfiniteQuery,
  useGetStudentSolutionsInfiniteQuery,
  useUpdateSolutionByTeacherMutation,
  useUpdateSolutionByStudentMutation,
  useGetSolutionsForStudentQuery,
  useDeleteSolutionMutation,
} from './api/solution-api'

export type { IGetSolutionsQuery } from './model/types'

export type {
  TSolutionTeacherSchema,
  TSolutionStudentSchema,
  TSolutionBaseSchema,
} from './model/schemas'

export type {
  IUpdateSolutionByTeacherRequest,
  IUpdateSolutionByStudentRequest,
} from './model/types'

export {
  solutionTeacherSchema,
  solutionStudentSchema,
  VALID_STATUSES,
} from './model/schemas'

export { getSchemaByRole } from './lib/get-schema-by-role'
export { toFormValuesByRole } from './lib/to-form-values-by-role'
export { useGradeLabel } from './model/use-grade-label'
export { parseGrade } from './lib/parse-grade'
