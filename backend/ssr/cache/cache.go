package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache 缓存接口
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
}

// RedisCache Redis 缓存实现
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache 创建 Redis 缓存
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) Clear(ctx context.Context) error {
	return r.client.FlushDB(ctx).Err()
}

// MemoryCache 内存缓存实现
type MemoryCache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
}

type cacheItem struct {
	value      string
	expiration time.Time
}

// NewMemoryCache 创建内存缓存
func NewMemoryCache() *MemoryCache {
	cache := &MemoryCache{
		data: make(map[string]cacheItem),
	}
	
	// 启动清理协程
	go cache.cleanup()
	
	return cache
}

func (m *MemoryCache) Get(ctx context.Context, key string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	item, exists := m.data[key]
	if !exists {
		return "", fmt.Errorf("key not found")
	}
	
	if time.Now().After(item.expiration) {
		return "", fmt.Errorf("key expired")
	}
	
	return item.value, nil
}

func (m *MemoryCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.data[key] = cacheItem{
		value:      string(data),
		expiration: time.Now().Add(expiration),
	}
	
	return nil
}

func (m *MemoryCache) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	delete(m.data, key)
	return nil
}

func (m *MemoryCache) Clear(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.data = make(map[string]cacheItem)
	return nil
}

func (m *MemoryCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		m.mu.Lock()
		now := time.Now()
		for key, item := range m.data {
			if now.After(item.expiration) {
				delete(m.data, key)
			}
		}
		m.mu.Unlock()
	}
}

// SSRCache SSR 缓存管理器
type SSRCache struct {
	cache Cache
}

// NewSSRCache 创建 SSR 缓存管理器
func NewSSRCache(cache Cache) *SSRCache {
	return &SSRCache{
		cache: cache,
	}
}

// GetPageCache 获取页面缓存
func (s *SSRCache) GetPageCache(ctx context.Context, path string, userID *uint) (string, error) {
	key := s.generatePageKey(path, userID)
	return s.cache.Get(ctx, key)
}

// SetPageCache 设置页面缓存
func (s *SSRCache) SetPageCache(ctx context.Context, path string, userID *uint, html string, expiration time.Duration) error {
	key := s.generatePageKey(path, userID)
	return s.cache.Set(ctx, key, html, expiration)
}

// GetDataCache 获取数据缓存
func (s *SSRCache) GetDataCache(ctx context.Context, key string) (map[string]interface{}, error) {
	data, err := s.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, err
	}
	
	return result, nil
}

// SetDataCache 设置数据缓存
func (s *SSRCache) SetDataCache(ctx context.Context, key string, data map[string]interface{}, expiration time.Duration) error {
	return s.cache.Set(ctx, key, data, expiration)
}

// generatePageKey 生成页面缓存键
func (s *SSRCache) generatePageKey(path string, userID *uint) string {
	if userID != nil {
		return fmt.Sprintf("ssr:page:%s:user:%d", path, userID)
	}
	return fmt.Sprintf("ssr:page:%s:guest", path)
}

// ClearUserCache 清除用户相关缓存
func (s *SSRCache) ClearUserCache(ctx context.Context, userID uint) error {
	// 这里可以实现更复杂的用户缓存清理逻辑
	// 例如清除用户相关的所有页面缓存
	return nil
}
