import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { resolve } from 'path'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    fs: {
      // Allow serving files from the ThreeUI package
      strict: false
    }
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    }
  }
})