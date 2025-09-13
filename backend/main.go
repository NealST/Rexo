package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/rexo/backend/api/v1"
	"github.com/rexo/backend/config"
	"github.com/rexo/backend/database"
	"github.com/rexo/backend/middleware"
	"github.com/rexo/backend/ssr/renderer"
)

// @title Rexo API
// @version 1.0
// @description 全栈 React 研发框架 API - 基于 Fiber + SSR
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 初始化配置
	cfg := config.Load()

	// 初始化数据库
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移数据库
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 获取项目根目录
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory:", err)
	}
	projectRoot = filepath.Dir(projectRoot) // 回到项目根目录

	// 初始化 SSR 渲染器
	ssrRenderer, err := renderer.NewRenderer(projectRoot, db)
	if err != nil {
		log.Printf("Failed to initialize SSR renderer: %v", err)
		log.Println("Continuing without SSR support...")
		ssrRenderer = nil
	}

	// 设置 Fiber 模式
	if cfg.Server.Environment == "production" {
		fiber.SetMode(fiber.ReleaseMode)
	}

	// 创建 Fiber 应用
	app := fiber.New(fiber.Config{
		AppName:      "Rexo API v1.0 with SSR",
		ErrorHandler: middleware.ErrorHandler,
	})

	// 添加中间件
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Server.CORSOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Swagger 文档
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 健康检查
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Rexo API is running with SSR support",
			"version": "1.0.0",
			"ssr":     ssrRenderer != nil,
		})
	})

	// 注册 API 路由
	v1.RegisterRoutes(app, db)

	// 注册 SSR 路由（如果 SSR 渲染器可用）
	if ssrRenderer != nil {
		registerSSRRoutes(app, ssrRenderer)
		log.Println("✅ SSR routes registered")
	} else {
		log.Println("⚠️  SSR routes not available")
	}

	// 启动服务器
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Rexo API server starting on port %s", port)
	log.Printf("📚 Swagger docs available at http://localhost:%s/swagger/", port)
	if ssrRenderer != nil {
		log.Printf("⚛️  SSR enabled - React components will be server-side rendered")
		log.Printf("🔧 SSR features: Data prefetching, SEO optimization, Performance monitoring")
	}
	
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// registerSSRRoutes 注册 SSR 路由
func registerSSRRoutes(app *fiber.App, renderer *renderer.Renderer) {
	// 创建 SSR 中间件
	ssrMiddleware := middleware.NewSSRMiddleware(renderer)

	// 首页
	app.Get("/", ssrMiddleware.RouteHandler("HomePage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))

	// 关于页面
	app.Get("/about", ssrMiddleware.RouteHandler("AboutPage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))

	// 登录页面
	app.Get("/login", ssrMiddleware.RouteHandler("LoginPage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))

	// 注册页面
	app.Get("/register", ssrMiddleware.RouteHandler("RegisterPage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))

	// 仪表板页面（需要认证）
	app.Get("/dashboard", middleware.AuthMiddleware(), ssrMiddleware.RouteHandler("DashboardPage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))

	// 个人资料页面（需要认证）
	app.Get("/profile", middleware.AuthMiddleware(), ssrMiddleware.RouteHandler("ProfilePage", func(c *fiber.Ctx) map[string]interface{} {
		return map[string]interface{}{
			"user": c.Locals("user"),
			"path": c.Path(),
		}
	}))
}
