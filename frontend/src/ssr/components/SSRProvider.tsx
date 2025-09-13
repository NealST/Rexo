import React, { createContext, useContext, useEffect, useState } from 'react'

// SSR 数据接口
interface SSRData {
  [key: string]: any
}

// SSR 上下文
interface SSRContextType {
  data: SSRData
  isServer: boolean
  isHydrated: boolean
}

const SSRContext = createContext<SSRContextType | null>(null)

// SSR Provider 组件
interface SSRProviderProps {
  children: React.ReactNode
  initialData?: SSRData
}

export function SSRProvider({ children, initialData = {} }: SSRProviderProps) {
  const [isHydrated, setIsHydrated] = useState(false)
  const [data, setData] = useState<SSRData>(initialData)

  useEffect(() => {
    // 客户端水合
    setIsHydrated(true)
    
    // 从 window.__SSR_DATA__ 获取服务端数据
    if (typeof window !== 'undefined' && (window as any).__SSR_DATA__) {
      setData((window as any).__SSR_DATA__)
    }
  }, [])

  const isServer = typeof window === 'undefined'

  return (
    <SSRContext.Provider value={{ data, isServer, isHydrated }}>
      {children}
    </SSRContext.Provider>
  )
}

// 使用 SSR 数据的 Hook
export function useSSRData<T = any>(key?: string): T | SSRData {
  const context = useContext(SSRContext)
  
  if (!context) {
    throw new Error('useSSRData must be used within SSRProvider')
  }

  if (key) {
    return context.data[key] as T
  }

  return context.data as T
}

// 检查是否在服务端
export function useIsServer(): boolean {
  const context = useContext(SSRContext)
  
  if (!context) {
    throw new Error('useIsServer must be used within SSRProvider')
  }

  return context.isServer
}

// 检查是否已水合
export function useIsHydrated(): boolean {
  const context = useContext(SSRContext)
  
  if (!context) {
    throw new Error('useIsHydrated must be used within SSRProvider')
  }

  return context.isHydrated
}
