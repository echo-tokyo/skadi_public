import {
  useState,
  useCallback,
  useMemo,
  useId,
  useEffect,
  useRef,
  type ReactNode,
} from 'react'
import * as Popover from '@radix-ui/react-popover'
import styles from '../styles/styles.module.scss'
import SelectContent from './SelectContent'
import ChevronIcon from '../../icons/ChevronIcon'
import type { MultipleProps, SelectOption } from '../types/types'

interface MultiSelectInternalProps<T extends string | number>
  extends MultipleProps<T> {
  wrapperClassName: string
  triggerClassName: string
}

const MultiSelect = <T extends string | number>({
  ref,
  options,
  selectedOptions,
  label,
  description,
  placeholder = 'Выберите',
  disabled,
  required,
  isValid = true,
  value,
  onChange,
  wrapperClassName,
  triggerClassName,
  searchable,
  searchPlaceholder,
  noResultsText,
  onSearchChange,
  onLoadMore,
  hasMore,
  isLoadingMore,
}: MultiSelectInternalProps<T>): ReactNode => {
  const [open, setOpen] = useState(false)
  // draftValue — локальная копия выбора. onChange вызывается только при "Сохранить".
  const [draftValue, setDraftValue] = useState<T[]>(value)

  // При открытии синхронизируем draft с внешним value
  useEffect(() => {
    if (open) setDraftValue(value)
  }, [open])

  // Карта всех известных options (value → SelectOption) для поиска label.
  // Ref, чтобы не вызывать лишних ре-рендеров при обновлении.
  const knownOptionsMap = useRef<Map<string | number, SelectOption<T>>>(
    new Map(selectedOptions?.map((o) => [o.value, o]) ?? []),
  )

  // Обновляем карту при каждом изменении options или selectedOptions
  useEffect(() => {
    selectedOptions?.forEach((o) => knownOptionsMap.current.set(o.value, o))
    options.forEach((o) => knownOptionsMap.current.set(o.value, o))
  }, [options, selectedOptions])

  // Ref для handleSave — чтобы всегда брать актуальный draftValue без пересоздания коллбэка
  const draftValueRef = useRef(draftValue)
  draftValueRef.current = draftValue

  const triggerId = useId()
  const labelId = `${triggerId}-label`
  const listboxId = useId()
  const descriptionId = `${triggerId}-description`

  // Триггер показывает сохранённое (внешнее) value, не draft
  const triggerLabel = useMemo(() => {
    if (value.length === 0) return null
    if (value.length <= 2) {
      const labels = value
        .map((v) => knownOptionsMap.current.get(v)?.label)
        .filter(Boolean)
      return labels.length > 0 ? labels.join(', ') : null
    }
    return `${value.length} выбрано`
  }, [value])

  const handleToggle = useCallback((optionValue: T): void => {
    setDraftValue((prev) =>
      prev.includes(optionValue)
        ? prev.filter((v) => v !== optionValue)
        : [...prev, optionValue],
    )
  }, [])

  const handleReset = useCallback(() => {
    onChange([])
    setDraftValue([])
    setOpen(false)
  }, [onChange])

  const handleSave = useCallback(() => {
    onChange(draftValueRef.current)
    setOpen(false)
  }, [onChange])

  // isSelected смотрит на draftValue, а не на внешний value
  const isSelected = useCallback((v: T) => draftValue.includes(v), [draftValue])

  // selectedOptions для SelectContent — текущий draft с labels из knownOptionsMap
  const enrichedSelectedOptions = useMemo(
    () =>
      draftValue
        .map((v) => knownOptionsMap.current.get(v))
        // eslint-disable-next-line eqeqeq
        .filter((o): o is SelectOption<T> => o != null),
    [draftValue],
  )

  return (
    <div className={wrapperClassName}>
      {label && (
        <span id={labelId} className={styles.label}>
          {label}
          {required && (
            <span className={styles.required} aria-label='обязательное поле'>
              *
            </span>
          )}
        </span>
      )}
      <Popover.Root open={open} onOpenChange={setOpen}>
        <Popover.Trigger
          ref={ref}
          id={triggerId}
          className={triggerClassName}
          disabled={disabled}
          aria-haspopup='listbox'
          aria-expanded={open}
          aria-controls={listboxId}
          aria-labelledby={label ? labelId : undefined}
          aria-invalid={!isValid}
          aria-describedby={description ? descriptionId : undefined}
          aria-required={required}
        >
          <span className={triggerLabel ? undefined : styles.placeholder}>
            {triggerLabel ?? placeholder}
          </span>
          <span className={styles.icon}>
            <ChevronIcon />
          </span>
        </Popover.Trigger>
        <Popover.Portal>
          <Popover.Content
            className={styles.content}
            style={{ width: 'var(--radix-popover-trigger-width)' }}
            sideOffset={4}
            align='start'
            onOpenAutoFocus={(e) => {
              if (!searchable) {
                e.preventDefault()
                const listbox = document.getElementById(listboxId)
                listbox?.focus()
              }
            }}
          >
            <SelectContent
              options={options}
              selectedOptions={enrichedSelectedOptions}
              isSelected={isSelected}
              onSelect={handleToggle}
              onReset={handleReset}
              onSave={handleSave}
              showButtons={true}
              committedValues={value}
              listboxId={listboxId}
              searchable={searchable}
              searchPlaceholder={searchPlaceholder}
              noResultsText={noResultsText}
              onSearchChange={onSearchChange}
              onLoadMore={onLoadMore}
              hasMore={hasMore}
              isLoadingMore={isLoadingMore}
            />
          </Popover.Content>
        </Popover.Portal>
      </Popover.Root>
      {description && (
        <p id={descriptionId} className={styles.description}>
          {description}
        </p>
      )}
    </div>
  )
}

export default MultiSelect
