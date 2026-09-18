import { ManagementLayout } from '@/shared/ui'
import { useState } from 'react'
import { CreateClassButton } from '@/features/create-class'
import { useInfiniteClasses } from '../model/use-infinite-classes'
import { ClassCardItem } from './ClassCardItem'

const ClassManagement = () => {
  const [search, setSearch] = useState('')
  const {
    classes,
    hasMore,
    isFetchingNextPage,
    loadMore,
    isLoading,
    isFetching,
  } = useInfiniteClasses({ search, 'per-page': 10 })

  return (
    <ManagementLayout
      isFetching={isFetching}
      title='Управление группами'
      searchTitle='Найти группу'
      items={classes}
      isLoading={isLoading}
      onDebouncedSearch={setSearch}
      infiniteScrollOptions={{ hasMore, isFetchingNextPage, loadMore }}
      getKey={(item) => item.id}
      renderItem={(item) => <ClassCardItem classData={item} />}
      actions={<CreateClassButton />}
    />
  )
}

export default ClassManagement
