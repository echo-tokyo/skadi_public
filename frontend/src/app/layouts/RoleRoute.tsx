import { selectAuthenticatedUser } from '@/entities/user'
import { useAppSelector } from '@/shared/lib'
import { TRole } from '@/shared/model'
import { Navigate, Outlet } from 'react-router'

interface IRoleRouteProps {
  role: TRole
  redirectTo?: string
}

const RoleRoute = ({
  role,
  redirectTo = '/personal-area',
}: IRoleRouteProps) => {
  const user = useAppSelector(selectAuthenticatedUser)

  if (user.role !== role) {
    return <Navigate to={redirectTo} replace />
  }

  return <Outlet />
}

export default RoleRoute
