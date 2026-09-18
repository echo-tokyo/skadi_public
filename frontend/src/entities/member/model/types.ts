import { TClass, TPagination, TProfile, TRole } from '@/shared/model'
import { TMemberFormData } from './member-fields-schema'

export interface IMember {
  class?: TClass
  id: number
  profile: TProfile
  role: TRole
  username: string
}

export interface IUpdateMemberRequest {
  class_id?: number
  password?: string
  profile: Omit<TProfile, 'id'>
}

export interface ICreateMemberRequest {
  class_id?: number
  password: string
  profile: Omit<TProfile, 'id'>
  role: TRole
  username: string
}

export interface IMembersResponse {
  data: IMember[]
  pagination?: TPagination
}

export interface IMembersQuery {
  free?: boolean
  role?: TRole[]
  'per-page'?: number
  search?: string
}

export interface IMemberFieldsRef {
  validate: () => Promise<boolean>
  getFieldsData: () => TMemberFormData
  reset: () => void
}
