import { memo, useEffect, useRef, useState } from 'react'
import { dropTargetForElements } from '@atlaskit/pragmatic-drag-and-drop/element/adapter'
import invariant from 'tiny-invariant'
import { TSolution, TStatusId, TStatusName } from '@/shared/model'
import { KanbanCard } from './KanbanCard'
import styles from '../styles.module.scss'
import { Text, TTextColor } from '@/shared/ui'

interface IKanbanColumnProps {
  id: TStatusId
  title: TStatusName
  cards: TSolution[]
}

interface IKanbanColumnBodyProps {
  id: TStatusId
  cards: TSolution[]
}

const KanbanColumnBody = memo(({ id, cards }: IKanbanColumnBodyProps) => {
  const ref = useRef<HTMLDivElement>(null)
  const [isDraggedOver, setIsDraggedOver] = useState(false)

  useEffect(() => {
    const el = ref.current
    invariant(el)

    return dropTargetForElements({
      element: el,
      getData: () => ({ type: 'kanban-column', columnId: id }),
      canDrop: ({ source }) =>
        source.data.type === 'kanban-card' && source.data.columnId !== id,
      onDragEnter: () => setIsDraggedOver(true),
      onDragLeave: () => setIsDraggedOver(false),
      onDrop: () => setIsDraggedOver(false),
    })
  }, [id])

  return (
    <div
      ref={ref}
      className={`${styles.columnBody} ${isDraggedOver ? styles.columnBodyOver : ''}`}
    >
      {cards.map((solution) => (
        <KanbanCard key={solution.id} solution={solution} columnId={id} />
      ))}
    </div>
  )
})

KanbanColumnBody.displayName = 'KanbanColumnBody'

export const KanbanColumn = memo(({ id, title, cards }: IKanbanColumnProps) => {
  const textColor = (): TTextColor => {
    if (id === 1) return '--color-red'
    if (id === 2) return '--color-primary'
    if (id === 3) return '--color-green'
    return '--color-secondary'
  }

  return (
    <div className={styles.column}>
      <Text size='20' weight='500' color={textColor()}>
        {`${title} (${cards.length})`}
      </Text>
      <KanbanColumnBody id={id} cards={cards} />
    </div>
  )
})

KanbanColumn.displayName = 'KanbanColumn'
