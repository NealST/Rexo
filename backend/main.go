package main

import (
	"log"
	"os"

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
)

// @title Rexo API
// @version 1.0
// @description 全栈 React 研发框架 API - 基于 Fiber
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

	// 创建 Fiber 应用
	app := fiber.New(fiber.Config{
		AppName:      "Rexo API v1.0",
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
			"message": "Rexo API is running",
			"version": "1.0.0",
		})
	})

	// 注册 API 路由
	v1.RegisterRoutes(app, db)

	// 启动服务器
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Rexo API server starting on port %s", port)
	log.Printf("📚 Swagger docs available at http://localhost:%s/swagger/", port)
	
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
