import {
  TClassSchema,
  transformToRequest,
  useCreateClassMutation,
} from '@/entities/class'
import { useMutationAction } from '@/shared/lib'

export const useCreateClass = () => {
  return useMutationAction<TClassSchema, ReturnType<typeof transformToRequest>>(
    {
      mutation: useCreateClassMutation(),
      prepare: transformToRequest,
      successMessage: 'Группа создана',
    },
  )
}
