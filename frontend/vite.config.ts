import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: '../backend/static',
    emptyOutDir: true,
  },
  server:{
    proxy: {
        '/api': {
          target: process.env.NODE_ENV === 'production' ? 'http://monitor.dn11.baimeow.cn' : 'http://localhost:8787',
        }
    }
  }
})
