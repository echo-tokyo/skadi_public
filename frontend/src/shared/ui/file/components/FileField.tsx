import { ReactNode } from 'react'
import styles from '../styles/styles.module.scss'
import commonStyles from '../../styles/common.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'
import Dropzone from './Dropzone'
import FileItem from './FileItem'

// TODO: нельзя скачать прикрепленные файлы

export interface IFileFieldProps {
  value?: File[]
  onChange: (files: File[]) => void
  attachments?: { id: number; name: string; size?: number }[]
  onRemoveAttachment?: (id: number) => void
  label?: string
  description?: string
  fluid?: boolean
  size?: 's' | 'm'
  disabled?: boolean
  isValid?: boolean
  required?: boolean
  accept?: string
  multiple?: boolean
}

const FileField = ({
  value = [],
  onChange,
  attachments = [],
  onRemoveAttachment,
  label,
  description,
  fluid,
  size = 'm',
  disabled,
  isValid = true,
  required,
  accept,
  multiple,
}: IFileFieldProps): ReactNode => {
  const totalCount = attachments.length + value.length

  const handleFiles = (incoming: FileList): void => {
    const existingNames = new Set([
      ...attachments.map((f) => f.name),
      ...value.map((f) => f.name),
    ])
    const added = Array.from(incoming).filter((f) => !existingNames.has(f.name))
    if (added.length === 0) return
    onChange([...value, ...added])
  }

  const handleRemoveLocal = (fileName: string): void => {
    onChange(value.filter((f) => f.name !== fileName))
  }

  const wrapperClassName = getUIClasses(
    styles.wrapper,
    { fluid, additionalClasses: isValid ? [] : [commonStyles.invalidText] },
    commonStyles,
  )

  return (
    <div className={wrapperClassName}>
      {label && (
        <label className={styles.title}>
          {label}
          {required && (
            <span className={styles.required} aria-label='обязательное поле'>
              *
            </span>
          )}
        </label>
      )}

      <Dropzone
        accept={accept}
        multiple={multiple}
        disabled={disabled}
        size={size}
        hasFiles={totalCount > 0}
        onFiles={handleFiles}
      />

      {description && <p className={styles.description}>{description}</p>}

      {totalCount > 0 && (
        <ul className={styles.fileList} aria-label='Выбранные файлы'>
          {attachments.map((file) => (
            <FileItem
              key={file.id}
              size={file.size}
              file={{ id: String(file.id), name: file.name }}
              onRemove={(id) => onRemoveAttachment?.(Number(id))}
            />
          ))}
          {value.map((file) => (
            <FileItem
              key={file.name}
              size={file.size}
              file={{ id: file.name, name: file.name }}
              onRemove={handleRemoveLocal}
            />
          ))}
        </ul>
      )}
    </div>
  )
}

FileField.displayName = 'FileField'
export default FileField
