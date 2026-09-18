import { Input, Select, SelectOption, Text } from '@/shared/ui'
import styles from '../styles.module.scss'
import { Controller, useFormContext, useWatch } from 'react-hook-form'
import { TaskCardMode, TDisplayValues } from '../../model/types'
import { TStatusId } from '@/shared/model'
import {
  TSolutionBaseSchema,
  TSolutionTeacherSchema,
} from '@/entities/solution'
import { CHECKED_STATUS_ID, GRADE_OPTIONS } from '@/shared/config'
import { noop } from '@/shared/lib'
import { useGradeLabel } from '@/entities/solution/model/use-grade-label'

interface ITaskGeneralSectionProps {
  displayValues: Pick<TDisplayValues, 'title' | 'teacher' | 'student' | 'grade'>
  statusOptions: SelectOption<TStatusId>[]
  disabled?: boolean
  mode: TaskCardMode
}

const GradeField = ({ disabled }: { disabled: boolean }) => {
  const { control } = useFormContext<TSolutionTeacherSchema>()
  return (
    <Controller
      control={control}
      name='grade'
      render={({ field, fieldState }) => (
        <Select
          label='Оценка'
          fluid
          value={field.value}
          options={GRADE_OPTIONS}
          disabled={disabled}
          onChange={field.onChange}
          isValid={!fieldState.error}
          description={fieldState.error?.message}
        />
      )}
    />
  )
}

const TaskGeneral = ({
  displayValues,
  statusOptions,
  disabled = false,
  mode,
}: ITaskGeneralSectionProps) => {
  const { control } = useFormContext<TSolutionBaseSchema>()
  const status = useWatch({ control, name: 'status' })

  const gradeLabel = useGradeLabel(displayValues.grade)

  return (
    <div className={styles.card}>
      <Text size='20' weight='600'>
        Общая информация
      </Text>
      <div className={styles.cardFields}>
        <Input
          title='Название задания'
          fluid
          value={displayValues.title}
          disabled
          onChange={noop}
        />
        <Input
          title='Проверяющий'
          fluid
          value={displayValues.teacher}
          disabled
          onChange={noop}
        />
        <Input
          title='Исполняющий'
          fluid
          value={displayValues.student}
          disabled
          onChange={noop}
        />
        <Controller
          control={control}
          name='status'
          render={({ field, fieldState }) => (
            <Select
              label='Статус'
              fluid
              required={mode !== 'student-view'}
              value={field.value}
              options={statusOptions}
              disabled={disabled}
              onChange={field.onChange}
              isValid={!fieldState.error}
              description={fieldState.error?.message}
            />
          )}
        />
        {mode === 'teacher' && status === CHECKED_STATUS_ID ? (
          <GradeField disabled={disabled} />
        ) : (
          mode === 'student-view' && (
            <Input
              title='Оценка'
              fluid
              disabled
              placeholder='Оценки нет'
              value={gradeLabel}
              onChange={noop}
            />
          )
        )}
      </div>
    </div>
  )
}

export default TaskGeneral
