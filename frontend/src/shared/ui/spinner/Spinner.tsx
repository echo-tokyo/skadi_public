import { FC } from 'react'
import styles from './styles.module.scss'

interface IProps {
  overlay?: boolean
  size?: 's' | 'm'
}

const Spinner: FC<IProps> = ({ overlay = false, size = 'm' }) => (
  <div className={overlay ? styles.overlay : styles.inline}>
    <div className={`${styles.spinner} ${styles[size]}`} />
  </div>
)

export default Spinner
