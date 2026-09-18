import UpdateSolutionByStudentButton from './components/UpdateSolutionByStudentButton'
import UpdateSolutionByTeacherButton from './components/UpdateSolutionByTeacherButton'

interface IUpdateSolutionButtonProps {
  id: number
  isTeacher: boolean
}

const UpdateSolutionButton = ({ id, isTeacher }: IUpdateSolutionButtonProps) => {
  return isTeacher ? (
    <UpdateSolutionByTeacherButton id={id} />
  ) : (
    <UpdateSolutionByStudentButton id={id} />
  )
}

export default UpdateSolutionButton
