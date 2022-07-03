package controller

import (
	"gin_api/common"
	"gin_api/model"
	"gin_api/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "數據驗證錯誤")
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.BindJSON(&requestCategory)

	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	c.DB.First(&updateCategory, categoryId)
	if updateCategory.ID == 0 {
		response.Fail(ctx, nil, "分類不存在")
		return
	}
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	c.DB.First(&category, categoryId)
	if category.ID == 0 {
		response.Fail(ctx, nil, "分類不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId); err != nil {
		response.Fail(ctx, nil, "刪除失敗")
		return
	}
	response.Success(ctx, gin.H{"category": ""}, "刪除成功")
}
