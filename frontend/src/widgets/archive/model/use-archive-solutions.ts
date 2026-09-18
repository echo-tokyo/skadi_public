import { useGetStudentSolutionsInfiniteQuery } from '@/entities/solution'
import { getErrorMessage } from '@/shared/api'
import { useEffect, useMemo } from 'react'
import { toast } from 'sonner'
import { IGetSolutionsQuery } from '@/entities/solution'
import { CHECKED_STATUS_ID } from '@/shared/config'

export const useArchiveSolutions = (params: IGetSolutionsQuery) => {
  const {
    data,
    isFetchingNextPage,
    fetchNextPage,
    hasNextPage,
    error,
    isError,
    isLoading,
    isFetching,
  } = useGetStudentSolutionsInfiniteQuery({
    ...params,
    status_id: CHECKED_STATUS_ID,
  })

  useEffect(() => {
    if (isError) toast.error(getErrorMessage(error))
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
