import clsx from 'clsx'

interface UIClassOptions {
  fluid?: boolean
  size?: 's' | 'm'
  additionalClasses?: string[]
}

export const getUIClasses = (
  baseClass: string,
  options: UIClassOptions,
  styles: Record<string, string>,
  sizeStyles?: Record<string, string>,
): string => {
  const { fluid, size, additionalClasses = [] } = options
  const sizeSource = sizeStyles ?? styles

  const classes = [
    baseClass,
    fluid && styles.fluid,
    size === 's' && sizeSource.size_s,
    size === 'm' && sizeSource.size_m,
    ...additionalClasses,
  ]

  return clsx(classes)
}
