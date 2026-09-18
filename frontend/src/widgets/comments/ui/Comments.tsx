import {
  Comment,
  MessageInput,
  PlugDefault,
  Sentinel,
  Skeleton,
  Text,
} from '@/shared/ui'
import { unixToDate, useInfiniteScroll, useShowSkeleton } from '@/shared/lib'
import styles from './styles.module.scss'
import { useGetComments } from '../model/use-get-comments'
import { useSendComment } from '../model/use-send-comment'
import { TRole } from '@/shared/model'
import { useCallback, useEffect, useMemo, useRef } from 'react'

interface ICommentsProps {
  role: TRole
  solutionId: number
  mode: 'edit' | 'view'
}

const Comments = (props: ICommentsProps) => {
  const { role, solutionId, mode } = props

  const { messages, hasMore, isFetchingNextPage, loadMore, isLoading } =
    useGetComments({
      id: solutionId,
    })

  const { submit, isLoading: isSending } = useSendComment({
    id: solutionId,
    role,
  })

  const chatRef = useRef<HTMLDivElement>(null)
  const isLoadingMoreRef = useRef(false)
  const scrollHeightBeforeRef = useRef(0)

  const handleLoadMore = useCallback(() => {
    if (chatRef.current) {
      scrollHeightBeforeRef.current = chatRef.current.scrollHeight
      isLoadingMoreRef.current = true
    }
    loadMore()
  }, [loadMore])

  const { sentinelRef } = useInfiniteScroll({
    hasMore,
    isFetchingNextPage,
    loadMore: handleLoadMore,
  })

  const showSkeleton = useShowSkeleton(isLoading)

  useEffect(() => {
    if (!chatRef.current) return
    if (isLoadingMoreRef.current) {
      isLoadingMoreRef.current = false
      chatRef.current.scrollTop +=
        chatRef.current.scrollHeight - scrollHeightBeforeRef.current
      return
    }
    chatRef.current.scrollTop = chatRef.current.scrollHeight
  }, [messages])

  const content = useMemo(() => {
    if (showSkeleton) {
      return <Skeleton height={'100%'} />
    }

    if (!messages.length && !isLoading) {
      return <PlugDefault title='Комментариев нет 🥲' />
    }

    return (
      <>
        <Sentinel ref={sentinelRef} />
        {messages.map((msg) => {
          const align = msg.role === role ? 'right' : 'left'
          return (
            <Comment
              key={msg.id}
              message={msg.message}
              time={unixToDate(msg.created_at)}
              sender={msg.role}
              align={align}
            />
          )
        })}
      </>
    )
  }, [showSkeleton, messages, role, sentinelRef])

  return (
    <div className={styles.comments}>
      <Text size='20' weight='600'>
        Комментарии
      </Text>
      <div className={styles.commentsContent} ref={chatRef}>
        {content}
      </div>
      <MessageInput onSubmit={submit} disabled={isSending || mode === 'view'} />
    </div>
  )
}

export default Comments
