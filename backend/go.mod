module github.com/rexo/backend

go 1.21

require (
	github.com/gofiber/fiber/v2 v2.50.0
	github.com/gofiber/swagger v0.1.12
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/joho/godotenv v1.4.0
	golang.org/x/crypto v0.9.0
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.2
	github.com/redis/go-redis/v9 v9.0.5
	github.com/swaggo/swag v1.16.1
	github.com/gofiber/contrib/cors v1.0.0
	github.com/go-playground/validator/v10 v10.14.0
	github.com/gofiber/contrib/jwt v1.0.0
	github.com/gofiber/contrib/logger v1.0.0
	// SSR 相关依赖
	github.com/robertkrimen/otto v0.2.1
	github.com/gin-gonic/gin v1.9.1
	github.com/gorilla/websocket v1.5.0
)
