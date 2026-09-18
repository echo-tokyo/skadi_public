import { ClassFields, IClassFieldsRef } from '@/entities/class'
import { useMemberSelectOptions } from '@/entities/member'
import { memo } from 'react'
import type { Ref } from 'react'

interface DialogContentProps {
  classFieldsRef: Ref<IClassFieldsRef>
  onDirtyChange: (isDirty: boolean) => void
}

const DialogContent = (props: DialogContentProps) => {
  const { classFieldsRef, onDirtyChange } = props

  const {
    options: studentOptions,
    fetchNextPage: fetchNextStudentsPage,
    hasNextPage: hasNextStudentsPage,
    isFetchingNextPage: isFetchingNextStudentsPage,
    onSearchChange: onStudentSearchChange,
  } = useMemberSelectOptions('student')
  const {
    options: teacherOptions,
    fetchNextPage: fetchNextTeachersPage,
    hasNextPage: hasNextTeachersPage,
    isFetchingNextPage: isFetchingNextTeachersPage,
    onSearchChange: onTeacherSearchChange,
  } = useMemberSelectOptions('teacher')

  return (
    <ClassFields
      ref={classFieldsRef}
      onDirtyChange={onDirtyChange}
      teacherField={{
        data: teacherOptions,
        hasMore: hasNextTeachersPage,
        isLoadingMore: isFetchingNextTeachersPage,
        onLoadMore: fetchNextTeachersPage,
        onSearchChange: onTeacherSearchChange,
      }}
      studentField={{
        data: studentOptions,
        hasMore: hasNextStudentsPage,
        isLoadingMore: isFetchingNextStudentsPage,
        onLoadMore: fetchNextStudentsPage,
        onSearchChange: onStudentSearchChange,
      }}
    />
  )
}

export default memo(DialogContent)
