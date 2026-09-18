import { parseGrade, useGradeLabel } from '@/entities/solution'
import { DeleteSolutionButton } from '@/features/delete-solution'
import { TSolution } from '@/shared/model'
import { AccordionCard, Button } from '@/shared/ui'
import { memo } from 'react'
import { useNavigate } from 'react-router'

interface ISolutionCardItemProps {
  solutionData: TSolution
}

export const SolutionCardItem = memo(
  ({ solutionData }: ISolutionCardItemProps) => {
    const nav = useNavigate()

    const parsedGrade = parseGrade(solutionData.grade)
    const gradeLabel = useGradeLabel(parsedGrade)

    return (
      <AccordionCard
        title={solutionData.student?.fullname ?? 'Нет ученика'}
        fields={[
          { label: 'Задание', value: solutionData.task.title },
          { label: 'Статус', value: solutionData.status.name },
          { label: 'Оценка', value: gradeLabel || '-' },
        ]}
        actions={
          <>
            <DeleteSolutionButton
              id={solutionData.id}
              studentName={solutionData.student?.fullname ?? 'Нет ученика'}
              taskTitle={solutionData.task.title}
            />
            <Button
              color='secondary'
              onClick={() => nav(`/personal-area/solutions/${solutionData.id}`)}
            >
              Просмотреть
            </Button>
          </>
        }
      />
    )
  },
)

SolutionCardItem.displayName = 'SolutionCardItem'
