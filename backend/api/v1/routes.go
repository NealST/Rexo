package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rexo/backend/api/v1/handlers"
	"github.com/rexo/backend/middleware"
	"gorm.io/gorm"
)

// RegisterRoutes 注册所有 API 路由
func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	// 创建 API v1 路由组
	api := app.Group("/api/v1")

	// 初始化处理器
	authHandler := handlers.NewAuthHandler(db)
	userHandler := handlers.NewUserHandler(db)

	// 公开路由（不需要认证）
	public := api.Group("/")
	public.Post("/auth/register", authHandler.Register)
	public.Post("/auth/login", authHandler.Login)
	public.Post("/auth/refresh", authHandler.RefreshToken)

	// 受保护的路由（需要认证）
	protected := api.Group("/", middleware.AuthMiddleware())
	protected.Get("/auth/profile", authHandler.Profile)
	protected.Put("/auth/profile", authHandler.UpdateProfile)
	protected.Post("/auth/logout", authHandler.Logout)

	// 用户管理路由
	protected.Get("/users", userHandler.GetUsers)
	protected.Get("/users/:id", userHandler.GetUser)
	protected.Put("/users/:id", userHandler.UpdateUser)
	protected.Delete("/users/:id", userHandler.DeleteUser)
}
