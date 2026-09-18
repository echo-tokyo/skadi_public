import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, expect, it, vi } from 'vitest'
import Button from '../Button'

vi.mock('../../spinner/Spinner', () => ({
  default: () => <div data-testid='spinner' />,
}))

describe('Button', () => {
  describe('рендеринг', () => {
    it('отображает children', () => {
      render(<Button>Нажми меня</Button>)
      expect(screen.getByRole('button').textContent).toContain('Нажми меня')
    })

    it('по умолчанию имеет type="button"', () => {
      render(<Button>Text</Button>)
      expect(screen.getByRole('button').getAttribute('type')).toBe('button')
    })

    it('рендерит type="submit"', () => {
      render(<Button type='submit'>Submit</Button>)
      expect(screen.getByRole('button').getAttribute('type')).toBe('submit')
    })

    it('при type="icon" рендерит type="button"', () => {
      render(<Button type='icon'>Icon</Button>)
      expect(screen.getByRole('button').getAttribute('type')).toBe('button')
    })

    it('применяет tabIndex', () => {
      render(<Button tabIndex={3}>Text</Button>)
      expect(screen.getByRole('button').getAttribute('tabIndex')).toBe('3')
    })
  })

  describe('клик', () => {
    it('вызывает onClick при клике', async () => {
      const onClick = vi.fn()
      render(<Button onClick={onClick}>Click</Button>)
      await userEvent.click(screen.getByRole('button'))
      expect(onClick).toHaveBeenCalledTimes(1)
    })

    it('не вызывает onClick когда disabled', async () => {
      const onClick = vi.fn()
      render(
        <Button onClick={onClick} disabled>
          Click
        </Button>,
      )
      await userEvent.click(screen.getByRole('button'))
      expect(onClick).not.toHaveBeenCalled()
    })

    it('не падает если onClick не передан', async () => {
      render(<Button>Click</Button>)
      await expect(
        userEvent.click(screen.getByRole('button')),
      ).resolves.not.toThrow()
    })
  })

  describe('disabled', () => {
    it('добавляет атрибут disabled', () => {
      render(<Button disabled>Text</Button>)
      expect((screen.getByRole('button') as HTMLButtonElement).disabled).toBe(
        true,
      )
    })
  })

  describe('isLoading', () => {
    it('показывает Spinner', () => {
      render(<Button isLoading>Text</Button>)
      expect(screen.getByTestId('spinner')).toBeTruthy()
    })

    it('скрывает children (hidden класс)', () => {
      render(<Button isLoading>Visible Text</Button>)
      const contentSpan = screen
        .getByRole('button')
        .querySelector('span:first-child')
      expect(contentSpan?.className).toContain('hidden')
    })

    it('блокирует кнопку во время загрузки', () => {
      render(<Button isLoading>Text</Button>)
      expect((screen.getByRole('button') as HTMLButtonElement).disabled).toBe(
        true,
      )
    })

    it('не вызывает onClick во время загрузки', async () => {
      const onClick = vi.fn()
      render(
        <Button isLoading onClick={onClick}>
          Text
        </Button>,
      )
      await userEvent.click(screen.getByRole('button'))
      expect(onClick).not.toHaveBeenCalled()
    })
  })
})
