import { FC } from 'react'
import { KanbanBoard } from './kanban/KanbanBoard'
import { PlugDefault, Skeleton } from '@/shared/ui'
import { useGetSolutions } from '../model/use-get-solutions'
import styles from './styles.module.scss'
import { useShowSkeleton } from '@/shared/lib'

const Kanban: FC = () => {
  const { solutions, isLoading } = useGetSolutions()
  const showSkeleton = useShowSkeleton(isLoading, 200)

  if (showSkeleton) {
    return <Skeleton height={'100%'} />
  }

  if (solutions.length === 0 && !isLoading) {
    return <PlugDefault />
  }

  if (isLoading) {
    return null
  }

  return (
    <div className={styles.wrapper}>
      <KanbanBoard solutions={solutions} />
    </div>
  )
}

export default Kanban
