import { useCreateCommentMutation } from '@/entities/comment'
import { useMutationAction } from '@/shared/lib'
import { TRole } from '@/shared/model'

export const useSendComment = ({ id, role }: { id: number; role: TRole }) =>
  useMutationAction({
    mutation: useCreateCommentMutation(),
    prepare: (message: string) => ({
      data: { message },
      id,
      role,
    }),
  })
