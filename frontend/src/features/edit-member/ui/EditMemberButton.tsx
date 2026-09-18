import { Button } from '@/shared/ui'
import { useEditMemberDialog } from '../model/use-edit-member-dialog'
import { IMember } from '@/entities/member'
import { memo } from 'react'

interface IEditMemberProps {
  member: IMember
}

export const EditMemberButton = memo((props: IEditMemberProps) => {
  const { showDialog } = useEditMemberDialog(props.member)

  const handleClick = () => {
    showDialog()
  }

  return (
    <Button color='secondary' onClick={handleClick}>
      Редактировать
    </Button>
  )
})

EditMemberButton.displayName = 'EditMemberButton'
