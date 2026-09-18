import { ReactNode, ChangeEvent, useId } from 'react'
import type { Ref } from 'react'
import styles from './styles.module.scss'
import commonStyles from '../styles/common.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'

type ResizeOption = 'none' | 'vertical' | 'horizontal' | 'both'

interface IProps {
  ref?: Ref<HTMLTextAreaElement>
  value: string
  label?: string
  description?: string
  placeholder?: string
  fluid?: boolean
  size?: 's' | 'm'
  disabled?: boolean
  isValid?: boolean
  required?: boolean
  maxLength?: number
  resize?: ResizeOption
  onChange: (value: string) => void
}

const Textarea = ({
  ref,
  placeholder = 'Ввод...',
  label,
  value,
  disabled,
  isValid = true,
  description,
  fluid,
  size = 'm',
  required,
  maxLength,
  resize = 'vertical',
  onChange,
}: IProps): ReactNode => {
  const textareaId = useId()
  const descriptionId = `${textareaId}-description`

  const handleChange = (e: ChangeEvent<HTMLTextAreaElement>): void => {
    onChange(e.target.value)
  }

  const wrapperClassName = getUIClasses(
    styles.wrapper,
    { fluid, additionalClasses: isValid ? [] : [commonStyles.invalidText] },
    commonStyles,
  )

  const textareaClassName = getUIClasses(
    styles.textarea,
    { size, fluid, additionalClasses: isValid ? [] : [commonStyles.invalid] },
    commonStyles,
    styles,
  )

  return (
    <div className={wrapperClassName}>
      {label && (
        <label htmlFor={textareaId} className={styles.textareaTitle}>
          {label}
          {required && (
            <span className={styles.required} aria-label='обязательное поле'>
              *
            </span>
          )}
        </label>
      )}
      <textarea
        ref={ref}
        id={textareaId}
        value={value}
        disabled={disabled}
        required={required}
        maxLength={maxLength}
        onChange={handleChange}
        placeholder={placeholder}
        className={textareaClassName}
        style={{ resize }}
        aria-invalid={!isValid}
        aria-describedby={description ? descriptionId : undefined}
        aria-required={required}
      />
      {description && (
        <p id={descriptionId} className={styles.textareaDescription}>
          {description}
        </p>
      )}
    </div>
  )
}

Textarea.displayName = 'Textarea'

export default Textarea
