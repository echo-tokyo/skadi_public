import { IMembersQuery, useGetMembersInfiniteQuery } from '@/entities/member'
import { getErrorMessage } from '@/shared/api'
import { useEffect, useMemo } from 'react'
import { toast } from 'sonner'

export const useInfiniteMembers = (params: IMembersQuery) => {
  const {
    data,
    isFetchingNextPage,
    fetchNextPage,
    hasNextPage,
    error,
    isError,
    isLoading,
    isFetching,
  } = useGetMembersInfiniteQuery(params)

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [isError])

  const members = useMemo(
    () => data?.pages.flatMap((page) => page.data) ?? [],
    [data?.pages],
  )

  return {
    members,
    isFetchingNextPage,
    loadMore: fetchNextPage,
    hasMore: hasNextPage,
    isLoading,
    isFetching,
  }
}
