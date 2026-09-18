import { Button } from '@/shared/ui'
import { useEditClassDialog } from '../model/useEditClassDialog'
import { IClass } from '@/entities/class'

export const EditClassButton = ({ classData }: { classData: IClass }) => {
  const { showDialog } = useEditClassDialog(classData)
  const handleClick = () => {
    showDialog()
  }

  return (
    <Button color='secondary' onClick={handleClick}>
      Редактировать
    </Button>
  )
}
