import { toast } from 'sonner'
import { getErrorMessage } from '@/shared/api'
import { useCallback, useRef } from 'react'

type MutationTuple<TArg> = readonly [
  (arg: TArg) => { unwrap: () => Promise<unknown> },
  { isLoading: boolean },
]

interface UseMutationActionConfig<TInput, TArg> {
  mutation: MutationTuple<TArg>
  prepare: (input: TInput) => TArg
  successMessage?: string
}

export const useMutationAction = <TInput, TArg>(
  config: UseMutationActionConfig<TInput, TArg>,
) => {
  const [mutate, { isLoading }] = config.mutation
  const configRef = useRef(config)
  configRef.current = config

  const submit = useCallback(
    async (input: TInput): Promise<boolean> => {
      try {
        await mutate(configRef.current.prepare(input)).unwrap()
        if (configRef.current.successMessage) {
          toast.info(configRef.current.successMessage)
        }
        return true
      } catch (err) {
        toast.error(getErrorMessage(err))
        return false
      }
    },
    [mutate],
  )

  return { submit, isLoading }
}
