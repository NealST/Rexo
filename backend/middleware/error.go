package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler(c *fiber.Ctx, err error) error {
	// 默认状态码
	code := fiber.StatusInternalServerError

	// 检查是否是 Fiber 错误
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// 记录错误日志
	log.Printf("Error: %v - Path: %s - Method: %s", err, c.Path(), c.Method())

	// 返回错误响应
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
		"code":    code,
	})
}
