import { getErrorMessage } from '@/shared/api'
import { toast } from 'sonner'
import { useLazyDownloadFileQuery } from '../api/file-api'

export const useDownloadFile = (fileName: string) => {
  const [trigger, { isLoading }] = useLazyDownloadFileQuery()

  const download = async (id: number) => {
    try {
      const blob = await trigger(id).unwrap()
      const url = URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.href = url
      a.download = fileName
      a.click()
      URL.revokeObjectURL(url)
    } catch (err) {
      toast.error(getErrorMessage(err))
    }
  }

  return { download, isLoading }
}
