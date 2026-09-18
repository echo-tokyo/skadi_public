import { useGetSolutionsForStudentQuery } from '@/entities/solution'
import { getErrorMessage } from '@/shared/api'
import { CHECKED_STATUS_ID } from '@/shared/config'
import { TSolution } from '@/shared/model'
import { useEffect } from 'react'
import { toast } from 'sonner'

export const useGetSolutions = () => {
  const { data, isError, error, isLoading } = useGetSolutionsForStudentQuery()

  useEffect(() => {
    if (isError) {
      toast.error(getErrorMessage(error))
    }
  }, [error, isError])

  const solutions: TSolution[] | undefined = data?.data.filter(
    (el) => el.status.id !== CHECKED_STATUS_ID,
  )

  return { solutions: solutions ?? [], isLoading }
}
