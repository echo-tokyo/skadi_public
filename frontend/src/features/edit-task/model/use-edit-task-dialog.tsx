import { useRef } from 'react'
import { useDialog } from '@/shared/lib'
import { ITaskFieldsRef } from '@/entities/task'
import { useEditTask } from './use-edit-task'
import DialogContent from '../ui/DialogContent'
import { TTaskWithStudents } from '@/shared/model'

export const useEditTaskDialog = (task: TTaskWithStudents) => {
  const { show, update } = useDialog()
  const formRef = useRef<ITaskFieldsRef>(null)
  const { submit } = useEditTask(task.task.id)
  const dialogIdRef = useRef<string | null>(null)

  const showDialog = () => {
    const id = show({
      title: 'Редактирование задания',
      isConfirmDisabled: true,
      content: (
        <DialogContent
          ref={formRef}
          taskData={task}
          onDirtyChange={(isDirty) => {
            if (dialogIdRef.current) {
              update(dialogIdRef.current, { isConfirmDisabled: !isDirty })
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
            throw new Error('Failed to update task')
          }
        }
      },
    })
    dialogIdRef.current = id
  }

  return { showDialog }
}
