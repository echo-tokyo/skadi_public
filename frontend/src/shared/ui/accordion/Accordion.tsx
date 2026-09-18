import { ReactNode, useState } from 'react'
import Button from '../button/Button'
import Text from '../text/Text'
import styles from './styles.module.scss'
import ChevronIcon from '../icons/ChevronIcon'

interface IAccordionProps {
  label: string
  content: ReactNode
}

const Accordion = (props: IAccordionProps) => {
  const { label, content } = props
  const { accordion, accordionHeader, accordionContent, chevron, chevronOpen } =
    styles
  const [isOpen, setIsOpen] = useState<boolean>(false)

  const toggle = () => {
    setIsOpen((prev) => !prev)
  }

  return (
    <div className={accordion}>
      <div className={accordionHeader} onClick={toggle}>
        <Text>{label}</Text>
        <Button color='secondary' type='icon' size='s'>
          <span className={`${chevron}${isOpen ? ` ${chevronOpen}` : ''}`}>
            <ChevronIcon />
          </span>
        </Button>
      </div>
      {isOpen && <div className={accordionContent}>{content}</div>}
    </div>
  )
}

export default Accordion
