# Rexo - 全栈 React 研发框架

> 基于 Go (Fiber) + React (TypeScript) 的现代化全栈研发框架

## 🚀 特性

### 后端 (Go + Fiber)
- ⚡ **高性能**: 基于 Fiber 框架，性能比 Gin 更优
- 🔐 **JWT 认证**: 完整的用户认证和授权系统
- 🗄️ **数据库 ORM**: 集成 GORM，支持 PostgreSQL/MySQL
- �� **API 文档**: 自动生成 Swagger 文档
- 🛡️ **中间件**: 认证、日志、限流、CORS 等
- 🔄 **热重载**: 开发时自动重启

### 前端 (React + TypeScript)
- ⚛️ **React 18**: 最新版本的 React 框架
- 📘 **TypeScript**: 完整的类型安全
- ⚡ **Vite**: 极速的构建工具
- �� **Tailwind CSS**: 现代化的样式框架
- 🗃️ **Zustand**: 轻量级状态管理
- 🛣️ **React Router**: 现代化路由管理

### 开发工具
- 🔧 **CLI 工具**: 项目初始化、代码生成
- 🔥 **热重载**: 前后端同时热重载
- 📝 **代码规范**: ESLint + Prettier + Go fmt
- 🧪 **测试框架**: Jest + React Testing Library + Go testing
- 🐳 **Docker**: 容器化部署

## 📁 项目结构

```
Rexo/
├── backend/                 # Go 后端 (Fiber)
│   ├── api/v1/             # API 路由和处理器
│   ├── models/              # 数据模型
│   ├── middleware/          # 中间件
│   ├── config/              # 配置管理
│   ├── database/            # 数据库连接和迁移
│   ├── utils/               # 工具函数
│   └── main.go              # 应用入口
├── frontend/                # React 前端
│   ├── src/
│   │   ├── components/      # 可复用组件
│   │   ├── pages/           # 页面组件
│   │   ├── hooks/           # 自定义 Hooks
│   │   ├── services/        # API 服务
│   │   ├── store/           # 状态管理 (Zustand)
│   │   ├── types/           # TypeScript 类型定义
│   │   └── styles/          # 样式文件
│   ├── package.json          # 依赖管理
│   └── vite.config.ts       # Vite 配置
├── shared/                  # 共享代码
├── cli/                     # 命令行工具
├── templates/               # 项目模板
├── docs/                    # 文档
├── examples/                # 示例项目
└── docker/                  # Docker 配置
```

## 🛠️ 技术栈

### 后端技术
- **语言**: Go 1.21+
- **框架**: Fiber v2
- **数据库**: PostgreSQL/MySQL + GORM
- **认证**: JWT + bcrypt
- **缓存**: Redis
- **文档**: Swagger

### 前端技术
- **语言**: TypeScript
- **框架**: React 18
- **构建工具**: Vite
- **状态管理**: Zustand
- **路由**: React Router v6
- **样式**: Tailwind CSS
- **HTTP 客户端**: Axios

### 开发工具
- **包管理**: Go Modules + npm/yarn
- **代码规范**: ESLint + Prettier + Go fmt
- **测试**: Jest + React Testing Library + Go testing
- **类型检查**: TypeScript + Go vet
- **容器化**: Docker + Docker Compose

## 🚀 快速开始

### 环境要求
- Go 1.21+
- Node.js 18+
- PostgreSQL 13+ 或 MySQL 8+
- Redis 6+

### 1. 克隆项目
```bash
git clone https://github.com/rexo/rexo.git
cd rexo
```

### 2. 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库和 Redis 连接信息
```

### 3. 启动后端
```bash
cd backend
go mod tidy
go run main.go
```

### 4. 启动前端
```bash
cd frontend
npm install
npm run dev
```

### 5. 访问应用
- 前端应用: http://localhost:3000
- 后端 API: http://localhost:8080
- API 文档: http://localhost:8080/swagger/

## 📚 API 文档

### 认证相关
- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/auth/profile` - 获取用户资料
- `PUT /api/v1/auth/profile` - 更新用户资料
- `POST /api/v1/auth/logout` - 用户登出

### 用户管理
- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取单个用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

## 🔧 开发指南

### 添加新的 API 端点
1. 在 `backend/api/v1/handlers/` 中创建处理器
2. 在 `backend/api/v1/routes.go` 中注册路由
3. 在 `frontend/src/services/` 中添加 API 调用
4. 在 `frontend/src/types/api.ts` 中定义类型

### 添加新的页面
1. 在 `frontend/src/pages/` 中创建页面组件
2. 在 `frontend/src/App.tsx` 中添加路由
3. 在 `frontend/src/components/layout/Header.tsx` 中添加导航

### 数据库迁移
```bash
# 添加新的模型到 backend/models/
# 运行应用时会自动迁移
go run main.go
```

## 🐳 Docker 部署

### 开发环境
```bash
docker-compose up -d
```

### 生产环境
```bash
docker-compose -f docker-compose.prod.yml up -d
```

## 🧪 测试

### 后端测试
```bash
cd backend
go test ./...
```

### 前端测试
```bash
cd frontend
npm test
```

## 📈 性能优化

### 后端优化
- 使用连接池管理数据库连接
- 实现 Redis 缓存
- 启用 Gzip 压缩
- 使用 CDN 加速静态资源

### 前端优化
- 代码分割和懒加载
- 图片优化和压缩
- 使用 React.memo 和 useMemo
- 启用 Service Worker

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Fiber](https://gofiber.io/) - Go Web 框架
- [React](https://reactjs.org/) - 前端框架
- [Vite](https://vitejs.dev/) - 构建工具
- [Tailwind CSS](https://tailwindcss.com/) - CSS 框架
- [Zustand](https://zustand-demo.pmnd.rs/) - 状态管理

## 📞 联系我们

- 项目主页: https://github.com/rexo/rexo
- 问题反馈: https://github.com/rexo/rexo/issues
- 邮箱: contact@rexo.dev

---

⭐ 如果这个项目对你有帮助，请给我们一个 Star！
