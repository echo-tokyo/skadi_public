import type { ReactNode } from 'react'
import styles from '../styles/styles.module.scss'
import commonStyles from '../../styles/common.module.scss'
import { getUIClasses } from '@/shared/lib/classNames/getUIClasses'
import SingleSelect from './SingleSelect'
import MultiSelect from './MultiSelect'

export type { SelectOption, SelectProps } from '../types/types'

import type { SelectProps } from '../types/types'

const Select = <T extends string | number = string>(
  props: SelectProps<T>,
): ReactNode => {
  const { fluid, size = 'm', isValid = true } = props

  const wrapperClassName = getUIClasses(
    styles.wrapper,
    { fluid, additionalClasses: isValid ? [] : [commonStyles.invalidText] },
    commonStyles,
  )
  const triggerClassName = getUIClasses(
    styles.trigger,
    { size, fluid, additionalClasses: isValid ? [] : [commonStyles.invalid] },
    commonStyles,
  )

  if (props.multiple) {
    return (
      <MultiSelect
        {...props}
        wrapperClassName={wrapperClassName}
        triggerClassName={triggerClassName}
      />
    )
  }

  return (
    <SingleSelect
      {...props}
      wrapperClassName={wrapperClassName}
      triggerClassName={triggerClassName}
    />
  )
}

export default Select
