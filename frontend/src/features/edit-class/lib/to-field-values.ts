import { IClass, TClassSchema } from '@/entities/class'

export const toFieldValues = (classData: IClass): TClassSchema => ({
  className: classData.name,
  schedule: classData.schedule ?? '',
  students: classData.students?.map((el) => String(el.id)).sort() ?? [],
  teacher: classData.teacher ? String(classData.teacher.id) : '',
})
