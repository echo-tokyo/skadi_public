import { memo, type SVGProps } from 'react'

const EyeIcon = memo((props: SVGProps<SVGSVGElement>) => (
  <svg
    width='16'
    height='16'
    viewBox='0 0 24 24'
    fill='none'
    stroke='#777777'
    strokeWidth='2'
    strokeLinecap='round'
    strokeLinejoin='round'
    aria-hidden='true'
    {...props}
  >
    <path d='M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z' />
    <circle cx='12' cy='12' r='3' />
  </svg>
))

EyeIcon.displayName = 'EyeIcon'

export default EyeIcon
