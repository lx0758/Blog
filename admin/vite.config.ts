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
  base: './',
  server: {
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
