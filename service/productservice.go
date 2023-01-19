package service

import (
	"financialproduct/common/response"
	"financialproduct/global"
	"financialproduct/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取产品信息
func GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products models.ProductSlice
		// 先从redis中获取数据
		global.Redis.Get(global.Ctx, "product:all").Scan(&products)
		if products != nil {
			response.Success(ctx, products, "获取产品信息成功")
			return
		}

		// 如果redis中没有数据，再从mysql中获取数据
		err := global.DB.Find(&products).Error
		if err != nil {
			response.Fail(ctx, "获取产品信息失败"+err.Error())
			return
		}

		// 将数据写入redis
		// 单个数据 将来根据名称可以获取到
		// for _, v := range products {
		// 	err = global.Redis.Set(global.Ctx, "product:list:"+v.Name, v, 0).Err()
		// 	if err != nil {
		// 		response.Fail(ctx, "数据缓存失败"+err.Error())
		// 		return
		// 	}
		// }

		// 全部数据
		err = global.Redis.Set(global.Ctx, "product:all", products, 0).Err()
		if err != nil {
			response.Fail(ctx, "数据缓存失败"+err.Error())
			return
		}

		response.Success(ctx, products, "获取产品信息成功")
	}
}

// 根据产品名称获取产品信息
func GetProductByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		var product models.ProductSlice
		var count int64

		// 已弃用 会产生脏数据
		// 从redis中查询数据
		// global.Redis.Get(global.Ctx, "product:name:"+name).Scan(&product)
		// if product != nil {
		// 	response.Success(ctx, gin.H{
		// 		"count": count,
		// 		"list":  product,
		// 	}, "获取产品信息成功")
		// 	return
		// }

		err := global.DB.Model(&models.Product{}).Where("name like ?", "%"+name+"%").Count(&count).Find(&product).Error
		if err != nil {
			response.Fail(ctx, "获取产品信息失败"+err.Error())
			return
		}

		// 将数据保存到redis中
		// err = global.Redis.Set(global.Ctx, "product:name:"+name, product, 0).Err()
		// if err != nil {
		// 	response.Fail(ctx, "数据缓存失败"+err.Error())
		// 	return
		// }

		response.Success(ctx, gin.H{
			"count": count,
			"list":  product,
		}, "获取产品信息成功")
	}
}

// 增加产品信息
func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			response.Fail(ctx, "参数错误"+err.Error())
			return
		}

		err = global.DB.Create(&product).Error
		if err != nil {
			response.Fail(ctx, "添加产品信息失败"+err.Error())
			return
		}

		// 删除redis中的数据
		global.Redis.Del(global.Ctx, "product:all")

		response.Success(ctx, nil, "添加产品信息成功")
	}
}

// 根据id删除产品信息
func DeleteProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		tx := global.DB.Delete(&models.Product{}, id)
		if tx.Error != nil {
			if tx.Error == gorm.ErrRecordNotFound {
				response.Fail(ctx, "删除产品信息失败，未找到该产品")
				return
			}
			response.Fail(ctx, "删除产品信息失败"+tx.Error.Error())
			return
		}

		if tx.RowsAffected == 0 {
			response.Fail(ctx, "删除产品信息失败，未找到该产品")
			return
		}

		// 删除redis中的数据
		global.Redis.Del(global.Ctx, "product:all")

		response.Success(ctx, nil, "删除产品信息成功")
	}
}

// 根据id更新产品信息
func UpdateProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			response.Fail(ctx, "参数错误"+err.Error())
			return
		}

		tx := global.DB.Model(&models.Product{}).Omit("id").Where("id = ?", product.ID).Updates(product)
		if tx.Error != nil {
			response.Fail(ctx, "更新产品信息失败"+tx.Error.Error())
			return
		}

		if tx.RowsAffected == 0 {
			response.Fail(ctx, "删除产品信息失败，未找到该产品")
			return
		}

		// 删除redis中的数据
		global.Redis.Del(global.Ctx, "product:all")

		response.Success(ctx, nil, "更新产品信息成功")
	}
}
