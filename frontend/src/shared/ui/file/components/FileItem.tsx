import { ReactNode } from 'react'
import Button from '../../button/Button'
import Text from '../../text/Text'
import styles from '../styles/styles.module.scss'
import { FileDisplayItem } from '../types/types'
import { DownloadIcon } from '../../icons'
import { formatFileSize } from '@/shared/lib'

interface IProps {
  file: FileDisplayItem
  size?: number
  onRemove?: (id: string) => void
  onDownload?: (id: string) => void
}

const FileItem = (props: IProps): ReactNode => {
  const { file, size, onDownload, onRemove } = props

  return (
    <li className={styles.fileItem}>
      <Text className={styles.fileItemName} size='14'>
        {file.name}
      </Text>
      {size && (
        <Text size='14' color='--color-gray'>
          {formatFileSize(size)}
        </Text>
      )}
      {onRemove && (
        <Button
          type='icon'
          size='s'
          color='secondary'
          aria-label={`Удалить файл ${file.name}`}
          onClick={() => onRemove(file.id)}
        >
          <Text size='20' color='--color-gray' className={styles.icon}>
            ×
          </Text>
        </Button>
      )}
      {onDownload && (
        <Button
          type='icon'
          size='s'
          color='secondary'
          aria-label={`Скачать файл ${file.name}`}
          onClick={() => onDownload(file.id)}
        >
          <span style={{ color: 'var(--color-gray)', display: 'flex' }}>
            <DownloadIcon />
          </span>
        </Button>
      )}
    </li>
  )
}

FileItem.displayName = 'FileItem'
export default FileItem
