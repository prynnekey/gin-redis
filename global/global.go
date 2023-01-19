package global

import (
	"context"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

var (
	// DB 数据库
	DB *gorm.DB
	// Redis
	Redis *redis.Client
	// redis 上下文
	Ctx = context.Background()
)
