import { describe, expect, it } from 'vitest'
import { getErrorMessage } from '../query-error'

describe('getErrorMessage', () => {
  it('возвращает fallback при отсутствии ошибки', () => {
    expect(getErrorMessage(null)).toBe('Произошла ошибка')
    expect(getErrorMessage(undefined)).toBe('Произошла ошибка')
    expect(getErrorMessage(0)).toBe('Произошла ошибка')
  })

  it('принимает кастомный fallback', () => {
    expect(getErrorMessage(null, 'Что-то пошло не так')).toBe('Что-то пошло не так')
  })

  it('достаёт message из FetchBaseQueryError', () => {
    const error = { data: { message: 'Пользователь не найден' } }
    expect(getErrorMessage(error)).toBe('Пользователь не найден')
  })

  it('возвращает fallback если data.message отсутствует', () => {
    expect(getErrorMessage({ data: {} })).toBe('Произошла ошибка')
    expect(getErrorMessage({ data: null })).toBe('Произошла ошибка')
  })

  it('достаёт message из SerializedError', () => {
    const error = { message: 'Network Error' }
    expect(getErrorMessage(error)).toBe('Network Error')
  })

  it('возвращает fallback если SerializedError.message undefined', () => {
    const error = { message: undefined }
    expect(getErrorMessage(error)).toBe('Произошла ошибка')
  })

  it('возвращает fallback для неизвестной формы ошибки', () => {
    expect(getErrorMessage('строка')).toBe('Произошла ошибка')
    expect(getErrorMessage(42)).toBe('Произошла ошибка')
    expect(getErrorMessage({ foo: 'bar' })).toBe('Произошла ошибка')
  })
})
