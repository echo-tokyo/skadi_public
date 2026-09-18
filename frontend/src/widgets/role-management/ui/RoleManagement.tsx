import { ManagementLayout, Select } from '@/shared/ui'
import { useInfiniteMembers } from '../model/use-infinite-members'
import { ROLE_OPTIONS, ROLE_VALUES } from '@/shared/config'
import { CreateRoleButton } from '@/features/create-member'
import { TRole } from '@/shared/model'
import { useState } from 'react'
import { MemberCardItem } from './MemberCardItem'

const RoleManagement = () => {
  const [role, setRole] = useState<TRole | ''>('')
  const [debouncedSearch, setDebouncedSearch] = useState<string>('')
  const {
    members,
    isFetchingNextPage,
    loadMore,
    hasMore,
    isLoading,
    isFetching,
  } = useInfiniteMembers({
    free: false,
    role: role ? [role] : ROLE_VALUES,
    search: debouncedSearch || undefined,
    'per-page': 10,
  })

  return (
    <ManagementLayout
      isFetching={isFetching}
      searchTitle='Найти пользователя'
      title='Управление ролями'
      renderItem={(item) => <MemberCardItem member={item} />}
      onDebouncedSearch={setDebouncedSearch}
      items={members}
      isLoading={isLoading}
      getKey={(item) => item.id}
      infiniteScrollOptions={{ hasMore, isFetchingNextPage, loadMore }}
      actions={
        <>
          <Select
            label='Фильтровать по:'
            placeholder='Выберите'
            options={ROLE_OPTIONS}
            value={role}
            onChange={(val) => setRole(val as TRole)}
          />
          <CreateRoleButton />
        </>
      }
    />
  )
}

export default RoleManagement
