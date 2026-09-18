import { baseApi, paginatedInfiniteQueryOptions } from '@/shared/api'
import {
  ICreateCommentRequest,
  IGetCommentsQuery,
  IGetCommentsResponse,
} from '../model/types'
import { TComment, TRole } from '@/shared/model'

const COMMENTS_PER_PAGE = 20

export const commentApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    createComment: builder.mutation<
      TComment,
      {
        data: ICreateCommentRequest
        id: number
        role: TRole
      }
    >({
      query: ({ data, id }) => ({
        url: `/solution/${id}/comment`,
        body: data,
        method: 'POST',
      }),

      async onQueryStarted({ data, id, role }, { dispatch, queryFulfilled }) {
        const queryArg: IGetCommentsQuery = { id }

        const optimisticComment: TComment = {
          id: -Date.now(),
          message: data.message,
          role,
          created_at: new Date().toISOString(),
        }

        const patchResult = dispatch(
          commentApi.util.updateQueryData('getComments', queryArg, (draft) => {
            if (draft.pages.length > 0) {
              draft.pages[0].data.unshift(optimisticComment)
            } else {
              draft.pages.push({
                data: [optimisticComment],
                pagination: {
                  page: 1,
                  per_page: COMMENTS_PER_PAGE,
                  pages: 1,
                  total: 1,
                },
              })
              draft.pageParams.push(1)
            }
          }),
        )

        try {
          await queryFulfilled
          dispatch(baseApi.util.invalidateTags(['Comment']))
        } catch {
          patchResult.undo()
        }
      },
    }),

    getComments: builder.infiniteQuery<
      IGetCommentsResponse,
      IGetCommentsQuery,
      number
    >({
      query: ({ pageParam, queryArg }) => ({
        url: `/solution/${queryArg.id}/comment`,
        params: {
          page: pageParam,
          'per-page': COMMENTS_PER_PAGE,
        },
      }),
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Comment'],
    }),
  }),
})

export const { useCreateCommentMutation, useGetCommentsInfiniteQuery } =
  commentApi
