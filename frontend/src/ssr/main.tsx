// SSR 入口文件
import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import App from '../App'
import '../styles/globals.css'

// 客户端水合
function hydrate() {
  const container = document.getElementById('root')
  if (!container) {
    console.error('Root container not found')
    return
  }

  // 检查是否已经有服务端渲染的内容
  const hasSSRContent = container.children.length > 0

  if (hasSSRContent) {
    // 水合现有的服务端渲染内容
    ReactDOM.hydrateRoot(container, (
      <React.StrictMode>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </React.StrictMode>
    ))
  } else {
    // 如果没有服务端渲染内容，进行客户端渲染
    ReactDOM.createRoot(container).render(
      <React.StrictMode>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </React.StrictMode>
    )
  }
}

// 等待 DOM 加载完成
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', hydrate)
} else {
  hydrate()
}

// 导出组件供服务端使用
export { App }
