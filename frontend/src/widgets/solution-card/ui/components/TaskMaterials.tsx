import { Text } from '@/shared/ui'
import styles from '../styles.module.scss'
import { TFile } from '@/shared/model'
import { FileDownload } from '@/features/file-download'

interface ITaskMaterialsProps {
  files: TFile[]
}

const TaskMaterials = ({ files }: ITaskMaterialsProps) => {
  return (
    <div className={styles.card}>
      <Text size='20' weight='600'>
        Материалы
      </Text>
      <div className={styles.cardMaterialsFields}>
        {files.length > 0 ? (
          files.map((el) => <FileDownload key={el.id} el={el} />)
        ) : (
          <Text>Пока пусто 🥲</Text>
        )}
      </div>
    </div>
  )
}

export default TaskMaterials
