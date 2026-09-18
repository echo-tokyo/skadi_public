import { TRole } from '@/shared/model'
import { useCallback, useEffect, useMemo, useState } from 'react'
import { useGetMembersInfiniteQuery } from '../api/member-api'
import { toast } from 'sonner'
import { getErrorMessage } from '@/shared/api'
import { useDebounce } from '@/shared/lib'

export const useMemberSelectOptions = (role: TRole, isFree = true) => {
  const [search, setSearch] = useState('')
  const debouncedSearch = useDebounce(search)

  const { data, fetchNextPage, hasNextPage, isFetchingNextPage, error } =
    useGetMembersInfiniteQuery({
      free: isFree,
      role: [role],
      'per-page': 10,
      search: debouncedSearch || undefined,
    })

  useEffect(() => {
    if (error) toast.error(getErrorMessage(error))
  }, [error])

  const options = useMemo(
    () =>
      data?.pages
        .flatMap((p) => p.data)
        .map((el) => ({
          label: el.profile.fullname,
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
