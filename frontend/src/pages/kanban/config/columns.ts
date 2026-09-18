import { TStatusId, TStatusName } from '@/shared/model'

export type KanbanColumnConfig = {
  id: TStatusId
  title: TStatusName
}

export const KANBAN_COLUMNS: KanbanColumnConfig[] = [
  { id: 1, title: 'Бэклог' },
  { id: 2, title: 'В работе' },
  { id: 3, title: 'На проверке' },
]
