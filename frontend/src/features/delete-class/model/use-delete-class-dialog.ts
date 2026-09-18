import { useDialog } from '@/shared/lib'
import { useDeleteClass } from './use-delete-class'

interface IProps {
  id: number
  name: string
}

export const useDeleteClassDialog = () => {
  const { show } = useDialog()
  const { deleteClassById } = useDeleteClass()

  const showDialog = (props: IProps): void => {
    show({
      title: 'Удаление группы',
      content: `Удалить группу ${props.name} ?`,
      positiveText: 'Удалить',
      onConfirm: async () => {
        const success = await deleteClassById(props.id)
        if (!success) {
          throw new Error('Не удалось удалить пользователя')
        }
      },
    })
  }

  return { showDialog }
}
