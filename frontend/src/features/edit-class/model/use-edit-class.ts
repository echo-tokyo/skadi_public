import { TClassSchema, transformToRequest, useEditClassMutation } from '@/entities/class'
import { useMutationAction } from '@/shared/lib'

export const useEditClass = (id: number) => {
  return useMutationAction<TClassSchema, { id: number; data: ReturnType<typeof transformToRequest> }>({
    mutation: useEditClassMutation(),
    prepare: (formData) => ({ id, data: transformToRequest(formData) }),
    successMessage: 'Группа обновлена',
  })
}
