import { Text, Textarea } from '@/shared/ui'
import styles from '../styles.module.scss'
import { noop } from '@/shared/lib'

interface ITaskDescriptionSectionProps {
  description: string
}

const TaskDescription = ({ description }: ITaskDescriptionSectionProps) => {
  return (
    <div className={styles.card}>
      <Text size='20' weight='600'>
        Описание
      </Text>
      <div className={styles.cardFields}>
        <Textarea
          label='Описание задания'
          fluid
          disabled
          value={description}
          onChange={noop}
        />
      </div>
    </div>
  )
}

export default TaskDescription
