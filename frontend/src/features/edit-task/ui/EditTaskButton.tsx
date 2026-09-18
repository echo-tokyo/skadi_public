import { Button } from '@/shared/ui'
import { memo } from 'react'
import { useEditTaskDialog } from '../model/use-edit-task-dialog'
import { TTaskWithStudents } from '@/shared/model'

interface IEditTaskButtonProps {
  task: TTaskWithStudents
}

const EditTaskButton = ({ task }: IEditTaskButtonProps) => {
  const { showDialog } = useEditTaskDialog(task)
  return (
    <Button color='secondary' onClick={showDialog}>
      Редактировать
    </Button>
  )
}

export default memo(EditTaskButton)
