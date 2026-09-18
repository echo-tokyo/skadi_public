export type TTextColor =
  | '--color-primary'
  | '--color-secondary'
  | '--color-white'
  | '--color-gray-ultra-light'
  | '--color-gray-light'
  | '--color-gray'
  | '--color-gray-hover'
  | '--color-red'
  | '--color-green'

type TextWeight =
  | '100'
  | '200'
  | '300'
  | '400'
  | '500'
  | '600'
  | '700'
  | '800'
  | '900'

interface IProps {
  children: string[] | string
  size?: '64' | '20' | '16' | '14' | '12'
  weight?: TextWeight
  color?: TTextColor
  className?: string
}

const Text = (props: IProps) => {
  const {
    children,
    size = '16',
    weight = '400',
    color = '--color-secondary',
    className,
  } = props

  return (
    <span
      className={className}
      style={{
        fontSize: size + 'px',
        fontWeight: weight,
        color: `var(${color})`,
      }}
    >
      {children}
    </span>
  )
}

export default Text
