import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, expect, it } from 'vitest'
import Accordion from '../Accordion'

describe('Accordion', () => {
  describe('рендеринг', () => {
    it('отображает label', () => {
      render(<Accordion label='Заголовок' content='Контент' />)
      expect(screen.getByText('Заголовок')).toBeTruthy()
    })

    it('скрывает content по умолчанию', () => {
      render(<Accordion label='Заголовок' content='Контент' />)
      expect(screen.queryByText('Контент')).toBeNull()
    })
  })

  describe('раскрытие', () => {
    it('показывает content при клике', async () => {
      render(<Accordion label='Заголовок' content='Контент' />)
      await userEvent.click(screen.getByText('Заголовок'))
      expect(screen.getByText('Контент')).toBeTruthy()
    })

    it('скрывает content при повторном клике', async () => {
      render(<Accordion label='Заголовок' content='Контент' />)
      await userEvent.click(screen.getByText('Заголовок'))
      await userEvent.click(screen.getByText('Заголовок'))
      expect(screen.queryByText('Контент')).toBeNull()
    })

    it('рендерит ReactNode как content', async () => {
      const content = <span data-testid='custom-content'>Кастомный контент</span>
      render(<Accordion label='Заголовок' content={content} />)
      await userEvent.click(screen.getByText('Заголовок'))
      expect(screen.getByTestId('custom-content')).toBeTruthy()
    })
  })
})
