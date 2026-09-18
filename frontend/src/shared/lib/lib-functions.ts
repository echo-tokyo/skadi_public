export const unixToDate = (unixDate: string) => {
  return new Date(unixDate).toLocaleDateString('ru')
}

export const formatFileSize = (bytes: number): string => {
  if (bytes >= 1024 * 1024) return `${(bytes / (1024 * 1024)).toFixed(2)} МБ`
  if (bytes >= 1024) return `${(bytes / 1024).toFixed(2)} КБ`
  return `${bytes} Б`
}

export const noop = (): void => {}
