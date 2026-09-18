import { TFile, TGrade } from '@/shared/model'

export type TaskCardMode = 'teacher' | 'student-edit' | 'student-view'

export type TDisplayValues = {
  title: string
  description: string
  teacher: string
  student: string
  answer: string
  files: TFile[]
  file_answer: TFile[]
  grade: TGrade | ''
}
