package renderer

import (
	"context"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rexo/backend/ssr/engine"
	"github.com/rexo/backend/ssr/services"
	"gorm.io/gorm"
)

// Renderer SSR 渲染器
type Renderer struct {
	engine      *engine.Engine
	templates   map[string]*template.Template
	basePath    string
	dataFetcher *services.DataFetcher
}

// NewRenderer 创建新的渲染器
func NewRenderer(basePath string, db *gorm.DB) (*Renderer, error) {
	// 创建 SSR 引擎
	ssrEngine, err := engine.NewEngine(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create SSR engine: %w", err)
	}

	// 创建数据预取服务
	dataFetcher := services.NewDataFetcher(db)

	renderer := &Renderer{
		engine:      ssrEngine,
		templates:   make(map[string]*template.Template),
		basePath:    basePath,
		dataFetcher: dataFetcher,
	}

	// 加载模板
	if err := renderer.loadTemplates(); err != nil {
		return nil, fmt.Errorf("failed to load templates: %w", err)
	}

	return renderer, nil
}

// loadTemplates 加载 HTML 模板
func (r *Renderer) loadTemplates() error {
	templatePath := filepath.Join(r.basePath, "ssr", "templates")
	
	// 默认模板
	defaultTemplate := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <meta name="description" content="{{.Description}}">
    <meta name="keywords" content="{{.Keywords}}">
    <link rel="icon" type="image/svg+xml" href="/vite.svg">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap" rel="stylesheet">
    {{if .CSS}}<style>{{.CSS}}</style>{{end}}
    <script>
        // 预加载关键资源
        window.__SSR_DATA__ = {{.Data}};
        window.__SSR_META__ = {
            title: "{{.Title}}",
            description: "{{.Description}}",
            path: "{{.Path}}",
            timestamp: {{.Timestamp}}
        };
    </script>
</head>
<body>
    <div id="root">{{.HTML}}</div>
    {{if .JS}}<script>{{.JS}}</script>{{end}}
    <script type="module" src="/src/main.tsx"></script>
    <script>
        // 性能监控
        window.addEventListener('load', function() {
            if (window.performance && window.performance.timing) {
                var timing = window.performance.timing;
                var loadTime = timing.loadEventEnd - timing.navigationStart;
                console.log('Page load time:', loadTime + 'ms');
            }
        });
    </script>
</body>
</html>`

	tmpl, err := template.New("default").Parse(defaultTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse default template: %w", err)
	}

	r.templates["default"] = tmpl
	return nil
}

// RenderPage 渲染页面
func (r *Renderer) RenderPage(c *fiber.Ctx, componentName string, props map[string]interface{}) error {
	// 获取路径和查询参数
	path := c.Path()
	query := make(map[string]string)
	for key, value := range c.Queries() {
		query[key] = value
	}

	// 创建上下文，设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 预取页面数据
	var pageData map[string]interface{}
	var userID *uint
	
	// 检查是否有用户信息
	if user := c.Locals("userID"); user != nil {
		if uid, ok := user.(uint); ok {
			userID = &uid
		}
	}

	// 获取页面数据
	var err error
	pageData, err = r.dataFetcher.FetchWithTimeout(ctx, 3*time.Second, func(ctx context.Context) (map[string]interface{}, error) {
		return r.dataFetcher.FetchPageData(ctx, path, userID)
	})

	if err != nil {
		// 如果数据获取失败，使用默认数据
		pageData = map[string]interface{}{
			"title":       "Rexo",
			"description": "基于 Go + React 的全栈研发框架",
			"path":        path,
		}
	}

	// 合并 props 和页面数据
	finalProps := make(map[string]interface{})
	for k, v := range props {
		finalProps[k] = v
	}
	for k, v := range pageData {
		finalProps[k] = v
	}

	// 渲染选项
	options := engine.RenderOptions{
		Component: componentName,
		Props:     finalProps,
		Path:      path,
		Query:     query,
	}

	// 执行 SSR 渲染
	result, err := r.engine.Render(options)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "SSR rendering failed",
			"details": err.Error(),
		})
	}

	// 准备模板数据
	templateData := map[string]interface{}{
		"Title":       pageData["title"],
		"Description": pageData["description"],
		"Keywords":    "rexo,react,go,fiber,ssr,fullstack",
		"HTML":        template.HTML(result.HTML),
		"CSS":        result.CSS,
		"JS":         result.JS,
		"Data":       result.Data,
		"Path":       path,
		"Timestamp":  time.Now().Unix(),
	}

	// 渲染 HTML 模板
	tmpl := r.templates["default"]
	if err := tmpl.Execute(c.Response().BodyWriter(), templateData); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Template rendering failed",
			"details": err.Error(),
		})
	}

	// 设置响应头
	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Set("Cache-Control", "public, max-age=300") // 5分钟缓存
	
	return nil
}

// RenderAPI 渲染 API 响应（用于 AJAX 请求）
func (r *Renderer) RenderAPI(c *fiber.Ctx, componentName string, props map[string]interface{}) error {
	// 获取路径和查询参数
	path := c.Path()
	query := make(map[string]string)
	for key, value := range c.Queries() {
		query[key] = value
	}

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 预取数据
	var userID *uint
	if user := c.Locals("userID"); user != nil {
		if uid, ok := user.(uint); ok {
			userID = &uid
		}
	}

	pageData, err := r.dataFetcher.FetchPageData(ctx, path, userID)
	if err != nil {
		pageData = map[string]interface{}{
			"path": path,
		}
	}

	// 合并数据
	finalProps := make(map[string]interface{})
	for k, v := range props {
		finalProps[k] = v
	}
	for k, v := range pageData {
		finalProps[k] = v
	}

	// 渲染选项
	options := engine.RenderOptions{
		Component: componentName,
		Props:     finalProps,
		Path:      path,
		Query:     query,
	}

	// 执行 SSR 渲染
	result, err := r.engine.Render(options)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error": "SSR rendering failed",
			"details": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"html": result.HTML,
			"css":  result.CSS,
			"js":   result.JS,
			"data": result.Data,
		},
	})
}

// IsSSRRequest 检查是否为 SSR 请求
func (r *Renderer) IsSSRRequest(c *fiber.Ctx) bool {
	// 检查请求头
	accept := c.Get("Accept")
	userAgent := c.Get("User-Agent")
	
	// 如果是 AJAX 请求，返回 false
	if strings.Contains(accept, "application/json") {
		return false
	}
	
	// 如果是爬虫或搜索引擎，返回 true
	botKeywords := []string{"bot", "crawler", "spider", "googlebot", "bingbot", "baiduspider"}
	for _, keyword := range botKeywords {
		if strings.Contains(strings.ToLower(userAgent), keyword) {
			return true
		}
	}
	
	// 检查是否为预渲染请求
	if c.Get("X-Prerender") != "" {
		return true
	}
	
	// 默认返回 true（启用 SSR）
	return true
}
