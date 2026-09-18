import {
  ReactNode,
  useRef,
  useState,
  DragEvent,
  ChangeEvent,
  KeyboardEvent,
} from 'react'
import styles from '../styles/styles.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'
import Text from '../../text/Text'

interface IProps {
  accept?: string
  multiple?: boolean
  disabled?: boolean
  size?: 's' | 'm'
  hasFiles?: boolean
  onFiles: (files: FileList) => void
}

const Dropzone = ({
  accept,
  multiple,
  disabled,
  size = 'm',
  hasFiles,
  onFiles,
}: IProps): ReactNode => {
  const inputRef = useRef<HTMLInputElement>(null)
  const [isDragging, setIsDragging] = useState(false)

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    if (e.target.files) onFiles(e.target.files)
    if (inputRef.current) inputRef.current.value = ''
  }

  const handleClick = (): void => {
    if (!disabled) inputRef.current?.click()
  }

  const handleKeyDown = (e: KeyboardEvent<HTMLDivElement>): void => {
    if (!disabled && (e.key === 'Enter' || e.key === ' ')) {
      e.preventDefault()
      inputRef.current?.click()
    }
  }

  const handleDragOver = (e: DragEvent<HTMLDivElement>): void => {
    e.preventDefault()
    if (!disabled) setIsDragging(true)
  }

  const handleDragLeave = (e: DragEvent<HTMLDivElement>): void => {
    e.preventDefault()
    if (!e.currentTarget.contains(e.relatedTarget as Node)) setIsDragging(false)
  }

  const handleDrop = (e: DragEvent<HTMLDivElement>): void => {
    e.preventDefault()
    setIsDragging(false)
    if (!disabled) onFiles(e.dataTransfer.files)
  }

  const className = getUIClasses(
    styles.dropzone,
    {
      size,
      additionalClasses: [
        isDragging ? styles.dropzoneDragging : '',
        disabled ? styles.dropzoneDisabled : '',
      ].filter(Boolean),
    },
    styles,
  )

  return (
    <>
      <input
        ref={inputRef}
        type='file'
        accept={accept}
        multiple={multiple}
        disabled={disabled}
        required={false}
        className={styles.inputHidden}
        onChange={handleChange}
      />
      <div
        role='button'
        tabIndex={disabled ? -1 : 0}
        aria-disabled={disabled}
        className={className}
        onClick={handleClick}
        onKeyDown={handleKeyDown}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
      >
        <Text size='12' color='--color-gray'>
          {hasFiles ? 'Изменить файлы' : 'Выберите файл'}
        </Text>
        <Text size='12' color='--color-gray'>
          или перетащите сюда
        </Text>
      </div>
    </>
  )
}

Dropzone.displayName = 'Dropzone'
export default Dropzone
