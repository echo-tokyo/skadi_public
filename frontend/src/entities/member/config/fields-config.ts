import { TMemberFormData } from '../model/member-fields-schema'

type TInputFieldName = Exclude<
  keyof TMemberFormData,
  'role' | 'extra' | 'class'
>

export type TFieldConfig =
  | { type: 'text'; name: TInputFieldName; title: string; required?: boolean }
  | { type: 'select'; name: 'role'; title: string; required?: boolean }
  | { type: 'select'; name: 'class'; title: string; required?: boolean }
  | { type: 'textarea'; name: 'extra'; title: string; required?: boolean }
  | { type: 'password'; name: 'password'; title: string; required?: boolean }

// Значения если не передан fieldData
export const INITIAL_FIELDS_VALUES: TMemberFormData = {
  fullname: '',
  role: 'student',
  class: '',
  username: '',
  password: '',
  address: '',
  email: '',
  phone: '',
  parentEmail: '',
  parentPhone: '',
  extra: '',
}

// Поля для маппинга
export const FIELD_CONFIG: TFieldConfig[] = [
  { type: 'text', name: 'fullname', title: 'ФИО', required: true },
  { type: 'select', name: 'role', title: 'Роль', required: true },
  { type: 'text', name: 'username', title: 'Логин', required: true },
  { type: 'password', name: 'password', title: 'Пароль', required: true },
  { type: 'select', name: 'class', title: 'Группа' },
  { type: 'text', name: 'address', title: 'Адрес проживания' },
  { type: 'text', name: 'email', title: 'Email' },
  { type: 'text', name: 'phone', title: 'Телефон' },
  { type: 'text', name: 'parentEmail', title: 'Email родителя' },
  { type: 'text', name: 'parentPhone', title: 'Телефон родителя' },
  { type: 'textarea', name: 'extra', title: 'Дополнительная информация' },
]

// Disabled поля при редактировании
export const BASE_DISABLED_FIELDS: Array<keyof TMemberFormData> = [
  'role',
  'username',
]
