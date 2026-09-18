import { Outlet } from 'react-router'
import { Header } from '@/widgets/header'
import { Breadcrumbs } from '@/shared/ui'

const AppLayout = () => (
  <>
    <Header />
    <div className='wrapper'>
      <Breadcrumbs />
      <Outlet />
    </div>
  </>
)

export default AppLayout
