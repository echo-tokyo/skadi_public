import { Button } from '@/shared/ui'
import { useFormContext } from 'react-hook-form'
import { TSolutionStudentSchema } from '@/entities/solution'
import { useStudentUpdateSolution } from '../../model/use-student-update-solution'

const UpdateSolutionByStudentButton = ({ id }: { id: number }) => {
  const {
    handleSubmit,
    reset,
    formState: { isDirty },
  } = useFormContext<TSolutionStudentSchema>()

  const { submit, isLoading } = useStudentUpdateSolution(id)

  const onSubmit = handleSubmit(async (data) => {
    const success = await submit(data)
    if (success) reset(data)
  })

  return (
    <Button disabled={!isDirty} isLoading={isLoading} onClick={onSubmit}>
      Сохранить решение
    </Button>
  )
}

export default UpdateSolutionByStudentButton
