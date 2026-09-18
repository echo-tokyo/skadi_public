import { TProfile } from '@/shared/model'
import { TClassSchema } from './class-form-schema'

type TClassPagination = {
  page: number
  pages: number
  per_page: number
  total: number
}

export interface IClassRequest {
  name: string
  schedule?: string
  students?: number[]
  teacher_id?: number
}

export interface IClass {
  id: number
  name: string
  schedule?: string
  students?: TProfile[]
  teacher?: TProfile
}

export interface IClassResponse {
  data: IClass[]
  pagination?: TClassPagination
}

export interface IClassQuery {
  'per-page'?: number
  search?: string
}

export interface IClassFieldsRef {
  validate: () => Promise<boolean>
  getFieldsData: () => TClassSchema
  reset: () => void
}
