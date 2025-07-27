import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
const backendHost = process.env.VITE_API_BASE_URL || 'http://localhost:8080'
console.log('Backend URL:', process.env.VITE_API_BASE_URL)
// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
    server: {
    host: true, 
    proxy: {
      '/api': {
        // target: 'http://localhost:8080', for local
        target: backendHost,
        changeOrigin: true,
        secure: false,
      }
    }
  }
})

