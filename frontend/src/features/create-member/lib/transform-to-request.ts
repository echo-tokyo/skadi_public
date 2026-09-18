import {
  ICreateMemberRequest,
  TMemberFormData,
  toProfile,
} from '@/entities/member'

export const transformToRequest = (
  data: TMemberFormData,
): ICreateMemberRequest => ({
  class_id: data.class ? Number(data.class) : undefined,
  username: data.username,
  password: data.password,
  role: data.role,
  profile: toProfile(data),
})
