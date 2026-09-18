import { useGetSolutionByIdQuery } from '@/entities/solution'
import { getErrorMessage } from '@/shared/api'
import { useEffect } from 'react'
import { toast } from 'sonner'

export const useGetSolution = (id: string | undefined) => {
  const { data, error, isError, isLoading, isFetching } =
    useGetSolutionByIdQuery(Number(id))

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [error, isError])

  return {
    data,
    isLoading,
    isFetching,
  }
}
