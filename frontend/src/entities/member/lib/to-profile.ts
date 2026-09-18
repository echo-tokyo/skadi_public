import { TProfile } from '@/shared/model'
import { TMemberFormData } from '../model/member-fields-schema'

export const toProfile = (data: TMemberFormData): Omit<TProfile, 'id'> => ({
  fullname: data.fullname,
  address: data.address,
  contact: { email: data.email, phone: data.phone },
  parent_contact: { email: data.parentEmail, phone: data.parentPhone },
  extra: data.extra,
})
