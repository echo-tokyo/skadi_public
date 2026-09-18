import { TGrade } from '@/shared/model'
import { useMemo } from 'react'
import { GRADE_OPTIONS } from '@/shared/config'

export const useGradeLabel = (raw: TGrade | '') => {
  return useMemo(
    () => GRADE_OPTIONS.find((el) => el.value === raw)?.label ?? '',
    [raw],
  )
}
