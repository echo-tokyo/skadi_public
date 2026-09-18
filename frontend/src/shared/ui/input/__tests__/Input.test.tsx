import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, expect, it, vi } from 'vitest'
import Input from '../Input'

describe('Input', () => {
  describe('рендеринг', () => {
    it('рендерит поле ввода', () => {
      render(<Input value='' onChange={vi.fn()} />)
      expect(screen.getByRole('textbox')).toBeTruthy()
    })

    it('отображает title как label', () => {
      render(<Input value='' onChange={vi.fn()} title='Email' />)
      expect(screen.getByLabelText('Email')).toBeTruthy()
    })

    it('показывает * при required', () => {
      render(<Input value='' onChange={vi.fn()} title='Email' required />)
      expect(screen.getByLabelText(/обязательное поле/i)).toBeTruthy()
    })

    it('отображает description', () => {
      render(<Input value='' onChange={vi.fn()} description='Подсказка' />)
      expect(screen.getByText('Подсказка')).toBeTruthy()
    })

    it('добавляет aria-invalid при isValid=false', () => {
      render(<Input value='' onChange={vi.fn()} isValid={false} />)
      expect(screen.getByRole('textbox').getAttribute('aria-invalid')).toBe('true')
    })

    it('не добавляет aria-invalid при isValid=true', () => {
      render(<Input value='' onChange={vi.fn()} isValid={true} />)
      expect(screen.getByRole('textbox').getAttribute('aria-invalid')).toBeNull()
    })

    it('отключает поле при disabled', () => {
      render(<Input value='' onChange={vi.fn()} disabled />)
      expect((screen.getByRole('textbox') as HTMLInputElement).disabled).toBe(true)
    })
  })

  describe('onChange', () => {
    it('вызывает onChange при вводе', async () => {
      const onChange = vi.fn()
      render(<Input value='' onChange={onChange} />)
      await userEvent.type(screen.getByRole('textbox'), 'a')
      expect(onChange).toHaveBeenCalledWith('a')
    })

    it('не вызывает onChange при disabled', async () => {
      const onChange = vi.fn()
      render(<Input value='' onChange={onChange} disabled />)
      await userEvent.type(screen.getByRole('textbox'), 'a')
      expect(onChange).not.toHaveBeenCalled()
    })
  })

  describe('password', () => {
    it('скрывает пароль по умолчанию', () => {
      render(<Input value='secret' onChange={vi.fn()} type='password' />)
      const input = screen.getByDisplayValue('secret') as HTMLInputElement
      expect(input.type).toBe('password')
    })

    it('показывает пароль при клике на toggle', async () => {
      render(<Input value='secret' onChange={vi.fn()} type='password' />)
      await userEvent.click(screen.getByRole('button'))
      const input = screen.getByDisplayValue('secret') as HTMLInputElement
      expect(input.type).toBe('text')
    })

    it('скрывает пароль при повторном клике на toggle', async () => {
      render(<Input value='secret' onChange={vi.fn()} type='password' />)
      await userEvent.click(screen.getByRole('button'))
      await userEvent.click(screen.getByRole('button'))
      const input = screen.getByDisplayValue('secret') as HTMLInputElement
      expect(input.type).toBe('password')
    })
  })
})
