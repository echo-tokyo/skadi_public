import { ITaskFieldsRef, TaskFields } from '@/entities/task'
import { useMemberSelectOptions } from '@/entities/member'
import { memo } from 'react'
import type { Ref } from 'react'
import { useClassSelectOptions } from '@/entities/class'
import { taskSchemaCreate } from '@/entities/task/model/schema'

interface IDialogContentProps {
  ref: Ref<ITaskFieldsRef>
  onDirtyChange?: (isDirty: boolean) => void
}

const DialogContent = ({ ref, onDirtyChange }: IDialogContentProps) => {
  const {
    options,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    onSearchChange,
  } = useMemberSelectOptions('student', false)

  const {
    options: classOptions,
    fetchNextPage: classFetchNextPage,
    hasNextPage: classHasNextPage,
    isFetchingNextPage: classIsFetchingNextPage,
    onSearchChange: classOnSearchChange,
  } = useClassSelectOptions()

  return (
    <TaskFields
      ref={ref}
      schema={taskSchemaCreate}
      onDirtyChange={onDirtyChange}
      studentField={{
        data: options,
        hasMore: hasNextPage,
        isLoadingMore: isFetchingNextPage,
        onLoadMore: fetchNextPage,
        onSearchChange,
      }}
      classField={{
        data: classOptions,
        hasMore: classHasNextPage,
        isLoadingMore: classIsFetchingNextPage,
        onLoadMore: classFetchNextPage,
        onSearchChange: classOnSearchChange,
      }}
    />
  )
}

export default memo(DialogContent)
