import { Button, Text } from '@/shared/ui'
import { FC, ReactNode } from 'react'
import { useNavigate } from 'react-router'
import styles from './styles.module.scss'

const MainPage: FC = (): ReactNode => {
  const nav = useNavigate()
  return (
    <div className={styles.wrapper}>
      <div className={styles.card}>
        <div className={styles.title}>
          <Text color='--color-primary' size='64' weight='800'>
            Привет !
          </Text>
          <Text color='--color-white' size='20'>
            Это твой личный кабинет, здесь есть информация о домашних заданиях и
            о преподавателях
          </Text>
        </div>
        <div className={styles.actions}>
          <Button
            onClick={() => window.open('https://itscadi.ru/', '_blank')}
            fluid
            color='white'
          >
            Главный сайт
          </Button>
          <Button onClick={() => nav('/personal-area')} fluid>
            Войти
          </Button>
        </div>
      </div>
    </div>
  )
}

export default MainPage
