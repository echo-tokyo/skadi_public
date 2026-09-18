import { TClass, TProfile, TRole } from '@/shared/model'

export interface IUser {
  class?: TClass
  id: string
  profile: TProfile
  role: TRole
  username: string
}

export type IUserResponse = IUser
