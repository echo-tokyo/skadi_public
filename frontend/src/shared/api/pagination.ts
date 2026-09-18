import type { TPagination } from '@/shared/model'

export const DEFAULT_PER_PAGE = 20

export type WithPagination = { pagination?: TPagination }

export const paginatedInfiniteQueryOptions = {
  initialPageParam: 1,
  getNextPageParam: ({ pagination }: WithPagination): number | undefined => {
    if (!pagination) return undefined
    return pagination.page < pagination.pages ? pagination.page + 1 : undefined
  },
}
