package test

import (
	"financialproduct/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestProduct(t *testing.T) {
	dsn := "root:prynnekey@tcp(127.0.0.1:3306)/financial_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Product{})
}
