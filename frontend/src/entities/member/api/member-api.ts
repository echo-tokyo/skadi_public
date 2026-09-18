import {
  baseApi,
  DEFAULT_PER_PAGE,
  paginatedInfiniteQueryOptions,
} from '@/shared/api'
import {
  ICreateMemberRequest,
  IMember,
  IMembersQuery,
  IMembersResponse,
  IUpdateMemberRequest,
} from '../model/types'

export const memberApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    getMembers: builder.infiniteQuery<IMembersResponse, IMembersQuery, number>({
      query: ({ queryArg, pageParam }) => {
        const { 'per-page': perPage, ...rest } = queryArg
        return {
          url: '/user',
          method: 'GET',
          params: {
            ...rest,
            page: pageParam,
            'per-page': perPage ?? DEFAULT_PER_PAGE,
          },
        }
      },
      infiniteQueryOptions: paginatedInfiniteQueryOptions,
      providesTags: ['Member'],
    }),

    createMember: builder.mutation<IMember, ICreateMemberRequest>({
      query: (data) => ({
        url: '/user',
        method: 'POST',
        body: data,
      }),
      invalidatesTags: (_result, _error, arg) => {
        const tags: ('Member' | 'Class')[] = ['Member']
        if (arg.class_id !== undefined) tags.push('Class')
        return tags
      },
    }),

    updateMember: builder.mutation<
      IMember,
      { id: number; data: IUpdateMemberRequest }
    >({
      query: ({ id, data }) => ({
        url: `/user/${id}`,
        method: 'PUT',
        body: data,
      }),
      invalidatesTags: (_result, _error, arg) => {
        const tags: ('Member' | 'Class')[] = ['Member']
        if (arg.data.class_id !== undefined) tags.push('Class')
        return tags
      },
    }),

    deleteMember: builder.mutation<void, number>({
      query: (id) => ({
        url: `/user/${id}`,
        method: 'DELETE',
      }),
      invalidatesTags: ['Member', 'Class'],
    }),
  }),
})

export const {
  useCreateMemberMutation,
  useGetMembersInfiniteQuery,
  useDeleteMemberMutation,
  useUpdateMemberMutation,
} = memberApi
