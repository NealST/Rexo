import React from 'react'
import { useSSRData, useIsServer } from '../components/SSRProvider'

interface AboutPageProps {
  user?: {
    id: number
    name: string
    email: string
  }
  path?: string
}

export function AboutPage(props: AboutPageProps = {}) {
  const ssrData = useSSRData<AboutPageProps>()
  const isServer = useIsServer()
  
  const data = { ...props, ...ssrData }

  return (
    <div className="min-h-screen bg-white">
      <div className="container mx-auto px-4 py-16">
        {/* Header */}
        <div className="text-center mb-16">
          <h1 className="text-4xl font-bold text-gray-900 mb-6">
            关于 Rexo
          </h1>
          <p className="text-xl text-gray-600 max-w-3xl mx-auto">
            Rexo 是一个现代化的全栈 React 研发框架，集成了 Go 后端和 React 前端，
            提供完整的开发工具链和最佳实践。
          </p>
          
          {isServer && (
            <div className="inline-flex items-center px-4 py-2 bg-green-100 text-green-800 rounded-full text-sm font-medium mt-6">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
              Server-Side Rendered
            </div>
          )}
        </div>

        {/* Features */}
        <div className="grid lg:grid-cols-2 gap-16 mb-16">
          <div>
            <h2 className="text-3xl font-bold text-gray-900 mb-6">核心特性</h2>
            <div className="space-y-6">
              <div className="flex items-start">
                <div className="w-8 h-8 bg-blue-100 rounded-lg flex items-center justify-center mr-4 flex-shrink-0">
                  <svg className="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                  </svg>
                </div>
                <div>
                  <h3 className="text-lg font-semibold mb-2">高性能后端</h3>
                  <p className="text-gray-600">
                    基于 Go Fiber 框架，提供卓越的性能和并发处理能力。
                  </p>
                </div>
              </div>

              <div className="flex items-start">
                <div className="w-8 h-8 bg-green-100 rounded-lg flex items-center justify-center mr-4 flex-shrink-0">
                  <svg className="w-4 h-4 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>
                <div>
                  <h3 className="text-lg font-semibold mb-2">类型安全</h3>
                  <p className="text-gray-600">
                    端到端 TypeScript 支持，确保代码质量和开发效率。
                  </p>
                </div>
              </div>

              <div className="flex items-start">
                <div className="w-8 h-8 bg-purple-100 rounded-lg flex items-center justify-center mr-4 flex-shrink-0">
                  <svg className="w-4 h-4 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                  </svg>
                </div>
                <div>
                  <h3 className="text-lg font-semibold mb-2">SSR 支持</h3>
                  <p className="text-gray-600">
                    内置服务端渲染支持，提升 SEO 和首屏加载性能。
                  </p>
                </div>
              </div>

              <div className="flex items-start">
                <div className="w-8 h-8 bg-yellow-100 rounded-lg flex items-center justify-center mr-4 flex-shrink-0">
                  <svg className="w-4 h-4 text-yellow-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
                <div>
                  <h3 className="text-lg font-semibold mb-2">开发工具</h3>
                  <p className="text-gray-600">
                    完整的 CLI 工具、代码生成器和开发环境配置。
                  </p>
                </div>
              </div>
            </div>
          </div>

          <div>
            <h2 className="text-3xl font-bold text-gray-900 mb-6">技术栈</h2>
            <div className="space-y-4">
              <div className="card">
                <div className="card-content">
                  <h3 className="font-semibold mb-2">后端技术</h3>
                  <ul className="text-sm text-gray-600 space-y-1">
                    <li>• Go 1.21+ (Fiber 框架)</li>
                    <li>• PostgreSQL/MySQL + GORM</li>
                    <li>• Redis 缓存</li>
                    <li>• JWT 认证</li>
                    <li>• Swagger API 文档</li>
                  </ul>
                </div>
              </div>

              <div className="card">
                <div className="card-content">
                  <h3 className="font-semibold mb-2">前端技术</h3>
                  <ul className="text-sm text-gray-600 space-y-1">
                    <li>• React 18 + TypeScript</li>
                    <li>• Vite 构建工具</li>
                    <li>• Tailwind CSS</li>
                    <li>• Zustand 状态管理</li>
                    <li>• React Router v6</li>
                  </ul>
                </div>
              </div>

              <div className="card">
                <div className="card-content">
                  <h3 className="font-semibold mb-2">开发工具</h3>
                  <ul className="text-sm text-gray-600 space-y-1">
                    <li>• ESLint + Prettier</li>
                    <li>• Jest + Testing Library</li>
                    <li>• Docker 容器化</li>
                    <li>• GitHub Actions CI/CD</li>
                    <li>• Hot Reload 热重载</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* CTA Section */}
        <div className="text-center bg-gray-50 rounded-2xl p-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-4">
            开始您的全栈开发之旅
          </h2>
          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
            使用 Rexo 框架，快速构建现代化的全栈应用程序。
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <a href="/docs" className="btn btn-primary btn-lg">
              查看文档
            </a>
            <a href="/examples" className="btn btn-outline btn-lg">
              示例项目
            </a>
            <a href="/github" className="btn btn-secondary btn-lg">
              GitHub
            </a>
          </div>
        </div>

        {/* Debug Info */}
        {process.env.NODE_ENV === 'development' && (
          <div className="mt-16 p-4 bg-gray-100 rounded-lg">
            <h4 className="font-semibold mb-2">Debug Info:</h4>
            <pre className="text-sm text-gray-600">
              {JSON.stringify({ 
                isServer, 
                path: data.path,
                user: data.user,
                ssrData: ssrData 
              }, null, 2)}
            </pre>
          </div>
        )}
      </div>
    </div>
  )
}

// 导出为全局函数，供服务端渲染使用
if (typeof window !== 'undefined') {
  (window as any).AboutPage = AboutPage
}
