import { UIMatch, useMatches, useParams } from 'react-router'

type BreadcrumbItem = {
  label: string | ((params: Record<string, string | undefined>) => string)
  to?: string
}

export interface Crumb {
  label: string
  to?: string
}

export const useBreadcrumbs = (): Crumb[] => {
  const matches = useMatches()
  const params = useParams()

  const match = [...matches].reverse().find((m: UIMatch) => {
    const handle = m.handle as { breadcrumbs?: BreadcrumbItem[] } | undefined
    return !!handle?.breadcrumbs
  })
  if (!match) return []

  const { breadcrumbs } = match.handle as { breadcrumbs: BreadcrumbItem[] }

  return breadcrumbs.map((crumb) => ({
    label:
      typeof crumb.label === 'function' ? crumb.label(params) : crumb.label,
    to: crumb.to,
  }))
}
