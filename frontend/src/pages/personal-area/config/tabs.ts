import { TRole } from '@/shared/model'
import { Archive } from '@/widgets/archive'
import { ClassManagement } from '@/widgets/class-management'
import { RoleManagement } from '@/widgets/role-management'
import { SolutionManagement } from '@/widgets/solution-management'
import { TaskManagement } from '@/widgets/task-management'
import { FC } from 'react'

export interface ITabConfig {
  name: string
  component?: FC
  role: TRole
  navigateTo?: string
}

export const TAB_CONFIG: ITabConfig[] = [
  {
    name: 'Управление ролями',
    component: RoleManagement,
    role: 'admin',
  },
  {
    name: 'Управление группами',
    component: ClassManagement,
    role: 'admin',
  },
  {
    name: 'Домашние задания',
    component: TaskManagement,
    role: 'teacher',
  },
  {
    name: 'Решения учеников',
    component: SolutionManagement,
    role: 'teacher',
  },
  {
    name: 'Канбан-доска',
    role: 'student',
    navigateTo: '/personal-area/kanban',
  },
  {
    name: 'Архив решений',
    role: 'student',
    component: Archive,
  },
]
