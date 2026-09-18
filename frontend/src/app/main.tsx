import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'
import { store } from './store/store'
import '@/app/styles/app.scss'
import { RouterProvider } from 'react-router'
import { router } from './routes/Routes'
import { DialogProvider } from '@/shared/lib'
import { Toaster } from 'sonner'

createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <Toaster position='bottom-center' invert />
    <DialogProvider>
      <RouterProvider router={router} />
    </DialogProvider>
  </Provider>,
)
