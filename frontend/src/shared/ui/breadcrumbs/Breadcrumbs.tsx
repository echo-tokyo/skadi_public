import { Link } from 'react-router'
import { useBreadcrumbs } from '@/shared/lib'
import styles from './styles.module.scss'
import Text from '../text/Text'
import { Crumb } from '@/shared/lib/hooks/use-breadcrumbs'

interface ICrumbItemProps {
  crumb: Crumb
  isLast: boolean
}

const CrumbItem = ({ crumb, isLast }: ICrumbItemProps) => (
  <div className={styles.crumb}>
    {!isLast && crumb.to ? (
      <Link to={crumb.to} className={styles.linkWrapper}>
        <Text size='16' color='--color-gray' className={styles.link}>
          {crumb.label}
        </Text>
      </Link>
    ) : (
      <Text size='16'>{crumb.label}</Text>
    )}
    {!isLast && (
      <Text size='16' color='--color-gray'>
        ›
      </Text>
    )}
  </div>
)

const Breadcrumbs = () => {
  const crumbs = useBreadcrumbs()

  if (!crumbs.length) return null

  return (
    <nav className={styles.breadcrumbs}>
      {crumbs.map((crumb, index) => (
        <CrumbItem
          key={index}
          crumb={crumb}
          isLast={index === crumbs.length - 1}
        />
      ))}
    </nav>
  )
}

export default Breadcrumbs
