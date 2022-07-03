package router

import (
	"gin_api/controller"
	"gin_api/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRouter := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRouter.POST("", categoryController.Create)
	categoryRouter.PUT("/:id", categoryController.Update)
	categoryRouter.GET("/:id", categoryController.Show)
	categoryRouter.DELETE("/:id", categoryController.Delete)

	postRouter := r.Group("/posts")
	postRouter.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRouter.POST("", postController.Create)
	postRouter.PUT("/:id", postController.Update)
	postRouter.GET("/:id", postController.Show)
	postRouter.DELETE("/:id", postController.Delete)
	postRouter.GET("", postController.PageList)
	return r
}
