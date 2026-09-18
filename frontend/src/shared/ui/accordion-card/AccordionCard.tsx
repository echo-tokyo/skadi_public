import { ReactNode } from 'react'
import Accordion from '../accordion/Accordion'
import Divider from '../divider/Divider'
import Text from '../text/Text'
import styles from './styles.module.scss'

export interface IAccordionCardField {
  label: string
  value?: string | number
}

interface IAccordionCardProps {
  title: string
  fields: IAccordionCardField[]
  actions?: ReactNode
}

const AccordionCard = ({ title, fields, actions }: IAccordionCardProps) => (
  <Accordion
    label={title}
    content={
      <div className={styles.content}>
        {fields.map(({ label, value }) => (
          <div key={label} className={styles.contentItem}>
            <Text size='14'>{`${label}: ${value ?? '-'}`}</Text>
            <Divider />
          </div>
        ))}
        <div className={styles.cardActions}>{actions}</div>
      </div>
    }
  />
)

export default AccordionCard
