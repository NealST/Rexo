import { createContext, useContext } from "react";

// SSR 数据接口
export interface SSRData {
  [key: string]: any;
}

// SSR 上下文
interface SSRContextType {
  data: SSRData;
  isServer: boolean;
  isHydrated: boolean;
}

export const SSRContext = createContext<SSRContextType | null>(null);

// 使用 SSR 数据的 Hook
export function useSSRData<T = any>(key?: string): T | SSRData {
  const context = useContext(SSRContext);

  if (!context) {
    throw new Error("useSSRData must be used within SSRProvider");
  }

  if (key) {
    return context.data[key] as T;
  }

  return context.data as T;
}

// 检查是否在服务端
export function useIsServer(): boolean {
  const context = useContext(SSRContext);

  if (!context) {
    throw new Error("useIsServer must be used within SSRProvider");
  }

  return context.isServer;
}

// 检查是否已水合
export function useIsHydrated(): boolean {
  const context = useContext(SSRContext);

  if (!context) {
    throw new Error("useIsHydrated must be used within SSRProvider");
  }

  return context.isHydrated;
}
