export function Footer() {
  return (
    <footer className="bg-gray-50 border-t">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          {/* Logo and Description */}
          <div className="col-span-1 md:col-span-2">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">Rexo</h3>
            <p className="text-gray-600 text-sm">
              基于 Go + React 的全栈研发框架，提供现代化的开发体验和高效的开发工具。
            </p>
          </div>

          {/* Quick Links */}
          <div>
            <h4 className="text-sm font-semibold text-gray-900 mb-3">快速链接</h4>
            <ul className="space-y-2">
              <li>
                <a href="/docs" className="text-sm text-gray-600 hover:text-primary-600">
                  文档
                </a>
              </li>
              <li>
                <a href="/examples" className="text-sm text-gray-600 hover:text-primary-600">
                  示例
                </a>
              </li>
              <li>
                <a href="/github" className="text-sm text-gray-600 hover:text-primary-600">
                  GitHub
                </a>
              </li>
            </ul>
          </div>

          {/* Community */}
          <div>
            <h4 className="text-sm font-semibold text-gray-900 mb-3">社区</h4>
            <ul className="space-y-2">
              <li>
                <a href="/discord" className="text-sm text-gray-600 hover:text-primary-600">
                  Discord
                </a>
              </li>
              <li>
                <a href="/twitter" className="text-sm text-gray-600 hover:text-primary-600">
                  Twitter
                </a>
              </li>
              <li>
                <a href="/blog" className="text-sm text-gray-600 hover:text-primary-600">
                  博客
                </a>
              </li>
            </ul>
          </div>
        </div>

        <div className="mt-8 pt-8 border-t border-gray-200">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <p className="text-sm text-gray-600">
              © 2024 Rexo Framework. All rights reserved.
            </p>
            <div className="flex space-x-6 mt-4 md:mt-0">
              <a href="/privacy" className="text-sm text-gray-600 hover:text-primary-600">
                隐私政策
              </a>
              <a href="/terms" className="text-sm text-gray-600 hover:text-primary-600">
                服务条款
              </a>
            </div>
          </div>
        </div>
      </div>
    </footer>
  )
}
