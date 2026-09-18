import { useDialog } from '@/shared/lib'
import { useDeleteMember } from './use-delete-member'

interface IProps {
  id: number
  fullname: string
}

export const useDeleteMemberDialog = () => {
  const { show } = useDialog()
  const { deleteMemberById } = useDeleteMember()

  const showDialog = (props: IProps): void => {
    show({
      title: 'Удаление пользователя',
      content: `Удалить пользователя ${props.fullname} ?`,
      positiveText: 'Удалить',
      onConfirm: async () => {
        const success = await deleteMemberById(props.id)
        if (!success) {
          throw new Error('Не удалось удалить пользователя')
        }
      },
    })
  }

  return { showDialog }
}
