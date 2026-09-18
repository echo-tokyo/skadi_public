import { ROLE_VALUES } from '@/shared/config'
import { z } from 'zod'

const phoneSchema = z
  .string()
  .regex(/^\+?[0-9\s\-()]{11,20}$/, 'Некорректный номер телефона')
  .or(z.literal(''))

// Базовая схема
const memberBaseSchema = z.object({
  fullname: z
    .string()
    .min(1, 'Обязательное поле')
    .max(150, 'Максимум 150 символов'),
  address: z.string().max(200, 'Максимум 200 символов'),
  email: z.email('Некорректный email').or(z.literal('')),
  phone: phoneSchema,
  extra: z.string().max(500, 'Максимум 500 символов'),
  password: z.union([
    z.literal(''),
    z.string().min(8, 'Минимум 8 символов').max(40, 'Максимум 40 символов'),
  ]),
})

// Схемы для редактирования
export const teacherSchema = memberBaseSchema.describe('TeacherSchema')
export const studentSchema = memberBaseSchema
  .extend({
    class: z.string().optional(),
    parentEmail: z.email('Некорректный email').or(z.literal('')),
    parentPhone: phoneSchema,
  })
  .describe('StudentSchema')

// Полные схемы для создания
const memberIdentitySchema = z.object({
  role: z.enum(ROLE_VALUES, { message: 'Выберите роль' }),
  username: z
    .string()
    .min(1, 'Обязательное поле')
    .min(3, 'Минимум 3 символа')
    .max(50, 'Максимум 50 символов')
    .regex(/^[a-zA-Z0-9_-]+$/, 'Только латиница, цифры, _ и -'),
  password: z
    .string()
    .min(1, 'Обязательное поле')
    .min(8, 'Минимум 8 символов')
    .max(40, 'Максимум 40 символов'),
})

export const studentFullSchema = z
  .object({
    ...studentSchema.shape,
    ...memberIdentitySchema.shape,
  })
  .describe('StudentFullSchema')
export const teacherFullSchema = z
  .object({
    ...teacherSchema.shape,
    ...memberIdentitySchema.shape,
  })
  .describe('TeacherFullSchema')

export type TTeacherSchema = z.infer<typeof teacherSchema>
export type TStudentSchema = z.infer<typeof studentSchema>
export type TTeacherFullSchema = z.infer<typeof teacherFullSchema>
export type TStudentFullSchema = z.infer<typeof studentFullSchema>
export type TMemberFormData = TStudentFullSchema
