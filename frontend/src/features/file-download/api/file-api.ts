import { baseApi } from '@/shared/api'

export const memberApi = baseApi.injectEndpoints({
  endpoints: (builder) => ({
    downloadFile: builder.query<Blob, number>({
      query: (id) => ({
        url: `/file/${id}`,
        method: 'GET',
        responseHandler: (response) => response.blob(),
      }),
    }),
  }),
})

export const { useLazyDownloadFileQuery } = memberApi
