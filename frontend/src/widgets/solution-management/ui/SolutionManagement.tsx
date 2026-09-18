import { ManagementLayout, Select } from '@/shared/ui'
import { useState } from 'react'
import { useInfiniteSolutions } from '../model/use-infinite-solutions'
import { SolutionCardItem } from './SolutionCardItem'
import { CHECKING_STATUS_ID, STATUS_OPTIONS } from '@/shared/config'
import { TStatusId } from '@/shared/model'

const SolutionManagement = () => {
  const [search, setSearch] = useState('')
  const [status, setStatus] = useState<TStatusId | undefined>(
    CHECKING_STATUS_ID,
  )
  const {
    solutions,
    hasMore,
    isFetchingNextPage,
    loadMore,
    isLoading,
    isFetching,
  } = useInfiniteSolutions({ 'per-page': 10, search, status_id: status })

  return (
    <ManagementLayout
      isFetching={isFetching}
      title='Решения учеников'
      searchTitle='Поиск по названию'
      items={solutions}
      isLoading={isLoading}
      onDebouncedSearch={setSearch}
      infiniteScrollOptions={{ hasMore, isFetchingNextPage, loadMore }}
      getKey={(item) => item.id}
      renderItem={(item) => <SolutionCardItem solutionData={item} />}
      actions={
        <Select
          label='Фильтровать по:'
          options={STATUS_OPTIONS}
          value={status ?? ''}
          onChange={(val) =>
            setStatus((Number(val) || undefined) as TStatusId | undefined)
          }
        />
      }
    />
  )
}

export default SolutionManagement
