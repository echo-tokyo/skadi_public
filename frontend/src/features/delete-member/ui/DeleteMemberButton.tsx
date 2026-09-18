import { Button } from '@/shared/ui'
import { useDeleteMemberDialog } from '../model/use-delete-member-dialog'

interface IDeleteMemberProps {
  id: number
  fullname?: string
}

export const DeleteMemberButton = (props: IDeleteMemberProps) => {
  const { showDialog } = useDeleteMemberDialog()
  const { fullname = 'пользователь', id } = props

  const handleClick = () => {
    showDialog({ fullname, id })
  }

  return (
    <Button color='inverted' onClick={handleClick}>
      Удалить пользователя
    </Button>
  )
}

DeleteMemberButton.displayName = 'DeleteMemberButton'
