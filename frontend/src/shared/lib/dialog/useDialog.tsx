import { useContext } from 'react'
import { DialogContext } from './context'
import { DialogContextType } from './types'

export const useDialog = (): DialogContextType => {
  const ctx = useContext(DialogContext)

  if (!ctx) {
    throw new Error('useDialog must be used within DialogProvider')
  }

  return ctx
}
