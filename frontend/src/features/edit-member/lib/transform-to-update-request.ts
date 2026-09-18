import {
  IUpdateMemberRequest,
  TMemberFormData,
  toProfile,
} from '@/entities/member'

export const transformToUpdateRequest = (
  data: TMemberFormData,
): IUpdateMemberRequest => ({
  class_id: Number(data.class),
  profile: toProfile(data),
  password: data.password,
})
