import { FC, ReactNode } from 'react'
import styles from './styles.module.scss'

const Header: FC = (): ReactNode => {
  const { logo } = styles

  return (
    <header>
      <img
        className={logo}
        src='../../../../public/skadi_logo.png'
        height='32'
        alt='IT-Школа "Скади"'
      ></img>
    </header>
  )
}

export default Header
