import { baseApi } from '@/shared/api'
import { IUserResponse } from '../model/types'

export const userApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    getMe: builder.query<IUserResponse, void>({
      query: () => ({
        url: '/user/me',
        method: 'GET',
      }),
      providesTags: ['User'],
    }),
  }),
})

export const { useGetMeQuery } = userApi
