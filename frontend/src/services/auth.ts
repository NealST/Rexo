import { apiService } from './api'
import type { LoginRequest, RegisterRequest, AuthResponse, User } from '@/types/api'

class AuthService {
  // 用户登录
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await apiService.post<AuthResponse>('/auth/login', credentials)
    
    if (response.success && response.data) {
      // 保存 token 和用户信息到本地存储
      localStorage.setItem('token', response.data.token)
      localStorage.setItem('user', JSON.stringify(response.data.user))
    }
    
    return response.data!
  }

  // 用户注册
  async register(userData: RegisterRequest): Promise<AuthResponse> {
    const response = await apiService.post<AuthResponse>('/auth/register', userData)
    
    if (response.success && response.data) {
      // 保存 token 和用户信息到本地存储
      localStorage.setItem('token', response.data.token)
      localStorage.setItem('user', JSON.stringify(response.data.user))
    }
    
    return response.data!
  }

  // 获取用户资料
  async getProfile(): Promise<User> {
    const response = await apiService.get<User>('/auth/profile')
    return response.data!
  }

  // 更新用户资料
  async updateProfile(userData: Partial<User>): Promise<User> {
    const response = await apiService.put<User>('/auth/profile', userData)
    return response.data!
  }

  // 用户登出
  async logout(): Promise<void> {
    try {
      await apiService.post('/auth/logout')
    } catch (error) {
      console.error('Logout error:', error)
    } finally {
      // 清除本地存储
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    }
  }

  // 刷新 token
  async refreshToken(): Promise<string> {
    const response = await apiService.post<{ token: string }>('/auth/refresh')
    const newToken = response.data!.token
    localStorage.setItem('token', newToken)
    return newToken
  }

  // 检查是否已登录
  isAuthenticated(): boolean {
    return !!localStorage.getItem('token')
  }

  // 获取当前用户
  getCurrentUser(): User | null {
    const userStr = localStorage.getItem('user')
    if (userStr) {
      try {
        return JSON.parse(userStr)
      } catch (error) {
        console.error('Error parsing user data:', error)
        return null
      }
    }
    return null
  }

  // 获取 token
  getToken(): string | null {
    return localStorage.getItem('token')
  }
}

export const authService = new AuthService()
