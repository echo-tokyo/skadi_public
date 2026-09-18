import { CHECKING_STATUS_ID, GRADE_OPTIONS } from '@/shared/config'
import { z } from 'zod'

export const VALID_STATUSES = [1, 2, 3, 4] as const
const GRADE_VALUES = GRADE_OPTIONS.map((o) => o.value)

export const solutionTeacherSchema = z.object({
  status: z.literal(VALID_STATUSES, 'Обязательное поле'),
  grade: z.enum(GRADE_VALUES).or(z.literal('')),
})
export const solutionStudentSchema = z
  .object({
    status: z.literal(VALID_STATUSES, 'Обязательное поле'),
    answer: z.string(),
    file_answer: z
      .array(z.file())
      .refine(
        (files) =>
          files.reduce((sum, f) => sum + f.size, 0) <= 30 * 1024 * 1024,
        'Суммарный размер файлов не должен превышать 30 МБ',
      ),
    deleted_file_ids: z.array(z.number()),
  })
  .superRefine((data, ctx) => {
    if (
      data.status === CHECKING_STATUS_ID &&
      data.file_answer.length === 0 &&
      !data.answer.trim()
    ) {
      ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: 'Укажите письменный ответ или прикрепите файл',
        path: ['answer'],
      })
    }
  })

export type TSolutionTeacherSchema = z.infer<typeof solutionTeacherSchema>
export type TSolutionStudentSchema = z.infer<typeof solutionStudentSchema>
export type TSolutionBaseSchema = Pick<TSolutionTeacherSchema, 'status'>
