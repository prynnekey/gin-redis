package config

import (
	"financialproduct/global"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() {
	// 初始化日志
	// initLog()
	// 初始化数据库
	initDB()
	// 初始化redis
	initRedis()
}

func initDB() {
	dsn := "root:prynnekey@tcp(127.0.0.1:3306)/financial_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
}

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.115.18:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	global.Redis = rdb
}
