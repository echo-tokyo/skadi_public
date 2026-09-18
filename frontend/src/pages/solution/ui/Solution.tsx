import { FC, useMemo } from 'react'
import styles from './styles.module.scss'
import {
  SolutionCard,
  TaskCardMode,
  toTaskValues,
} from '@/widgets/solution-card'
import { useParams, Navigate } from 'react-router'
import { useGetSolution } from '../model/use-get-solution'
import { selectAuthenticatedUser } from '@/entities/user'
import { useAppSelector, useShowSkeleton } from '@/shared/lib'
import { Skeleton } from '@/shared/ui'
import { getSchemaByRole, toFormValuesByRole } from '@/entities/solution'
import { TFile } from '@/shared/model'
import { Comments } from '@/widgets/comments'
import Spinner from '@/shared/ui/spinner/Spinner'
import { CHECKED_STATUS_ID } from '@/shared/config'

const Solution: FC = () => {
  const { id } = useParams()
  const user = useAppSelector(selectAuthenticatedUser)
  const role = user.role
  const { data, isLoading, isFetching } = useGetSolution(id)
  const schema = getSchemaByRole(role)
  const showSkeleton = useShowSkeleton(isLoading, 200)
  const showSpinner = useShowSkeleton(isFetching, 200)

  if (role === 'admin') {
    return <Navigate to='/personal-area' replace />
  }

  const solution = data?.solution
  const serverFiles: TFile[] = data?.solution.files ?? []

  const mode: TaskCardMode =
    role === 'teacher'
      ? 'teacher'
      : solution?.status.id === CHECKED_STATUS_ID
        ? 'student-view'
        : 'student-edit'

  const modeForComments =
    solution?.status.id === CHECKED_STATUS_ID ? 'view' : 'edit'

  const displayValues = useMemo(() => toTaskValues(solution), [solution])
  const editableValues = useMemo(
    () => toFormValuesByRole(solution, role),
    [solution, role],
  )

  if (showSkeleton) {
    return <Skeleton height={'100%'} />
  }

  if (isLoading) {
    return null
  }

  return (
    <>
      <div className={styles.wrapper}>
        <SolutionCard
          solutionId={Number(id)}
          schema={schema}
          isFetching={isFetching}
          serverFiles={serverFiles}
          editableValues={editableValues}
          displayValues={displayValues}
          mode={mode}
          sideBar={
            <Comments
              solutionId={Number(id)}
              role={role}
              mode={modeForComments}
            />
          }
        />
      </div>
      {showSpinner && <Spinner overlay />}
    </>
  )
}

export default Solution
