import { Button } from '@/shared/ui'
import { useCreateMemberDialog } from '../model/useCreateMemberDialog'
import { memo } from 'react'

const CreateRoleButton = () => {
  const { showDialog } = useCreateMemberDialog()
  return <Button onClick={showDialog}>Создать роль</Button>
}

export default memo(CreateRoleButton)
