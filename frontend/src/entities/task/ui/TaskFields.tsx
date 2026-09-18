import { ReactNode, useEffect, useImperativeHandle, useRef } from 'react'
import type { Ref } from 'react'
import { Controller, useController, useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { FileField, Input, Select, Textarea } from '@/shared/ui'
import { TPaginatedSelectField } from '@/shared/ui'
import { TFile } from '@/shared/model'
import { ITaskFieldsRef } from '../model/types'
import styles from './styles.module.scss'
import {
  taskSchemaCreate,
  taskSchemaUpdate,
  TTaskSchemaUpdate,
} from '../model/schema'
import { defaultValues } from '../config/default-values'
import z from 'zod'

interface ITaskFieldsProps {
  ref?: Ref<ITaskFieldsRef>
  studentField: TPaginatedSelectField
  classField?: TPaginatedSelectField
  onDirtyChange?: (isDirty: boolean) => void
  fieldValues?: TTaskSchemaUpdate
  schema: typeof taskSchemaCreate | typeof taskSchemaUpdate
  serverFiles?: TFile[]
}

const TaskFields = ({
  schema,
  ref,
  studentField,
  classField,
  onDirtyChange,
  fieldValues,
  serverFiles = [],
}: ITaskFieldsProps): ReactNode => {
  const {
    control,
    reset,
    getValues,
    handleSubmit,
    formState: { isDirty },
  } = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema),
    defaultValues: fieldValues ?? defaultValues,
  })

  const { field: deletedField } = useController({
    control,
    name: 'deletedFileIds',
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
    [reset, getValues, handleSubmit],
  )

  return (
    <div className={styles.content}>
      <Controller
        control={control}
        name='title'
        render={({ field, fieldState }) => (
          <Input
            fluid
            title='Название'
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
        name='students'
        render={({ field, fieldState }) => (
          <Select
            fluid
            label='Ученики'
            multiple
            placeholder='Выберите'
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            options={studentField.data}
            selectedOptions={studentField.selectedOptions}
            value={field.value}
            searchable
            onLoadMore={studentField.onLoadMore}
            hasMore={studentField.hasMore}
            isLoadingMore={studentField.isLoadingMore}
            onSearchChange={studentField.onSearchChange}
            onChange={field.onChange}
          />
        )}
      />
      {classField && (
        <Controller
          control={control}
          name='classes'
          render={({ field, fieldState }) => (
            <Select
              fluid
              label='Группы'
              multiple
              placeholder='Выберите'
              isValid={!fieldState.error}
              description={fieldState.error?.message}
              options={classField.data}
              selectedOptions={classField.selectedOptions}
              value={field.value}
              searchable
              onLoadMore={classField.onLoadMore}
              hasMore={classField.hasMore}
              isLoadingMore={classField.isLoadingMore}
              onSearchChange={classField.onSearchChange}
              onChange={field.onChange}
            />
          )}
        />
      )}
      <Controller
        control={control}
        name='description'
        render={({ field, fieldState }) => (
          <Textarea
            label='Описание'
            fluid
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
        name='files'
        render={({ field, fieldState }) => (
          <FileField
            label='Файлы'
            fluid
            multiple
            isValid={!fieldState.error}
            description={fieldState.error?.message}
            attachments={serverFiles
              .filter((f) => !deletedField.value.includes(f.id))
              .map((f) => ({ id: f.id, name: f.name, size: f.size }))}
            onRemoveAttachment={(id) =>
              deletedField.onChange([...deletedField.value, id])
            }
            value={field.value}
            onChange={field.onChange}
          />
        )}
      />
    </div>
  )
}

export default TaskFields
