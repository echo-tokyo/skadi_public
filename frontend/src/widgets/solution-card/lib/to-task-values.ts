import { TSolution } from '@/shared/model'
import { TDisplayValues } from '../model/types'
import { parseGrade } from '@/entities/solution'

export const toTaskValues = (
  solutionData: TSolution | undefined,
): TDisplayValues => ({
  description: solutionData?.task.description ?? '',
  student: solutionData?.student?.fullname ?? '',
  title: solutionData?.task.title ?? '',
  answer: solutionData?.answer ?? '',
  teacher: solutionData?.task.teacher?.fullname ?? 'Преподаватель',
  files: solutionData?.task.files ?? [],
  file_answer: solutionData?.files ?? [],
  grade: parseGrade(solutionData?.grade),
})
