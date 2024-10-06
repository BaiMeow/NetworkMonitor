import { defineConfig } from 'vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import vue from '@vitejs/plugin-vue'
import IconsResolver from 'unplugin-icons/resolver'
import Icons from 'unplugin-icons/vite'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  plugins: [
    vue(),
    AutoImport({
      resolvers: [
        ElementPlusResolver(),
        IconsResolver({
          prefix: 'Icon',
        }),
      ],
    }),
    Components({
      resolvers: [
        IconsResolver({
          enabledCollections: ['ep'],
        }),
        ElementPlusResolver(),
      ],
    }),
    Icons({
      autoInstall: true,
    }),
  ],
  build: {
    outDir: '../backend/static',
    emptyOutDir: true,
  },
  server: {
    proxy: {
      '/api': {
        target: 'https://monitor.dn11.baimeow.cn',
        changeOrigin: true,
      },
      '/monitor-metadata.json': {
        target: 'https://monitor.dn11.baimeow.cn',
        changeOrigin: true,
      },
    },
  },
})
