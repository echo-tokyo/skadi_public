import { ReactNode, ChangeEvent, useId, useState, useMemo } from 'react'
import type { Ref } from 'react'
import styles from './styles.module.scss'
import commonStyles from '../styles/common.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'
import Button from '../button/Button'
import { EyeIcon, EyeOffIcon } from '../icons'

interface IProps {
  ref?: Ref<HTMLInputElement>
  value: string
  title?: string
  description?: string
  placeholder?: string
  fluid?: boolean
  size?: 's' | 'm'
  disabled?: boolean
  isValid?: boolean
  type?: 'text' | 'password'
  required?: boolean
  onChange: (value: string) => void
}

interface PasswordToggleProps {
  isVisible: boolean
  onToggle: () => void
}

const PasswordToggle = ({ isVisible, onToggle }: PasswordToggleProps): ReactNode => (
  <i className={styles.passwordToggle}>
    <Button type='icon' size='s' color='secondary' onClick={onToggle} tabIndex={-1}>
      {isVisible ? <EyeOffIcon /> : <EyeIcon />}
    </Button>
  </i>
)

const Input = ({
  ref,
  placeholder = 'Ввод...',
  title,
  value,
  disabled,
  isValid = true,
  description,
  type = 'text',
  fluid,
  size = 'm',
  required,
  onChange,
}: IProps): ReactNode => {
  const inputId = useId()
  const descriptionId = `${inputId}-description`
  const [isVisible, setIsVisible] = useState<boolean>(false)

  const types = type === 'password' ? (isVisible ? 'text' : 'password') : 'text'

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    onChange(e.target.value)
  }

  const wrapperClassName = useMemo(
    () =>
      getUIClasses(
        styles.wrapper,
        { fluid, additionalClasses: isValid ? [] : [commonStyles.invalidText] },
        commonStyles,
      ),
    [fluid, isValid],
  )

  const inputClassName = useMemo(
    () =>
      getUIClasses(
        styles.input,
        { size, fluid, additionalClasses: isValid ? [] : [commonStyles.invalid] },
        commonStyles,
      ),
    [size, fluid, isValid],
  )

  return (
    <div className={wrapperClassName}>
      {title && (
        <label htmlFor={inputId} className={styles.inputTitle}>
          {title}
          {required && (
            <span className={styles.required} aria-label='обязательное поле'>
              *
            </span>
          )}
        </label>
      )}
      <div className={styles.inputField}>
        {type === 'password' && (
          <PasswordToggle isVisible={isVisible} onToggle={() => setIsVisible(!isVisible)} />
        )}
        <input
          autoComplete='off'
          ref={ref}
          id={inputId}
          value={value}
          type={types}
          disabled={disabled}
          required={required}
          onChange={handleChange}
          placeholder={placeholder}
          className={inputClassName}
          aria-invalid={!isValid || undefined}
          aria-describedby={description ? descriptionId : undefined}
          aria-required={required}
        />
      </div>
      {description && (
        <p
          id={descriptionId}
          className={isValid ? styles.inputDescription : styles.inputError}
        >
          {description}
        </p>
      )}
    </div>
  )
}

Input.displayName = 'Input'

export default Input
