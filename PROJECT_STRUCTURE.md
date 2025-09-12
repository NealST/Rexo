# Rexo 全栈 React 研发框架 - 工程结构

## 整体架构

```
Rexo/                          # 项目根目录
├── backend/                   # Go 后端框架
│   ├── api/                   # API 路由和控制器
│   │   ├── v1/               # API 版本管理
│   │   │   ├── auth.go       # 认证相关 API
│   │   │   ├── user.go       # 用户管理 API
│   │   │   └── common.go     # 通用 API
│   │   └── middleware/       # API 中间件
│   │       ├── auth.go       # 认证中间件
│   │       ├── cors.go       # CORS 处理
│   │       ├── logging.go    # 日志中间件
│   │       └── rate_limit.go # 限流中间件
│   ├── models/               # 数据模型
│   │   ├── user.go          # 用户模型
│   │   ├── auth.go          # 认证模型
│   │   └── base.go          # 基础模型
│   ├── services/            # 业务逻辑层
│   │   ├── auth_service.go  # 认证服务
│   │   ├── user_service.go  # 用户服务
│   │   └── email_service.go # 邮件服务
│   ├── database/            # 数据库相关
│   │   ├── connection.go    # 数据库连接
│   │   ├── migrations/      # 数据库迁移
│   │   └── seeds/           # 种子数据
│   ├── config/              # 配置管理
│   │   ├── config.go        # 配置结构
│   │   ├── database.go      # 数据库配置
│   │   └── server.go        # 服务器配置
│   ├── utils/               # 工具函数
│   │   ├── jwt.go          # JWT 工具
│   │   ├── password.go     # 密码处理
│   │   ├── validation.go   # 数据验证
│   │   └── response.go     # 响应格式化
│   ├── handlers/            # HTTP 处理器
│   ├── main.go             # 应用入口
│   ├── go.mod              # Go 模块文件
│   └── go.sum              # Go 依赖锁定
│
├── frontend/                 # React 前端框架
│   ├── src/
│   │   ├── components/      # 可复用组件
│   │   │   ├── ui/         # 基础 UI 组件
│   │   │   ├── forms/      # 表单组件
│   │   │   ├── layout/     # 布局组件
│   │   │   └── common/     # 通用组件
│   │   ├── pages/          # 页面组件
│   │   │   ├── auth/       # 认证页面
│   │   │   ├── dashboard/  # 仪表板
│   │   │   └── profile/    # 用户资料
│   │   ├── hooks/          # 自定义 Hooks
│   │   │   ├── useAuth.ts  # 认证 Hook
│   │   │   ├── useApi.ts   # API 调用 Hook
│   │   │   └── useLocalStorage.ts
│   │   ├── services/       # API 服务
│   │   │   ├── api.ts      # API 客户端
│   │   │   ├── auth.ts     # 认证服务
│   │   │   └── user.ts     # 用户服务
│   │   ├── store/          # 状态管理
│   │   │   ├── index.ts    # Store 配置
│   │   │   ├── auth.ts     # 认证状态
│   │   │   └── user.ts     # 用户状态
│   │   ├── utils/          # 工具函数
│   │   │   ├── constants.ts # 常量定义
│   │   │   ├── helpers.ts  # 辅助函数
│   │   │   └── validation.ts # 前端验证
│   │   ├── types/          # TypeScript 类型定义
│   │   │   ├── api.ts      # API 类型
│   │   │   ├── auth.ts     # 认证类型
│   │   │   └── user.ts     # 用户类型
│   │   ├── styles/         # 样式文件
│   │   │   ├── globals.css # 全局样式
│   │   │   ├── components/ # 组件样式
│   │   │   └── themes/     # 主题配置
│   │   ├── App.tsx         # 应用根组件
│   │   └── index.tsx       # 应用入口
│   ├── public/             # 静态资源
│   ├── package.json        # 依赖管理
│   ├── tsconfig.json       # TypeScript 配置
│   ├── tailwind.config.js  # Tailwind CSS 配置
│   └── vite.config.ts      # Vite 构建配置
│
├── shared/                  # 共享代码
│   ├── types/              # 共享类型定义
│   │   ├── api.ts          # API 接口类型
│   │   ├── auth.ts         # 认证相关类型
│   │   └── common.ts       # 通用类型
│   ├── constants/          # 共享常量
│   │   ├── api.ts          # API 常量
│   │   └── auth.ts         # 认证常量
│   └── utils/              # 共享工具函数
│       ├── validation.ts   # 验证规则
│       └── formatting.ts   # 格式化函数
│
├── cli/                     # 命令行工具
│   ├── commands/           # 命令实现
│   │   ├── init.go         # 项目初始化
│   │   ├── generate.go     # 代码生成
│   │   ├── dev.go          # 开发服务器
│   │   └── build.go        # 项目构建
│   ├── templates/          # 代码模板
│   │   ├── component/      # 组件模板
│   │   ├── page/           # 页面模板
│   │   ├── api/            # API 模板
│   │   └── model/          # 模型模板
│   ├── main.go             # CLI 入口
│   └── go.mod              # CLI 依赖
│
├── templates/               # 项目模板
│   ├── basic/              # 基础模板
│   ├── auth/               # 带认证的模板
│   ├── admin/              # 管理后台模板
│   └── api-only/           # 纯 API 模板
│
├── docs/                    # 文档
│   ├── getting-started.md  # 快速开始
│   ├── architecture.md     # 架构说明
│   ├── api/                # API 文档
│   ├── frontend/           # 前端文档
│   └── deployment.md       # 部署指南
│
├── examples/                # 示例项目
│   ├── todo-app/           # Todo 应用示例
│   ├── blog/               # 博客系统示例
│   └── e-commerce/         # 电商系统示例
│
├── scripts/                 # 构建脚本
│   ├── build.sh            # 构建脚本
│   ├── dev.sh              # 开发脚本
│   └── deploy.sh           # 部署脚本
│
├── .github/                 # GitHub 配置
│   ├── workflows/          # CI/CD 工作流
│   │   ├── ci.yml          # 持续集成
│   │   └── deploy.yml      # 部署流程
│   └── ISSUE_TEMPLATE/     # Issue 模板
│
├── docker/                  # Docker 配置
│   ├── Dockerfile.backend  # 后端 Dockerfile
│   ├── Dockerfile.frontend # 前端 Dockerfile
│   └── docker-compose.yml  # 开发环境
│
├── .gitignore              # Git 忽略文件
├── .env.example            # 环境变量示例
├── LICENSE                 # 开源协议
├── README.md               # 项目说明
└── rexo.yaml               # 框架配置文件
```

## 核心特性

### 后端 (Go)
- **高性能**: 基于 Gin 框架，支持高并发
- **类型安全**: 完整的类型定义和验证
- **ORM**: 集成 GORM，支持多种数据库
- **认证**: JWT + Refresh Token 机制
- **中间件**: 认证、日志、限流、CORS 等
- **API 版本管理**: 支持多版本 API
- **热重载**: 开发时自动重启

### 前端 (React)
- **现代化**: React 18 + TypeScript + Vite
- **状态管理**: Zustand 轻量级状态管理
- **路由**: React Router v6
- **UI 框架**: Tailwind CSS + Headless UI
- **类型安全**: 与后端共享类型定义
- **代码分割**: 自动代码分割和懒加载
- **PWA 支持**: 渐进式 Web 应用

### 开发工具
- **CLI 工具**: 项目初始化、代码生成
- **热重载**: 前后端同时热重载
- **类型生成**: 自动生成前后端类型定义
- **代码规范**: ESLint + Prettier + Go fmt
- **测试**: Jest + React Testing Library + Go testing

### 工程化
- **Docker**: 容器化部署
- **CI/CD**: GitHub Actions 自动化
- **环境管理**: 多环境配置
- **监控**: 日志和性能监控
- **文档**: 自动生成 API 文档

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Gin
- **数据库**: PostgreSQL/MySQL + GORM
- **认证**: JWT + bcrypt
- **缓存**: Redis
- **文档**: Swagger

### 前端
- **语言**: TypeScript
- **框架**: React 18
- **构建工具**: Vite
- **状态管理**: Zustand
- **路由**: React Router v6
- **样式**: Tailwind CSS
- **UI 组件**: Headless UI

### 开发工具
- **包管理**: Go Modules + npm/yarn
- **代码规范**: ESLint + Prettier + Go fmt
- **测试**: Jest + React Testing Library + Go testing
- **类型检查**: TypeScript + Go vet
- **容器化**: Docker + Docker Compose

## 项目初始化流程

1. **安装 CLI**: `go install github.com/rexo/cli@latest`
2. **创建项目**: `rexo init my-project`
3. **选择模板**: 基础模板、认证模板、管理后台等
4. **配置环境**: 数据库、Redis、JWT 密钥等
5. **启动开发**: `rexo dev`

## 开发工作流

1. **API 优先**: 先定义 API 接口和类型
2. **类型共享**: 前后端共享类型定义
3. **热重载开发**: 前后端同时热重载
4. **代码生成**: 使用 CLI 生成 CRUD 代码
5. **测试驱动**: 单元测试 + 集成测试
6. **持续集成**: 自动测试和部署
