import { useRef } from 'react'
import { useDialog } from '@/shared/lib'
import { ITaskFieldsRef, TTaskSchemaCreate } from '@/entities/task'
import { useCreateTask } from './use-create-task'
import DialogContent from '../ui/DialogContent'

export const useCreateTaskDialog = () => {
  const { show, update } = useDialog()
  const { submit } = useCreateTask()
  const formRef = useRef<ITaskFieldsRef>(null)
  const dialogIdRef = useRef<string | null>(null)

  const showDialog = (): void => {
    const id = show({
      title: 'Создание задания',
      content: (
        <DialogContent
          ref={formRef}
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

        const formData = formRef.current?.getFieldsData() as TTaskSchemaCreate

        if (formData) {
          const success = await submit(formData)
          if (!success) {
            throw new Error('Failed to create task')
          }
        }
      },
    })

    dialogIdRef.current = id
  }

  return { showDialog }
}
