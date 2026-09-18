import { Button } from '@/shared/ui'
import { useDeleteClassDialog } from '../model/use-delete-class-dialog'

interface IDeleteClassButtonProps {
  id: number
  name: string
}

export const DeleteClassButton = (props: IDeleteClassButtonProps) => {
  const { id, name } = props
  const { showDialog } = useDeleteClassDialog()

  const handleClassDelete = () => {
    showDialog({ id, name })
  }

  return (
    <Button color='inverted' onClick={handleClassDelete}>
      Удалить группу
    </Button>
  )
}
