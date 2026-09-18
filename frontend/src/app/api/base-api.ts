import { initializeAuthActions } from '@/shared/api'
import { router } from '../routes/Routes'
import { store } from '../store/store'
import { clearUserData } from '@/entities/user'
import { toast } from 'sonner'
import { closeAllDialogs } from '@/shared/lib/dialog/dialogActions'

initializeAuthActions({
  onAuthFailure: () => {
    store.dispatch(clearUserData())
    router.navigate('/authorization', { replace: true })
    toast.error('Требуется авторизация')
    closeAllDialogs()
  },
})

export { baseApi } from '@/shared/api'
