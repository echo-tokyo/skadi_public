import { ReactNode, useEffect, useImperativeHandle, useRef } from 'react'
import type { Ref } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Input, Select } from '@/shared/ui'
import styles from './styles.module.scss'
import { classSchema } from '../model/class-form-schema'
import type { TClassSchema } from '../model/class-form-schema'
import { INITIAL_FIELDS_VALUES } from '../config/fields-config'
import { IClassFieldsRef } from '../model/types'
import { TPaginatedSelectField } from '@/shared/ui'

interface IClassFieldsProps {
  ref?: Ref<IClassFieldsRef>
  fieldValues?: TClassSchema
  teacherField: TPaginatedSelectField
  studentField: TPaginatedSelectField
  onDirtyChange?: (isDirty: boolean) => void
}

const ClassFields = ({
  fieldValues = INITIAL_FIELDS_VALUES,
  teacherField,
  studentField,
  onDirtyChange,
  ref,
}: IClassFieldsProps): ReactNode => {
  const {
    control,
    trigger,
    reset,
    getValues,
    handleSubmit,
    formState: { isDirty },
  } = useForm<TClassSchema>({
    resolver: zodResolver(classSchema),
    defaultValues: fieldValues,
  })

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
    [trigger, reset, getValues],
  )

  return (
    <div className={styles.wrapper}>
      <Controller
        control={control}
        name='className'
        render={({ field, fieldState }) => (
          <Input
            ref={field.ref}
            fluid
            title='Название группы'
            required
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            value={field.value}
            onChange={field.onChange}
          />
        )}
      />
      <Controller
        control={control}
        name='teacher'
        render={({ field, fieldState }) => (
          <Select
            ref={field.ref}
            fluid
            label='Преподаватель'
            placeholder='Выберите'
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            options={teacherField.data}
            selectedOptions={teacherField.selectedOptions}
            value={field.value ?? ''}
            searchable
            onLoadMore={teacherField.onLoadMore}
            hasMore={teacherField.hasMore}
            isLoadingMore={teacherField.isLoadingMore}
            onSearchChange={teacherField.onSearchChange}
            onChange={field.onChange}
          />
        )}
      />
      <Controller
        control={control}
        name='students'
        render={({ field, fieldState }) => (
          <Select
            ref={field.ref}
            fluid
            label='Ученики'
            multiple
            placeholder='Выберите'
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            options={studentField.data}
            selectedOptions={studentField.selectedOptions}
            value={field.value ?? []}
            searchable
            onLoadMore={studentField.onLoadMore}
            hasMore={studentField.hasMore}
            isLoadingMore={studentField.isLoadingMore}
            onSearchChange={studentField.onSearchChange}
            onChange={(val) => field.onChange([...val].sort())}
          />
        )}
      />
      <Controller
        control={control}
        name='schedule'
        render={({ field, fieldState }) => (
          <Input
            ref={field.ref}
            fluid
            title='Расписание'
            placeholder='Каждый четверг, 18:00 - 19:00'
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            value={field.value}
            onChange={field.onChange}
          />
        )}
      />
    </div>
  )
}

export default ClassFields
