import { baseApi } from '@/shared/api'
import { ISignInRequest } from '../model/types'
import { IUserResponse } from '@/entities/user'

export const authApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    signIn: builder.mutation<IUserResponse, ISignInRequest>({
      query: (credentials) => ({
        url: '/auth/login',
        method: 'POST',
        body: credentials,
        credentials: 'include',
      }),
      invalidatesTags: ['Auth'],
    }),

    logout: builder.mutation<void, void>({
      query: () => ({
        url: 'auth/private/logout',
        method: 'POST',
        credentials: 'include',
      }),
    }),

    // Пока не используется отсюда
    // getAccess: builder.mutation<void, void>({
    //   query: () => ({
    //     url: 'auth/private/obtain',
    //     method: 'POST',
    //     credentials: 'include',
    //   }),
    //   invalidatesTags: ['Auth'],
    // }),
  }),
})

export const {
  useSignInMutation,
  useLogoutMutation,
  // useGetAccessMutation,
} = authApi
