import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: process.env.NODE_ENV === 'production' ? 'http://monitor.dn11.baimeow.cn/api' : 'http://localhost:8787',
      }
    },
  },
  build: {
    outDir: '../backend/static',
    emptyOutDir: true,
  }
})
