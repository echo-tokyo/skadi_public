import { IMember } from '@/entities/member'
import { AccordionCard, IAccordionCardField } from '@/shared/ui'
import { DeleteMemberButton } from '@/features/delete-member'
import { EditMemberButton } from '@/features/edit-member'
import { memo, useMemo } from 'react'

export const MemberCardItem = memo(({ member }: { member: IMember }) => {
  const accordionFields = useMemo(() => {
    const fields: IAccordionCardField[] = [
      { label: 'Логин', value: member.username },
      { label: 'Статус', value: member.role },
    ]

    if (member.role === 'student') {
      fields.push({ label: 'Группа', value: member.class?.name })
    }

    return fields
  }, [member.username, member.role, member.class?.name])

  const actions = useMemo(
    () => (
      <>
        <DeleteMemberButton fullname={member.profile.fullname} id={member.id} />
        <EditMemberButton member={member} />
      </>
    ),
    [member],
  )

  return (
    <AccordionCard
      title={member.profile.fullname}
      fields={accordionFields}
      actions={actions}
    />
  )
})

MemberCardItem.displayName = 'MemberCardItem'
