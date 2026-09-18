import { CSSProperties, FC } from 'react'
import styles from './styles.module.scss'

interface SkeletonProps {
  width?: CSSProperties['width']
  height?: CSSProperties['height']
}

const Skeleton: FC<SkeletonProps> = ({ width = '100%', height = '1rem' }) => (
  <div className={styles.skeleton} style={{ width, height }} />
)

export default Skeleton
