import { Text, Textarea } from '@/shared/ui'
import styles from '../styles.module.scss'
import { TFile } from '@/shared/model'
import { FileDownload } from '@/features/file-download'
import { noop } from '@/shared/lib'

interface ITaskAnswerTeacherProps {
  answer: string
  fileAnswer: TFile[]
}

const TaskViewAnswer = ({ answer, fileAnswer }: ITaskAnswerTeacherProps) => {
  return (
    <div className={styles.card}>
      <Text size='20' weight='600'>
        Ответ
      </Text>
      <div className={styles.cardAnswerFields}>
        <Textarea
          label='Письменный ответ'
          placeholder='Письменного ответа нет'
          fluid
          disabled
          value={answer}
          onChange={noop}
        />
        <div className={styles.cardAnswerFiles}>
          {fileAnswer.length > 0 &&
            fileAnswer.map((el) => <FileDownload key={el.id} el={el} />)}
        </div>
      </div>
    </div>
  )
}

export default TaskViewAnswer
