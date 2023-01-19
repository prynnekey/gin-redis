package routers

import (
	"financialproduct/service"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	// 获取所有产品
	r.GET("/", service.GetProduct())

	// 根据名称获取产品
	r.GET("/:name", service.GetProductByName())

	// 添加产品
	r.POST("/", service.AddProduct())

	// 根据id删除产品
	r.DELETE("/:id", service.DeleteProductById())

	// 根据id更新产品
	r.PUT("/", service.UpdateProductById())

	return r
}
