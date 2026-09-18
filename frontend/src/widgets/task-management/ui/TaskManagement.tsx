import { ManagementLayout } from '@/shared/ui'
import { useState } from 'react'
import { useInfiniteTasks } from '../model/use-infinite-tasks'
import { TaskCardItem } from './TaskCardItem'
import { CreateTaskButton } from '@/features/create-task'

const TaskManagement = () => {
  const [search, setSearch] = useState('')
  const {
    tasks,
    hasMore,
    isFetchingNextPage,
    loadMore,
    isLoading,
    isFetching,
  } = useInfiniteTasks({ 'per-page': 10, search })

  return (
    <ManagementLayout
      isFetching={isFetching}
      title='Домашние задания'
      searchTitle='Поиск по названию'
      items={tasks}
      isLoading={isLoading}
      onDebouncedSearch={setSearch}
      infiniteScrollOptions={{ hasMore, isFetchingNextPage, loadMore }}
      getKey={(item) => item.task.id}
      renderItem={(item) => <TaskCardItem taskData={item} />}
      actions={<CreateTaskButton />}
    />
  )
}

export default TaskManagement
