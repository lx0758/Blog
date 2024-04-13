import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  // 使用 WebHistory 必须和实际部署地址一致
  // 使用 WebHashHistory/MemoryHistory 则可以配置为 ./ 保持动态
  base: '/admin/',
  server: {
    host: '0.0.0.0',
    proxy: {
      '^/(api|files)/.*': {
          target: 'http://127.0.0.1:8080/',
          ws: true,
          secure: false,
          changeOrigin: true,
      },
    }
  }
})
