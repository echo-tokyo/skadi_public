import { useRef } from 'react'
import { useDialog } from '@/shared/lib'
import { IClassFieldsRef } from '@/entities/class'
import DialogContent from '../ui/DialogContent'
import { useCreateClass } from './use-create-class'

export const useCreateClassDialog = () => {
  const { show, update } = useDialog()
  const formRef = useRef<IClassFieldsRef>(null)
  const dialogIdRef = useRef<string | null>(null)
  const { submit } = useCreateClass()

  const showDialog = (): void => {
    const id = show({
      title: 'Создание группы',
      content: (
        <DialogContent
          classFieldsRef={formRef}
          onDirtyChange={(isDirty) => {
            if (dialogIdRef.current) {
              update(dialogIdRef.current, { isConfirmDisabled: !isDirty })
            }
          }}
        />
      ),
      positiveText: 'Создать',
      negativeText: 'Отмена',
      size: 'm',
      isConfirmDisabled: true,
      onConfirm: async () => {
        const isValid = await formRef.current?.validate()

        if (!isValid) {
          throw new Error('Validation failed')
        }

        const formData = formRef.current?.getFieldsData()

        if (formData) {
          const success = await submit(formData)
          if (!success) {
            throw new Error('Failed to create member')
          }
        }
      },
    })

    dialogIdRef.current = id
  }

  return { showDialog }
}
