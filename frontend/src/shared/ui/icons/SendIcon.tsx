import { memo } from 'react'

const SendIcon = memo(() => (
  <svg
    width='16'
    height='16'
    viewBox='0 0 16 16'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    aria-hidden='true'
  >
    <path
      d='M8 13V3M8 3L4 7M8 3L12 7'
      stroke='currentColor'
      strokeWidth='1.5'
      strokeLinecap='round'
      strokeLinejoin='round'
    />
  </svg>
))

SendIcon.displayName = 'SendIcon'

export default SendIcon
