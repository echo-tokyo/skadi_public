import { TRole, TStatusId } from '../model'
import { TGrade } from '../model/types/types'
import { SelectOption } from '../ui'

// Селекту не нужен admin
export const ROLE_OPTIONS: SelectOption[] = [
  { value: 'teacher', label: 'Преподаватель' },
  { value: 'student', label: 'Студент' },
]

export const ROLE_VALUES: TRole[] = ['teacher', 'student']

export const STATUS_OPTIONS: SelectOption<TStatusId>[] = [
  { value: 1, label: 'Бэклог' },
  { value: 2, label: 'В работе' },
  { value: 3, label: 'На проверке' },
  { value: 4, label: 'Проверено' },
]

export const CHECKED_STATUS_ID = 4
export const CHECKING_STATUS_ID = 3

export const GRADE_OPTIONS: SelectOption<TGrade>[] = [
  { value: '5', label: 'Отлично' },
  { value: '4', label: 'Хорошо' },
  { value: '3', label: 'Можешь лучше' },
  { value: '2', label: 'Не старался' },
]
