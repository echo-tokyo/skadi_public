import { EditTaskButton } from '@/features/edit-task'
import { AccordionCard } from '@/shared/ui'
import { memo } from 'react'
import { TTaskWithStudents } from '@/shared/model'
import { DeleteTaskButton } from '@/features/delete-task'

interface ITaskCardItemProps {
  taskData: TTaskWithStudents
}

export const TaskCardItem = memo(({ taskData }: ITaskCardItemProps) => {
  return (
    <AccordionCard
      title={taskData.task.title}
      fields={[
        { label: 'Описание', value: taskData.task.description },
        {
          label: 'Ученики',
          value: taskData.students?.map((el) => el.fullname).join(', '),
        },
      ]}
      actions={
        <>
          <DeleteTaskButton
            id={taskData.task.id}
            taskName={taskData.task.title}
          />
          <EditTaskButton task={taskData} />
        </>
      }
    />
  )
})

TaskCardItem.displayName = 'TaskCardItem'
