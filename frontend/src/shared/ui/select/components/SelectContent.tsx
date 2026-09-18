import {
  useState,
  useCallback,
  useMemo,
  useRef,
  useEffect,
  type ReactNode,
} from 'react'
import styles from '../styles/styles.module.scss'
import SelectItem from './SelectItem'
import Button from '../../button/Button'
import type { SelectOption } from '../types/types'
import Input from '../../input/Input'

const SEARCH_DEBOUNCE_MS = 300

interface SelectContentProps<T extends string | number> {
  options: SelectOption<T>[]
  /** Предзагруженные выбранные элементы — seed для форм редактирования */
  selectedOptions?: SelectOption<T>[]
  isSelected: (value: T) => boolean
  onSelect: (value: T) => void
  onReset: () => void
  onSave: () => void
  /** Показывать ли кнопки "Сбросить" / "Сохранить" */
  showButtons: boolean
  /** Сохранённые значения для вычисления isDirty (внешний value из Multi/Single) */
  committedValues?: readonly (string | number)[]
  listboxId: string
  searchable?: boolean
  searchPlaceholder?: string
  noResultsText?: string
  onSearchChange?: (query: string) => void
  onLoadMore?: () => void
  hasMore?: boolean
  isLoadingMore?: boolean
}

const SelectContent = <T extends string | number>({
  options,
  selectedOptions,
  isSelected,
  onSelect,
  onReset,
  onSave,
  showButtons,
  committedValues,
  listboxId,
  searchable,
  searchPlaceholder = 'Поиск...',
  noResultsText = 'Ничего не найдено',
  onSearchChange,
  onLoadMore,
  hasMore,
  isLoadingMore,
}: SelectContentProps<T>): ReactNode => {
  const [searchQuery, setSearchQuery] = useState('')
  const listRef = useRef<HTMLDivElement>(null)

  // Серверный поиск: дебаунс перед вызовом колбэка
  useEffect(() => {
    if (!onSearchChange) return
    const timer = setTimeout(
      () => onSearchChange(searchQuery),
      SEARCH_DEBOUNCE_MS,
    )
    return () => clearTimeout(timer)
  }, [searchQuery, onSearchChange])

  // Объединяем seed и серверные данные: seed идёт первым, сервер перезаписывает
  const allKnownOptions = useMemo(() => {
    const map = new Map<string | number, SelectOption<T>>()
    selectedOptions?.forEach((o) => map.set(o.value, o))
    options.forEach((o) => map.set(o.value, o))
    return [...map.values()]
  }, [options, selectedOptions])

  // Клиентская фильтрация — только когда нет серверного поиска
  const clientFiltered = useMemo<SelectOption<T>[] | null>(() => {
    if (onSearchChange) return null
    if (!searchQuery.trim()) return null
    const query = searchQuery.toLowerCase()
    return options.filter((o) => o.label.toLowerCase().includes(query))
  }, [options, searchQuery, onSearchChange])

  // Секция выбранных — из allKnownOptions (включает seed)
  const selectedSection = useMemo(
    () => allKnownOptions.filter((o) => isSelected(o.value)),
    [allKnownOptions, isSelected],
  )

  // Основной список — результаты поиска/сервера без уже выбранных
  const mainList = useMemo(
    () => (clientFiltered ?? options).filter((o) => !isSelected(o.value)),
    [clientFiltered, options, isSelected],
  )

  const isDirty = useMemo(() => {
    if (!committedValues) return true
    const draftValues = selectedOptions?.map((o) => o.value) ?? []
    return (
      draftValues.length !== committedValues.length ||
      draftValues.some((v) => !committedValues.includes(v))
    )
  }, [selectedOptions, committedValues])

  const handleSearchChange = useCallback((value: string) => {
    setSearchQuery(value)
  }, [])

  return (
    <div className={styles.contentInner}>
      {searchable && (
        <div className={styles.searchWrapper}>
          <Input
            fluid
            placeholder={searchPlaceholder}
            value={searchQuery}
            onChange={handleSearchChange}
          />
        </div>
      )}
      <div
        ref={listRef}
        role='listbox'
        id={listboxId}
        className={styles.viewport}
      >
        {selectedSection.length === 0 && mainList.length === 0 ? (
          <div className={styles.noResults}>{noResultsText}</div>
        ) : (
          <>
            {selectedSection.map((option) => (
              <SelectItem
                key={option.value}
                label={option.label}
                selected={true}
                disabled={option.disabled}
                onSelect={() => onSelect(option.value)}
              />
            ))}

            {selectedSection.length > 0 && mainList.length > 0 && (
              <div className={styles.sectionDivider} />
            )}

            {mainList.map((option) => (
              <SelectItem
                key={option.value}
                label={option.label}
                selected={false}
                disabled={option.disabled}
                onSelect={() => onSelect(option.value)}
              />
            ))}
          </>
        )}
      </div>
      {onLoadMore && hasMore && (
        <div className={styles.loadMoreWrapper}>
          <Button
            fluid
            onClick={onLoadMore}
            color='ghost'
            isLoading={isLoadingMore}
          >
            {isLoadingMore ? 'Загрузка...' : 'Загрузить ещё'}
          </Button>
        </div>
      )}
      {showButtons && (
        <div className={styles.footer}>
          <Button color='secondary' fluid onClick={onReset}>
            Сбросить
          </Button>
          <Button color='primary' fluid onClick={onSave} disabled={!isDirty}>
            Сохранить
          </Button>
        </div>
      )}
    </div>
  )
}

export default SelectContent
