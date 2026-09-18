import { IClassQuery, useGetClassesInfiniteQuery } from '@/entities/class'
import { getErrorMessage } from '@/shared/api'
import { useEffect, useMemo } from 'react'
import { toast } from 'sonner'

export const useInfiniteClasses = (params: IClassQuery) => {
  const {
    data,
    isFetchingNextPage,
    fetchNextPage,
    hasNextPage,
    error,
    isError,
    isLoading,
    isFetching,
  } = useGetClassesInfiniteQuery(params)

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [error, isError])

  const classes = useMemo(
    () => data?.pages.flatMap((page) => page.data) ?? [],
    [data?.pages],
  )

  return {
    classes,
    isFetchingNextPage,
    loadMore: fetchNextPage,
    hasMore: hasNextPage,
    isLoading,
    isFetching,
  }
}
