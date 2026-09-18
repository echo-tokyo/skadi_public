export type TRole = 'admin' | 'teacher' | 'student'

type TContact = { email?: string; phone?: string }

type TTeacherProfile = Omit<TProfile, 'class'>

export type TClass = {
  id: number
  name: string
  schedule?: string
  students?: TProfile[]
  teacher?: TTeacherProfile
}

export type TProfile = {
  id: number
  address?: string
  contact?: TContact
  extra?: string
  fullname: string
  parent_contact?: TContact
}

export type TTask = {
  description?: string
  id: number
  teacher?: TProfile
  title: string
  files: TFile[]
}

export type TFile = {
  id: number
  mime_type: string
  name: string
  size: number
}

export type TComment = {
  created_at: string
  id: number
  message: string
  role: TRole
}

export type TTaskWithStudents = {
  students?: TProfile[]
  task: TTask
}

export type TStatus = {
  id: 1 | 2 | 3 | 4
  name: 'Бэклог' | 'В работе' | 'На проверке' | 'Проверено'
}

export type TStatusId = TStatus['id']
export type TStatusName = TStatus['name']

export type TGrade = '2' | '3' | '4' | '5'

export type TSolution = {
  answer?: string
  grade?: string
  id: number
  status: TStatus
  files: TFile[]
  student?: TProfile
  task: TTask
  updated_at?: string
}

export type TPagination = {
  page: number
  per_page: number
  pages: number
  total: number
}
