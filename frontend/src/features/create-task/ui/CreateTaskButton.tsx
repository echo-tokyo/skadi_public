import { Button } from '@/shared/ui'
import { memo } from 'react'
import { useCreateTaskDialog } from '../model/useCreateTaskDialog'

const CreateTaskButton = () => {
  const { showDialog } = useCreateTaskDialog()
  return <Button onClick={showDialog}>Создать задание</Button>
}

export default memo(CreateTaskButton)
