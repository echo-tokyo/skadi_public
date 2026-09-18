import type { FetchBaseQueryError } from '@reduxjs/toolkit/query'
import type { SerializedError } from '@reduxjs/toolkit'

type ApiErrorData = { message?: string }

const isFetchError = (error: unknown): error is FetchBaseQueryError =>
  typeof error === 'object' && error !== null && 'data' in error

const isSerializedError = (error: unknown): error is SerializedError =>
  typeof error === 'object' && error !== null && 'message' in error

export const getErrorMessage = (
  error: unknown,
  fallback = 'Произошла ошибка',
): string => {
  if (!error) {
    return fallback
  }
  if (isFetchError(error)) {
    return (error.data as ApiErrorData)?.message ?? fallback
  }
  if (isSerializedError(error)) {
    return error.message ?? fallback
  }
  return fallback
}
