import { useCallback, useEffect, useMemo, useState } from 'react'
import { toast } from 'sonner'
import { getErrorMessage } from '@/shared/api'
import { useDebounce } from '@/shared/lib'
import { useGetClassesInfiniteQuery } from '../api/class-api'

export const useClassSelectOptions = () => {
  const [search, setSearch] = useState('')
  const debouncedSearch = useDebounce(search)

  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, error } =
    useGetClassesInfiniteQuery({
      'per-page': 10,
      search: debouncedSearch,
    })

  useEffect(() => {
    if (error) toast.error(getErrorMessage(error))
  }, [error])

  const options = useMemo(
    () =>
      data?.pages
        .flatMap((p) => p.data)
        .map((el) => ({
          label: el.name,
          value: String(el.id),
        })) ?? [],
    [data?.pages],
  )

  const onSearchChange = useCallback((query: string) => {
    setSearch(query)
  }, [])

  return {
    options,
    hasNextPage,
    isFetchingNextPage,
    fetchNextPage,
    onSearchChange,
  }
}
