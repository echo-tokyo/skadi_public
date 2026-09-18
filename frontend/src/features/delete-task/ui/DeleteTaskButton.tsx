import { Button } from '@/shared/ui'
import { useDeleteTaskDialog } from '../model/use-delete-task-dialog'

interface IDeleteTaskProps {
  id: number
  taskName: string
}

export const DeleteTaskButton = (props: IDeleteTaskProps) => {
  const { showDialog } = useDeleteTaskDialog()
  const { taskName, id } = props

  const handleClick = () => {
    showDialog({ taskName, id })
  }

  return (
    <Button color='inverted' onClick={handleClick}>
      Удалить задание
    </Button>
  )
}

DeleteTaskButton.displayName = 'DeleteTaskButton'
