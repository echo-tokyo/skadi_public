import {
  TPagination,
  TProfile,
  TSolution,
  TStatus,
  TStatusId,
  TTask,
} from '@/shared/model'

export interface IGetSolutionByIdResponse {
  other_students: TProfile[]
  solution: TSolution
}

export interface IGetSolutionsResponse {
  data: TSolution[]
  pagination: TPagination
}

export interface IUpdateSolutionByTeacherRequest {
  grade?: string
  status_id?: number
}

export interface IUpdateSolutionByTeacherResponse {
  answer?: string
  grade?: string
  id: number
  status: TStatus
  student?: TProfile
  task: TTask
  updated_at?: string
}

export interface IUpdateSolutionByStudentRequest {
  status_id?: TStatusId
  answer?: string
  delete_files?: number[]
  file?: File[]
}

export interface IUpdateSolutionByStudentResponse {
  answer: string
  id: string
  grade?: string
  status?: TStatus
  student: TProfile
  task?: TTask
  updated_at: string
}
export interface IGetSolutionsQuery {
  status_id?: TStatusId
  'per-page'?: number
  search?: string
}
