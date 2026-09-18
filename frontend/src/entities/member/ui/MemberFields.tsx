import {
  ReactNode,
  useCallback,
  useEffect,
  useImperativeHandle,
  useRef,
} from 'react'
import type { Ref } from 'react'
import { Controller, useForm, useWatch } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import type { Resolver } from 'react-hook-form'
import { Input, Select, Textarea } from '@/shared/ui'
import {
  TMemberFormData,
  studentSchema,
  teacherSchema,
  studentFullSchema,
  teacherFullSchema,
} from '../model/member-fields-schema'
import {
  BASE_DISABLED_FIELDS,
  FIELD_CONFIG,
  INITIAL_FIELDS_VALUES,
} from '../config/fields-config'
import { FIELD_OVERRIDES } from '../config/field-overrides'
import styles from './styles.module.scss'
import { ROLE_OPTIONS } from '@/shared/config'
import { IMemberFieldsRef } from '../model/types'
import { TPaginatedSelectField } from '@/shared/ui'

interface IMemberFormProps {
  ref?: Ref<IMemberFieldsRef>
  mode: 'create' | 'update'
  fieldData?: TMemberFormData
  classField: TPaginatedSelectField
  onDirtyChange?: (isDirty: boolean) => void
}

const MemberFields = ({
  mode,
  fieldData = INITIAL_FIELDS_VALUES,
  onDirtyChange,
  ref,
  classField,
}: IMemberFormProps): ReactNode => {
  const resolver = useCallback(
    async (
      data: TMemberFormData,
      context: unknown,
      options: Parameters<Resolver<TMemberFormData>>[2],
    ) => {
      const schema =
        mode === 'update'
          ? data.role === 'teacher'
            ? teacherSchema
            : studentSchema
          : data.role === 'teacher'
            ? teacherFullSchema
            : studentFullSchema

      return (zodResolver(schema) as unknown as Resolver<TMemberFormData>)(
        data,
        context,
        options,
      )
    },
    [mode],
  )

  const {
    control,
    handleSubmit,
    reset,
    resetField,
    getValues,
    formState: { isDirty },
  } = useForm<TMemberFormData>({
    resolver,
    defaultValues: fieldData,
  })

  const role = useWatch({ control, name: 'role' })
  const disabledFields = mode === 'update' ? BASE_DISABLED_FIELDS : []

  const onDirtyChangeRef = useRef(onDirtyChange)
  useEffect(() => {
    onDirtyChangeRef.current = onDirtyChange
  })

  useEffect(() => {
    onDirtyChangeRef.current?.(isDirty)
  }, [isDirty])

  useImperativeHandle(
    ref,
    () => ({
      validate: () =>
        new Promise((resolve) =>
          handleSubmit(
            () => resolve(true),
            () => resolve(false),
          )(),
        ),
      getFieldsData: () => getValues(),
      reset: () => reset(),
    }),
    [handleSubmit, reset, getValues],
  )

  const visibleFields = FIELD_CONFIG.filter((f) => {
    if (FIELD_OVERRIDES[mode][f.name]?.hidden) return false
    return (
      !(f.name === 'class' && role === 'teacher') &&
      !(f.name === 'parentEmail' && role === 'teacher') &&
      !(f.name === 'parentPhone' && role === 'teacher')
    )
  })

  return (
    <div className={styles.wrapper}>
      {visibleFields.map((field) => {
        const disabled = disabledFields.includes(field.name)
        const override = FIELD_OVERRIDES[mode][field.name] ?? {}
        const title = override.title ?? field.title
        const required = override.required ?? field.required

        if (field.type === 'select' && field.name === 'class') {
          return (
            <Controller
              key={field.name}
              control={control}
              name='class'
              render={({ field: f, fieldState }) => (
                <Select
                  ref={f.ref}
                  label={title}
                  placeholder='Выберите'
                  fluid
                  required={required}
                  disabled={disabled}
                  isValid={!fieldState.error}
                  description={fieldState.error?.message}
                  options={classField.data}
                  searchable
                  hasMore={classField.hasMore}
                  isLoadingMore={classField.isLoadingMore}
                  onLoadMore={classField.onLoadMore}
                  onSearchChange={classField.onSearchChange}
                  value={f.value ?? ''}
                  onChange={f.onChange}
                />
              )}
            />
          )
        }

        if (field.type === 'select' && field.name === 'role') {
          return (
            <Controller
              key={field.name}
              control={control}
              name='role'
              render={({ field: f, fieldState }) => (
                <Select
                  ref={f.ref}
                  label={title}
                  placeholder='Выберите'
                  fluid
                  required={required}
                  disabled={disabled}
                  isValid={!fieldState.error}
                  description={fieldState.error?.message}
                  options={ROLE_OPTIONS}
                  value={f.value ?? ''}
                  onChange={(v) => {
                    if (v === 'teacher') resetField('class')
                    f.onChange(v)
                  }}
                />
              )}
            />
          )
        }

        if (field.type === 'textarea') {
          return (
            <Controller
              key={field.name}
              control={control}
              name={field.name}
              render={({ field: f, fieldState }) => (
                <Textarea
                  ref={f.ref}
                  label={title}
                  placeholder='Ввод..'
                  fluid
                  required={required}
                  disabled={disabled}
                  isValid={!fieldState.error}
                  description={fieldState.error?.message}
                  resize='none'
                  value={(f.value as string) ?? ''}
                  onChange={f.onChange}
                />
              )}
            />
          )
        }

        return (
          <Controller
            key={field.name}
            control={control}
            name={field.name}
            render={({ field: f, fieldState }) => (
              <Input
                ref={f.ref}
                title={title}
                fluid
                type={field.type}
                required={required}
                disabled={disabled}
                isValid={!fieldState.error}
                description={fieldState.error?.message}
                value={(f.value as string) ?? ''}
                onChange={f.onChange}
              />
            )}
          />
        )
      })}
    </div>
  )
}

export default MemberFields
