import { ReactNode } from 'react'

type TSize = 's' | 'm' | 'l'
export type TPatch = Pick<DialogParams, 'isConfirmDisabled'>
export interface DialogParams {
  title?: string
  content: ReactNode
  positiveText?: string
  negativeText?: string
  onConfirm?: () => Promise<void>
  onClose?: () => void
  isConfirmDisabled?: boolean
  size?: TSize
}

export interface DialogState extends DialogParams {
  id: string
  isClosing: boolean
  isLoading: boolean
}

export interface DialogContextType {
  show: (params: DialogParams) => string
  update: (id: string, patch: TPatch) => void
}
