import { ITaskFieldsRef, TaskFields } from '@/entities/task'
import { useMemberSelectOptions } from '@/entities/member'
import { memo, useMemo } from 'react'
import type { Ref } from 'react'
import { SelectOption } from '@/shared/ui'
import { TTaskWithStudents } from '@/shared/model'
import { toFormData } from '../lib/to-form-data'
import { taskSchemaUpdate } from '@/entities/task/model/schema'

interface IDialogContentProps {
  ref: Ref<ITaskFieldsRef>
  taskData: TTaskWithStudents
  onDirtyChange?: (isDirty: boolean) => void
}

const DialogContent = ({
  ref,
  taskData,
  onDirtyChange,
}: IDialogContentProps) => {
  const {
    options,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    onSearchChange,
  } = useMemberSelectOptions('student', false)

  const fieldValues = useMemo(() => toFormData(taskData), [taskData])

  // FIXME: такой же код есть в другом месте
  const seedStudentOptions = useMemo<SelectOption[]>(
    () =>
      taskData.students?.map((s) => ({
        value: String(s.id),
        label: s.fullname,
      })) ?? [],
    [taskData.students],
  )

  const mergedStudentOptions = useMemo<SelectOption[]>(() => {
    const map = new Map(seedStudentOptions.map((o) => [o.value, o]))
    options.forEach((o) => map.set(o.value, o))
    return [...map.values()]
  }, [seedStudentOptions, options])

  return (
    <TaskFields
      ref={ref}
      schema={taskSchemaUpdate}
      fieldValues={fieldValues}
      onDirtyChange={onDirtyChange}
      serverFiles={taskData.task.files}
      studentField={{
        data: mergedStudentOptions,
        selectedOptions: seedStudentOptions,
        hasMore: hasNextPage,
        isLoadingMore: isFetchingNextPage,
        onLoadMore: fetchNextPage,
        onSearchChange,
      }}
    />
  )
}

export default memo(DialogContent)
