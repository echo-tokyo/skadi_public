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
import type { SingleProps, SelectOption } from '../types/types'

interface SingleSelectInternalProps<T extends string | number>
  extends SingleProps<T> {
  wrapperClassName: string
  triggerClassName: string
}

const SingleSelect = <T extends string | number>({
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
  showButtons,
  wrapperClassName,
  triggerClassName,
  searchable,
  searchPlaceholder,
  noResultsText,
  onSearchChange,
  onLoadMore,
  hasMore,
  isLoadingMore,
}: SingleSelectInternalProps<T>): ReactNode => {
  const [open, setOpen] = useState(false)
  // draftValue используется только в режиме showButtons
  const [draftValue, setDraftValue] = useState<T | ''>(value)

  // При открытии синхронизируем draft с внешним value
  useEffect(() => {
    if (open && showButtons) setDraftValue(value)
  }, [open])

  const knownOptionsMap = useRef<Map<string | number, SelectOption<T>>>(
    new Map(selectedOptions?.map((o) => [o.value, o]) ?? []),
  )

  useEffect(() => {
    selectedOptions?.forEach((o) => knownOptionsMap.current.set(o.value, o))
    options.forEach((o) => knownOptionsMap.current.set(o.value, o))
  }, [options, selectedOptions])

  const draftValueRef = useRef(draftValue)
  draftValueRef.current = draftValue

  const triggerId = useId()
  const labelId = `${triggerId}-label`
  const listboxId = useId()
  const descriptionId = `${triggerId}-description`

  const selectedLabel = useMemo(() => {
    return (
      knownOptionsMap.current.get(value)?.label ??
      [...(selectedOptions ?? []), ...options].find((o) => o.value === value)
        ?.label
    )
  }, [value, options, selectedOptions])

  const handleSelect = useCallback(
    (optionValue: T): void => {
      if (showButtons) {
        setDraftValue((prev) => (prev === optionValue ? '' : optionValue))
      } else {
        onChange(value === optionValue ? '' : optionValue)
        setOpen(false)
      }
    },
    [value, onChange, showButtons],
  )

  const handleReset = useCallback(() => {
    onChange('' as T | '')
    setDraftValue('')
    setOpen(false)
  }, [onChange])

  const handleSave = useCallback(() => {
    onChange(draftValueRef.current)
    setOpen(false)
  }, [onChange])

  const isSelected = useCallback(
    (v: T) => (showButtons ? draftValue === v : value === v),
    [showButtons, draftValue, value],
  )

  const enrichedSelectedOptions = useMemo(() => {
    if (!showButtons) return selectedOptions
    const known = knownOptionsMap.current.get(draftValue as string | number)
    return known
      ? [known]
      : selectedOptions?.filter((o) => o.value === draftValue)
  }, [showButtons, draftValue, selectedOptions])

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
          <span className={selectedLabel ? undefined : styles.placeholder}>
            {selectedLabel ?? placeholder}
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
              onSelect={handleSelect}
              onReset={handleReset}
              onSave={handleSave}
              showButtons={!!showButtons}
              committedValues={showButtons ? (value ? [value] : []) : undefined}
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

export default SingleSelect
