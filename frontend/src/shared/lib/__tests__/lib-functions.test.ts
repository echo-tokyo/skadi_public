import { describe, expect, it } from 'vitest'
import { formatFileSize, unixToDate } from '../lib-functions'

describe('formatFileSize', () => {
  it('форматирует байты', () => {
    expect(formatFileSize(0)).toBe('0 Б')
    expect(formatFileSize(500)).toBe('500 Б')
    expect(formatFileSize(1023)).toBe('1023 Б')
  })

  it('переходит в КБ ровно на 1024', () => {
    expect(formatFileSize(1024)).toBe('1.00 КБ')
  })

  it('форматирует килобайты', () => {
    expect(formatFileSize(1536)).toBe('1.50 КБ')
    expect(formatFileSize(1024 * 1024 - 1)).toBe('1024.00 КБ')
  })

  it('переходит в МБ ровно на 1024 * 1024', () => {
    expect(formatFileSize(1024 * 1024)).toBe('1.00 МБ')
  })

  it('форматирует мегабайты', () => {
    expect(formatFileSize(1024 * 1024 * 2.5)).toBe('2.50 МБ')
    expect(formatFileSize(1024 * 1024 * 30)).toBe('30.00 МБ')
  })
})

describe('unixToDate', () => {
  it('форматирует дату в русской локали', () => {
    expect(unixToDate('2024-01-15')).toBe('15.01.2024')
  })
})
