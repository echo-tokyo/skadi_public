import z from 'zod'

export const authSchema = z.object({
  login: z.string().min(1, 'Обязательное поле').max(50, 'Максимум 50 символов'),
  password: z
    .string()
    .min(1, 'Обязательное поле')
    .min(8, 'Минимум 8 символов')
    .max(40, 'Максимум 40 символов'),
})

export type TAuthSchema = z.infer<typeof authSchema>
