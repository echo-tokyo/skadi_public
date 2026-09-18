import { TComment, TPagination } from '@/shared/model'

export interface ICreateCommentRequest {
  message: string
}

export interface IGetCommentsQuery {
  id: number
  'per-page'?: number
}

export interface IGetCommentsResponse {
  data: TComment[]
  pagination: TPagination
}
