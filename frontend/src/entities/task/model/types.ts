import {
  TPagination,
  TSolution,
  TTask,
  TTaskWithStudents,
} from '@/shared/model'
import { TTaskSchemaCreate, TTaskSchemaUpdate } from './schema'

export interface ITaskFieldsRef {
  validate: () => Promise<boolean>
  getFieldsData: () => TTaskSchemaCreate | TTaskSchemaUpdate
  reset: () => void
}

export interface ICreateTaskRequest {
  classes?: number[]
  description: string
  students?: number[]
  title: string
  file?: File[]
}

export interface ICreateTaskResponse {
  solutions?: TSolution[]
  task: TTask
}

export interface IGetTasksResponse {
  data: TTaskWithStudents[]
  pagination?: TPagination
}

export interface IGetTasksQuery {
  'per-page'?: number
  search?: string
}

export interface IUpdateTaskRequest {
  id: number
  description: string
  students?: number[]
  title: string
  file?: File[]
  delete_files?: number[]
}

export interface IUpdateTaskResponse {
  task: TTask
}
