import {
  baseApi,
  DEFAULT_PER_PAGE,
  paginatedInfiniteQueryOptions,
} from '@/shared/api'
import {
  IClass,
  IClassQuery,
  IClassResponse,
  IClassRequest,
} from '../model/types'

export const classApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    getClasses: builder.infiniteQuery<IClassResponse, IClassQuery, number>({
      query: ({ queryArg, pageParam }) => {
        const { 'per-page': perPage, ...rest } = queryArg
        return {
          url: '/class',
          method: 'GET',
          params: {
            ...rest,
            page: pageParam,
            'per-page': perPage ?? DEFAULT_PER_PAGE,
          },
        }
      },
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Class'],
    }),

    createClass: builder.mutation<IClass, IClassRequest>({
      query: (data) => ({
        url: '/class',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: ['Class', 'Member'],
    }),

    editClass: builder.mutation<IClass, { id: number; data: IClassRequest }>({
      query: ({ id, data }) => ({
        url: `/class/${id}`,
        method: 'PATCH',
        body: data,
      }),
      invalidatesTags: ['Class', 'Member'],
    }),

    deleteClass: builder.mutation<void, number>({
      query: (id) => ({
        url: `/class/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Class', 'Member'],
    }),
  }),
})

export const {
  useGetClassesInfiniteQuery,
  useCreateClassMutation,
  useDeleteClassMutation,
  useEditClassMutation,
} = classApi
