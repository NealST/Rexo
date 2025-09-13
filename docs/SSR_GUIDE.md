# Rexo SSR 服务端渲染指南

## 概述

Rexo 框架集成了完整的服务端渲染 (SSR) 功能，基于 Go + React 实现，提供：

- ⚡ **高性能渲染**: 基于 Fiber 框架和 V8 引擎
- 🔄 **数据预取**: 服务端数据预取和客户端水合
- 💾 **智能缓存**: Redis 和内存缓存支持
- 🎯 **SEO 优化**: 完整的 SEO 元数据支持
- 📊 **性能监控**: 内置性能监控和优化

## 架构设计

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Client        │    │   Go Server     │    │   Database      │
│   (React)       │    │   (Fiber)       │    │   (PostgreSQL)  │
├─────────────────┤    ├─────────────────┤    ├─────────────────┤
│ • Hydration     │◄──►│ • SSR Engine    │◄──►│ • User Data     │
│ • Client Routes │    │ • Data Fetcher  │    │ • Page Data     │
│ • State Mgmt    │    │ • Cache Layer   │    │ • Content       │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 核心组件

### 1. SSR 引擎 (Engine)

负责在服务端执行 React 组件渲染：

```go
// 创建 SSR 引擎
engine, err := engine.NewEngine(basePath)

// 渲染组件
result, err := engine.Render(engine.RenderOptions{
    Component: "HomePage",
    Props:     map[string]interface{}{"user": userData},
    Path:      "/",
    Query:     queryParams,
})
```

### 2. 数据预取服务 (DataFetcher)

在服务端预取页面所需数据：

```go
// 创建数据预取服务
fetcher := services.NewDataFetcher(db)

// 获取页面数据
data, err := fetcher.FetchPageData(ctx, "/dashboard", &userID)
```

### 3. 缓存系统 (Cache)

提供多层缓存支持：

```go
// Redis 缓存
redisCache := cache.NewRedisCache(redisClient)
ssrCache := cache.NewSSRCache(redisCache)

// 内存缓存
memoryCache := cache.NewMemoryCache()
ssrCache := cache.NewSSRCache(memoryCache)
```

## 使用方法

### 1. 创建 SSR 组件

```tsx
// src/ssr/pages/HomePage.tsx
import React from 'react'
import { useSSRData, useIsServer } from '../components/SSRProvider'

export function HomePage(props: HomePageProps = {}) {
  const ssrData = useSSRData<HomePageProps>()
  const isServer = useIsServer()
  
  const data = { ...props, ...ssrData }

  return (
    <div className="min-h-screen">
      <h1>Welcome to Rexo</h1>
      {isServer && (
        <div className="ssr-indicator">
          Server-Side Rendered
        </div>
      )}
    </div>
  )
}

// 注册为全局组件
if (typeof window !== 'undefined') {
  (window as any).HomePage = HomePage
}
```

### 2. 配置 SSR 路由

```go
// backend/main.go
func registerSSRRoutes(app *fiber.App, renderer *renderer.Renderer) {
    ssrMiddleware := middleware.NewSSRMiddleware(renderer)

    // 首页
    app.Get("/", ssrMiddleware.RouteHandler("HomePage", func(c *fiber.Ctx) map[string]interface{} {
        return map[string]interface{}{
            "user": c.Locals("user"),
            "path": c.Path(),
        }
    }))
}
```

### 3. 数据预取

```go
// 在 DataFetcher 中添加自定义数据获取逻辑
func (df *DataFetcher) FetchPageData(ctx context.Context, path string, userID *uint) (map[string]interface{}, error) {
    data := make(map[string]interface{})
    
    switch path {
    case "/dashboard":
        // 获取仪表板数据
        if userID != nil {
            user, _ := df.FetchUserData(ctx, *userID)
            data["user"] = user.ToResponse()
            
            // 获取统计数据
            data["stats"] = map[string]interface{}{
                "totalProjects": 5,
                "activeTasks":   12,
            }
        }
    }
    
    return data, nil
}
```

## 性能优化

### 1. 缓存策略

```go
// 页面缓存
func (r *Renderer) RenderPage(c *fiber.Ctx, componentName string, props map[string]interface{}) error {
    // 检查缓存
    if cachedHTML, err := r.ssrCache.GetPageCache(ctx, path, userID); err == nil {
        return c.SendString(cachedHTML)
    }
    
    // 渲染页面
    result, err := r.engine.Render(options)
    
    // 缓存结果
    r.ssrCache.SetPageCache(ctx, path, userID, result.HTML, 5*time.Minute)
    
    return nil
}
```

### 2. 数据预取优化

```go
// 并行数据获取
func (df *DataFetcher) FetchPageDataParallel(ctx context.Context, path string, userID *uint) (map[string]interface{}, error) {
    var wg sync.WaitGroup
    data := make(map[string]interface{})
    
    // 并行获取用户数据和统计数据
    wg.Add(2)
    
    go func() {
        defer wg.Done()
        if userID != nil {
            if user, err := df.FetchUserData(ctx, *userID); err == nil {
                data["user"] = user.ToResponse()
            }
        }
    }()
    
    go func() {
        defer wg.Done()
        data["stats"] = df.fetchStats(ctx)
    }()
    
    wg.Wait()
    return data, nil
}
```

### 3. 组件懒加载

```tsx
// 动态导入组件
const LazyComponent = React.lazy(() => import('./LazyComponent'))

function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <LazyComponent />
    </Suspense>
  )
}
```

## SEO 优化

### 1. 元数据管理

```go
// 在渲染器中设置 SEO 元数据
templateData := map[string]interface{}{
    "Title":       pageData["title"],
    "Description": pageData["description"],
    "Keywords":    "rexo,react,go,fiber,ssr",
    "OpenGraph": map[string]string{
        "og:title":       pageData["title"],
        "og:description": pageData["description"],
        "og:image":       "/images/og-image.jpg",
    },
}
```

### 2. 结构化数据

```tsx
// 在组件中添加结构化数据
function HomePage() {
  const structuredData = {
    "@context": "https://schema.org",
    "@type": "WebSite",
    "name": "Rexo",
    "description": "全栈 React 研发框架",
  }

  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
      />
      <div>...</div>
    </>
  )
}
```

## 监控和调试

### 1. 性能监控

```go
// 添加性能监控
func (r *Renderer) RenderPage(c *fiber.Ctx, componentName string, props map[string]interface{}) error {
    start := time.Now()
    
    // 渲染逻辑...
    
    duration := time.Since(start)
    log.Printf("SSR render time: %v for %s", duration, componentName)
    
    // 记录到监控系统
    metrics.RecordSSRRenderTime(componentName, duration)
    
    return nil
}
```

### 2. 调试信息

```tsx
// 开发环境调试信息
{process.env.NODE_ENV === 'development' && (
  <div className="debug-info">
    <h4>SSR Debug Info:</h4>
    <pre>{JSON.stringify({ 
      isServer, 
      path: data.path,
      user: data.user,
      ssrData: ssrData 
    }, null, 2)}</pre>
  </div>
)}
```

## 部署配置

### 1. 环境变量

```bash
# .env
SSR_ENABLED=true
SSR_CACHE_TTL=300
SSR_REDIS_URL=redis://localhost:6379
SSR_TIMEOUT=5s
```

### 2. Docker 配置

```dockerfile
# Dockerfile.ssr
FROM node:18-alpine AS frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build:ssr

FROM golang:1.21-alpine AS backend-build
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
COPY --from=frontend-build /app/frontend/dist ./dist
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=backend-build /app/backend/main .
CMD ["./main"]
```

## 最佳实践

### 1. 组件设计

- 确保组件在服务端和客户端都能正常渲染
- 避免使用浏览器特定的 API
- 使用 `useEffect` 处理客户端特定的逻辑

### 2. 数据获取

- 在服务端预取关键数据
- 使用缓存减少数据库查询
- 实现数据获取的超时和错误处理

### 3. 性能优化

- 启用页面缓存
- 使用 CDN 加速静态资源
- 实现组件级别的缓存

### 4. 错误处理

- 实现优雅的错误降级
- 记录详细的错误日志
- 提供用户友好的错误页面

## 故障排除

### 常见问题

1. **水合不匹配**: 确保服务端和客户端渲染结果一致
2. **内存泄漏**: 定期清理 V8 引擎实例
3. **缓存问题**: 检查缓存键的生成和过期策略
4. **性能问题**: 监控渲染时间和内存使用

### 调试工具

```bash
# 启用详细日志
export SSR_DEBUG=true

# 检查缓存状态
redis-cli keys "ssr:*"

# 性能分析
go tool pprof http://localhost:8080/debug/pprof/profile
```

## 总结

Rexo 的 SSR 功能提供了完整的服务端渲染解决方案，通过合理的架构设计和性能优化，能够显著提升应用的 SEO 表现和首屏加载速度。通过遵循最佳实践和充分利用缓存机制，可以构建高性能的全栈应用。
