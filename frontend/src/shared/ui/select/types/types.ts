import type { Ref } from 'react'

export interface SelectOption<T extends string | number = string> {
  label: string
  value: T
  disabled?: boolean
}

interface BaseProps<T extends string | number = string> {
  ref?: Ref<HTMLButtonElement>
  options: SelectOption<T>[]
  /** Предзагруженные выбранные элементы (seed) — для форм редактирования,
   *  когда серверные данные ещё не пришли */
  selectedOptions?: SelectOption<T>[]
  placeholder?: string
  label?: string
  description?: string
  disabled?: boolean
  fluid?: boolean
  size?: 's' | 'm'
  required?: boolean
  isValid?: boolean
  /** Включить поиск по опциям */
  searchable?: boolean
  /** Placeholder для поля поиска */
  searchPlaceholder?: string
  /** Текст когда ничего не найдено */
  noResultsText?: string
  /**
   * Callback для серверного поиска. Вызывается с дебаунсом (300ms).
   * Когда передан — клиентская фильтрация отключается.
   */
  onSearchChange?: (query: string) => void
  /** Callback для загрузки следующей страницы */
  onLoadMore?: () => void
  /** Есть ли ещё страницы для загрузки */
  hasMore?: boolean
  /** Идёт ли загрузка следующей страницы */
  isLoadingMore?: boolean
}

export interface SingleProps<T extends string | number = string>
  extends BaseProps<T> {
  multiple?: false
  value: T | ''
  onChange: (value: T | '') => void
  /** Показывать ли кнопки "Сбросить" / "Сохранить" в single-режиме */
  showButtons?: boolean
}

export interface MultipleProps<T extends string | number = string>
  extends BaseProps<T> {
  multiple: true
  value: T[]
  onChange: (value: T[]) => void
  // кнопки в multi всегда показываются — без отдельного пропа
}

export type SelectProps<T extends string | number = string> =
  | SingleProps<T>
  | MultipleProps<T>

export type TPaginatedSelectField = {
  data: SelectOption[]
  selectedOptions?: SelectOption[]
  onSearchChange?: (query: string) => void
  onLoadMore?: () => void
  hasMore?: boolean
  isLoadingMore?: boolean
}
