import {
  createApi,
  fetchBaseQuery,
  type BaseQueryFn,
  type FetchArgs,
  type FetchBaseQueryError,
} from '@reduxjs/toolkit/query/react'
import { Mutex } from 'async-mutex'

interface IAuthActions {
  onAuthFailure: () => void
}

const mutex = new Mutex()
let authActions: IAuthActions | null = null

const baseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_API_BASE_URL || '',
  credentials: 'include',
  paramsSerializer: (params) => {
    const searchParams = new URLSearchParams()
    Object.entries(params).forEach(([key, value]) => {
      if (Array.isArray(value)) {
        value.forEach((item) => searchParams.append(key, item))
      } else if (value !== undefined && value !== null) {
        searchParams.append(key, String(value))
      }
    })
    return searchParams.toString()
  },
})

export const initializeAuthActions = (actions: IAuthActions): void => {
  authActions = actions
}

const baseQueryWithReauth: BaseQueryFn<
  string | FetchArgs,
  unknown,
  FetchBaseQueryError
> = async (args, api, extraOptions) => {
  if (!authActions) {
    throw new Error(
      'Auth actions not initialized. Call initializeAuthActions() in app layer.',
    )
  }

  await mutex.waitForUnlock()

  let result = await baseQuery(args, api, extraOptions)

  if (result.error?.status === 401) {
    if (!mutex.isLocked()) {
      const release = await mutex.acquire()

      try {
        const refreshResult = await baseQuery(
          {
            url: '/auth/private/obtain',
            method: 'POST',
          },
          api,
          extraOptions,
        )

        if (refreshResult.error?.status === 401) {
          authActions.onAuthFailure()
        } else if (!refreshResult.error) {
          result = await baseQuery(args, api, extraOptions)
        }
      } finally {
        release()
      }
    } else {
      await mutex.waitForUnlock()
      result = await baseQuery(args, api, extraOptions)
    }
  }
  return result
}

export const baseApi = createApi({
  reducerPath: 'api',
  tagTypes: [
    'Auth',
    'User',
    'Member',
    'Class',
    'Task',
    'Solution',
    'StudentSolution',
    'Comment',
  ],
  baseQuery: baseQueryWithReauth,
  endpoints: () => ({}),
})
