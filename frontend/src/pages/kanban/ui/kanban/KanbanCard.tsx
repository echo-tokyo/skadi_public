import { memo, useEffect, useRef, useState } from 'react'
import { draggable } from '@atlaskit/pragmatic-drag-and-drop/element/adapter'
import invariant from 'tiny-invariant'
import { TSolution, TStatusId } from '@/shared/model'
import styles from '../styles.module.scss'
import { useNavigate } from 'react-router'
import { Text } from '@/shared/ui'
import { unixToDate } from '@/shared/lib'

interface IKanbanCardProps {
  solution: TSolution
  columnId: TStatusId
}

export const KanbanCard = memo(({ solution, columnId }: IKanbanCardProps) => {
  const ref = useRef<HTMLDivElement>(null)
  const [dragging, setDragging] = useState(false)

  useEffect(() => {
    const el = ref.current
    invariant(el)

    return draggable({
      element: el,
      getInitialData: () => ({
        type: 'kanban-card',
        cardId: solution.id,
        columnId,
      }),
      onDragStart: () => setDragging(true),
      onDrop: () => setDragging(false),
    })
  }, [solution.id, columnId])

  const nav = useNavigate()

  return (
    <div
      ref={ref}
      onClick={() => nav(`/personal-area/kanban/solutions/${solution.id}`)}
      className={`${styles.card} ${dragging ? styles.dragging : ''}`}
    >
      <div className={styles.cardHeader}>
        <Text className={styles.cardTitle}>{solution.task.title}</Text>
        <Text className={styles.cardTitle}>
          {unixToDate(solution.updated_at ?? '-')}
        </Text>
      </div>
      {solution.task.description && (
        <Text
          size='14'
          color='--color-gray'
          weight='300'
          className={styles.cardDescription}
        >
          Описание: {solution.task.description}
        </Text>
      )}
    </div>
  )
})

KanbanCard.displayName = 'KanbanCard'
