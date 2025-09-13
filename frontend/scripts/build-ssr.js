const { build } = require('vite')
const path = require('path')
const fs = require('fs')

async function buildSSR() {
  console.log('🚀 Building SSR bundle...')
  
  try {
    // 构建 SSR 版本
    await build({
      configFile: path.resolve(__dirname, '../vite.ssr.config.ts'),
      build: {
        ssr: true,
        outDir: 'dist',
        rollupOptions: {
          input: './src/ssr/main.tsx',
          output: {
            format: 'cjs',
            entryFileNames: 'ssr.js',
          },
        },
      },
    })

    console.log('✅ SSR bundle built successfully')
    
    // 复制必要的文件到后端
    const backendPath = path.resolve(__dirname, '../../backend/dist')
    if (!fs.existsSync(backendPath)) {
      fs.mkdirSync(backendPath, { recursive: true })
    }

    // 复制 SSR 文件
    fs.copyFileSync(
      path.resolve(__dirname, '../dist/ssr.js'),
      path.resolve(backendPath, 'ssr.js')
    )

    console.log('✅ SSR files copied to backend')
    
  } catch (error) {
    console.error('❌ SSR build failed:', error)
    process.exit(1)
  }
}

buildSSR()
