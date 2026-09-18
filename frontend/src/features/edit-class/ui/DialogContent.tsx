import { ClassFields, IClass, IClassFieldsRef } from '@/entities/class'
import { useMemberSelectOptions } from '@/entities/member'
import { memo, useMemo } from 'react'
import type { Ref } from 'react'
import { toFieldValues } from '../lib/to-field-values'
import type { SelectOption } from '@/shared/ui'

interface DialogContentProps {
  classData: IClass
  classFieldsRef: Ref<IClassFieldsRef>
  onDirtyChange: (isDirty: boolean) => void
}

const DialogContent = (props: DialogContentProps) => {
  const { classFieldsRef, onDirtyChange, classData } = props
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

  const fieldValues = useMemo(() => toFieldValues(classData), [classData])

  const seedTeacherOptions = useMemo<SelectOption[]>(
    () =>
      classData.teacher
        ? [
            {
              value: String(classData.teacher.id),
              label: classData.teacher.fullname,
            },
          ]
        : [],
    [classData.teacher],
  )

  const seedStudentOptions = useMemo<SelectOption[]>(
    () =>
      classData.students?.map((s) => ({
        value: String(s.id),
        label: s.fullname,
      })) ?? [],
    [classData.students],
  )

  const mergedStudentOptions = useMemo<SelectOption[]>(() => {
    const map = new Map(seedStudentOptions.map((o) => [o.value, o]))
    studentOptions.forEach((o) => map.set(o.value, o))
    return [...map.values()]
  }, [seedStudentOptions, studentOptions])

  return (
    <ClassFields
      ref={classFieldsRef}
      fieldValues={fieldValues}
      onDirtyChange={onDirtyChange}
      teacherField={{
        data: teacherOptions,
        selectedOptions: seedTeacherOptions,
        hasMore: hasNextTeachersPage,
        isLoadingMore: isFetchingNextTeachersPage,
        onLoadMore: fetchNextTeachersPage,
        onSearchChange: onTeacherSearchChange,
      }}
      studentField={{
        data: mergedStudentOptions,
        selectedOptions: seedStudentOptions,
        hasMore: hasNextStudentsPage,
        isLoadingMore: isFetchingNextStudentsPage,
        onLoadMore: fetchNextStudentsPage,
        onSearchChange: onStudentSearchChange,
      }}
    />
  )
}

export default memo(DialogContent)
