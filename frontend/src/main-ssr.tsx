// ssr render entry
import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import { SSRContext, type SSRData } from "./context/ssr";
import "./styles/globals.css";

// SSR Provider 组件
interface SSRProviderProps {
  children: React.ReactNode;
  initialData?: SSRData;
}

export function SSRProvider({ children, initialData = {} }: SSRProviderProps) {
  const [isHydrated, setIsHydrated] = useState(false);
  const [data, setData] = useState<SSRData>(initialData);

  useEffect(() => {
    // 客户端水合
    setIsHydrated(true);

    // 从 window.__SSR_DATA__ 获取服务端数据
    if (typeof window !== "undefined" && (window as any).__SSR_DATA__) {
      setData((window as any).__SSR_DATA__);
    }
  }, []);

  const isServer = typeof window === "undefined";

  return (
    <SSRContext.Provider value={{ data, isServer, isHydrated }}>
      {children}
    </SSRContext.Provider>
  );
}

// 客户端水合
function hydrate() {
  const container = document.getElementById("root");
  if (!container) {
    console.error("Root container not found");
    return;
  }

  // 检查是否已经有服务端渲染的内容
  const hasSSRContent = container.children.length > 0;

  if (hasSSRContent) {
    // 水合现有的服务端渲染内容
    ReactDOM.hydrateRoot(
      container,
      <React.StrictMode>
        <SSRProvider>
          <BrowserRouter>
            <App />
          </BrowserRouter>
        </SSRProvider>
      </React.StrictMode>
    );
  } else {
    // 如果没有服务端渲染内容，进行客户端渲染
    ReactDOM.createRoot(container).render(
      <React.StrictMode>
        <SSRProvider>
          <BrowserRouter>
            <App />
          </BrowserRouter>
        </SSRProvider>
      </React.StrictMode>
    );
  }
}

// 等待 DOM 加载完成
if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", hydrate);
} else {
  hydrate();
}
