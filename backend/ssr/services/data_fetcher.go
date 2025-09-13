package services

import (
	"context"
	"fmt"
	"time"

	"github.com/rexo/backend/models"
	"gorm.io/gorm"
)

// DataFetcher 数据预取服务
type DataFetcher struct {
	db *gorm.DB
}

// NewDataFetcher 创建数据预取服务
func NewDataFetcher(db *gorm.DB) *DataFetcher {
	return &DataFetcher{
		db: db,
	}
}

// FetchUserData 获取用户数据
func (df *DataFetcher) FetchUserData(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	
	if err := df.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	
	return &user, nil
}

// FetchPageData 根据路径获取页面数据
func (df *DataFetcher) FetchPageData(ctx context.Context, path string, userID *uint) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	
	// 根据路径获取不同的数据
	switch path {
	case "/":
		data["pageType"] = "home"
		data["title"] = "Rexo - 全栈 React 研发框架"
		data["description"] = "基于 Go + React 的全栈研发框架，支持服务端渲染"
		
		// 如果有用户ID，获取用户信息
		if userID != nil {
			if user, err := df.FetchUserData(ctx, *userID); err == nil {
				data["user"] = user.ToResponse()
			}
		}
		
	case "/about":
		data["pageType"] = "about"
		data["title"] = "关于 Rexo"
		data["description"] = "了解 Rexo 框架的特性和优势"
		
	case "/dashboard":
		data["pageType"] = "dashboard"
		data["title"] = "仪表板"
		data["description"] = "管理您的项目和应用程序"
		
		// 获取仪表板数据
		if userID != nil {
			if user, err := df.FetchUserData(ctx, *userID); err == nil {
				data["user"] = user.ToResponse()
			}
			
			// 这里可以添加更多仪表板相关的数据
			data["stats"] = map[string]interface{}{
				"totalProjects": 5,
				"activeTasks":   12,
				"completedTasks": 8,
			}
		}
		
	case "/profile":
		data["pageType"] = "profile"
		data["title"] = "个人资料"
		data["description"] = "管理您的个人资料和设置"
		
		if userID != nil {
			if user, err := df.FetchUserData(ctx, *userID); err == nil {
				data["user"] = user.ToResponse()
			}
		}
		
	default:
		data["pageType"] = "unknown"
		data["title"] = "Rexo"
		data["description"] = "基于 Go + React 的全栈研发框架"
	}
	
	// 添加通用数据
	data["timestamp"] = time.Now().Unix()
	data["path"] = path
	
	return data, nil
}

// FetchWithTimeout 带超时的数据获取
func (df *DataFetcher) FetchWithTimeout(ctx context.Context, timeout time.Duration, fetchFunc func(context.Context) (map[string]interface{}, error)) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	
	return fetchFunc(ctx)
}
