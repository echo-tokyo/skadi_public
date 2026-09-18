import {
  baseApi,
  DEFAULT_PER_PAGE,
  paginatedInfiniteQueryOptions,
} from '@/shared/api'
import {
  ICreateTaskResponse,
  IGetTasksQuery,
  IGetTasksResponse,
  IUpdateTaskResponse,
} from '../model/types'

export const taskApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    createTask: builder.mutation<ICreateTaskResponse, FormData>({
      query: (data) => ({
        url: '/task',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: ['Task', 'Solution'],
    }),

    updateTask: builder.mutation<IUpdateTaskResponse, FormData>({
      query: (data) => ({
        url: `/task/${data.get('id')}`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: ['Task', 'Solution'],
    }),

    getTasks: builder.infiniteQuery<IGetTasksResponse, IGetTasksQuery, number>({
      query: ({ queryArg, pageParam }) => {
        const { 'per-page': perPage, ...rest } = queryArg
        return {
          url: '/task',
          method: 'GET',
          params: {
            ...rest,
            page: pageParam,
            'per-page': perPage ?? DEFAULT_PER_PAGE,
          },
        }
      },
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Task'],
    }),

    deleteTask: builder.mutation<void, number>({
      query: (id) => ({
        url: `/task/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Task', 'Solution'],
    }),
  }),
})

export const {
  useCreateTaskMutation,
  useUpdateTaskMutation,
  useGetTasksInfiniteQuery,
  useDeleteTaskMutation,
} = taskApi
