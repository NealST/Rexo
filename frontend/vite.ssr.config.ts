import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// SSR 构建配置
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@/components': path.resolve(__dirname, './src/components'),
      '@/pages': path.resolve(__dirname, './src/pages'),
      '@/hooks': path.resolve(__dirname, './src/hooks'),
      '@/services': path.resolve(__dirname, './src/services'),
      '@/store': path.resolve(__dirname, './src/store'),
      '@/utils': path.resolve(__dirname, './src/utils'),
      '@/types': path.resolve(__dirname, './src/types'),
      '@/styles': path.resolve(__dirname, './src/styles'),
      '@/ssr': path.resolve(__dirname, './src/ssr'),
    },
  },
  build: {
    outDir: 'dist',
    rollupOptions: {
      input: {
        main: './index.html',
        ssr: './src/ssr/main.tsx',
      },
      output: {
        entryFileNames: (chunkInfo) => {
          return chunkInfo.name === 'ssr' ? 'ssr.js' : '[name]-[hash].js'
        },
      },
    },
    ssr: true,
  },
  ssr: {
    noExternal: ['react', 'react-dom'],
  },
})
