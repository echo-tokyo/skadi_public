import { useDialog } from '@/shared/lib'
import { useDeleteTask } from './use-delete-task'

interface IProps {
  id: number
  taskName: string
}

export const useDeleteTaskDialog = () => {
  const { show } = useDialog()
  const { deleteTaskById } = useDeleteTask()

  const showDialog = (props: IProps): void => {
    show({
      title: 'Удаление задания',
      content: `Удалить задание '${props.taskName}' со всеми его решениями ?`,
      positiveText: 'Удалить',
      size: 's',
      onConfirm: async () => {
        const success = await deleteTaskById(props.id)
        if (!success) {
          throw new Error('Не удалось удалить задание')
        }
      },
    })
  }

  return { showDialog }
}
