import { FormProvider, useForm } from 'react-hook-form'
import { ReactNode, useEffect, useMemo } from 'react'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import styles from './styles.module.scss'
import {
  TSolutionTeacherSchema,
  TSolutionStudentSchema,
  solutionTeacherSchema,
  solutionStudentSchema,
} from '@/entities/solution'
import { TaskCardMode, TDisplayValues } from '../model/types'
import { CHECKED_STATUS_ID, STATUS_OPTIONS } from '@/shared/config'
import { Button } from '@/shared/ui'
import { TFile } from '@/shared/model'
import { UpdateSolutionButton } from '@/features/update-solution'
import TaskGeneral from './components/TaskGeneral'
import TaskDescription from './components/TaskDescription'
import TaskMaterials from './components/TaskMaterials'
import TaskViewAnswer from './components/TaskViewAnswer'
import TaskEditAnswer from './components/TaskEditAnswer'

interface ITaskCardProps {
  mode: TaskCardMode
  editableValues: TSolutionTeacherSchema | TSolutionStudentSchema
  displayValues: TDisplayValues
  schema: typeof solutionTeacherSchema | typeof solutionStudentSchema
  solutionId: number
  serverFiles: TFile[]
  sideBar: ReactNode
  isFetching: boolean
}

const SolutionCard = (props: ITaskCardProps) => {
  const {
    mode,
    editableValues,
    schema,
    displayValues,
    solutionId,
    serverFiles,
    sideBar,
    isFetching,
  } = props

  const isViewOnly = mode === 'student-view'
  const isTeacher = mode === 'teacher'

  const statusOptions = useMemo(() => {
    const validStatuses = schema.shape.status.values
    return STATUS_OPTIONS.filter(
      ({ value }) => validStatuses && Array.from(validStatuses).includes(value),
    ).map((opt) =>
      mode === 'student-edit' && opt.value === CHECKED_STATUS_ID
        ? { ...opt, disabled: true }
        : opt,
    )
  }, [schema, mode])

  const methods = useForm<z.infer<typeof schema>>({
    defaultValues: editableValues,
    resolver: zodResolver(schema),
  })

  const {
    reset,
    formState: { isDirty, isSubmitting },
  } = methods

  useEffect(() => {
    reset(editableValues)
  }, [editableValues, reset])

  return (
    <FormProvider {...methods}>
      {!isViewOnly && (
        <div className={styles.actions}>
          <Button
            onClick={() => reset(editableValues)}
            disabled={!isDirty || isFetching || isSubmitting}
            color='secondary'
          >
            Сбросить
          </Button>
          <UpdateSolutionButton id={solutionId} isTeacher={isTeacher} />
        </div>
      )}
      <div className={styles.content}>
        <div className={styles.cards}>
          <TaskGeneral
            displayValues={displayValues}
            statusOptions={statusOptions}
            mode={mode}
            disabled={isViewOnly}
          />
          <TaskDescription description={displayValues.description} />
          <TaskMaterials files={displayValues.files} />
          {mode === 'teacher' || mode === 'student-view' ? (
            <TaskViewAnswer
              answer={displayValues.answer}
              fileAnswer={displayValues.file_answer}
            />
          ) : (
            <TaskEditAnswer serverFiles={serverFiles} disabled={isViewOnly} />
          )}
        </div>
        {sideBar}
      </div>
    </FormProvider>
  )
}

export default SolutionCard
