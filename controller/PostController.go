package controller

import (
	"gin_api/common"
	"gin_api/model"
	"gin_api/response"
	"gin_api/vo"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})

	return PostController{DB: db}
}

func (c PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest

	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil, "数据验证不通过！")
		return
	}

	// 获取登陆用户user 需要加上中间件auth
	user, _ := ctx.Get("user")

	// 创建post文章
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	// 插入数据
	if err := c.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	} else {
		response.Success(ctx, nil, "创建成功！")
	}
}

func (c PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	if err := c.DB.Preload("Category").Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章属于您，请勿非法操作")
		return
	}
	update_post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	// 更新文章
	if err := c.DB.Model(&post).Updates(update_post).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		log.Print(err.Error())
		return
	}

	response.Success(ctx, gin.H{"post": requestPost}, "更新成功")
}

func (c PostController) Show(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if err := c.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}
	response.Success(ctx, gin.H{"post": post}, "更新成功")
}

func (c PostController) Delete(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	if err := c.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章属于您，请勿非法操作")
		return
	}

	c.DB.Delete(&post)

	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (c PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	c.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int64
	c.DB.Model(model.Post{}).Count(&total)

	// 返回数据
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}
