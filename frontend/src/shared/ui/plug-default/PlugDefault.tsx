import Text from '../text/Text'
import styles from './styles.module.scss'

interface IPlugDefaultProps {
  title?: string
}

const PlugDefault = (props: IPlugDefaultProps) => {
  const { title = 'Ничего не найдено 🥲' } = props
  return (
    <div className={styles.wrapper}>
      <Text size='20'>{title}</Text>
    </div>
  )
}

export default PlugDefault
