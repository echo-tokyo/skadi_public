import { useEffect, useMemo } from 'react'
import { monitorForElements } from '@atlaskit/pragmatic-drag-and-drop/element/adapter'
import { useStudentUpdateSolutionStatus } from '@/features/update-solution'
import { TSolution, TStatusId } from '@/shared/model'
import { KANBAN_COLUMNS } from '../../config/columns'
import { KanbanColumn } from './KanbanColumn'
import styles from '../styles.module.scss'

interface IKanbanBoardProps {
  solutions: TSolution[]
}

const EMPTY_ARRAY: TSolution[] = []
export const KanbanBoard = ({ solutions }: IKanbanBoardProps) => {
  const { submit: updateStatus } = useStudentUpdateSolutionStatus()

  useEffect(() => {
    return monitorForElements({
      canMonitor: ({ source }) => source.data.type === 'kanban-card',
      onDrop({ source, location }) {
        const dropTarget = location.current.dropTargets[0]
        if (!dropTarget) return

        const cardId = source.data.cardId as number
        const newColumnId = dropTarget.data.columnId as TStatusId

        updateStatus({ id: cardId, status: newColumnId })
      },
    })
  }, [])

  const solutionsByColumn = useMemo(
    () =>
      solutions.reduce<Partial<Record<TStatusId, TSolution[]>>>((acc, s) => {
        const id = s.status.id
        if (!id) return acc
        if (!acc[id]) acc[id] = []
        acc[id]!.push(s)
        return acc
      }, {}),
    [solutions],
  )

  return (
    <div className={styles.board}>
      {KANBAN_COLUMNS.map((col) => (
        <KanbanColumn
          key={col.id}
          id={col.id}
          title={col.title}
          cards={solutionsByColumn[col.id] ?? EMPTY_ARRAY}
        />
      ))}
    </div>
  )
}
