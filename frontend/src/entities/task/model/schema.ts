import z from 'zod'

export const taskSchemaUpdate = z
  .object({
    title: z.string().min(1, 'Обязательное поле'),
    description: z.string().min(1, 'Обязательное поле'),
    students: z.array(z.string()),
    files: z
      .array(z.file())
      .refine(
        (files) =>
          files.reduce((sum, f) => sum + f.size, 0) <= 30 * 1024 * 1024,
        'Суммарный размер файлов не должен превышать 30 МБ',
      ),
    deletedFileIds: z.array(z.number()),
  })
  .describe('taskSchemaUpdate')

export const taskSchemaCreate = z
  .object({
    ...taskSchemaUpdate.shape,
    classes: z.array(z.string()),
  })
  .describe('taskSchemaCreate')

export type TTaskSchemaCreate = z.infer<typeof taskSchemaCreate>
export type TTaskSchemaUpdate = z.infer<typeof taskSchemaUpdate>
