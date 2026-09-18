import { createContext } from 'react'
import { DialogContextType } from './types'

export const DialogContext = createContext<DialogContextType | null>(null)
