import { useDialog } from '@/shared/lib'
import { useDeleteSolution } from './use-delete-solution'

interface IProps {
  id: number
  studentName: string
  taskTitle: string
}

export const useDeleteSolutionDialog = () => {
  const { show } = useDialog()
  const { deleteSolutionById } = useDeleteSolution()

  const showDialog = (props: IProps): void => {
    show({
      title: 'Удаление решения',
      content: `Удалить решение студента '${props.studentName}' по заданию '${props.taskTitle}' ?`,
      positiveText: 'Удалить',
      size: 's',
      onConfirm: async () => {
        const success = await deleteSolutionById(props.id)
        if (!success) {
          throw new Error('Не удалось удалить решение')
        }
      },
    })
  }

  return { showDialog }
}
