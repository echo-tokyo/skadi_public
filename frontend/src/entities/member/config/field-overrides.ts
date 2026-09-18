import type { TMemberFormData } from '../model/member-fields-schema'

type TFieldOverride = {
  title?: string
  required?: boolean
  hidden?: boolean
}

type TFieldOverridesMap = Partial<Record<keyof TMemberFormData, TFieldOverride>>

// Переопределение полей в зависимости от режима
export const FIELD_OVERRIDES: Record<'create' | 'update', TFieldOverridesMap> =
  {
    create: {},
    update: {
      password: { title: 'Новый пароль', required: false },
    },
  }
