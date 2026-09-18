import { createBrowserRouter } from 'react-router'
import Authorization from '@/pages/authorization'
import MainPage from '@/pages/main'
import { PersonalArea } from '@/pages/personal-area'
import { Solution } from '@/pages/solution'
import AppLayout from '../layouts/AppLayout'
import ProtectedRoute from '../layouts/ProtectedRoute'
import RoleRoute from '../layouts/RoleRoute'
import { Kanban } from '@/pages/kanban'

const crumb = {
  home: { label: 'Главная', to: '/' },
  authorization: { label: 'Авторизация', to: '/authorization' },
  personalArea: { label: 'Личный кабинет', to: '/personal-area' },
  kanban: { label: 'Канбан-доска', to: '/personal-area/kanban' },
  task: { label: 'Задача' },
} as const

const studentRoutes = [
  {
    path: '/personal-area/kanban',
    Component: Kanban,
    handle: {
      breadcrumbs: [crumb.home, crumb.personalArea, { label: 'Канбан-доска' }],
    },
  },
  {
    path: '/personal-area/kanban/solutions/:id',
    Component: Solution,
    handle: {
      breadcrumbs: [crumb.home, crumb.personalArea, crumb.kanban, crumb.task],
    },
  },
]

export const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
      {
        path: '/',
        Component: MainPage,
      },
      {
        path: '/authorization',
        Component: Authorization,
        handle: { breadcrumbs: [crumb.home, crumb.authorization] },
      },
      {
        element: <ProtectedRoute />,
        children: [
          {
            path: '/personal-area',
            Component: PersonalArea,
            handle: {
              breadcrumbs: [crumb.home, { label: 'Личный кабинет' }],
            },
          },
          {
            path: '/personal-area/solutions/:id',
            Component: Solution,
            handle: {
              breadcrumbs: [crumb.home, crumb.personalArea, crumb.task],
            },
          },
          {
            element: <RoleRoute role='student' />,
            children: studentRoutes,
          },
        ],
      },
    ],
  },
])
