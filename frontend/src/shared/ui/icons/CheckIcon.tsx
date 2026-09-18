import { memo, type SVGProps } from 'react'

const CheckIcon = memo((props: SVGProps<SVGSVGElement>) => (
  <svg
    width='12'
    height='12'
    viewBox='0 0 24 24'
    fill='none'
    stroke='currentColor'
    strokeWidth='2.5'
    strokeLinecap='round'
    strokeLinejoin='round'
    aria-hidden='true'
    {...props}
  >
    <polyline points='20 6 9 17 4 12' />
  </svg>
))

CheckIcon.displayName = 'Check'

export default CheckIcon
