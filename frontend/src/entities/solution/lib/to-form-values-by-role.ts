import { TRole, TSolution, TStatusId } from '@/shared/model'
import {
  VALID_STATUSES,
  TSolutionTeacherSchema,
  TSolutionStudentSchema,
} from '../model/schemas'
import { parseGrade } from './parse-grade'

const DEFAULT_STATUS = VALID_STATUSES[0]

const resolveStatus = (
  raw: TStatusId | undefined,
): (typeof VALID_STATUSES)[number] =>
  raw !== undefined && VALID_STATUSES.includes(raw) ? raw : DEFAULT_STATUS

const toTeacherFormValues = (
  solutionData: TSolution | undefined,
): TSolutionTeacherSchema => ({
  status: resolveStatus(solutionData?.status.id),
  grade: parseGrade(solutionData?.grade),
})

const toStudentFormValues = (
  solutionData: TSolution | undefined,
): TSolutionStudentSchema => ({
  status: resolveStatus(solutionData?.status.id),
  answer: solutionData?.answer ?? '',
  file_answer: [],
  deleted_file_ids: [],
})

export const toFormValuesByRole = (
  solutionData: TSolution | undefined,
  role: TRole,
): TSolutionTeacherSchema | TSolutionStudentSchema => {
  if (role === 'teacher') {
    return toTeacherFormValues(solutionData)
  }
  if (role === 'student') {
    return toStudentFormValues(solutionData)
  }
  throw new Error(`Unexpected role ${role}`)
}
