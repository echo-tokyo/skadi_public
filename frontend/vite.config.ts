import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': '/src',
      '@/shared': '/src/shared/',
      '@/entities': '/src/entities/',
      '@/features': '/src/features/',
      '@/widgets': '/src/widgets/',
      '@/pages': '/src/pages/',
      '@/app': '/src/app/',
    },
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8000',
        changeOrigin: true,
        secure: false,
        timeout: 60000,
        proxyTimeout: 60000,
      },
    },
  },
})
