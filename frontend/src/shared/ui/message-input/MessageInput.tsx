import {
  useState,
  useRef,
  useEffect,
  useCallback,
  useMemo,
  KeyboardEvent,
  JSX,
  memo,
} from 'react'
import { SendIcon } from '../icons'
import styles from './styles.module.scss'
import Button from '../button/Button'

interface IMessageInputProps {
  onSubmit: (value: string) => void
  placeholder?: string
  disabled?: boolean
}

const IS_MOBILE = /Mobi|Android/i.test(navigator.userAgent)

const MessageInput = ({
  onSubmit,
  placeholder = 'Ввод...',
  disabled = false,
}: IMessageInputProps): JSX.Element => {
  const [value, setValue] = useState('')
  const textareaRef = useRef<HTMLTextAreaElement>(null)

  const trimmedValue = useMemo(() => value.trim(), [value])

  useEffect(() => {
    const el = textareaRef.current
    if (!el) return

    el.style.height = 'auto'
    el.style.height = `${el.scrollHeight}px`

    el.scrollTop = el.scrollHeight
  }, [value])

  const handleSubmit = useCallback((): void => {
    if (!trimmedValue || disabled) return
    onSubmit(trimmedValue)
    setValue('')
  }, [trimmedValue, disabled, onSubmit])

  const handleKeyDown = useCallback(
    (e: KeyboardEvent<HTMLTextAreaElement>): void => {
      if (e.key !== 'Enter' || IS_MOBILE) return
      if (!trimmedValue) {
        e.preventDefault()
        return
      }
      if (!e.shiftKey) {
        e.preventDefault()
        handleSubmit()
      }
    },
    [trimmedValue, handleSubmit],
  )

  return (
    <div className={styles.root}>
      <textarea
        ref={textareaRef}
        className={styles.textarea}
        value={value}
        rows={1}
        placeholder={placeholder}
        disabled={disabled}
        onChange={(e) => setValue(e.target.value)}
        onKeyDown={handleKeyDown}
      />
      <Button
        type='icon'
        aria-label='Отправить'
        disabled={disabled || !trimmedValue}
        onClick={handleSubmit}
      >
        <SendIcon />
      </Button>
    </div>
  )
}

MessageInput.displayName = 'MessageInput'

export default memo(MessageInput)
