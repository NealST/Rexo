// API 响应基础结构
export interface ApiResponse<T = any> {
  success: boolean
  message?: string
  data?: T
  error?: string
  code?: number
}

// 分页响应
export interface PaginatedResponse<T> extends ApiResponse<T[]> {
  pagination?: {
    page: number
    limit: number
    total?: number
    pages?: number
  }
}

// 用户相关类型
export interface User {
  id: number
  email: string
  username: string
  first_name: string
  last_name: string
  avatar: string
  is_active: boolean
  is_admin: boolean
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  username: string
  password: string
  first_name?: string
  last_name?: string
}

export interface AuthResponse {
  user: User
  token: string
}

// 错误类型
export interface ApiError {
  success: false
  error: string
  code: number
}
