import styles from './styles.module.scss'

interface IDividerProps {
  vertical?: boolean
}

const Divider = (props: IDividerProps) => {
  const { divider, vertical } = styles
  return <div className={`${divider} ${props.vertical ? vertical : ''}`}></div>
}

export default Divider
