import React from 'react'
import { useSSRData, useIsServer } from '../components/SSRProvider'

interface HomePageProps {
  user?: {
    id: number
    name: string
    email: string
  }
  path?: string
}

export function HomePage(props: HomePageProps = {}) {
  const ssrData = useSSRData<HomePageProps>()
  const isServer = useIsServer()
  
  // 合并 props 和 SSR 数据
  const data = { ...props, ...ssrData }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="container mx-auto px-4 py-16">
        {/* Hero Section */}
        <div className="text-center mb-16">
          <h1 className="text-5xl font-bold text-gray-900 mb-6">
            Welcome to{' '}
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-purple-600">
              Rexo
            </span>
          </h1>
          <p className="text-xl text-gray-600 mb-8 max-w-3xl mx-auto">
            基于 Go (Fiber) + React 的全栈研发框架，支持服务端渲染 (SSR)，
            提供现代化的开发体验和高效的开发工具。
          </p>
          
          {/* SSR 状态指示器 */}
          {isServer && (
            <div className="inline-flex items-center px-4 py-2 bg-green-100 text-green-800 rounded-full text-sm font-medium mb-8">
              <div className="w-2 h-2 bg-green-500 rounded-full mr-2"></div>
              Server-Side Rendered
            </div>
          )}
          
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <a
              href="/docs"
              className="btn btn-primary btn-lg"
            >
              查看文档
            </a>
            <a
              href="/examples"
              className="btn btn-outline btn-lg"
            >
              示例项目
            </a>
          </div>
        </div>

        {/* Features Grid */}
        <div className="grid md:grid-cols-3 gap-8 mb-16">
          <div className="card">
            <div className="card-content">
              <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
                <svg className="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2">高性能</h3>
              <p className="text-gray-600">
                基于 Fiber 框架，性能比传统框架更优，支持高并发处理。
              </p>
            </div>
          </div>

          <div className="card">
            <div className="card-content">
              <div className="w-12 h-12 bg-green-100 rounded-lg flex items-center justify-center mb-4">
                <svg className="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2">类型安全</h3>
              <p className="text-gray-600">
                端到端 TypeScript 支持，确保代码质量和开发效率。
              </p>
            </div>
          </div>

          <div className="card">
            <div className="card-content">
              <div className="w-12 h-12 bg-purple-100 rounded-lg flex items-center justify-center mb-4">
                <svg className="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2">SSR 支持</h3>
              <p className="text-gray-600">
                内置服务端渲染支持，提升 SEO 和首屏加载性能。
              </p>
            </div>
          </div>
        </div>

        {/* User Info */}
        {data.user ? (
          <div className="text-center">
            <div className="card max-w-md mx-auto">
              <div className="card-content">
                <h3 className="text-lg font-semibold mb-2">欢迎回来！</h3>
                <p className="text-gray-600">
                  你好，{data.user.name}！你已经登录到 Rexo 平台。
                </p>
                <div className="mt-4">
                  <a href="/dashboard" className="btn btn-primary btn-sm mr-2">
                    进入仪表板
                  </a>
                  <a href="/profile" className="btn btn-outline btn-sm">
                    个人资料
                  </a>
                </div>
              </div>
            </div>
          </div>
        ) : (
          <div className="text-center">
            <div className="card max-w-md mx-auto">
              <div className="card-content">
                <h3 className="text-lg font-semibold mb-2">开始使用</h3>
                <p className="text-gray-600 mb-4">
                  注册账户开始您的全栈开发之旅。
                </p>
                <div className="flex gap-2 justify-center">
                  <a href="/register" className="btn btn-primary btn-sm">
                    立即注册
                  </a>
                  <a href="/login" className="btn btn-outline btn-sm">
                    登录
                  </a>
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Debug Info (仅开发环境) */}
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
  (window as any).HomePage = HomePage
}
