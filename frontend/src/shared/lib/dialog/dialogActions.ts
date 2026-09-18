type CloseAll = () => void

let _closeAll: CloseAll | null = null

export const registerCloseAll = (fn: CloseAll) => {
  _closeAll = fn
}

export const closeAllDialogs = () => {
  _closeAll?.()
}
