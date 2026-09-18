import { TClassSchema } from '../model/class-form-schema'
import { IClassRequest } from '../model/types'

export const transformToRequest = (data: TClassSchema): IClassRequest => {
  return {
    name: data.className,
    schedule: data.schedule,
    students: data.students ? data.students.map((el) => Number(el)) : undefined,
    teacher_id: Number(data.teacher),
  }
}
