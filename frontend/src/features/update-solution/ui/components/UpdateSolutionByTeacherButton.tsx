import { Button } from '@/shared/ui'
import { useFormContext } from 'react-hook-form'
import { TSolutionTeacherSchema } from '@/entities/solution'
import { useTeacherUpdateSolution } from '../../model/use-teacher-update-solution'
import { CHECKED_STATUS_ID } from '@/shared/config'

const UpdateSolutionByTeacherButton = ({ id }: { id: number }) => {
  const {
    watch,
    handleSubmit,
    reset,
    formState: { isDirty },
  } = useFormContext<TSolutionTeacherSchema>()

  const { submit, isLoading } = useTeacherUpdateSolution(id)
  const statusValue = watch('status')

  const onSubmit = handleSubmit(async (data) => {
    const success = await submit(data)
    if (success) reset(data)
  })

  return (
    <Button disabled={!isDirty} isLoading={isLoading} onClick={onSubmit}>
      {statusValue === CHECKED_STATUS_ID
        ? 'Одобрить выполнение'
        : 'Сохранить решение'}
    </Button>
  )
}

export default UpdateSolutionByTeacherButton
