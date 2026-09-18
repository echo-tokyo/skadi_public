import { Button } from '@/shared/ui'
import { useDeleteSolutionDialog } from '../model/use-delete-solution-dialog'

interface IDeleteSolutionProps {
  id: number
  studentName: string
  taskTitle: string
}

export const DeleteSolutionButton = (props: IDeleteSolutionProps) => {
  const { showDialog } = useDeleteSolutionDialog()
  const { id, studentName, taskTitle } = props

  const handleClick = () => {
    showDialog({ id, studentName, taskTitle })
  }

  return (
    <Button color='inverted' onClick={handleClick}>
      Удалить решение
    </Button>
  )
}

DeleteSolutionButton.displayName = 'DeleteSolutionButton'
