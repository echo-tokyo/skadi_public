import { TFile } from '@/shared/model'
import { FileItem, Text } from '@/shared/ui'
import { useDownloadFile } from '../model/use-download-file'

interface IFileProps {
  el: TFile
}

export const FileDownload = ({ el }: IFileProps) => {
  const { download } = useDownloadFile(el.name)

  return (
    <div>
      <Text size='12'>Файл</Text>
      <FileItem
        file={{ id: el.id.toString(), name: el.name }}
        size={el.size}
        onDownload={(id) => download(Number(id))}
      />
    </div>
  )
}
