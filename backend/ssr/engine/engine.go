package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/robertkrimen/otto"
)

// SSR Engine 用于在服务端渲染 React 组件
type Engine struct {
	vm       *otto.Otto
	reactJS  string
	appJS    string
	mu       sync.RWMutex
	basePath string
}

// RenderOptions 渲染选项
type RenderOptions struct {
	Component string                 `json:"component"`
	Props     map[string]interface{} `json:"props"`
	Path      string                 `json:"path"`
	Query     map[string]string      `json:"query"`
}

// RenderResult 渲染结果
type RenderResult struct {
	HTML     string                 `json:"html"`
	CSS      string                 `json:"css"`
	JS       string                 `json:"js"`
	Data     map[string]interface{} `json:"data"`
	Error    string                 `json:"error,omitempty"`
}

// NewEngine 创建新的 SSR 引擎
func NewEngine(basePath string) (*Engine, error) {
	engine := &Engine{
		basePath: basePath,
	}

	// 创建新的 JavaScript 虚拟机
	engine.vm = otto.New()

	// 加载 React 和应用的 JavaScript 代码
	if err := engine.loadScripts(); err != nil {
		return nil, fmt.Errorf("failed to load scripts: %w", err)
	}

	// 设置全局变量和函数
	if err := engine.setupGlobals(); err != nil {
		return nil, fmt.Errorf("failed to setup globals: %w", err)
	}

	return engine, nil
}

// loadScripts 加载必要的 JavaScript 文件
func (e *Engine) loadScripts() error {
	// 加载 React 库
	reactPath := filepath.Join(e.basePath, "dist", "react.js")
	if data, err := ioutil.ReadFile(reactPath); err == nil {
		e.reactJS = string(data)
	} else {
		// 如果文件不存在，使用内嵌的简化版本
		e.reactJS = e.getEmbeddedReact()
	}

	// 加载应用代码
	appPath := filepath.Join(e.basePath, "dist", "app.js")
	if data, err := ioutil.ReadFile(appPath); err == nil {
		e.appJS = string(data)
	} else {
		// 如果文件不存在，使用内嵌的简化版本
		e.appJS = e.getEmbeddedApp()
	}

	// 执行 React 库
	if _, err := e.vm.Run(e.reactJS); err != nil {
		return fmt.Errorf("failed to load React: %w", err)
	}

	// 执行应用代码
	if _, err := e.vm.Run(e.appJS); err != nil {
		return fmt.Errorf("failed to load app: %w", err)
	}

	return nil
}

// setupGlobals 设置全局变量和函数
func (e *Engine) setupGlobals() error {
	// 设置全局变量
	e.vm.Set("__SSR_ENV__", "server")
	e.vm.Set("__SSR_PATH__", "")
	e.vm.Set("__SSR_QUERY__", map[string]string{})
	e.vm.Set("__SSR_PROPS__", map[string]interface{}{})

	// 设置全局函数
	e.vm.Set("__SSR_SET_PATH__", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) > 0 {
			path, _ := call.Argument(0).ToString()
			e.vm.Set("__SSR_PATH__", path)
		}
		return otto.UndefinedValue()
	})

	e.vm.Set("__SSR_SET_QUERY__", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) > 0 {
			query, _ := call.Argument(0).Export()
			e.vm.Set("__SSR_QUERY__", query)
		}
		return otto.UndefinedValue()
	})

	e.vm.Set("__SSR_SET_PROPS__", func(call otto.FunctionCall) otto.Value {
		if len(call.ArgumentList) > 0 {
			props, _ := call.Argument(0).Export()
			e.vm.Set("__SSR_PROPS__", props)
		}
		return otto.UndefinedValue()
	})

	return nil
}

// Render 渲染 React 组件
func (e *Engine) Render(options RenderOptions) (*RenderResult, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// 设置渲染参数
	e.vm.Set("__SSR_PATH__", options.Path)
	e.vm.Set("__SSR_QUERY__", options.Query)
	e.vm.Set("__SSR_PROPS__", options.Props)

	// 构建渲染脚本
	renderScript := fmt.Sprintf(`
		try {
			// 设置渲染参数
			__SSR_SET_PATH__('%s');
			__SSR_SET_QUERY__(%s);
			__SSR_SET_PROPS__(%s);
			
			// 渲染组件
			var result = __SSR_RENDER__('%s');
			
			// 返回结果
			JSON.stringify({
				html: result.html,
				css: result.css || '',
				js: result.js || '',
				data: result.data || {}
			});
		} catch (error) {
			JSON.stringify({
				error: error.message
			});
		}
	`, 
		options.Path,
		e.marshalToJSON(options.Query),
		e.marshalToJSON(options.Props),
		options.Component,
	)

	// 执行渲染
	value, err := e.vm.Run(renderScript)
	if err != nil {
		return nil, fmt.Errorf("failed to execute render script: %w", err)
	}

	// 解析结果
	resultStr, err := value.ToString()
	if err != nil {
		return nil, fmt.Errorf("failed to convert result to string: %w", err)
	}

	var result RenderResult
	if err := json.Unmarshal([]byte(resultStr), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result: %w", err)
	}

	if result.Error != "" {
		return nil, fmt.Errorf("render error: %s", result.Error)
	}

	return &result, nil
}

// marshalToJSON 将 Go 对象转换为 JSON 字符串
func (e *Engine) marshalToJSON(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Failed to marshal object: %v", err)
		return "{}"
	}
	return string(data)
}

// getEmbeddedReact 获取内嵌的简化 React 库
func (e *Engine) getEmbeddedReact() string {
	return `
		// 简化的 React 实现用于 SSR
		var React = {
			createElement: function(type, props, ...children) {
				// 处理 children
				var allChildren = [];
				for (var i = 2; i < arguments.length; i++) {
					if (arguments[i] !== null && arguments[i] !== undefined) {
						allChildren.push(arguments[i]);
					}
				}
				
				return {
					type: type,
					props: props || {},
					children: allChildren
				};
			},
			Component: function(props) {
				this.props = props;
				this.state = {};
			},
			Fragment: function(props) {
				return props.children;
			},
			StrictMode: function(props) {
				return props.children;
			}
		};
		
		// 简化的渲染函数
		function renderToString(element) {
			if (typeof element === 'string' || typeof element === 'number') {
				return String(element);
			}
			
			if (element === null || element === undefined) {
				return '';
			}
			
			if (!element || !element.type) {
				return '';
			}
			
			var tag = element.type;
			var props = element.props || {};
			var children = element.children || [];
			
			// 处理 Fragment
			if (tag === React.Fragment) {
				var fragmentHtml = '';
				for (var i = 0; i < children.length; i++) {
					fragmentHtml += renderToString(children[i]);
				}
				return fragmentHtml;
			}
			
			// 处理函数组件
			if (typeof tag === 'function') {
				try {
					var componentProps = props || {};
					var componentElement = tag(componentProps);
					return renderToString(componentElement);
				} catch (error) {
					return '<div class="error">Component Error: ' + error.message + '</div>';
				}
			}
			
			var html = '<' + tag;
			
			// 添加属性
			for (var key in props) {
				if (key !== 'children' && props[key] !== undefined && props[key] !== null) {
					var value = String(props[key]);
					// 转义 HTML 属性
					value = value.replace(/"/g, '&quot;').replace(/'/g, '&#39;');
					html += ' ' + key + '="' + value + '"';
				}
			}
			
			html += '>';
			
			// 添加子元素
			for (var i = 0; i < children.length; i++) {
				html += renderToString(children[i]);
			}
			
			html += '</' + tag + '>';
			
			return html;
		}
		
		// 全局渲染函数
		window.__SSR_RENDER__ = function(componentName) {
			var component = window[componentName];
			if (!component) {
				throw new Error('Component not found: ' + componentName);
			}
			
			try {
				var element = component(__SSR_PROPS__);
				var html = renderToString(element);
				
				return {
					html: html,
					css: '',
					js: '',
					data: {}
				};
			} catch (error) {
				return {
					html: '<div class="error">Render Error: ' + error.message + '</div>',
					css: '',
					js: '',
					data: {}
				};
			}
		};
	`
}

// getEmbeddedApp 获取内嵌的简化应用代码
func (e *Engine) getEmbeddedApp() string {
	return `
		// 示例组件
		function HomePage(props) {
			var user = props.user || {};
			var path = props.path || '';
			
			return React.createElement('div', {
				className: 'min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100'
			}, [
				React.createElement('div', {
					className: 'container mx-auto px-4 py-16'
				}, [
					React.createElement('div', {
						className: 'text-center mb-16'
					}, [
						React.createElement('h1', {
							className: 'text-5xl font-bold text-gray-900 mb-6'
						}, [
							'Welcome to ',
							React.createElement('span', {
								className: 'text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-purple-600'
							}, 'Rexo')
						]),
						React.createElement('p', {
							className: 'text-xl text-gray-600 mb-8 max-w-3xl mx-auto'
						}, '基于 Go (Fiber) + React 的全栈研发框架，支持服务端渲染 (SSR)，提供现代化的开发体验和高效的开发工具。'),
						React.createElement('div', {
							className: 'inline-flex items-center px-4 py-2 bg-green-100 text-green-800 rounded-full text-sm font-medium mb-8'
						}, [
							React.createElement('div', {
								className: 'w-2 h-2 bg-green-500 rounded-full mr-2'
							}),
							'Server-Side Rendered'
						])
					])
				])
			]);
		}
		
		function AboutPage(props) {
			return React.createElement('div', {
				className: 'min-h-screen bg-white'
			}, [
				React.createElement('div', {
					className: 'container mx-auto px-4 py-16'
				}, [
					React.createElement('div', {
						className: 'text-center mb-16'
					}, [
						React.createElement('h1', {
							className: 'text-4xl font-bold text-gray-900 mb-6'
						}, '关于 Rexo'),
						React.createElement('p', {
							className: 'text-xl text-gray-600 max-w-3xl mx-auto'
						}, 'Rexo 是一个现代化的全栈 React 研发框架，集成了 Go 后端和 React 前端，提供完整的开发工具链和最佳实践。')
					])
				])
			]);
		}
		
		// 注册组件
		window.HomePage = HomePage;
		window.AboutPage = AboutPage;
	`
}
