import {
  IGetSolutionsQuery,
  useGetSolutionsInfiniteQuery,
} from '@/entities/solution'
import { getErrorMessage } from '@/shared/api'
import { useEffect, useMemo } from 'react'
import { toast } from 'sonner'

export const useInfiniteSolutions = (params: IGetSolutionsQuery) => {
  const {
    data,
    isFetchingNextPage,
    fetchNextPage,
    hasNextPage,
    error,
    isError,
    isLoading,
    isFetching,
  } = useGetSolutionsInfiniteQuery(params)

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [isError])

  const solutions = useMemo(
    () => data?.pages.flatMap((page) => page.data) ?? [],
    [data?.pages],
  )

  return {
    solutions,
    isFetchingNextPage,
    loadMore: fetchNextPage,
    hasMore: hasNextPage,
    isLoading,
    isFetching,
  }
}
