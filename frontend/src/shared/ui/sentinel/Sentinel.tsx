import { Ref } from 'react'
import styles from './styles.module.scss'

interface IProps {
  ref: Ref<HTMLDivElement>
}

const Sentinel = ({ ref }: IProps) => {
  return <div ref={ref} className={styles.sentinel}></div>
}

export default Sentinel
