import { ReactNode, useCallback, useEffect, useMemo, useState } from 'react'
import { Dialog } from '@/shared/ui'
import { DialogContext } from './context'
import { DialogContextType, DialogParams, DialogState, TPatch } from './types'
import { registerCloseAll } from './dialogActions'

const ANIMATION_DURATION = 150

export const DialogProvider = ({
  children,
}: {
  children: ReactNode
}): ReactNode => {
  const [dialogs, setDialogs] = useState<DialogState[]>([])

  const show = useCallback((params: DialogParams): string => {
    const id = crypto.randomUUID()
    setDialogs((prev) => [
      ...prev,
      { ...params, id, isClosing: false, isLoading: false },
    ])
    return id
  }, [])

  const hide = (id: string): void => {
    setDialogs((prev) =>
      prev.map((d) => (d.id === id ? { ...d, isClosing: true } : d)),
    )
    setTimeout(() => {
      setDialogs((prev) => prev.filter((m) => m.id !== id))
    }, ANIMATION_DURATION)
  }

  const update = useCallback((id: string, patch: TPatch) => {
    setDialogs((prev) =>
      prev.map((el) => (el.id === id ? { ...el, ...patch } : el)),
    )
  }, [])

  const setLoading = (id: string, isLoading: boolean): void => {
    setDialogs((prev) =>
      prev.map((d) => (d.id === id ? { ...d, isLoading } : d)),
    )
  }

  const handleClose = (dialog: DialogState): void => {
    dialog.onClose?.()
    hide(dialog.id)
  }

  const handleConfirm = async (dialog: DialogState): Promise<void> => {
    if (!dialog.onConfirm) {
      hide(dialog.id)
      return
    }

    setLoading(dialog.id, true)
    try {
      await dialog.onConfirm()
      hide(dialog.id)
    } catch {
      setLoading(dialog.id, false)
    }
  }

  const close = useCallback(() => {
    setDialogs([])
  }, [])

  useEffect(() => {
    registerCloseAll(close)
  }, [close])

  const contextValue: DialogContextType = useMemo(
    () => ({ show, update }),
    [show, update],
  )

  return (
    <DialogContext value={contextValue}>
      {children}
      {dialogs.map((dialog) => (
        <Dialog
          key={dialog.id}
          onClose={() => handleClose(dialog)}
          onConfirm={() => handleConfirm(dialog)}
          title={dialog.title}
          positiveText={dialog.positiveText}
          negativeText={dialog.negativeText}
          isClosing={dialog.isClosing}
          isConfirmDisabled={dialog.isConfirmDisabled}
          isConfirmLoading={dialog.isLoading}
          size={dialog.size}
        >
          {dialog.content}
        </Dialog>
      ))}
    </DialogContext>
  )
}
