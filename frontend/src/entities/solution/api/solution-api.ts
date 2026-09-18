import {
  baseApi,
  DEFAULT_PER_PAGE,
  paginatedInfiniteQueryOptions,
} from '@/shared/api'
import { TStatusId } from '@/shared/model'
import {
  IGetSolutionByIdResponse,
  IGetSolutionsQuery,
  IGetSolutionsResponse,
  IUpdateSolutionByStudentResponse,
  IUpdateSolutionByTeacherRequest,
  IUpdateSolutionByTeacherResponse,
} from '../model/types'

export const solutionApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    updateSolutionByStudent: builder.mutation<
      IUpdateSolutionByStudentResponse,
      { id: number; data: FormData }
    >({
      query: ({ data, id }) => ({
        url: `/solution/for-student/${id}`,
        method: 'PATCH',
        body: data,
      }),
      async onQueryStarted({ id, data }, { dispatch, queryFulfilled }) {
        const statusId = data.get('status_id')
        const patch = dispatch(
          solutionApi.util.updateQueryData(
            'getSolutionsForStudent',
            undefined,
            (draft) => {
              if (statusId === null) return
              const solution = draft.data.find((s) => s.id === id)
              if (solution) solution.status.id = Number(statusId) as TStatusId
            },
          ),
        )
        try {
          await queryFulfilled
          dispatch(solutionApi.util.invalidateTags([{ type: 'Solution', id }]))
        } catch {
          patch.undo()
        }
      },
    }),

    updateSolutionByTeacher: builder.mutation<
      IUpdateSolutionByTeacherResponse,
      { id: number; data: IUpdateSolutionByTeacherRequest }
    >({
      query: ({ data, id }) => ({
        url: `/solution/for-teacher/${id}`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: ['Solution'],
    }),

    getSolutionById: builder.query<IGetSolutionByIdResponse, number>({
      query: (id) => ({
        url: `solution/${id}`,
        method: 'GET',
      }),
      providesTags: (_result, _error, id) => [{ type: 'Solution', id }],
    }),

    getSolutions: builder.infiniteQuery<
      IGetSolutionsResponse,
      IGetSolutionsQuery,
      number
    >({
      query: ({ queryArg, pageParam }) => {
        const { 'per-page': perPage, ...rest } = queryArg
        return {
          url: '/solution/for-teacher',
          method: 'GET',
          params: {
            ...rest,
            page: pageParam,
            'per-page': perPage ?? DEFAULT_PER_PAGE,
          },
        }
      },
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Solution'],
    }),

    getSolutionsForStudent: builder.query<IGetSolutionsResponse, void>({
      query: () => ({
        url: '/solution/for-student',
        method: 'GET',
      }),
      providesTags: ['Solution'],
    }),

    getStudentSolutions: builder.infiniteQuery<
      IGetSolutionsResponse,
      IGetSolutionsQuery,
      number
    >({
      query: ({ queryArg, pageParam }) => {
        const { 'per-page': perPage, ...rest } = queryArg
        return {
          url: '/solution/for-student',
          method: 'GET',
          params: {
            ...rest,
            page: pageParam,
            'per-page': perPage ?? DEFAULT_PER_PAGE,
          },
        }
      },
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Solution'],
    }),

    deleteSolution: builder.mutation<void, number>({
      query: (id) => ({
        url: `/solution/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Solution'],
    }),
  }),
})

export const {
  useGetSolutionByIdQuery,
  useGetSolutionsInfiniteQuery,
  useGetStudentSolutionsInfiniteQuery,
  useUpdateSolutionByTeacherMutation,
  useUpdateSolutionByStudentMutation,
  useGetSolutionsForStudentQuery,
  useDeleteSolutionMutation,
} = solutionApi
