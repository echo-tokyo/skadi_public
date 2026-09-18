import Text from '../text/Text'
import styles from './styles.module.scss'

interface IProps {
  message: string
  sender: string
  time: string
  align?: 'left' | 'right'
}

const Comment = (props: IProps) => {
  const { message, sender, time, align = 'left' } = props

  return (
    <div className={`${styles.comment} ${styles[align]}`}>
      <Text size='14' weight='500'>
        {sender}
      </Text>
      <div className={styles.bubble}>
        <Text size='14'>{message}</Text>
        <Text size='12' color='--color-gray' className={styles.time}>
          {time}
        </Text>
      </div>
    </div>
  )
}

export default Comment
