import { useRef, useMemo } from 'react'
import { useDialog } from '@/shared/lib'
import { IMember, IMemberFieldsRef } from '@/entities/member'
import { toFormData } from '../lib/to-form-data'
import { useEditMember } from './use-edit-member'
import DialogContent from '../ui/DialogContent'

export const useEditMemberDialog = (member: IMember) => {
  const { show, update } = useDialog()
  const formRef = useRef<IMemberFieldsRef>(null)
  const { submit } = useEditMember(member.id)
  const fieldData = useMemo(() => toFormData(member), [member])
  const dialogIdRef = useRef<string | null>(null)

  const showDialog = () => {
    const id = show({
      title: 'Редактирование пользователя',
      isConfirmDisabled: true,

      content: (
        <DialogContent
          ref={formRef}
          fieldData={fieldData}
          onDirtyChange={(isDirty) => {
            if (dialogIdRef.current) {
              update(dialogIdRef.current, {
                isConfirmDisabled: !isDirty,
              })
            }
          }}
        />
      ),
      positiveText: 'Сохранить',
      negativeText: 'Отмена',
      size: 'm',
      onConfirm: async () => {
        const isValid = await formRef.current?.validate()

        if (!isValid) {
          throw new Error('Validation failed')
        }

        const formData = formRef.current?.getFieldsData()

        if (formData) {
          const success = await submit(formData)
          if (!success) {
            throw new Error('Failed to update member')
          }
        }
      },
    })
    dialogIdRef.current = id
  }

  return { showDialog }
}
