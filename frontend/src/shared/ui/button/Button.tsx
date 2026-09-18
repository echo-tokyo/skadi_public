import { MouseEvent, ReactNode } from 'react'
import styles from './styles.module.scss'
import commonStyles from '../styles/common.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'
import Spinner from '../spinner/Spinner'

interface IProps {
  children: ReactNode
  onClick?: () => void
  fluid?: boolean
  size?: 's' | 'm'
  disabled?: boolean
  isLoading?: boolean
  type?: 'button' | 'submit' | 'icon'
  color?: 'primary' | 'secondary' | 'ghost' | 'inverted' | 'white'
  tabIndex?: number
}

const Button = (props: IProps): ReactNode => {
  const {
    children,
    onClick,
    fluid = false,
    size = 'm',
    isLoading = false,
    disabled,
    type = 'button',
    color = 'primary',
    tabIndex = 0,
  } = props

  const handleClick = (
    e: MouseEvent<HTMLButtonElement, globalThis.MouseEvent>,
  ): void => {
    if (type === 'button') {
      e.preventDefault()
    }
    if (onClick) {
      onClick()
    }
  }

  const buttonClassName = getUIClasses(
    styles.button,
    {
      fluid,
      size,
      additionalClasses: [commonStyles[color], styles[type]],
    },
    commonStyles,
  )

  return (
    <button
      disabled={disabled || isLoading}
      onClick={(e) => handleClick(e)}
      className={buttonClassName}
      type={type !== 'icon' ? type : 'button'}
      tabIndex={tabIndex}
    >
      <span className={isLoading ? styles.hidden : undefined}>{children}</span>
      {isLoading && (
        <span className={styles.spinnerOverlay}>
          <Spinner size='s' />
        </span>
      )}
    </button>
  )
}

export default Button
