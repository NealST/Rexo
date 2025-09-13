package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rexo/backend/ssr/renderer"
)

// SSRMiddleware SSR 中间件
type SSRMiddleware struct {
	renderer *renderer.Renderer
}

// NewSSRMiddleware 创建 SSR 中间件
func NewSSRMiddleware(renderer *renderer.Renderer) *SSRMiddleware {
	return &SSRMiddleware{
		renderer: renderer,
	}
}

// Handle SSR 处理函数
func (m *SSRMiddleware) Handle(componentName string, getProps func(*fiber.Ctx) map[string]interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 检查是否为 SSR 请求
		if !m.renderer.IsSSRRequest(c) {
			// 不是 SSR 请求，返回 JSON 数据
			return m.renderer.RenderAPI(c, componentName, getProps(c))
		}

		// 是 SSR 请求，渲染完整页面
		return m.renderer.RenderPage(c, componentName, getProps(c))
	}
}

// RouteHandler 路由处理器
func (m *SSRMiddleware) RouteHandler(componentName string, getProps func(*fiber.Ctx) map[string]interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取路径
		path := c.Path()
		
		// 检查是否为 API 路径
		if strings.HasPrefix(path, "/api/") {
			return c.Next()
		}

		// 检查是否为静态资源
		if m.isStaticResource(path) {
			return c.Next()
		}

		// 执行 SSR 渲染
		return m.Handle(componentName, getProps)(c)
	}
}

// isStaticResource 检查是否为静态资源
func (m *SSRMiddleware) isStaticResource(path string) bool {
	staticExtensions := []string{
		".js", ".css", ".png", ".jpg", ".jpeg", ".gif", ".svg", ".ico",
		".woff", ".woff2", ".ttf", ".eot", ".pdf", ".zip", ".mp4", ".mp3",
	}
	
	for _, ext := range staticExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	
	return false
}

// DefaultProps 默认属性获取函数
func (m *SSRMiddleware) DefaultProps(c *fiber.Ctx) map[string]interface{} {
	return map[string]interface{}{
		"path":  c.Path(),
		"query": c.Queries(),
		"user":  c.Locals("user"),
	}
}
