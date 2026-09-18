import { Button } from '@/shared/ui'
import { useCreateClassDialog } from '../model/useCreateClassDialog'

const CreateClassButton = () => {
  const { showDialog } = useCreateClassDialog()
  return <Button onClick={showDialog}>Создать группу</Button>
}

export default CreateClassButton
