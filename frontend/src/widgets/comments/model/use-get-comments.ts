import { useGetCommentsInfiniteQuery } from '@/entities/comment'
import { getErrorMessage } from '@/shared/api'
import { useEffect, useMemo } from 'react'
import { toast } from 'sonner'

export const useGetComments = ({ id }: { id: number }) => {
  const {
    data,
    isFetchingNextPage,
    fetchNextPage,
    hasNextPage,
    error,
    isError,
    isLoading,
  } = useGetCommentsInfiniteQuery({ id })

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [error, isError])

  const messages = useMemo(
    () =>
      data?.pages
        .flatMap((page) => page.data)
        .slice()
        .reverse() ?? [],
    [data?.pages],
  )

  return {
    messages,
    isFetchingNextPage,
    loadMore: fetchNextPage,
    hasMore: hasNextPage,
    isLoading,
  }
}
