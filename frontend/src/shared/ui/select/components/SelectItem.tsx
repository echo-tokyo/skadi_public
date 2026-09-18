import { memo, type ReactNode } from 'react'
import clsx from 'clsx'
import styles from '../styles/styles.module.scss'
import Check from '../../icons/CheckIcon'

interface SelectItemProps {
  label: string
  selected: boolean
  disabled?: boolean
  onSelect: () => void
}

const SelectItem = memo(
  ({ label, selected, disabled, onSelect }: SelectItemProps): ReactNode => (
    <div
      role='option'
      aria-selected={selected}
      aria-disabled={disabled}
      data-disabled={disabled || undefined}
      className={clsx(styles.item, selected && styles.itemChecked)}
      onClick={disabled ? undefined : onSelect}
    >
      <span className={styles.checkIndicator}>{selected && <Check />}</span>
      {label}
    </div>
  ),
)
SelectItem.displayName = 'SelectItem'

export default SelectItem
